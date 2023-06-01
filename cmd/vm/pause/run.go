package pause

import (
	"github.com/pkg/errors"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/vm"
	"log"
)

type runConfig struct {
	project   string
	zone      string
	name      string
	terminate bool
}

func run(cfg runConfig) error {
	instance, err := vm.GetInstance(cfg.project, cfg.zone, cfg.name)
	if err != nil {
		return err
	}
	if instance == nil {
		return errors.Errorf("instance '%s/%s/%s' not found", cfg.project, cfg.zone, cfg.name)
	}

	if err = vm.StopInstance(instance, cfg.terminate); err != nil {
		return err
	}
	log.Printf("INFO: instance %s has been stopped!", instance.GetName())
	return nil
}
