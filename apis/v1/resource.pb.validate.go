// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: apis/v1/resource.proto

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

// Validate checks the field values on Resource with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Resource) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Resource with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResourceMultiError, or nil
// if none found.
func (m *Resource) ValidateAll() error {
	return m.validate(true)
}

func (m *Resource) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPausedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ResourceValidationError{
					field:  "PausedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ResourceValidationError{
					field:  "PausedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPausedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResourceValidationError{
				field:  "PausedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch v := m.Specifier.(type) {
	case *Resource_Cluster:
		if v == nil {
			err := ResourceValidationError{
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
					errors = append(errors, ResourceValidationError{
						field:  "Cluster",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResourceValidationError{
						field:  "Cluster",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCluster()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourceValidationError{
					field:  "Cluster",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Resource_Sql:
		if v == nil {
			err := ResourceValidationError{
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
					errors = append(errors, ResourceValidationError{
						field:  "Sql",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResourceValidationError{
						field:  "Sql",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSql()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourceValidationError{
					field:  "Sql",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Resource_Vm:
		if v == nil {
			err := ResourceValidationError{
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
					errors = append(errors, ResourceValidationError{
						field:  "Vm",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResourceValidationError{
						field:  "Vm",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetVm()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResourceValidationError{
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
		return ResourceMultiError(errors)
	}

	return nil
}

// ResourceMultiError is an error wrapping multiple validation errors returned
// by Resource.ValidateAll() if the designated constraints aren't met.
type ResourceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResourceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResourceMultiError) AllErrors() []error { return m }

// ResourceValidationError is the validation error returned by
// Resource.Validate if the designated constraints aren't met.
type ResourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResourceValidationError) ErrorName() string { return "ResourceValidationError" }

// Error satisfies the builtin error interface
func (e ResourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResourceValidationError{}

// Validate checks the field values on Cluster with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Cluster) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cluster with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ClusterMultiError, or nil if none found.
func (m *Cluster) ValidateAll() error {
	return m.validate(true)
}

func (m *Cluster) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Project

	// no validation rules for Name

	// no validation rules for Location

	for idx, item := range m.GetNodePools() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClusterValidationError{
						field:  fmt.Sprintf("NodePools[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClusterValidationError{
						field:  fmt.Sprintf("NodePools[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterValidationError{
					field:  fmt.Sprintf("NodePools[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ClusterMultiError(errors)
	}

	return nil
}

// ClusterMultiError is an error wrapping multiple validation errors returned
// by Cluster.ValidateAll() if the designated constraints aren't met.
type ClusterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClusterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClusterMultiError) AllErrors() []error { return m }

// ClusterValidationError is the validation error returned by Cluster.Validate
// if the designated constraints aren't met.
type ClusterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterValidationError) ErrorName() string { return "ClusterValidationError" }

// Error satisfies the builtin error interface
func (e ClusterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCluster.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterValidationError{}

// Validate checks the field values on Sql with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Sql) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Sql with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SqlMultiError, or nil if none found.
func (m *Sql) ValidateAll() error {
	return m.validate(true)
}

func (m *Sql) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return SqlMultiError(errors)
	}

	return nil
}

// SqlMultiError is an error wrapping multiple validation errors returned by
// Sql.ValidateAll() if the designated constraints aren't met.
type SqlMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SqlMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SqlMultiError) AllErrors() []error { return m }

// SqlValidationError is the validation error returned by Sql.Validate if the
// designated constraints aren't met.
type SqlValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SqlValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SqlValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SqlValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SqlValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SqlValidationError) ErrorName() string { return "SqlValidationError" }

// Error satisfies the builtin error interface
func (e SqlValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSql.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SqlValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SqlValidationError{}

// Validate checks the field values on Vm with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Vm) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Vm with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in VmMultiError, or nil if none found.
func (m *Vm) ValidateAll() error {
	return m.validate(true)
}

func (m *Vm) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Zone

	if len(errors) > 0 {
		return VmMultiError(errors)
	}

	return nil
}

// VmMultiError is an error wrapping multiple validation errors returned by
// Vm.ValidateAll() if the designated constraints aren't met.
type VmMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VmMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VmMultiError) AllErrors() []error { return m }

// VmValidationError is the validation error returned by Vm.Validate if the
// designated constraints aren't met.
type VmValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VmValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VmValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VmValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VmValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VmValidationError) ErrorName() string { return "VmValidationError" }

// Error satisfies the builtin error interface
func (e VmValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVm.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VmValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VmValidationError{}

// Validate checks the field values on Cluster_NodePool with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Cluster_NodePool) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cluster_NodePool with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Cluster_NodePoolMultiError, or nil if none found.
func (m *Cluster_NodePool) ValidateAll() error {
	return m.validate(true)
}

func (m *Cluster_NodePool) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for InitialNodeCount

	// no validation rules for CurrentSize

	if all {
		switch v := interface{}(m.GetAutoscaling()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Cluster_NodePoolValidationError{
					field:  "Autoscaling",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Cluster_NodePoolValidationError{
					field:  "Autoscaling",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAutoscaling()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Cluster_NodePoolValidationError{
				field:  "Autoscaling",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Cluster_NodePoolMultiError(errors)
	}

	return nil
}

// Cluster_NodePoolMultiError is an error wrapping multiple validation errors
// returned by Cluster_NodePool.ValidateAll() if the designated constraints
// aren't met.
type Cluster_NodePoolMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Cluster_NodePoolMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Cluster_NodePoolMultiError) AllErrors() []error { return m }

// Cluster_NodePoolValidationError is the validation error returned by
// Cluster_NodePool.Validate if the designated constraints aren't met.
type Cluster_NodePoolValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Cluster_NodePoolValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Cluster_NodePoolValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Cluster_NodePoolValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Cluster_NodePoolValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Cluster_NodePoolValidationError) ErrorName() string { return "Cluster_NodePoolValidationError" }

// Error satisfies the builtin error interface
func (e Cluster_NodePoolValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCluster_NodePool.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Cluster_NodePoolValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Cluster_NodePoolValidationError{}

// Validate checks the field values on Cluster_NodePool_AutoScaling with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Cluster_NodePool_AutoScaling) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cluster_NodePool_AutoScaling with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Cluster_NodePool_AutoScalingMultiError, or nil if none found.
func (m *Cluster_NodePool_AutoScaling) ValidateAll() error {
	return m.validate(true)
}

func (m *Cluster_NodePool_AutoScaling) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enabled

	// no validation rules for MinNodeCount

	// no validation rules for MaxNodeCount

	// no validation rules for Autoprovisioned

	// no validation rules for LocationPolicy

	// no validation rules for TotalMinNodeCount

	// no validation rules for TotalMaxNodeCount

	if len(errors) > 0 {
		return Cluster_NodePool_AutoScalingMultiError(errors)
	}

	return nil
}

// Cluster_NodePool_AutoScalingMultiError is an error wrapping multiple
// validation errors returned by Cluster_NodePool_AutoScaling.ValidateAll() if
// the designated constraints aren't met.
type Cluster_NodePool_AutoScalingMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Cluster_NodePool_AutoScalingMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Cluster_NodePool_AutoScalingMultiError) AllErrors() []error { return m }

// Cluster_NodePool_AutoScalingValidationError is the validation error returned
// by Cluster_NodePool_AutoScaling.Validate if the designated constraints
// aren't met.
type Cluster_NodePool_AutoScalingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Cluster_NodePool_AutoScalingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Cluster_NodePool_AutoScalingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Cluster_NodePool_AutoScalingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Cluster_NodePool_AutoScalingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Cluster_NodePool_AutoScalingValidationError) ErrorName() string {
	return "Cluster_NodePool_AutoScalingValidationError"
}

// Error satisfies the builtin error interface
func (e Cluster_NodePool_AutoScalingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCluster_NodePool_AutoScaling.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Cluster_NodePool_AutoScalingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Cluster_NodePool_AutoScalingValidationError{}