package refresh

import (
	"log"

	"github.com/pkg/errors"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/gke"
)

type runConfig struct {
	name     string
	project  string
	location string
}

func run(cfg runConfig) error {
	gkeClient := gke.NewClient(gke.Options{})
	cluster, err := gkeClient.GetCluster(cfg.project, cfg.location, cfg.name)
	if err != nil {
		return err
	}
	if cluster == nil {
		return errors.Errorf("cluster '%s/%s/%s' not found", cfg.project, cfg.location, cfg.name)
	}
	if err = gkeClient.RefreshCluster(cluster); err != nil {
		return err
	}
	log.Printf("INFO: cluster '%s' has been refreshed!", cfg.name)
	return nil
}
