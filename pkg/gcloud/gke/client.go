package gke

import (
	"cloud.google.com/go/container/apiv1/containerpb"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/vietanhduong/pause-gcp/pkg/utils/exec"
	"github.com/vietanhduong/pause-gcp/pkg/utils/protoutil"
	"github.com/vietanhduong/pause-gcp/pkg/utils/sets"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"log"
	"strconv"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Client struct{}

func NewClient() *Client {
	client := &Client{}
	return client
}

func (c *Client) ListClusters(project string) ([]*Cluster, error) {
	raw, err := exec.Run(exec.Command("gcloud", "container", "clusters", "list", "--project", project, "--format", "json"))
	if err != nil {
		return nil, errors.Wrapf(err, "list clusters")
	}

	var tmp []any
	_ = json.UnmarshalFromString(raw, &tmp)

	clusters := make([]*containerpb.Cluster, len(tmp))

	for i, e := range tmp {
		b, _ := json.Marshal(e)
		clusters[i] = &containerpb.Cluster{}
		_ = protoutil.Unmarshal(b, clusters[i])
	}

	out := make([]*Cluster, len(clusters))
	var convert = func(i int, e *containerpb.Cluster) error {
		out[i] = &Cluster{
			Project:  project,
			Name:     e.GetName(),
			Location: e.GetLocation(),
		}
		out[i].NodePools = make([]*NodePool, len(e.GetNodePools()))
		for j, p := range e.GetNodePools() {
			out[i].NodePools[j] = &NodePool{
				Name:             p.GetName(),
				InstanceGroups:   p.GetInstanceGroupUrls(),
				Locations:        p.GetLocations(),
				InitialNodeCount: p.GetInitialNodeCount(),
				CurrentSize:      int32(getNodePoolSize(project, e.Name, p.Name)),
			}
			if a := p.GetAutoscaling(); a != nil {
				out[i].NodePools[j].Autoscaling = &Autoscaling{
					Enabled:           a.GetEnabled(),
					MinNodeCount:      a.GetMinNodeCount(),
					MaxNodeCount:      a.GetMaxNodeCount(),
					Autoprovisioned:   a.GetAutoprovisioned(),
					LocationPolicy:    a.GetLocationPolicy().String(),
					TotalMinNodeCount: a.GetTotalMinNodeCount(),
					TotalMaxNodeCount: a.GetTotalMaxNodeCount(),
				}
			}
		}
		return nil
	}
	var eg errgroup.Group
	for i, e := range clusters {
		eg.Go(func() error { return convert(i, e) })
	}
	_ = eg.Wait()
	return out, nil
}

func (c *Client) GetCluster(project, location, name string) (*Cluster, error) {
	raw, err := exec.Run(exec.Command("gcloud",
		"container",
		"clusters",
		"describe", name,
		"--project", project,
		"--location", location,
		"--format", "json"))
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("No cluster named '%s' in %s.", name, project)) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "get cluster")
	}

	var cluster containerpb.Cluster
	_ = protoutil.Unmarshal([]byte(raw), &cluster)

	out := &Cluster{
		Project:   project,
		Name:      cluster.GetName(),
		Location:  cluster.GetLocation(),
		NodePools: make([]*NodePool, len(cluster.GetNodePools())),
	}
	for i, p := range cluster.GetNodePools() {
		out.NodePools[i] = &NodePool{
			Name:             p.GetName(),
			InstanceGroups:   p.GetInstanceGroupUrls(),
			Locations:        p.GetLocations(),
			InitialNodeCount: p.GetInitialNodeCount(),
			CurrentSize:      int32(getNodePoolSize(project, cluster.Name, p.Name)),
		}
		if a := p.GetAutoscaling(); a != nil {
			out.NodePools[i].Autoscaling = &Autoscaling{
				Enabled:           a.GetEnabled(),
				MinNodeCount:      a.GetMinNodeCount(),
				MaxNodeCount:      a.GetMaxNodeCount(),
				Autoprovisioned:   a.GetAutoprovisioned(),
				LocationPolicy:    a.GetLocationPolicy().String(),
				TotalMinNodeCount: a.GetTotalMinNodeCount(),
				TotalMaxNodeCount: a.GetTotalMaxNodeCount(),
			}
		}
	}
	return out, nil
}

