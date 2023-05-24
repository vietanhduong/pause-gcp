package unpause

import (
	"github.com/pkg/errors"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/sql"
	"log"
)

type runConfig struct {
	project string
	name    string
}

func run(cfg runConfig) error {
	instance, err := sql.GetInstance(cfg.project, cfg.name)
	if err != nil {
		return err
	}
	if instance == nil {
		return errors.Errorf("sql instance '%s/%s' not found", cfg.project, cfg.name)
	}

	if err = sql.StartInstance(instance); err != nil {
		return err
	}
	log.Printf("INFO: sql instance %s has been started!", instance.GetName())
	return nil
}
