package pause

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/pkg/gcloud/gke"
	"github.com/vietanhduong/pause-gcp/pkg/utils/sets"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
	"strings"
	"sync"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// now is time.Now function
var now = time.Now

type runConfig struct {
	configFile string
	configDir  string
	force      bool
	dryRun     bool
}

func run(runCfg runConfig) error {
	// parse config
	cfg, err := parseConfigFile(runCfg.configFile)
	if err != nil {
		return err
	}
	// validate the config
	if err = validateConfig(cfg); err != nil {
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
	state := readBackupState(cfg.configDir, schedule)
	if !shouldExecute(schedule, state) && !cfg.force {
		return nil
	}

	newState := &apis.BackupState{
		PauseAt:  timestamppb.New(time.Now()),
		Project:  schedule.GetProject(),
		Schedule: schedule,
		DryRun:   cfg.dryRun,
	}

	// pause clusters
	clusters, err := pauseCluster(schedule, cfg)
	if err != nil {
		return err
	}

	for _, c := range clusters {
		newState.PausedResources = append(newState.PausedResources, &apis.Resource{Specifier: &apis.Resource_Cluster{Cluster: c}})
	}

	// pause vm

	// pause sql

	return writeBackupState(cfg.configDir, newState)
}

func pauseCluster(schedule *apis.Schedule, cfg runConfig) ([]*apis.Cluster, error) {
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

	tmp := make([]*apis.Cluster, len(clusters))
	var wg sync.WaitGroup
	wg.Add(len(clusters))
	for i, c := range clusters {
		go func(i int, c *apis.Cluster) {
			defer wg.Done()
			t := time.Now()
			log.Printf("INFO: prepare to pause cluster '%s/%s'...", c.GetLocation(), c.GetName())
			defer func() { log.Printf("INFO: pause cluster '%s/%s' complated (took=%v)!", c.GetLocation(), c.GetName(), time.Since(t)) }()
			if !cfg.dryRun {
				if err := client.PauseCluster(c.DeepCopy(), schedule.GetExcept()); err != nil {
					log.Printf("WARN: pause cluster '%s/%s' got error: %v", c.GetLocation(), c.GetName(), err)
					return
				}
			}
			tmp[i] = c
		}(i, c)
	}
	wg.Wait()
	var out []*apis.Cluster
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

func shouldExecute(schedule *apis.Schedule, state *apis.BackupState) bool {
	// this means, the job already done and no repeat is specified. We don't need to re-run
	// if the state is exists, but in dry-run mode, lets execute it again
	if (schedule.GetRepeat() == nil || !schedule.GetRepeat().GetEveryDay()) && state != nil && !state.DryRun {
		return false
	}

	repeat := schedule.GetRepeat()
	day := now().Weekday()

	if repeat.GetWeekDays() && (day == time.Saturday || day == time.Sunday) {
		return false
	}

	if repeat.GetWeekends() && (day != time.Saturday && day != time.Sunday) {
		return false
	}

	if len(repeat.GetOther().GetDays()) > 0 {
		days := sets.New(repeat.GetOther().GetDays()...)
		if !days.Contains(apis.Repeat_Day(int32(day))) {
			return false
		}
	}
	stopAt, _ := time.Parse("2006-01-02 15:04",
		fmt.Sprintf("%s %s", now().Format("2006-01-02"), schedule.GetStopAt()))
	return stopAt.Before(now())
}

func readBackupState(configDir string, schedule *apis.Schedule) *apis.BackupState {
	b, _ := os.ReadFile(buildBackupStateFilename(configDir, schedule))
	if len(b) == 0 {
		return nil
	}
	var out apis.BackupState
	if err := json.Unmarshal(b, &out); err != nil {
		return nil
	}
	return &out
}

var marshaler = protojson.MarshalOptions{Indent: "    "}

func writeBackupState(configDir string, state *apis.BackupState) error {
	b, _ := marshaler.Marshal(state)
	_ = os.MkdirAll(fmt.Sprintf("%s/.backup-state", configDir), 0755)
	filename := buildBackupStateFilename(configDir, state.GetSchedule())
	log.Printf("INFO: prepare to write backup state to file %q with content:\n%v", filename, string(b))
	return os.WriteFile(filename, b, 0644)
}

func parseConfigFile(path string) (*apis.Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if b, err = yaml.YAMLToJSON(b); err != nil {
		return nil, err
	}

	var cfg apis.Config
	err = json.Unmarshal(b, &cfg)
	return &cfg, nil
}

func validateConfig(cfg *apis.Config) error {
	if err := cfg.ValidateAll(); err != nil {
		return err
	}
	ids := sets.New[string]()
	for _, s := range cfg.GetSchedules() {
		if id := buildScheduleId(s); ids.Contains(id) {
			return errors.Errorf("duplicate id %q", id)
		} else {
			ids.Insert(id)
		}
	}
	return nil
}

// buildScheduleId return an id by format: <project_id>_<stop_at>-<start_at>
func buildScheduleId(schedule *apis.Schedule) string {
	return fmt.Sprintf("%s_%s",
		schedule.GetProject(),
		strings.ReplaceAll(
			fmt.Sprintf("%s-%s", schedule.GetStopAt(), schedule.GetStartAt()),
			":", ""),
	)
}

func buildBackupStateFilename(configDir string, schedule *apis.Schedule) string {
	return fmt.Sprintf("%s/.backup-state/schedule_%s.json", configDir, buildScheduleId(schedule))
}
