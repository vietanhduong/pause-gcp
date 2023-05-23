package pause

import (
	"fmt"
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/cmd/utils"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/gke"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"path/filepath"
	"sync"
	"time"
)

type runConfig struct {
	configFile string
	configDir  string
	force      bool
	dryRun     bool
}

func run(runCfg runConfig) error {
	// parse config
	cfg, err := utils.ParseConfigFile(runCfg.configFile)
	if err != nil {
		return err
	}
	// validate the config
	if err = utils.ValidateConfig(cfg); err != nil {
		return err
	}

	runCfg.configDir = filepath.Dir(runCfg.configFile)

	var wg sync.WaitGroup
	wg.Add(len(cfg.Schedules))
	for _, s := range cfg.Schedules {
		go func(s *apis.Schedule) {
			defer wg.Done()
			t := time.Now()
			log.Printf("INFO: prepare to execute schedule(project=%s)...\n", s.GetProject())
			defer func() { log.Printf("INFO: execute schedule(project=%s) completed (took=%v)!\n", s.GetProject(), time.Since(t)) }()
			if err := execute(s, runCfg); err != nil {
				log.Printf("WARN: execute schedule (project=%s) got error: %v\n", s.GetProject(), err)
			}
		}(s)
	}
	wg.Wait()
	return nil
}

func execute(schedule *apis.Schedule, cfg runConfig) error {
	state := utils.ReadBackupState(utils.BuildBackupStateFilename(cfg.configDir, schedule))
	if !utils.ShouldExecute(true, schedule, state) && !cfg.force {
		return nil
	}

	newState := &apis.BackupState{
		Project:  schedule.GetProject(),
		Schedule: schedule,
		DryRun:   cfg.dryRun,
	}

	// pause clusters
	clusters, err := pauseCluster(schedule, cfg)
	if err != nil {
		return err
	}

	newState.PausedResources = append(newState.PausedResources, clusters...)

	// pause vm

	// pause sql

	return utils.WriteBackupState(cfg.configDir, newState)
}

func pauseCluster(schedule *apis.Schedule, cfg runConfig) ([]*apis.Resource, error) {
	client := gke.NewClient()
	clusters, err := client.ListClusters(schedule.GetProject())
	if err != nil {
		return nil, err
	}

	currentClusters := make(map[string]*apis.Cluster)
	for _, c := range clusters {
		currentClusters[fmt.Sprintf("%s/%s", c.GetLocation(), c.GetName())] = c
	}

	var pauseClusters []*apis.Cluster

	for _, r := range schedule.GetResources() {
		if pc := r.GetCluster(); pc != nil {
			if c, ok := currentClusters[fmt.Sprintf("%s/%s", pc.GetLocation(), pc.GetName())]; !ok {
				return nil, errors.Errorf("cluster '%s/%s' not found!", pc.GetLocation(), pc.GetName())
			} else {
				// if no pool is specified, we implicitly select all existing pools.
				if len(pc.GetNodePools()) == 0 {
					pauseClusters = append(pauseClusters, c)
					continue
				}

				var pausePools []*apis.Cluster_NodePool
				pools := make(map[string]*apis.Cluster_NodePool)
				for _, p := range c.GetNodePools() {
					pools[p.GetName()] = p
				}

				for _, pp := range pc.GetNodePools() {
					if p, ok := pools[pp.GetName()]; !ok {
						return nil, errors.Errorf("not found pool '%s' in cluster '%s/%s'", pp.GetName(), pc.GetLocation(), pc.GetName())
					} else {
						pausePools = append(pausePools, p)
					}
				}
				c = c.DeepCopy()
				c.NodePools = pausePools
				pauseClusters = append(pauseClusters, c)
			}
		}
	}

	tmp := make([]*apis.Resource, len(pauseClusters))
	var wg sync.WaitGroup
	wg.Add(len(pauseClusters))
	for i, c := range pauseClusters {
		go func(i int, c *apis.Cluster) {
			defer wg.Done()
			t := time.Now()
			log.Printf("INFO: prepare to pause cluster '%s/%s'...", c.GetLocation(), c.GetName())
			defer func() { log.Printf("INFO: pause cluster '%s/%s' complated (took=%v)!", c.GetLocation(), c.GetName(), time.Since(t)) }()

			var cluster *apis.Cluster
			var err error
			if cluster, err = client.PauseCluster(c, cfg.dryRun); err != nil {
				log.Printf("WARN: pause cluster '%s/%s' got error: %v", c.GetLocation(), c.GetName(), err)
				return
			}
			tmp[i] = &apis.Resource{
				Specifier: &apis.Resource_Cluster{Cluster: cluster},
				PausedAt:  timestamppb.New(time.Now()),
			}
		}(i, c)
	}
	wg.Wait()
	var out []*apis.Resource
	for _, c := range tmp {
		if c != nil {
			out = append(out, c)
		}
	}
	return out, nil
}

func pauseVm(schedule *apis.Schedule) ([]*apis.Vm, error) {
	return nil, nil
}

func pauseSql(schedule *apis.Schedule) ([]*apis.Vm, error) {
	return nil, nil
}
