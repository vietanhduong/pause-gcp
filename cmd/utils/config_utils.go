package utils

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"github.com/vietanhduong/pause-gcp/pkg/utils/sets"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"os"
	"sigs.k8s.io/yaml"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var now = time.Now

func ShouldExecute(schedule *apis.Schedule, state *apis.BackupState) bool {
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

func ReadBackupState(path string) *apis.BackupState {
	b, _ := os.ReadFile(path)
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

func WriteBackupState(configDir string, state *apis.BackupState) error {
	b, _ := marshaler.Marshal(state)
	_ = os.MkdirAll(fmt.Sprintf("%s/.backup-state", configDir), 0755)
	filename := BuildBackupStateFilename(configDir, state.GetSchedule())
	log.Printf("INFO: prepare to write backup state to file %q with content:\n%v", filename, string(b))
	return os.WriteFile(filename, b, 0644)
}

func ParseConfigFile(path string) (*apis.Config, error) {
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

func ValidateConfig(cfg *apis.Config) error {
	if err := cfg.ValidateAll(); err != nil {
		return err
	}
	ids := sets.New[string]()
	for _, s := range cfg.GetSchedules() {
		if id := BuildScheduleId(s); ids.Contains(id) {
			return errors.Errorf("duplicate id %q", id)
		} else {
			ids.Insert(id)
		}
	}
	return nil
}

// BuildScheduleId return an id by format: <project_id>_<stop_at>-<start_at>
func BuildScheduleId(schedule *apis.Schedule) string {
	return fmt.Sprintf("%s_%s",
		schedule.GetProject(),
		strings.ReplaceAll(
			fmt.Sprintf("%s-%s", schedule.GetStopAt(), schedule.GetStartAt()),
			":", ""),
	)
}

func BuildBackupStateFilename(configDir string, schedule *apis.Schedule) string {
	return fmt.Sprintf("%s/.backup-state/schedule_%s.json", configDir, BuildScheduleId(schedule))
}
