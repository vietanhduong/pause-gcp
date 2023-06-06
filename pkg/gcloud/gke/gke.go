package gke

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/container/apiv1/containerpb"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/pkg/utils/exec"
	"github.com/vietanhduong/pause-gcp/pkg/utils/protoutil"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ListClusters(project string) ([]*apis.Cluster, error) {
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

	out := make([]*apis.Cluster, len(clusters))
	convert := func(i int, e *containerpb.Cluster) error {
		out[i] = &apis.Cluster{
			Project:  project,
			Name:     e.GetName(),
			Location: e.GetLocation(),
		}
		out[i].NodePools = make([]*apis.Cluster_NodePool, len(e.GetNodePools()))
		for j, p := range e.GetNodePools() {
			out[i].NodePools[j] = &apis.Cluster_NodePool{
				Name:             p.GetName(),
				InstanceGroups:   p.GetInstanceGroupUrls(),
				Locations:        p.GetLocations(),
				InitialNodeCount: p.GetInitialNodeCount(),
				CurrentSize:      int32(getNodePoolSize(project, e.Name, p.Name)),
				Spot:             p.GetConfig().GetSpot(),
				Preemptible:      p.GetConfig().GetPreemptible(),
			}
			if a := p.GetAutoscaling(); a != nil {
				out[i].NodePools[j].Autoscaling = &apis.Cluster_NodePool_AutoScaling{
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
		i, e := i, e
		eg.Go(func() error { return convert(i, e) })
	}
	_ = eg.Wait()
	return out, nil
}

func GetCluster(project, location, name string) (*apis.Cluster, error) {
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

	out := &apis.Cluster{
		Project:   project,
		Name:      cluster.GetName(),
		Location:  cluster.GetLocation(),
		NodePools: make([]*apis.Cluster_NodePool, len(cluster.GetNodePools())),
	}
	for i, p := range cluster.GetNodePools() {
		out.NodePools[i] = &apis.Cluster_NodePool{
			Name:             p.GetName(),
			InstanceGroups:   p.GetInstanceGroupUrls(),
			Locations:        p.GetLocations(),
			InitialNodeCount: p.GetInitialNodeCount(),
			CurrentSize:      int32(getNodePoolSize(project, cluster.Name, p.Name)),
			Spot:             p.GetConfig().GetSpot(),
			Preemptible:      p.GetConfig().GetPreemptible(),
		}
		if a := p.GetAutoscaling(); a != nil {
			out.NodePools[i].Autoscaling = &apis.Cluster_NodePool_AutoScaling{
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

func PauseCluster(cluster *apis.Cluster, dryRun bool) error {
	resize := func(cluster *apis.Cluster, pool *apis.Cluster_NodePool) error {
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

	pause := func(cluster *apis.Cluster, pool *apis.Cluster_NodePool) error {
		if pool.GetAutoscaling() != nil || pool.GetAutoscaling().GetEnabled() {
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
		}

		if dryRun || pool.GetCurrentSize() == 0 {
			return nil
		}

		ticker := time.NewTicker(time.Second)

		// resize node pool isn't completed at the first time. After disable the autoscaling, GCP set the nodeCount is
		// the initialNodeCount. Currently, we can't change the initialNodeCount setting.
		if err := resize(cluster, pool); err != nil {
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
					log.Printf("INFO: current size of '%s/%s': %d\n", cluster.Name, pool.Name, size)
				}
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
			case <-context.Background().Done():
				return nil
			}
		}
	}

	var err error
	defer func() {
		if err != nil {
			log.Printf("WARN: pause cluster '%s/%s' failed: %v\n", cluster.GetLocation(), cluster.GetName(), err)
			log.Printf("INFO: error detected, prepare to rollback cluster '%s/%s'.", cluster.GetLocation(), cluster.GetName())
			_ = UnpauseCluster(cluster)
		}
	}()

	for _, p := range cluster.NodePools {
		if err = pause(cluster, p); err != nil {
			return err
		}
	}

	return nil
}

func UnpauseCluster(cluster *apis.Cluster) error {
	unpause := func(cluster *apis.Cluster, p *apis.Cluster_NodePool) error {
		if p.GetAutoscaling() != nil && p.GetAutoscaling().GetEnabled() {
			_, err := exec.Run(exec.Command("gcloud",
				"container",
				"clusters",
				"update", cluster.Name,
				"--project", cluster.Project,
				"--location", cluster.Location,
				"--node-pool", p.GetName(),
				"--enable-autoscaling",
				"--min-nodes", strconv.Itoa(int(p.GetAutoscaling().GetMinNodeCount())),
				"--max-nodes", strconv.Itoa(int(p.GetAutoscaling().GetMaxNodeCount())),
				"--quiet",
			))
			if err != nil {
				return errors.Wrapf(err, "unpause cluster: enable autoscaling '%s/%s'", cluster.Name, p.Name)
			}
			log.Printf("INFO: enabled autoscaling for '%s/%s'\n", cluster.Name, p.Name)
		}

		_, err := exec.Run(exec.Command("gcloud",
			"container",
			"clusters",
			"resize", cluster.Name,
			"--project", cluster.Project,
			"--location", cluster.Location,
			"--node-pool", p.GetName(),
			"--num-nodes", strconv.Itoa(int(p.GetCurrentSize())),
			"--quiet",
		))
		if err != nil {
			return errors.Wrapf(err, "unpause cluster: resize pool '%s/%s'", cluster.Name, p.Name)
		}
		log.Printf("INFO: resized '%s/%s' to %d\n", cluster.Name, p.GetName(), p.GetCurrentSize())
		return nil
	}

	for _, p := range cluster.NodePools {
		if err := unpause(cluster, p); err != nil {
			return err
		}
	}
	log.Printf("INFO: unpause cluster '%s' is completed!\n", cluster.Name)
	return nil
}

type instanceGroup struct {
	Name      string `json:"name,omitempty"`
	Zone      string `json:"zone,omitempty"`
	Region    string `json:"region,omitempty"`
	IsManaged string `json:"isManaged,omitempty"`
}

func RefreshCluster(cluster *apis.Cluster, recreate bool) error {
	refresh := func(project string, ig instanceGroup) error {
		if ig.IsManaged != "Yes" {
			log.Printf("INFO: instance group %s has been ignored becase it's not managed\n", ig.Name)
			return nil
		}
		locationFlag := "--zone"
		location := ig.Zone
		if ig.Region != "" {
			locationFlag = "--region"
			location = ig.Region
		}

		replaceMethod := "substitute"
		if recreate {
			replaceMethod = "recreate"
		}

		_, err := exec.Run(exec.Command("gcloud",
			"compute",
			"instance-groups",
			"managed",
			"rolling-action",
			"replace",
			ig.Name,
			"--project", project,
			locationFlag, location,
			"--replacement-method", replaceMethod,
			"--max-surge", "0",
			"--max-unavailable", "1"))
		if err != nil {
			log.Printf("WARN: refresh instance group %q got error: %v\n", ig.Name, err)
			return err
		}
		return nil
	}

	execute := func(cluster *apis.Cluster, pool *apis.Cluster_NodePool) error {
		log.Printf("INFO: prepare to refresh pool %s...\n", pool.GetName())
		defer log.Printf("INFO: refresh pool %s completed!\n", pool.GetName())
		if !pool.GetPreemptible() && !pool.GetSpot() {
			log.Printf("INFO: pool %q has been ignored because this pool is on-demand pool.", pool.GetName())
			return nil
		}
		raw, err := exec.Run(exec.Command("gcloud",
			"compute",
			"instance-groups",
			"list",
			"--project", cluster.GetProject(),
			"--filter", fmt.Sprintf("name:gke-%s-%s-*", cluster.GetName(), pool.GetName()),
			"--format", "json(name, zone.basename(), region.basename(), isManaged)"))
		if err != nil {
			log.Printf("WARN: list instance group (cluster=%s, pool=%s) got error: %v\n", cluster.GetName(), pool.GetName(), err)
			return err
		}

		var igs []instanceGroup
		if err = json.UnmarshalFromString(raw, &igs); err != nil {
			log.Printf("WARN: unmarshal instance groups (cluster=%s, pool=%s) got error: %v\n", cluster.GetName(), pool.GetName(), err)
			return err
		}

		var eg errgroup.Group
		for _, ig := range igs {
			ig := ig
			eg.Go(func() error { return refresh(cluster.GetProject(), ig) })
		}

		return eg.Wait()
	}

	var eg errgroup.Group
	for _, p := range cluster.GetNodePools() {
		p := p
		eg.Go(func() error { return execute(cluster, p) })
	}
	return eg.Wait()
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
