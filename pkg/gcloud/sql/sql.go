package sql

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/sqladmin/v1"
)

type Client struct{}

func NewClient() *Client { return &Client{} }

func (c *Client) GetInstance(project, name string) (*apis.Sql, error) {
	svc, err := c.newSqlService()
	if err != nil {
		return nil, errors.Wrap(err, "get instance")
	}
	instance, err := svc.Instances.Get(project, name).Do()
	if err != nil {
		var apiErr *googleapi.Error
		if errors.As(err, &apiErr) && apiErr.Code == 404 {
			return nil, nil
		}
		return nil, errors.Wrap(err, "get instance")
	}

	return &apis.Sql{
		Name:            instance.Name,
		Project:         project,
		GceZone:         instance.GceZone,
		IsRunning:       instance.Settings.ActivationPolicy == "ALWAYS",
		DatabaseVersion: instance.DatabaseVersion,
		Region:          instance.Region,
	}, nil
}

func (c *Client) StopInstance(instance *apis.Sql) error {
	svc, err := c.newSqlService()
	if err != nil {
		return errors.Wrap(err, "stop instance")
	}

	if err = waitOperations(svc, instance); err != nil {
		return errors.Wrap(err, "stop instance")
	}

	// retrieve the latest state of the input instance
	var sql *apis.Sql
	sql, err = c.GetInstance(instance.GetProject(), instance.GetName())
	if err != nil {
		return errors.Wrap(err, "stop instance")
	}

	if sql == nil {
		return errors.Errorf("instance %q already be deleted", instance.GetName())
	}

	if !sql.IsRunning {
		log.Printf("INFO: instance %s is stopping, no action needed", instance.GetName())
		return nil
	}

	req := &sqladmin.DatabaseInstance{Settings: &sqladmin.Settings{ActivationPolicy: "NEVER"}}
	_, err = svc.Instances.Patch(instance.GetProject(), instance.GetName(), req).Do()
	if err != nil {
		return errors.Wrap(err, "start instance")
	}
	return nil
}

func (c *Client) StartInstance(instance *apis.Sql) error {
	svc, err := c.newSqlService()
	if err != nil {
		return errors.Wrap(err, "start instance")
	}

	if err = waitOperations(svc, instance); err != nil {
		return errors.Wrap(err, "start instance")
	}

	// retrieve the latest state of the input instance
	var sql *apis.Sql
	sql, err = c.GetInstance(instance.GetProject(), instance.GetName())
	if err != nil {
		return errors.Wrap(err, "start instance")
	}

	if sql == nil {
		return errors.Errorf("instance %q already be deleted", instance.GetName())
	}

	if sql.IsRunning {
		log.Printf("INFO: instance %s is running, no action needed", instance.GetName())
		return nil
	}

	req := &sqladmin.DatabaseInstance{Settings: &sqladmin.Settings{ActivationPolicy: "ALWAYS"}}
	_, err = svc.Instances.Patch(instance.GetProject(), instance.GetName(), req).Do()
	if err != nil {
		return errors.Wrap(err, "start instance")
	}
	return nil
}

func (c *Client) newSqlService() (*sqladmin.Service, error) {
	return sqladmin.NewService(context.Background())
}

func isRunning(state string) bool {
	switch state {
	case "PENDING_CREATE", "RUNNABLE":
		return true
	default:
		return false
	}
}

func waitOperations(svc *sqladmin.Service, instance *apis.Sql) error {
	resp, err := svc.Operations.List(instance.GetProject()).Instance(instance.GetName()).Do()
	if err != nil {
		return errors.Wrap(err, "wait operations")
	}
	wait := func(op *sqladmin.Operation) error {
		if op.Status == "DONE" {
			return nil
		}
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ticker.C:
				log.Printf("INFO: waiting operation %q...\n", instance.GetName())
				tmp, err := svc.Operations.Get(instance.GetProject(), op.Name).Do()
				if err != nil {
					log.Printf("WARN: wait operation %q got error: %v\n", op.Name, err)
					return err
				}
				if tmp.Status == "DONE" {
					return nil
				}
			case <-context.Background().Done():
				return nil
			}
		}
	}
	var eg errgroup.Group
	for _, op := range resp.Items {
		op := op
		eg.Go(func() error { return wait(op) })
	}
	return eg.Wait()
}

func isStopping(state string) bool {
	switch state {
	case "STOPPED", "SUSPENDED", "PENDING_DELETE":
		return true
	default:
		return false
	}
}
