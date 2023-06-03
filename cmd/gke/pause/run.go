package pause

import (
	"fmt"
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/gke"
	"github.com/vietanhduong/pause-gcp/pkg/utils/exec"
	"github.com/vietanhduong/pause-gcp/pkg/utils/sets"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"os"
	"strings"
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
		dst := fmt.Sprintf("%s/gke_%s_%s_%s.state.json", strings.TrimSuffix(cfg.outputDir, "/"), cfg.project, cfg.location, cfg.clusterName)
		if strings.HasPrefix(strings.ToLower(cfg.outputDir), "gs://") {
			_, err = exec.Run(exec.Command("bash", "-c", fmt.Sprintf(`echo '%s' | gsutil cp -L /dev/null - %s`, string(b), dst)))
			if err != nil {
				return err
			}
		} else {
			_ = os.MkdirAll(cfg.outputDir, 0755)
			if err = os.WriteFile(dst, b, 0644); err != nil {
				return err
			}
		}
		log.Printf("INFO: Cluster's state has been written to %q", dst)
	}
	log.Printf("INFO: cluster %q has been paused!", cluster.GetName())
	return nil
}
