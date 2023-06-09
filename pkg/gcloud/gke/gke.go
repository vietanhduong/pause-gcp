package gke

import (
	"fmt"
	"log"
	"strings"
	"time"

	compute_v1 "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	container_v1 "cloud.google.com/go/container/apiv1"
	"cloud.google.com/go/container/apiv1/containerpb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"

	"github.com/googleapis/gax-go/v2/apierror"
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

type Client struct {
	options Options
}

func NewClient(o Options) *Client {
	return &Client{options: o}
}

// ListClusters return all clusters is running in the input project
func (c *Client) ListClusters(project string) ([]*apis.Cluster, error) {
	conn, err := c.newClusterConn()
	if err != nil {
		return nil, errors.Wrap(err, "list clusters")
	}
	defer conn.Close()

	igConn, err := c.newInstanceGroupsConn()
	if err != nil {
		return nil, errors.Wrap(err, "list clusters")
	}
	defer igConn.Close()

	resp, err := conn.ListClusters(context.TODO(), &containerpb.ListClustersRequest{
		Parent: fmt.Sprintf("projects/%s/locations/-", project),
	})
	if err != nil {
		return nil, errors.Wrap(err, "list clusters")
	}
	out := make([]*apis.Cluster, len(resp.GetClusters()))

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
				CurrentSize:      int32(getNodePoolSize(igConn, project, e.GetName(), p.GetName(), p.GetLocations())),
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
	for i, e := range resp.GetClusters() {
		i, e := i, e
		eg.Go(func() error { return convert(i, e) })
	}
	_ = eg.Wait()
	return out, nil
}

// GetCluster return a cluster if exists or nil otherwise
func (c *Client) GetCluster(project, location, name string) (*apis.Cluster, error) {
	conn, err := c.newClusterConn()
	if err != nil {
		return nil, errors.Wrap(err, "get cluster")
	}
	defer conn.Close()

	igConn, err := c.newInstanceGroupsConn()
	if err != nil {
		return nil, errors.Wrap(err, "get cluster")
	}
	defer igConn.Close()

	req := &containerpb.GetClusterRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/clusters/%s", project, location, name),
	}
	cluster, err := conn.GetCluster(context.TODO(), req)
	if err != nil {
		var apiErr *apierror.APIError
		if errors.As(err, &apiErr) && apiErr.GRPCStatus().Code() == codes.NotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "get cluster")
	}
	out := &apis.Cluster{
		Project:  project,
		Name:     cluster.GetName(),
		Location: cluster.GetLocation(),
	}
	out.NodePools = make([]*apis.Cluster_NodePool, len(cluster.GetNodePools()))
	for i, p := range cluster.GetNodePools() {
		out.NodePools[i] = &apis.Cluster_NodePool{
			Name:             p.GetName(),
			InstanceGroups:   p.GetInstanceGroupUrls(),
			Locations:        p.GetLocations(),
			InitialNodeCount: p.GetInitialNodeCount(),
			CurrentSize:      int32(getNodePoolSize(igConn, project, cluster.GetName(), p.GetName(), p.GetLocations())),
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

// PauseCluster pasue the input cluster. This function will pause all node pools in the input cluster
// regradless the current pools in the current cluster. This will helpful when you can ignore the pause
// action with some pools.
func (c *Client) PauseCluster(cluster *apis.Cluster) error {
	conn, err := c.newClusterConn()
	if err != nil {
		return errors.Wrap(err, "list clusters")
	}
	defer conn.Close()

	igConn, err := c.newInstanceGroupsConn()
	if err != nil {
		return errors.Wrap(err, "list clusters")
	}
	defer igConn.Close()

	pause := func(cluster *apis.Cluster, pool *apis.Cluster_NodePool) error {
		if pool.GetAutoscaling() != nil || pool.GetAutoscaling().GetEnabled() {
			op, err := conn.SetNodePoolAutoscaling(context.Background(), &containerpb.SetNodePoolAutoscalingRequest{
				Name:        fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", cluster.GetProject(), cluster.GetLocation(), cluster.GetName(), pool.GetName()),
				Autoscaling: &containerpb.NodePoolAutoscaling{Enabled: false},
			})
			if err != nil {
				log.Printf("WARN: disable node pool autoscaling for '%s/%s' failed\n", cluster.GetName(), cluster.GetLocation())
				return errors.Wrapf(err, "pause '%s/%s'", cluster.GetName(), cluster.GetLocation())
			}
			if err = waitOp(conn, op); err != nil {
				return errors.Wrapf(err, "pause '%s/%s'", cluster.GetName(), cluster.GetLocation())
			}
			log.Printf("INFO: disabled autoscaling for '%s/%s'\n", cluster.Name, pool.Name)
		}

		if err := resize(conn, igConn, cluster, pool, 0); err != nil {
			return errors.Wrap(err, "pause cluster")
		}
		return nil
	}

	defer func() {
		if err != nil {
			log.Printf("WARN: pause cluster '%s/%s' failed: %v\n", cluster.GetLocation(), cluster.GetName(), err)
			log.Printf("INFO: error detected, prepare to rollback cluster '%s/%s'.", cluster.GetLocation(), cluster.GetName())
			_ = c.UnpauseCluster(cluster)
		}
	}()

	for _, p := range cluster.NodePools {
		if err = pause(cluster, p); err != nil {
			return err
		}
	}

	return nil
}

// UnpauseCluster unpause the input cluster. Similar with the PauseCluster function. This function
// just unpause the pools are presented in the input clusters.
func (c *Client) UnpauseCluster(cluster *apis.Cluster) error {
	conn, err := c.newClusterConn()
	if err != nil {
		return errors.Wrap(err, "list clusters")
	}
	defer conn.Close()

	igConn, err := c.newInstanceGroupsConn()
	if err != nil {
		return errors.Wrap(err, "list clusters")
	}
	defer igConn.Close()
	unpause := func(cluster *apis.Cluster, p *apis.Cluster_NodePool) error {
		if err := waitClusterOperation(conn, cluster); err != nil {
			return errors.Wrap(err, "unpaise cluster")
		}

		// re-enable autoscaling if the setting is presented
		if p.GetAutoscaling() != nil && p.GetAutoscaling().GetEnabled() {
			op, err := conn.SetNodePoolAutoscaling(context.Background(), &containerpb.SetNodePoolAutoscalingRequest{
				Name: fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", cluster.GetProject(), cluster.GetLocation(), cluster.GetName(), p.GetName()),
				Autoscaling: &containerpb.NodePoolAutoscaling{
					Enabled:      true,
					MinNodeCount: p.GetAutoscaling().GetMinNodeCount(),
					MaxNodeCount: p.GetAutoscaling().GetMaxNodeCount(),
				},
			})
			if err != nil {
				log.Printf("WARN: enable node pool autoscaling for '%s/%s' failed\n", cluster.GetName(), cluster.GetLocation())
				return errors.Wrapf(err, "unpause '%s/%s'", cluster.GetName(), cluster.GetLocation())
			}
			if err = waitOp(conn, op); err != nil {
				return errors.Wrapf(err, "unpause '%s/%s'", cluster.GetName(), cluster.GetLocation())
			}
			log.Printf("INFO: enabled autoscaling for '%s/%s'\n", cluster.Name, p.Name)
		}

		if err := resize(conn, igConn, cluster, p, int(p.GetCurrentSize())); err != nil {
			return errors.Wrap(err, "unpause cluster")
		}
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

// RefreshCluster recreate the worker nodes of the input clusters. This function will ignore
// with a node pool has not `spot` or `preemptible` type
func (c *Client) RefreshCluster(cluster *apis.Cluster) error {
	conn, err := c.newManagedInstanceGroupConn()
	if err != nil {
		return errors.Wrap(err, "refresh cluster")
	}
	defer conn.Close()

	refresh := func(mig *computepb.InstanceGroupManager) error {
		instances, err := getMangedInstanceNames(conn, cluster.GetProject(), mig)
		if err != nil {
			return err
		}

		req := &computepb.RecreateInstancesInstanceGroupManagerRequest{
			InstanceGroupManager: mig.GetName(),
			Project:              cluster.GetProject(),
			Zone:                 basename(mig.GetZone()),
			InstanceGroupManagersRecreateInstancesRequestResource: &computepb.InstanceGroupManagersRecreateInstancesRequest{
				Instances: instances,
			},
		}
		_, err = conn.RecreateInstances(context.Background(), req)
		if err != nil {
			log.Printf("WARN: refresh MIG %q got error: %v\n", mig.GetName(), err)
		}
		return err
	}

	getMigs := func(pool, z string) ([]*computepb.InstanceGroupManager, error) {
		filterQuery := fmt.Sprintf("name:gke-%s-%s-*", cluster.GetName(), pool)
		req := &computepb.ListInstanceGroupManagersRequest{
			Project: cluster.GetProject(),
			Filter:  &filterQuery,
			Zone:    z,
		}
		it := conn.List(context.Background(), req)
		var out []*computepb.InstanceGroupManager
		for {
			ig, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}
			out = append(out, ig)
		}
		return out, nil
	}

	// find manged instance groups of the cluster
	var igs []*computepb.InstanceGroupManager
	for _, p := range cluster.GetNodePools() {
		if !p.GetPreemptible() && !p.GetSpot() {
			log.Printf("INFO: pool %q has been ignored because this pool is on-demand pool.", p.GetName())
			continue
		}
		for _, z := range p.GetLocations() {
			tmp, err := getMigs(p.GetName(), z)
			if err != nil {
				return err
			}
			igs = append(igs, tmp...)
		}
	}
	var eg errgroup.Group
	for _, ig := range igs {
		ig := ig
		eg.Go(func() error { return refresh(ig) })
	}
	return eg.Wait()
}

// resize the input node pool
func resize(clusterConn *container_v1.ClusterManagerClient, igConn *compute_v1.InstanceGroupsClient, cluster *apis.Cluster, pool *apis.Cluster_NodePool, size int) error {
	_resize := func() (bool, error) {
		currentSize := getNodePoolSize(igConn, cluster.GetProject(), cluster.GetName(), pool.GetName(), pool.GetLocations())
		if size == 0 && size == currentSize {
			log.Printf("INFO: node pool '%s/%s' has been resized to 0!", cluster.GetName(), pool.GetName())
			return true, nil
		}
		if currentSize >= size && size > 0 {
			log.Printf("INFO: node pool '%s/%s' has been resized to %d (current=%d)!\n", cluster.GetName(), pool.GetName(), size, currentSize)
			return true, nil
		} else {
			log.Printf("INFO: resizing node pool '%s/%s'! current=%d; expect=%d\n", cluster.GetName(), pool.GetName(), currentSize, size)
		}
		if err := waitClusterOperation(clusterConn, cluster); err != nil {
			return false, errors.Wrap(err, "resize")
		}

		req := &containerpb.SetNodePoolSizeRequest{
			Name:      fmt.Sprintf("projects/%s/locations/%s/clusters/%s/nodePools/%s", cluster.GetProject(), cluster.GetLocation(), cluster.GetName(), pool.GetName()),
			NodeCount: int32(size),
		}
		op, err := clusterConn.SetNodePoolSize(context.Background(), req)
		if err != nil {
			return false, errors.Wrap(err, "resize")
		}
		return false, waitOp(clusterConn, op)
	}
	var stop bool
	var err error
	if stop, err = _resize(); err != nil || stop {
		return err
	}

	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			if stop, err = _resize(); err != nil || stop {
				return err
			}
		case <-context.Background().Done():
			return nil
		}
	}
}

func (c *Client) newClusterConn() (*container_v1.ClusterManagerClient, error) {
	var opts []option.ClientOption
	if c.options.Credentials != "" {
		opts = append(opts, option.WithCredentialsFile(c.options.Credentials))
	}
	return container_v1.NewClusterManagerClient(context.Background(), opts...)
}

func (c *Client) newManagedInstanceGroupConn() (*compute_v1.InstanceGroupManagersClient, error) {
	var opts []option.ClientOption
	if c.options.Credentials != "" {
		opts = append(opts, option.WithCredentialsFile(c.options.Credentials))
	}
	return compute_v1.NewInstanceGroupManagersRESTClient(context.Background(), opts...)
}

func (c *Client) newInstanceGroupsConn() (*compute_v1.InstanceGroupsClient, error) {
	var opts []option.ClientOption
	if c.options.Credentials != "" {
		opts = append(opts, option.WithCredentialsFile(c.options.Credentials))
	}
	return compute_v1.NewInstanceGroupsRESTClient(context.Background(), opts...)
}

// getMangedInstanceNames return the url of managed instances by the input MIG
func getMangedInstanceNames(conn *compute_v1.InstanceGroupManagersClient, project string, mig *computepb.InstanceGroupManager) ([]string, error) {
	req := &computepb.ListManagedInstancesInstanceGroupManagersRequest{
		InstanceGroupManager: mig.GetName(),
		Project:              project,
		Zone:                 basename(mig.GetZone()),
	}
	var names []string
	it := conn.ListManagedInstances(context.Background(), req)
	for {
		instance, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, errors.Wrap(err, "get managed instance name")
		}
		names = append(names, instance.GetInstance())
	}
	return names, nil
}

