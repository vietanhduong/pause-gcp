package unpause

import (
	"github.com/pkg/errors"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/vm"
	"log"
)

type runConfig struct {
	project string
	zone    string
	name    string
}

func run(cfg runConfig) error {
	instance, err := vm.GetInstance(cfg.project, cfg.zone, cfg.name)
	if err != nil {
		return err
	}
	if instance == nil {
		return errors.Errorf("instance '%s/%s/%s' not found", cfg.project, cfg.zone, cfg.name)
	}

	if err = vm.StartInstance(instance); err != nil {
		return err
	}
	log.Printf("INFO: instance %s has been started!", instance.GetName())
	return nil
}