func (c *Client) PauseCluster(cluster *Cluster, exceptPools []string) error {
	var resize = func(cluster *Cluster, pool *NodePool) error {
		_, err := exec.Run(exec.Command("gcloud",
			"container",
			"clusters",
			"resize", cluster.Name,
			"--project", cluster.Project,
			"--location", cluster.Location,
			"--node-pool", pool.Name,
			"--num-nodes", "0",
			"--quiet",
		))
		if err != nil {
			return errors.Wrapf(err, "pause cluster: resize pool '%s/%s'", cluster.Name, pool.Name)
		}
		return nil
	}

	var pause = func(cluster *Cluster, pool *NodePool) error {
		_, err := exec.Run(exec.Command("gcloud",
			"container",
			"clusters",
			"update", cluster.Name,
			"--project", cluster.Project,
			"--location", cluster.Location,
			"--node-pool", pool.Name,
			"--no-enable-autoscaling",
			"--quiet",
		))
		if err != nil {
			return errors.Wrapf(err, "pause cluster: disable autoscaling '%s/%s'", cluster.Name, pool.Name)
		}
		log.Printf("INFO: disabled autoscaling of '%s/%s'\n", cluster.Name, pool.Name)

		ticker := time.NewTicker(time.Second)

		// resize node pool isn't completed at the first time. After disable the autoscaling, GCP set the nodeCount is
		// the initialNodeCount. Currently, we can't change the initialNodeCount setting.
		if err = resize(cluster, pool); err != nil {
			return err
		}
		// this will ensure that the nodeCount will be 0
		for {
			select {
			case <-ticker.C:
				if size := getNodePoolSize(cluster.Project, cluster.Name, pool.Name); size == 0 {
					log.Printf("INFO: resized pool for '%s/%s' is completed!\n", cluster.Name, pool.Name)
					return nil
				} else {
					log.Printf("INFO: current size of '%s/%s':%d\n", cluster.Name, pool.Name, size)
				}
				_, err = exec.Run(exec.Command("gcloud",
					"container",
					"clusters",
					"resize", cluster.Name,
					"--project", cluster.Project,
					"--location", cluster.Location,
					"--node-pool", pool.Name,
					"--num-nodes", "0",
					"--quiet",
				))
				if err != nil {
					return errors.Wrapf(err, "pause cluster: resize pool '%s/%s'", cluster.Name, pool.Name)
				}
			case <-context.Background().Done():
				return nil
			}
		}
	}

	except := sets.New(exceptPools...)

	var wg errgroup.Group
	for _, p := range cluster.NodePools {
		if except.Contains(p.Name) {
			continue
		}
		wg.Go(func() error { return pause(cluster, p) })
	}

	return wg.Wait()
}

func (c *Client) UnpauseCluster(cluster *Cluster, exceptPools []string) error {
	except := sets.New(exceptPools...)
	for _, p := range cluster.NodePools {
		if except.Contains(p.Name) {
			continue
		}

		_, err := exec.Run(exec.Command("gcloud",
			"container",
			"clusters",
			"update", cluster.Name,
			"--project", cluster.Project,
			"--location", cluster.Location,
			"--node-pool", p.Name,
			"--enable-autoscaling",
			"--min-nodes", strconv.Itoa(int(p.Autoscaling.MinNodeCount)),
			"--max-nodes", strconv.Itoa(int(p.Autoscaling.MaxNodeCount)),
			"--quiet",
		))
		if err != nil {
			return errors.Wrapf(err, "unpause cluster: enable autoscaling '%s/%s'", cluster.Name, p.Name)
		}
		log.Printf("INFO: enabled autoscaling for '%s/%s'\n", cluster.Name, p.Name)

		_, err = exec.Run(exec.Command("gcloud",
			"container",
			"clusters",
			"resize", cluster.Name,
			"--project", cluster.Project,
			"--location", cluster.Location,
			"--node-pool", p.Name,
			"--num-nodes", strconv.Itoa(int(p.CurrentSize)),
			"--quiet",
		))
		if err != nil {
			return errors.Wrapf(err, "unpause cluster: resize pool '%s/%s'", cluster.Name, p.Name)
		}
		log.Printf("INFO: resized '%s/%s' to %d\n", cluster.Name, p.Name, p.CurrentSize)
	}

	log.Printf("INFO: unpause cluster '%s' is completed!\n", cluster.Name)
	return nil
}

func getNodePoolSize(project, cluster, pool string) int {
	out, err := exec.Run(exec.Command("gcloud",
		"compute",
		"instance-groups",
		"list",
		"--filter", fmt.Sprintf("name:gke-%s-%s-*", cluster, pool),
		"--project", project,
		"--format", "value(size)",
	))

	if err != nil {
		log.Printf("WARN: get pool size of '%s/%s' got error: %v\n", cluster, pool, err)
		return 0
	}
	val, _ := strconv.Atoi(out)
	return val
}
