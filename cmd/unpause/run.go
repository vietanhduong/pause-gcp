package unpause

import (
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/cmd/utils"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/gke"
	"github.com/vietanhduong/pause-gcp/pkg/utils/sets"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type runConfig struct {
	configFile string
	force      bool

	gkeClient gke.Interface
}

func run(runCfg runConfig) error {
	cfg, err := utils.ParseConfigFile(runCfg.configFile)
	if err != nil {
		return err
	}
	// validate the config
	if err = utils.ValidateConfig(cfg); err != nil {
		return err
	}

	backupDir := filepath.Join(filepath.Dir(runCfg.configFile), ".backup-state")
	backups, err := getBackupStateFiles(backupDir)
	if err != nil {
		return err
	}

	scheduleKeys := buildScheduleKeys(cfg)

	var wg sync.WaitGroup
	wg.Add(len(backups))
	for _, p := range backups {
		path := filepath.Join(backupDir, p)
		state := utils.ReadBackupState(path)
		if state == nil {
			continue
		}

		removed := !scheduleKeys.Contains(utils.BuildScheduleId(state.GetSchedule()))

		go func(state *apis.BackupState) {
			defer wg.Done()
			t := time.Now()
			log.Printf("INFO: prepare to execute backup(project=%s)...\n", state.GetProject())
			defer func() { log.Printf("INFO: execute backup(project=%s) completed (took=%v)!\n", state.GetProject(), time.Since(t)) }()
			newState, err := execute(state, runCfg.force || removed, runCfg)
			if err != nil {
				log.Printf("WARN: execute backup(project=%s) got error: %v\n", state.GetProject(), err)
				_ = utils.WriteBackupState(backupDir, newState)
				return
			}

			// skip execute
			if len(newState.GetPausedResources()) > 0 {
				return
			}

			// if the unpause process complete, we can remove the state file
			_ = os.Remove(path)
		}(state)
	}
	wg.Wait()
	return nil
}

func execute(state *apis.BackupState, force bool, cfg runConfig) (*apis.BackupState, error) {
	if !utils.ShouldExecute(false, state.GetSchedule(), nil) && !force {
		return state, nil
	}

	var wg sync.WaitGroup
	wg.Add(len(state.PausedResources))
	for i, pr := range state.GetPausedResources() {
		go func(i int, pr *apis.Resource) {
			defer wg.Done()

			switch r := pr.GetSpecifier().(type) {
			case *apis.Resource_Cluster:
				if err := unpauseCluster(r.Cluster, cfg); err != nil {
					log.Printf("WARN: unpause cluster '%s/%s' got error: %v", r.Cluster.GetLocation(), r.Cluster.GetName(), err)
				} else {
					state.PausedResources[i] = nil
				}
			case *apis.Resource_Sql:
			case *apis.Resource_Vm:
			}
		}(i, pr)

	}
	wg.Wait()

	var pausedResources []*apis.Resource
	for _, r := range state.GetPausedResources() {
		if r != nil {
			pausedResources = append(pausedResources, r)
		}
	}

	state.PausedResources = pausedResources
	if len(pausedResources) > 0 {
		return state, errors.Errorf("unpause incomplete (%s resources left)", len(pausedResources))
	}
	return state, nil
}

func unpauseCluster(c *apis.Cluster, cfg runConfig) error {
	t := time.Now()
	log.Printf("INFO: prepare to unpause cluster '%s/%s'...", c.GetLocation(), c.GetName())
	defer func() { log.Printf("INFO: unpause cluster '%s/%s' complete (took=%v)!", c.GetLocation(), c.GetName(), time.Since(t)) }()
	return cfg.gkeClient.UnpauseCluster(c)
}

func getBackupStateFiles(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var out []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			out = append(out, file.Name())
		}
	}
	return out, nil
}

func buildScheduleKeys(cfg *apis.Config) sets.String {
	keys := sets.New[string]()
	for _, s := range cfg.GetSchedules() {
		keys.Insert(utils.BuildScheduleId(s))
	}
	return keys
}
