package pause

import (
	"fmt"
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/gke"
	"github.com/vietanhduong/pause-gcp/pkg/utils/sets"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"os"
	"path"
)

type runConfig struct {
	project     string
	location    string
	clusterName string
	outputDir   string
	exceptPools []string
}

func run(cfg runConfig) error {
	cluster, err := gke.GetCluster(cfg.project, cfg.location, cfg.clusterName)
	if err != nil {
		return err
	}
	if cluster == nil {
		return errors.Errorf("cluster '%s/%s/%s' not found", cfg.project, cfg.location, cfg.clusterName)
	}

	exceptPools := sets.New(cfg.exceptPools...)
	var pools []*apis.Cluster_NodePool
	for _, p := range cluster.GetNodePools() {
		if !exceptPools.Contains(p.GetName()) {
			pools = append(pools, p)
		}
	}
	cluster.NodePools = pools
	if err = gke.PauseCluster(cluster, false); err != nil {
		return err
	}

	b, _ := protojson.MarshalOptions{Indent: "    "}.Marshal(cluster)
	_, _ = fmt.Fprintf(os.Stdout, "%s\n", b)
	log.Printf("Recommend: please keep this information. You can use it to restore (unpause) your cluster.")

	if cfg.outputDir != "" {
		_ = os.MkdirAll(cfg.outputDir, 0755)
		dst := path.Join(cfg.outputDir, fmt.Sprintf("gke_%s_%s_%s.state.json", cfg.project, cfg.location, cfg.clusterName))
		if err = os.WriteFile(dst, b, 0644); err != nil {
			return err
		}

		log.Printf("INFO: Cluster's state has been written to %q", dst)
	}
	log.Printf("INFO: cluster %q has been paused!", cluster.GetName())
	return nil
}
