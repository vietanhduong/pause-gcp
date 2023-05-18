package gke

import (
	container "cloud.google.com/go/container/apiv1"
	"cloud.google.com/go/container/apiv1/containerpb"
	"fmt"
	"github.com/googleapis/gax-go/v2/apierror"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/vietanhduong/pause-gke-cluster/pkg/exec"
	"github.com/vietanhduong/pause-gke-cluster/pkg/gcloud/options"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"log"
	"os"
	"path"
	"strconv"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const (
	DefaultBackupStateSuffix = ".backup_state.json"
)

type Client struct {
	Opts options.Options
}

func NewClient(opt ...options.Option) *Client {
	client := &Client{}
	client.Opts.BackupStatePath = "."
	for _, o := range opt {
		o(&client.Opts)
	}
	return client
}

func (c *Client) ListClusters(project string) ([]*Cluster, error) {
	conn, err := c.newClusterClient()
	if err != nil {
		return nil, errors.Wrapf(err, "list clusters: new client")
	}
	defer conn.Close()

	resp, err := conn.ListClusters(context.TODO(), &containerpb.ListClustersRequest{
		Parent: fmt.Sprintf("projects/%s/locations/-", project),
	})

	if err != nil {
		return nil, errors.Wrapf(err, "list clusters: make request")
	}

	out := make([]*Cluster, len(resp.GetClusters()))
	for i, e := range resp.GetClusters() {
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
	}

	return out, nil
}

func (c *Client) GetCluster(project, location, name string) (*Cluster, error) {
	conn, err := c.newClusterClient()
	if err != nil {
		return nil, errors.Wrapf(err, "get cluster: new client")
	}
	defer conn.Close()

	resp, err := conn.GetCluster(context.TODO(), &containerpb.GetClusterRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/clusters/%s", project, location, name),
	})

	if err != nil {
		var apiErr *apierror.APIError
		if errors.As(err, &apiErr) {
			if apiErr.GRPCStatus().Code() == codes.NotFound {
				log.Printf("WARN: cluster %q not found", name)
				return nil, nil
			}
			return nil, errors.New(fmt.Sprintf("get cluster request: %s", apiErr.GRPCStatus().Message()))
		}
		return nil, errors.Wrapf(err, "get cluster: make request")
	}

	cluster := &Cluster{
		Project:  project,
		Name:     resp.GetName(),
		Location: resp.GetLocation(),
	}
	cluster.NodePools = make([]*NodePool, len(resp.GetNodePools()))
	for i, p := range resp.GetNodePools() {
		cluster.NodePools[i] = &NodePool{
			Name:             p.GetName(),
			InstanceGroups:   p.GetInstanceGroupUrls(),
			Locations:        p.GetLocations(),
			InitialNodeCount: p.GetInitialNodeCount(),
			CurrentSize:      int32(getNodePoolSize(cluster.Project, cluster.Name, p.Name)),
		}
		if a := p.GetAutoscaling(); a != nil {
			cluster.NodePools[i].Autoscaling = &Autoscaling{
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

	return cluster, nil
}

func (c *Client) PauseCluster(req PauseClusterRequest) error {
	cluster, err := c.GetCluster(req.Project, req.Location, req.ClusterName)
	if err != nil {
		return errors.Wrapf(err, "pause cluster: get cluster")
	}

	if cluster == nil {
		return errors.New(fmt.Sprintf("cluster %q not found", req.ClusterName))
	}

	conn, err := c.newClusterClient()
	if err != nil {
		return errors.Wrapf(err, "pause cluster: new connection")
	}
	defer conn.Close()

	var resize = func(cluster *Cluster, pool *NodePool) error {
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
		return nil
	}

	var pause = func(cluster *Cluster, pool *NodePool) error {
		_, err = exec.Run(exec.Command("gcloud",
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

	var wg errgroup.Group
	for _, p := range cluster.NodePools {
		wg.Go(func() error { return pause(cluster, p) })
	}

	if err = wg.Wait(); err != nil {
		return err
	}

	// write the previous state to the backup state file
	b, _ := json.Marshal(cluster)
	log.Printf("INFO: Current State of the cluster %q (this message will be used in case the writing process has a problem):\n%s", cluster.Name, string(b))

	if err = os.WriteFile(path.Join(c.Opts.BackupStatePath, genBackupFile(cluster)), b, 0644); err != nil {
		return errors.Wrapf(err, "pause cluster: write backup state:")
	}
	log.Printf("INFO: pause cluster '%s' is completed!\n", cluster.Name)
	return nil
}

func (c *Client) UnpauseCluster(req UnpauseClusterRequest) error {
	// load backup state
	backupPath := path.Join(c.Opts.BackupStatePath, genBackupFile(&Cluster{Project: req.Project, Location: req.Location, Name: req.ClusterName}))
	b, err := os.ReadFile(backupPath)
	if err != nil {
		log.Printf("WARN: read backup state file %q got error: %v\n", c.Opts.BackupStatePath, err)
	}

	var previous Cluster
	if err = json.Unmarshal(b, &previous); err != nil {
		log.Printf("WARN: unmarshal backup state got error: %v\n", err)
	}

	conn, err := c.newClusterClient()
	if err != nil {
		return errors.Wrapf(err, "unpause cluster: new connection")
	}
	defer conn.Close()

	for _, p := range previous.NodePools {
		_, err = exec.Run(exec.Command("gcloud",
			"container",
			"clusters",
			"update", previous.Name,
			"--project", previous.Project,
			"--location", previous.Location,
			"--node-pool", p.Name,
			"--enable-autoscaling",
			"--min-nodes", strconv.Itoa(int(p.Autoscaling.MinNodeCount)),
			"--max-nodes", strconv.Itoa(int(p.Autoscaling.MaxNodeCount)),
			"--quiet",
		))
		if err != nil {
			return errors.Wrapf(err, "unpause cluster: enable autoscaling '%s/%s'", previous.Name, p.Name)
		}
		log.Printf("INFO: enabled autoscaling for '%s/%s'\n", previous.Name, p.Name)

		_, err = exec.Run(exec.Command("gcloud",
			"container",
			"clusters",
			"resize", previous.Name,
			"--project", previous.Project,
			"--location", previous.Location,
			"--node-pool", p.Name,
			"--num-nodes", strconv.Itoa(int(p.CurrentSize)),
			"--quiet",
		))
		if err != nil {
			return errors.Wrapf(err, "unpause cluster: resize pool '%s/%s'", previous.Name, p.Name)
		}
		log.Printf("INFO: resized '%s/%s' to %d\n", previous.Name, p.Name, p.CurrentSize)
	}

	log.Printf("INFO: unpause cluster '%s' is completed!\n", previous.Name)
	_ = os.Remove(backupPath)
	return nil
}

func (c *Client) newClusterClient() (*container.ClusterManagerClient, error) {
	var opts []option.ClientOption
	if c.Opts.CredentialsFilePath != "" {
		opts = append(opts, option.WithCredentialsFile(c.Opts.CredentialsFilePath))
	}
	return container.NewClusterManagerClient(context.TODO(), opts...)
}

func genBackupFile(cluster *Cluster) string {
	return fmt.Sprintf("%s_%s_%s%s", cluster.Project, cluster.Location, cluster.Name, DefaultBackupStateSuffix)
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
