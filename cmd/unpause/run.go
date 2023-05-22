package unpause

import (
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/cmd/utils"
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
		if state == nil || state.DryRun {
			continue
		}

		removed := !scheduleKeys.Contains(utils.BuildScheduleId(state.GetSchedule()))

		go func(state *apis.BackupState) {
			defer wg.Done()
			t := time.Now()
			log.Printf("INFO: prepare to execute backup(project=%s)...\n", state.GetProject())
			defer func() { log.Printf("INFO: execute backup(project=%s) completed (took=%v)!\n", state.GetProject(), time.Since(t)) }()
			if err := execute(state, runCfg.force || removed); err != nil {
				log.Printf("WARN: execute backup(project=%s) got error: %v\n", state.GetProject(), err)
			} else {
				if removed {
					_ = os.Remove(path)
				}
			}
		}(state)
	}
	wg.Wait()
	return nil
}

func execute(state *apis.BackupState, force bool) error {
	if !utils.ShouldExecute(state.GetSchedule(), nil) && !force {
		return nil
	}
	return nil
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
