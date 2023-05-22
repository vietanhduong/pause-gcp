package pause

import (
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/cmd/utils"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/gke"
	"github.com/vietanhduong/pause-gcp/pkg/utils/sets"
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
	if !utils.ShouldExecute(schedule, state) && !cfg.force {
		return nil
	}

	newState := &apis.BackupState{
		Project:  schedule.GetProject(),
		Schedule: schedule,
		DryRun:   cfg.dryRun,
	}

	// pause clusters
	clusters, err := pauseCluster(schedule, state, cfg)
	if err != nil {
		return err
	}

	newState.PausedResources = append(newState.PausedResources, clusters...)

	// pause vm

	// pause sql

	return utils.WriteBackupState(cfg.configDir, newState)
}

func pauseCluster(schedule *apis.Schedule, state *apis.BackupState, cfg runConfig) ([]*apis.Resource, error) {
	var skip bool
	for _, e := range schedule.GetExcept() {
		if c := e.GetCluster(); c == (&apis.Cluster{}) {
			skip = true
			break
		}
	}

	if skip {
		log.Printf("INFO: ignore GKE resource!\n")
		return nil, nil
	}

	client := gke.NewClient()
	clusters, err := client.ListClusters(schedule.GetProject())
	if err != nil {
		return nil, err
	}

	tmp := make([]*apis.Resource, len(clusters))
	var wg sync.WaitGroup
	wg.Add(len(clusters))
	for i, c := range clusters {
		// ensure that the cluster presents is paused or not (the state). If not, we should pause it.
		if isClusterPaused(c, schedule) {
			continue
		}

		go func(i int, c *apis.Cluster) {
			defer wg.Done()
			t := time.Now()
			log.Printf("INFO: prepare to pause cluster '%s/%s'...", c.GetLocation(), c.GetName())
			defer func() { log.Printf("INFO: pause cluster '%s/%s' complated (took=%v)!", c.GetLocation(), c.GetName(), time.Since(t)) }()

			var cluster *apis.Cluster
			var err error
			if cluster, err = client.PauseCluster(c, schedule.GetExcept(), cfg.dryRun); err != nil {
				log.Printf("WARN: pause cluster '%s/%s' got error: %v", c.GetLocation(), c.GetName(), err)
				return
			}
			tmp[i] = &apis.Resource{
				Specifier:     &apis.Resource_Cluster{Cluster: cluster},
				TimeSpecifier: &apis.Resource_PausedAt{PausedAt: timestamppb.New(time.Now())},
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

// isClusterPaused ensure that if all node pools of the input cluster is paused. If a single node pool is not paused,
// this still return true. PauseCluster function will do nothing if the node pool has size 0
func isClusterPaused(cluster *apis.Cluster, schedule *apis.Schedule) bool {
	exceptPools := sets.New[string]()

	for _, e := range schedule.GetExcept() {
		if exceptCluster := e.GetCluster(); exceptCluster != nil && exceptCluster.GetName() == cluster.GetName() {
			if exceptCluster.GetLocation() != "" && exceptCluster.GetLocation() != cluster.GetLocation() {
				continue
			}
			for _, p := range exceptCluster.GetNodePools() {
				exceptPools.Insert(p.GetName())
			}
			break
		}
	}

	for _, p := range cluster.GetNodePools() {
		// if a pool is not contained in the except pool, but it's current size > 0 then we should pause it.
		if !exceptPools.Contains(p.GetName()) && p.GetCurrentSize() > 0 {
			return false
		}
	}

	return false
}

func pauseVm(schedule *apis.Schedule) ([]*apis.Vm, error) {
	return nil, nil
}

func pauseSql(schedule *apis.Schedule) ([]*apis.Vm, error) {
	return nil, nil
}
