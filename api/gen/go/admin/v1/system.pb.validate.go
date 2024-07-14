// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: admin/v1/system.proto

package adminv1

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

// Validate checks the field values on System with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *System) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on System with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SystemMultiError, or nil if none found.
func (m *System) ValidateAll() error {
	return m.validate(true)
}

func (m *System) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.Theme != nil {
		// no validation rules for Theme
	}

	if m.Title != nil {
		// no validation rules for Title
	}

	if m.Keywords != nil {
		// no validation rules for Keywords
	}

	if m.Description != nil {
		// no validation rules for Description
	}

	if m.RecordNumber != nil {
		// no validation rules for RecordNumber
	}

	if len(errors) > 0 {
		return SystemMultiError(errors)
	}

	return nil
}

// SystemMultiError is an error wrapping multiple validation errors returned by
// System.ValidateAll() if the designated constraints aren't met.
type SystemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SystemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SystemMultiError) AllErrors() []error { return m }

// SystemValidationError is the validation error returned by System.Validate if
// the designated constraints aren't met.
type SystemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SystemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SystemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SystemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SystemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SystemValidationError) ErrorName() string { return "SystemValidationError" }

// Error satisfies the builtin error interface
func (e SystemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSystem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SystemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SystemValidationError{}

// Validate checks the field values on ListSystemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListSystemResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListSystemResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListSystemResponseMultiError, or nil if none found.
func (m *ListSystemResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListSystemResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListSystemResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListSystemResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSystemResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return ListSystemResponseMultiError(errors)
	}

	return nil
}

// ListSystemResponseMultiError is an error wrapping multiple validation errors
// returned by ListSystemResponse.ValidateAll() if the designated constraints
// aren't met.
type ListSystemResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListSystemResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListSystemResponseMultiError) AllErrors() []error { return m }

// ListSystemResponseValidationError is the validation error returned by
// ListSystemResponse.Validate if the designated constraints aren't met.
type ListSystemResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSystemResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSystemResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSystemResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSystemResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSystemResponseValidationError) ErrorName() string {
	return "ListSystemResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListSystemResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSystemResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSystemResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSystemResponseValidationError{}

// Validate checks the field values on GetSystemRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetSystemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSystemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSystemRequestMultiError, or nil if none found.
func (m *GetSystemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSystemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetSystemRequestMultiError(errors)
	}

	return nil
}

// GetSystemRequestMultiError is an error wrapping multiple validation errors
// returned by GetSystemRequest.ValidateAll() if the designated constraints
// aren't met.
type GetSystemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSystemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSystemRequestMultiError) AllErrors() []error { return m }

// GetSystemRequestValidationError is the validation error returned by
// GetSystemRequest.Validate if the designated constraints aren't met.
type GetSystemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSystemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSystemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSystemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSystemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSystemRequestValidationError) ErrorName() string { return "GetSystemRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetSystemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSystemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSystemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSystemRequestValidationError{}

// Validate checks the field values on CreateSystemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateSystemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateSystemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateSystemRequestMultiError, or nil if none found.
func (m *CreateSystemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateSystemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSystem()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateSystemRequestValidationError{
					field:  "System",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateSystemRequestValidationError{
					field:  "System",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSystem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSystemRequestValidationError{
				field:  "System",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return CreateSystemRequestMultiError(errors)
	}

	return nil
}

// CreateSystemRequestMultiError is an error wrapping multiple validation
// errors returned by CreateSystemRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateSystemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateSystemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateSystemRequestMultiError) AllErrors() []error { return m }

// CreateSystemRequestValidationError is the validation error returned by
// CreateSystemRequest.Validate if the designated constraints aren't met.
type CreateSystemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSystemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSystemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSystemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSystemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSystemRequestValidationError) ErrorName() string {
	return "CreateSystemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSystemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSystemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSystemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSystemRequestValidationError{}

// Validate checks the field values on UpdateSystemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateSystemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateSystemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateSystemRequestMultiError, or nil if none found.
func (m *UpdateSystemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateSystemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetSystem()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateSystemRequestValidationError{
					field:  "System",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateSystemRequestValidationError{
					field:  "System",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSystem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateSystemRequestValidationError{
				field:  "System",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return UpdateSystemRequestMultiError(errors)
	}

	return nil
}

// UpdateSystemRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateSystemRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateSystemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateSystemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateSystemRequestMultiError) AllErrors() []error { return m }

// UpdateSystemRequestValidationError is the validation error returned by
// UpdateSystemRequest.Validate if the designated constraints aren't met.
type UpdateSystemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateSystemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateSystemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateSystemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateSystemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateSystemRequestValidationError) ErrorName() string {
	return "UpdateSystemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateSystemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateSystemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateSystemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateSystemRequestValidationError{}

// Validate checks the field values on DeleteSystemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteSystemRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteSystemRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteSystemRequestMultiError, or nil if none found.
func (m *DeleteSystemRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteSystemRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if m.OperatorId != nil {
		// no validation rules for OperatorId
	}

	if len(errors) > 0 {
		return DeleteSystemRequestMultiError(errors)
	}

	return nil
}

// DeleteSystemRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteSystemRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteSystemRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteSystemRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteSystemRequestMultiError) AllErrors() []error { return m }

// DeleteSystemRequestValidationError is the validation error returned by
// DeleteSystemRequest.Validate if the designated constraints aren't met.
type DeleteSystemRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteSystemRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteSystemRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteSystemRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteSystemRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteSystemRequestValidationError) ErrorName() string {
	return "DeleteSystemRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteSystemRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteSystemRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteSystemRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteSystemRequestValidationError{}