func getNodePoolSize(conn *compute_v1.InstanceGroupsClient, project, cluster, pool string, zones []string) int {
	ret := make([]int, len(zones))
	getSize := func(i int, z string) error {
		filterQuery := fmt.Sprintf("name:gke-%s-%s-*", cluster, pool)
		req := &computepb.ListInstanceGroupsRequest{
			Project: project,
			Filter:  &filterQuery,
			Zone:    z,
		}
		it := conn.List(context.Background(), req)
		for {
			resp, err := it.Next()
			if err == nil || err == iterator.Done {
				ret[i] = int(resp.GetSize())
				return nil
			}
			if err != nil {
				log.Printf("WARN: get pool size '%s/%s/%s' got error: %v\n", project, cluster, pool, err)
				return nil
			}
		}
	}
	var eg errgroup.Group
	for i, z := range zones {
		i, z := i, z
		eg.Go(func() error { return getSize(i, z) })
	}
	_ = eg.Wait()

	var size int
	for _, val := range ret {
		size += val
	}
	return size
}

func waitClusterOperation(conn *container_v1.ClusterManagerClient, cluster *apis.Cluster) error {
	req := &containerpb.ListOperationsRequest{Parent: fmt.Sprintf("projects/%s/locations/%s", cluster.GetProject(), cluster.GetLocation())}
	resp, err := conn.ListOperations(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "wait cluster operation")
	}
	wait := func(op *containerpb.Operation) error {
		if !strings.Contains(op.GetTargetLink(), fmt.Sprintf("clusters/%s", cluster.GetName())) || op.GetStatus() == containerpb.Operation_DONE {
			return nil
		}
		if err := waitOp(conn, op); err != nil {
			return errors.Wrapf(err, "wait operation %q", op.GetName())
		}
		log.Printf("INFO: handle operation '%s/%s' has been completed\n", cluster.GetName(), op.GetName())
		return nil
	}

	var eg errgroup.Group
	for _, op := range resp.GetOperations() {
		op := op
		eg.Go(func() error { return wait(op) })
	}
	return eg.Wait()
}

func waitOp(conn *container_v1.ClusterManagerClient, op *containerpb.Operation) error {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			name := op.GetSelfLink()
			name = strings.TrimPrefix(name, "https://container.googleapis.com/v1alpha1/")
			name = strings.TrimPrefix(name, "https://container.googleapis.com/v1/")

			tmp, err := conn.GetOperation(context.Background(), &containerpb.GetOperationRequest{Name: name})
			if err != nil {
				return errors.Wrapf(err, "wait operation %q", op.GetName())
			}
			if tmp.GetStatus() == containerpb.Operation_DONE {
				return nil
			}
		case <-context.Background().Done():
			return nil
		}
	}
}

func basename(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}
