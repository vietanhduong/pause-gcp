package pause

import (
	"log"

	"github.com/pkg/errors"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/vm"
)

type runConfig struct {
	project   string
	zone      string
	name      string
	terminate bool
}

func run(cfg runConfig) error {
	client := vm.NewClient()
	instance, err := client.GetInstance(cfg.project, cfg.zone, cfg.name)
	if err != nil {
		return err
	}
	if instance == nil {
		return errors.Errorf("instance '%s/%s/%s' not found", cfg.project, cfg.zone, cfg.name)
	}

	if err = client.StopInstance(instance, cfg.terminate); err != nil {
		return err
	}
	log.Printf("INFO: instance %s has been stopped!", instance.GetName())
	return nil
}
