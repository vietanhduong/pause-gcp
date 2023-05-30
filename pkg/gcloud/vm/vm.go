package vm

import (
	"cloud.google.com/go/compute/apiv1/computepb"
	"fmt"
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/pkg/utils/exec"
	"github.com/vietanhduong/pause-gcp/pkg/utils/protoutil"
	"log"
	"strings"
)

func GetInstance(project, zone, name string) (*apis.Vm, error) {
	raw, err := exec.Run(exec.Command("gcloud",
		"compute",
		"instances",
		"describe", name,
		"--project", project,
		"--zone", zone,
		"--format", "json"))
	if err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("%s' was not found", name)) {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "get instance")
	}

	var instance computepb.Instance
	if err = protoutil.Unmarshal([]byte(raw), &instance); err != nil {
		return nil, err
	}

	return &apis.Vm{
		Name:    instance.GetName(),
		Zone:    instance.GetZone(),
		State:   instance.GetStatus(),
		Project: project,
	}, nil
}

func StopInstance(vm *apis.Vm, terminate bool) error {
	if vm.GetState() == "TERMINATED" || vm.GetState() == "SUSPENDED" {
		log.Printf("INFO: instance %s already stopped", vm.GetName())
		return nil
	}
	if vm.GetState() != "RUNNING" {
		return errors.Errorf("instance is incorrect state (%s)", vm.GetState())
	}

	cmd := "suspend"
	if terminate {
		cmd = "stop"
	}
	_, err := exec.Run(exec.Command("gcloud",
		"compute",
		"instances",
		cmd, vm.GetName(),
		"--project", vm.GetProject(),
		"--zone", vm.GetZone()))
	return err
}

func StartInstance(vm *apis.Vm) error {
	if vm.GetState() == "RUNNING" {
		log.Printf("INFO: instance %s already started", vm.GetName())
		return nil
	}
	if vm.GetState() != "TERMINATED" && vm.GetState() != "SUSPENDED" {
		return errors.Errorf("instance is incorrect state (%s)", vm.GetState())
	}

	var cmd = "start"
	if vm.GetState() == "SUSPENDED" {
		cmd = "resume"
	}
	_, err := exec.Run(exec.Command("gcloud",
		"compute",
		"instances",
		cmd, vm.GetName(),
		"--project", vm.GetProject(),
		"--zone", vm.GetZone()))
	return err
}
