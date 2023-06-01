package refresh

import (
	"github.com/pkg/errors"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/gke"
	"log"
)

type runConfig struct {
	name     string
	project  string
	location string
	recreate bool
}

func run(cfg runConfig) error {
	cluster, err := gke.GetCluster(cfg.project, cfg.location, cfg.name)
	if err != nil {
		return err
	}
	if cluster == nil {
		return errors.Errorf("cluster '%s/%s/%s' not found", cfg.project, cfg.location, cfg.name)
	}
	if err = gke.RefreshCluster(cluster, cfg.recreate); err != nil {
		return err
	}
	log.Printf("INFO: cluster '%s' has been refreshed!", cfg.name)
	return nil
}
