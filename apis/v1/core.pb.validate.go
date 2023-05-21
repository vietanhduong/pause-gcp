// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: apis/v1/core.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Config) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ConfigMultiError, or nil if none found.
func (m *Config) ValidateAll() error {
	return m.validate(true)
}

func (m *Config) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StateOutput

	for idx, item := range m.GetSchedules() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfigValidationError{
						field:  fmt.Sprintf("Schedules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfigValidationError{
						field:  fmt.Sprintf("Schedules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  fmt.Sprintf("Schedules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ConfigMultiError(errors)
	}

	return nil
}

// ConfigMultiError is an error wrapping multiple validation errors returned by
// Config.ValidateAll() if the designated constraints aren't met.
type ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigMultiError) AllErrors() []error { return m }

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

// Validate checks the field values on Schedule with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Schedule) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Schedule with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ScheduleMultiError, or nil
// if none found.
func (m *Schedule) ValidateAll() error {
	return m.validate(true)
}

func (m *Schedule) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetProject()) < 1 {
		err := ScheduleValidationError{
			field:  "Project",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Schedule_StopAt_Pattern.MatchString(m.GetStopAt()) {
		err := ScheduleValidationError{
			field:  "StopAt",
			reason: "value does not match regex pattern \"^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Schedule_StartAt_Pattern.MatchString(m.GetStartAt()) {
		err := ScheduleValidationError{
			field:  "StartAt",
			reason: "value does not match regex pattern \"^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRepeat()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ScheduleValidationError{
					field:  "Repeat",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ScheduleValidationError{
					field:  "Repeat",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRepeat()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ScheduleValidationError{
				field:  "Repeat",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetExcept() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ScheduleValidationError{
						field:  fmt.Sprintf("Except[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ScheduleValidationError{
						field:  fmt.Sprintf("Except[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ScheduleValidationError{
					field:  fmt.Sprintf("Except[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ScheduleMultiError(errors)
	}

	return nil
}

// ScheduleMultiError is an error wrapping multiple validation errors returned
// by Schedule.ValidateAll() if the designated constraints aren't met.
type ScheduleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ScheduleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ScheduleMultiError) AllErrors() []error { return m }

// ScheduleValidationError is the validation error returned by
// Schedule.Validate if the designated constraints aren't met.
type ScheduleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ScheduleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ScheduleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ScheduleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ScheduleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ScheduleValidationError) ErrorName() string { return "ScheduleValidationError" }

// Error satisfies the builtin error interface
func (e ScheduleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchedule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ScheduleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ScheduleValidationError{}

var _Schedule_StopAt_Pattern = regexp.MustCompile("^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$")

var _Schedule_StartAt_Pattern = regexp.MustCompile("^([0-1]?[0-9]|2[0-3]):[0-5][0-9]$")

// Validate checks the field values on Except with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Except) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Except with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ExceptMultiError, or nil if none found.
func (m *Except) ValidateAll() error {
	return m.validate(true)
}

func (m *Except) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Specifier.(type) {
	case *Except_Cluster_:
		if v == nil {
			err := ExceptValidationError{
				field:  "Specifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetCluster()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ExceptValidationError{
						field:  "Cluster",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ExceptValidationError{
						field:  "Cluster",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCluster()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExceptValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Except_Sql_:
		if v == nil {
			err := ExceptValidationError{
				field:  "Specifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetSql()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ExceptValidationError{
						field:  "Sql",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ExceptValidationError{
						field:  "Sql",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSql()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExceptValidationError{
					field:  "Sql",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Except_Vm_:
		if v == nil {
			err := ExceptValidationError{
				field:  "Specifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetVm()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ExceptValidationError{
						field:  "Vm",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ExceptValidationError{
						field:  "Vm",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetVm()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExceptValidationError{
					field:  "Vm",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return ExceptMultiError(errors)
	}

	return nil
}

// ExceptMultiError is an error wrapping multiple validation errors returned by
// Except.ValidateAll() if the designated constraints aren't met.
type ExceptMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExceptMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExceptMultiError) AllErrors() []error { return m }

// ExceptValidationError is the validation error returned by Except.Validate if
// the designated constraints aren't met.
type ExceptValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExceptValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExceptValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExceptValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExceptValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExceptValidationError) ErrorName() string { return "ExceptValidationError" }

// Error satisfies the builtin error interface
func (e ExceptValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExcept.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExceptValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExceptValidationError{}

// Validate checks the field values on Repeat with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Repeat) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Repeat with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RepeatMultiError, or nil if none found.
func (m *Repeat) ValidateAll() error {
	return m.validate(true)
}

func (m *Repeat) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.Specifier.(type) {
	case *Repeat_EveryDay:
		if v == nil {
			err := RepeatValidationError{
				field:  "Specifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for EveryDay
	case *Repeat_WeekDays:
		if v == nil {
			err := RepeatValidationError{
				field:  "Specifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for WeekDays
	case *Repeat_Weekends:
		if v == nil {
			err := RepeatValidationError{
				field:  "Specifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		// no validation rules for Weekends
	case *Repeat_Other_:
		if v == nil {
			err := RepeatValidationError{
				field:  "Specifier",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetOther()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RepeatValidationError{
						field:  "Other",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RepeatValidationError{
						field:  "Other",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetOther()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RepeatValidationError{
					field:  "Other",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return RepeatMultiError(errors)
	}

	return nil
}

// RepeatMultiError is an error wrapping multiple validation errors returned by
// Repeat.ValidateAll() if the designated constraints aren't met.
type RepeatMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RepeatMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RepeatMultiError) AllErrors() []error { return m }

// RepeatValidationError is the validation error returned by Repeat.Validate if
// the designated constraints aren't met.
type RepeatValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RepeatValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RepeatValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RepeatValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RepeatValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RepeatValidationError) ErrorName() string { return "RepeatValidationError" }

// Error satisfies the builtin error interface
func (e RepeatValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRepeat.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RepeatValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RepeatValidationError{}

// Validate checks the field values on Except_Cluster with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Except_Cluster) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Except_Cluster with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Except_ClusterMultiError,
// or nil if none found.
func (m *Except_Cluster) ValidateAll() error {
	return m.validate(true)
}

func (m *Except_Cluster) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Zone

	// no validation rules for Region

	if len(errors) > 0 {
		return Except_ClusterMultiError(errors)
	}

	return nil
}

// Except_ClusterMultiError is an error wrapping multiple validation errors
// returned by Except_Cluster.ValidateAll() if the designated constraints
// aren't met.
type Except_ClusterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Except_ClusterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Except_ClusterMultiError) AllErrors() []error { return m }

// Except_ClusterValidationError is the validation error returned by
// Except_Cluster.Validate if the designated constraints aren't met.
type Except_ClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Except_ClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Except_ClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Except_ClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Except_ClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Except_ClusterValidationError) ErrorName() string { return "Except_ClusterValidationError" }

// Error satisfies the builtin error interface
func (e Except_ClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExcept_Cluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Except_ClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Except_ClusterValidationError{}

// Validate checks the field values on Except_Sql with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Except_Sql) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Except_Sql with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Except_SqlMultiError, or
// nil if none found.
func (m *Except_Sql) ValidateAll() error {
	return m.validate(true)
}

func (m *Except_Sql) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return Except_SqlMultiError(errors)
	}

	return nil
}

// Except_SqlMultiError is an error wrapping multiple validation errors
// returned by Except_Sql.ValidateAll() if the designated constraints aren't met.
type Except_SqlMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Except_SqlMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Except_SqlMultiError) AllErrors() []error { return m }

// Except_SqlValidationError is the validation error returned by
// Except_Sql.Validate if the designated constraints aren't met.
type Except_SqlValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Except_SqlValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Except_SqlValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Except_SqlValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Except_SqlValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Except_SqlValidationError) ErrorName() string { return "Except_SqlValidationError" }

// Error satisfies the builtin error interface
func (e Except_SqlValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExcept_Sql.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Except_SqlValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Except_SqlValidationError{}

// Validate checks the field values on Except_Vm with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Except_Vm) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Except_Vm with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Except_VmMultiError, or nil
// if none found.
func (m *Except_Vm) ValidateAll() error {
	return m.validate(true)
}

func (m *Except_Vm) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Zone

	if len(errors) > 0 {
		return Except_VmMultiError(errors)
	}

	return nil
}

// Except_VmMultiError is an error wrapping multiple validation errors returned
// by Except_Vm.ValidateAll() if the designated constraints aren't met.
type Except_VmMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Except_VmMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Except_VmMultiError) AllErrors() []error { return m }

// Except_VmValidationError is the validation error returned by
// Except_Vm.Validate if the designated constraints aren't met.
type Except_VmValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Except_VmValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Except_VmValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Except_VmValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Except_VmValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Except_VmValidationError) ErrorName() string { return "Except_VmValidationError" }

// Error satisfies the builtin error interface
func (e Except_VmValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExcept_Vm.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Except_VmValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Except_VmValidationError{}

// Validate checks the field values on Repeat_Other with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Repeat_Other) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Repeat_Other with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Repeat_OtherMultiError, or
// nil if none found.
func (m *Repeat_Other) ValidateAll() error {
	return m.validate(true)
}

func (m *Repeat_Other) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return Repeat_OtherMultiError(errors)
	}

	return nil
}

// Repeat_OtherMultiError is an error wrapping multiple validation errors
// returned by Repeat_Other.ValidateAll() if the designated constraints aren't met.
type Repeat_OtherMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Repeat_OtherMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Repeat_OtherMultiError) AllErrors() []error { return m }

// Repeat_OtherValidationError is the validation error returned by
// Repeat_Other.Validate if the designated constraints aren't met.
type Repeat_OtherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Repeat_OtherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Repeat_OtherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Repeat_OtherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Repeat_OtherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Repeat_OtherValidationError) ErrorName() string { return "Repeat_OtherValidationError" }

// Error satisfies the builtin error interface
func (e Repeat_OtherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRepeat_Other.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Repeat_OtherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Repeat_OtherValidationError{}