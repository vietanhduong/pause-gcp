package unpause

import (
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/gke"
	"log"
)

type runConfig struct {
	cluster *apis.Cluster
}

func run(cfg runConfig) error {
	cluster, err := gke.GetCluster(cfg.cluster.GetProject(), cfg.cluster.GetLocation(), cfg.cluster.GetName())
	if err != nil {
		return err
	}
	if cluster == nil {
		return errors.Errorf("cluster '%s/%s/%s' not found", cfg.cluster.GetProject(), cfg.cluster.GetLocation(), cfg.cluster.GetName())
	}
	if err = gke.UnpauseCluster(cfg.cluster); err != nil {
		return err
	}
	log.Printf("INFO: cluster '%s' is running now!", cfg.cluster.GetName())
	return nil
}
