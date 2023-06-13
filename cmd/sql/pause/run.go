package pause

import (
	"log"

	"github.com/pkg/errors"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/sql"
)

type runConfig struct {
	project string
	name    string
}

func run(cfg runConfig) error {
	client := sql.NewClient()
	instance, err := client.GetInstance(cfg.project, cfg.name)
	if err != nil {
		return err
	}
	if instance == nil {
		return errors.Errorf("sql instance '%s/%s' not found", cfg.project, cfg.name)
	}

	if err = client.StopInstance(instance); err != nil {
		return err
	}
	log.Printf("INFO: sql instance %s has been stopped!", instance.GetName())
	return nil
}
