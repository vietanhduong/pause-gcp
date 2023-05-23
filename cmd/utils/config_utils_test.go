package utils

import (
	"github.com/stretchr/testify/assert"
	apis "github.com/vietanhduong/pause-gcp/apis/v1"
	"testing"
	"time"
)

func Test_shouldExecute(t *testing.T) {
	var tests = []struct {
		name     string
		pause    bool
		time     func() time.Time
		schedule func() *apis.Schedule
		backup   func() *apis.BackupState
		expected bool
	}{
		{
			name: "TEST SUCCESS: no repeat setup and no backup, expect true",
			schedule: func() *apis.Schedule {
				return &apis.Schedule{
					StopAt:  "20:00",
					StartAt: "08:00",
				}
			},
			time: func() time.Time {
				mt, _ := time.Parse("2006-01-02 15:04:05", "2023-05-22 21:00:00")
				return mt
			},
			pause: true,
			backup: func() *apis.BackupState {
				return nil
			},
			expected: true,
		},
		{
			name: "TEST SUCCESS: no repeat setup and backup existed, expect false",
			schedule: func() *apis.Schedule {
				return &apis.Schedule{
					StopAt:  "20:00",
					StartAt: "08:00",
				}
			},
			time: func() time.Time {
				mt, _ := time.Parse("2006-01-02 15:04:05", "2023-05-22 21:00:00")
				return mt
			},
			pause: true,
			backup: func() *apis.BackupState {
				return &apis.BackupState{}
			},
			expected: false,
		},
		{
			name: "TEST SUCCESS: repeat on weekdays, expect true",
			schedule: func() *apis.Schedule {
				return &apis.Schedule{
					StopAt:  "20:00",
					StartAt: "08:00",
					Repeat: &apis.Repeat{
						Specifier: &apis.Repeat_WeekDays{WeekDays: true},
					},
				}
			},
			backup: func() *apis.BackupState { return nil },
			time: func() time.Time {
				mt, _ := time.Parse("2006-01-02 15:04:05", "2023-05-22 12:00:00")
				return mt
			},
			expected: true,
		},
		{
			name: "TEST SUCCESS: repeat on weekdays, but today is saturday, expect false",
			schedule: func() *apis.Schedule {
				return &apis.Schedule{
					StopAt:  "20:00",
					StartAt: "08:00",
					Repeat: &apis.Repeat{
						Specifier: &apis.Repeat_WeekDays{WeekDays: true},
					},
				}
			},
			backup: func() *apis.BackupState { return nil },
			time: func() time.Time {
				mt, _ := time.Parse("2006-01-02 15:04:05", "2023-05-27 12:00:00")
				return mt
			},
			expected: false,
		},
		{
			name: "TEST SUCCESS: repeat on weekends, expect true",
			schedule: func() *apis.Schedule {
				return &apis.Schedule{
					StopAt:  "12:00",
					StartAt: "08:00",
					Repeat: &apis.Repeat{
						Specifier: &apis.Repeat_Weekends{Weekends: true},
					},
				}
			},
			backup: func() *apis.BackupState { return nil },
			time: func() time.Time {
				mt, _ := time.Parse("2006-01-02 15:04:05", "2023-05-27 14:00:00")
				return mt
			},
			expected: true,
		},
		{
			name: "TEST SUCCESS: repeat on weekends, but today is monday, expect false",
			schedule: func() *apis.Schedule {
				return &apis.Schedule{
					StopAt:  "20:00",
					StartAt: "08:00",
					Repeat: &apis.Repeat{
						Specifier: &apis.Repeat_Weekends{Weekends: true},
					},
				}
			},
			backup: func() *apis.BackupState { return nil },
			time: func() time.Time {
				mt, _ := time.Parse("2006-01-02 15:04:05", "2023-05-22 12:00:00")
				return mt
			},
			expected: false,
		},
		{
			name: "TEST SUCCESS: repeat on monday, expect false",
			schedule: func() *apis.Schedule {
				return &apis.Schedule{
					StopAt:  "20:00",
					StartAt: "08:00",
					Repeat: &apis.Repeat{
						Specifier: &apis.Repeat_Other_{
							Other: &apis.Repeat_Other{Days: []apis.Repeat_Day{apis.Repeat_MONDAY}},
						},
					},
				}
			},
			backup: func() *apis.BackupState { return nil },
			time: func() time.Time {
				mt, _ := time.Parse("2006-01-02 15:04:05", "2023-05-22 12:00:00")
				return mt
			},
			expected: true,
		},
		{
			name: "TEST SUCCESS: repeat on monday, but today is tuesday, expect false",
			schedule: func() *apis.Schedule {
				return &apis.Schedule{
					StopAt:  "20:00",
					StartAt: "08:00",
					Repeat: &apis.Repeat{
						Specifier: &apis.Repeat_Other_{
							Other: &apis.Repeat_Other{Days: []apis.Repeat_Day{apis.Repeat_MONDAY}},
						},
					},
				}
			},
			backup: func() *apis.BackupState { return nil },
			time: func() time.Time {
				mt, _ := time.Parse("2006-01-02 15:04:05", "2023-05-23 12:00:00")
				return mt
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.time == nil || tt.time().IsZero() {
				now = time.Now
			} else {
				now = tt.time
			}
			actual := ShouldExecute(tt.pause, tt.schedule(), tt.backup())
			assert.Equal(t, tt.expected, actual)
		})
	}
}
