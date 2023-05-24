package sql

import (
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/pkg/utils/exec"
	"github.com/vietanhduong/pause-gcp/pkg/utils/protoutil"
	"log"
	"strings"
)

func GetInstance(project, name string) (*apis.Sql, error) {
	raw, err := exec.Run(exec.Command("gcloud",
		"sql",
		"instances",
		"describe", name,
		"--project", project,
		"--format", "json"))
	if err != nil {
		if strings.Contains(err.Error(), "The Cloud SQL instance does not exist") {
			return nil, nil
		}
		return nil, errors.Wrapf(err, "get instance")
	}

	var instance apis.Sql
	if err = protoutil.UnmarshalAllowUnknown([]byte(raw), &instance); err != nil {
		return nil, err
	}
	return &instance, nil
}

func StopInstance(instance *apis.Sql) error {
	if instance.GetState() == "SUSPENDED" || instance.GetState() == "STOPPED" {
		log.Printf("INFO: instance %s already stopped", instance.GetName())
		return nil
	}
	if instance.GetState() != "RUNNABLE" {
		return errors.Errorf("instance is incorrect state (%s)", instance.GetState())
	}

	_, err := exec.Run(exec.Command("gcloud",
		"sql",
		"instances",
		"patch", instance.GetName(),
		"--project", instance.GetProject(),
		"--activation-policy", "NEVER"))
	return err
}

func StartInstance(instance *apis.Sql) error {
	if instance.GetState() == "RUNNABLE" {
		log.Printf("INFO: instance %s already started", instance.GetName())
		return nil
	}
	if instance.GetState() != "STOPPED" && instance.GetState() != "SUSPENDED" {
		return errors.Errorf("instance is incorrect state (%s)", instance.GetState())
	}
	_, err := exec.Run(exec.Command("gcloud",
		"sql",
		"instances",
		"patch", instance.GetName(),
		"--project", instance.GetProject(),
		"--activation-policy", "ALWAYS"))
	return err
}
