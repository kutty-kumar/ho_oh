// Code generated by protoc-gen-validate
// source: ho-oh/ditto_v1/ditto.proto
// DO NOT EDIT!!!

package ditto_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on PrinterDto with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *PrinterDto) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for SerialNumber

	// no validation rules for Status

	// no validation rules for ExternalId

	// no validation rules for ProductNumber

	if v, ok := interface{}(m.GetFromDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PrinterDtoValidationError{
				field:  "FromDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetToDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PrinterDtoValidationError{
				field:  "ToDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FromIndex

	// no validation rules for ToIndex

	return nil
}

// PrinterDtoValidationError is the validation error returned by
// PrinterDto.Validate if the designated constraints aren't met.
type PrinterDtoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrinterDtoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrinterDtoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrinterDtoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrinterDtoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrinterDtoValidationError) ErrorName() string { return "PrinterDtoValidationError" }

// Error satisfies the builtin error interface
func (e PrinterDtoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrinterDto.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrinterDtoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrinterDtoValidationError{}

// Validate checks the field values on CreatePrinterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreatePrinterRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePrinterRequestValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreatePrinterRequestValidationError is the validation error returned by
// CreatePrinterRequest.Validate if the designated constraints aren't met.
type CreatePrinterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePrinterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePrinterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePrinterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePrinterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePrinterRequestValidationError) ErrorName() string {
	return "CreatePrinterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePrinterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePrinterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePrinterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePrinterRequestValidationError{}

// Validate checks the field values on CreatePrinterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreatePrinterResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreatePrinterResponseValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreatePrinterResponseValidationError is the validation error returned by
// CreatePrinterResponse.Validate if the designated constraints aren't met.
type CreatePrinterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreatePrinterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreatePrinterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreatePrinterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreatePrinterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreatePrinterResponseValidationError) ErrorName() string {
	return "CreatePrinterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreatePrinterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreatePrinterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreatePrinterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreatePrinterResponseValidationError{}

// Validate checks the field values on UpdatePrinterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdatePrinterRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PrinterId

	if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePrinterRequestValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdatePrinterRequestValidationError is the validation error returned by
// UpdatePrinterRequest.Validate if the designated constraints aren't met.
type UpdatePrinterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePrinterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePrinterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePrinterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePrinterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePrinterRequestValidationError) ErrorName() string {
	return "UpdatePrinterRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePrinterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePrinterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePrinterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePrinterRequestValidationError{}

// Validate checks the field values on UpdatePrinterResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdatePrinterResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdatePrinterResponseValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdatePrinterResponseValidationError is the validation error returned by
// UpdatePrinterResponse.Validate if the designated constraints aren't met.
type UpdatePrinterResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdatePrinterResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdatePrinterResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdatePrinterResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdatePrinterResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdatePrinterResponseValidationError) ErrorName() string {
	return "UpdatePrinterResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdatePrinterResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdatePrinterResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdatePrinterResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdatePrinterResponseValidationError{}

// Validate checks the field values on GetPrinterByExternalIdRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPrinterByExternalIdRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PrinterId

	return nil
}

// GetPrinterByExternalIdRequestValidationError is the validation error
// returned by GetPrinterByExternalIdRequest.Validate if the designated
// constraints aren't met.
type GetPrinterByExternalIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPrinterByExternalIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPrinterByExternalIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPrinterByExternalIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPrinterByExternalIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPrinterByExternalIdRequestValidationError) ErrorName() string {
	return "GetPrinterByExternalIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetPrinterByExternalIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPrinterByExternalIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPrinterByExternalIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPrinterByExternalIdRequestValidationError{}

// Validate checks the field values on GetPrinterByExternalIdResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetPrinterByExternalIdResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetPrinterByExternalIdResponseValidationError{
				field:  "Response",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetPrinterByExternalIdResponseValidationError is the validation error
// returned by GetPrinterByExternalIdResponse.Validate if the designated
// constraints aren't met.
type GetPrinterByExternalIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetPrinterByExternalIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetPrinterByExternalIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetPrinterByExternalIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetPrinterByExternalIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetPrinterByExternalIdResponseValidationError) ErrorName() string {
	return "GetPrinterByExternalIdResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetPrinterByExternalIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetPrinterByExternalIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetPrinterByExternalIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetPrinterByExternalIdResponseValidationError{}

// Validate checks the field values on MultiGetPrintersByExternalIdRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *MultiGetPrintersByExternalIdRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// MultiGetPrintersByExternalIdRequestValidationError is the validation error
// returned by MultiGetPrintersByExternalIdRequest.Validate if the designated
// constraints aren't met.
type MultiGetPrintersByExternalIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiGetPrintersByExternalIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiGetPrintersByExternalIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiGetPrintersByExternalIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiGetPrintersByExternalIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiGetPrintersByExternalIdRequestValidationError) ErrorName() string {
	return "MultiGetPrintersByExternalIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiGetPrintersByExternalIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiGetPrintersByExternalIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiGetPrintersByExternalIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiGetPrintersByExternalIdRequestValidationError{}

// Validate checks the field values on MultiGetPrintersByExternalIdResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *MultiGetPrintersByExternalIdResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResult() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiGetPrintersByExternalIdResponseValidationError{
					field:  fmt.Sprintf("Result[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiGetPrintersByExternalIdResponseValidationError is the validation error
// returned by MultiGetPrintersByExternalIdResponse.Validate if the designated
// constraints aren't met.
type MultiGetPrintersByExternalIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiGetPrintersByExternalIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiGetPrintersByExternalIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiGetPrintersByExternalIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiGetPrintersByExternalIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiGetPrintersByExternalIdResponseValidationError) ErrorName() string {
	return "MultiGetPrintersByExternalIdResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiGetPrintersByExternalIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiGetPrintersByExternalIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiGetPrintersByExternalIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiGetPrintersByExternalIdResponseValidationError{}
