package env

import (
	"os"
	"strconv"
	"strings"
	timelib "time"
)

// StringFromEnv returns the env variable for the given key
// and falls back to the given defaultValue if not set
func StringFromEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// ParseNumFromEnv helper function to parse a number from an environment variable. Returns a
// default if env is not set, is not parseable to a number, exceeds max (if
// max is greater than 0) or is less than min.
func ParseNumFromEnv(env string, defaultValue, min, max int) int {
	str := os.Getenv(env)
	if str == "" {
		return defaultValue
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return defaultValue
	}
	if num < min {
		return defaultValue
	}
	if num > max {
		return defaultValue
	}
	return num
}

// ParseDurationFromEnv helper function to parse a time duration from an environment variable. Returns a
// default if env is not set, is not parseable to a duration, exceeds max (if
// max is greater than 0) or is less than min.
func ParseDurationFromEnv(env string, defaultValue, min, max timelib.Duration) timelib.Duration {
	str := os.Getenv(env)
	if str == "" {
		return defaultValue
	}
	dur, err := timelib.ParseDuration(strings.ToLower(str))
	if err != nil {
		return defaultValue
	}

	if dur < min {
		return defaultValue
	}
	if dur > max {
		return defaultValue
	}
	return dur
}

// ParseBoolFromEnv retrieves a boolean value from given environment envVar.
// Returns default value if envVar is not set.
func ParseBoolFromEnv(envVar string, defaultValue bool) bool {
	if val := os.Getenv(envVar); val != "" {
		val = strings.ToLower(val)
		switch val {
		case "true":
			return true
		case "1":
			return true
		case "false":
			return false
		case "0":
			return false
		}
	}
	return defaultValue
}
