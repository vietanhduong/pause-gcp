package vm

import (
	"context"
	"log"
	"strings"

	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
)

type Client struct{}

func NewClient() *Client { return &Client{} }

func (c *Client) GetInstance(project, zone, name string) (*apis.Vm, error) {
	svc, err := c.newComputeService()
	if err != nil {
		return nil, errors.Wrap(err, "get instance")
	}

	instance, err := svc.Instances.Get(project, zone, name).Do()
	if err != nil {
		var apiErr *googleapi.Error
		if errors.As(err, &apiErr) && apiErr.Code == 404 {
			return nil, nil
		}
		return nil, err
	}

	return &apis.Vm{
		Name:    instance.Name,
		Zone:    basename(instance.Zone),
		State:   instance.Status,
		Project: project,
	}, nil
}

func (c *Client) StopInstance(instance *apis.Vm, terminate bool) error {
	if isStopping(instance.GetState()) {
		log.Printf("INFO: instance %s already stopped", instance.GetName())
		return nil
	}

	if !isRunning(instance.GetState()) {
		return errors.Errorf("instance is incorrect state (%s)", instance.GetState())
	}

	svc, err := c.newComputeService()
	if err != nil {
		return errors.Wrap(err, "stop instance")
	}

	if terminate {
		_, err = svc.Instances.Stop(instance.GetProject(), instance.GetZone(), instance.GetName()).Do()
	} else {
		_, err = svc.Instances.Suspend(instance.GetProject(), instance.GetZone(), instance.GetName()).Do()
	}

	if err != nil {
		return errors.Wrap(err, "stop instance")
	}
	return nil
}

func (c *Client) StartInstance(instance *apis.Vm) error {
	if isRunning(instance.GetState()) {
		log.Printf("INFO: instance %s already started", instance.GetName())
		return nil
	}
	if !isStopping(instance.GetState()) {
		return errors.Errorf("instance is incorrect state (%s)", instance.GetState())
	}

	svc, err := c.newComputeService()
	if err != nil {
		return errors.Wrap(err, "start instance")
	}

	if instance.GetState() == "SUSPENDED" {
		_, err = svc.Instances.Resume(instance.GetProject(), instance.GetZone(), instance.GetName()).Do()
	} else {
		_, err = svc.Instances.Start(instance.GetProject(), instance.GetZone(), instance.GetName()).Do()
	}

	if err != nil {
		return errors.Wrap(err, "start instance")
	}
	return nil
}

func (c *Client) newComputeService() (*compute.Service, error) {
	return compute.NewService(context.Background())
}

func basename(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func isRunning(state string) bool {
	switch state {
	case "PROVISIONING", "RUNNING":
		return true
	default:
		return false
	}
}

func isStopping(state string) bool {
	switch state {
	case "STOPPED", "STOPPING", "SUSPENDED", "SUSPENDING", "TERMINATED":
		return true
	default:
		return false
	}
}
