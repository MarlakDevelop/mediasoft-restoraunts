// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: restaurant_product.proto

package restaurant

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

// Validate checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProductRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductRequestMultiError, or nil if none found.
func (m *CreateProductRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Type

	// no validation rules for Weight

	// no validation rules for Price

	if len(errors) > 0 {
		return CreateProductRequestMultiError(errors)
	}

	return nil
}

// CreateProductRequestMultiError is an error wrapping multiple validation
// errors returned by CreateProductRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateProductRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductRequestMultiError) AllErrors() []error { return m }

// CreateProductRequestValidationError is the validation error returned by
// CreateProductRequest.Validate if the designated constraints aren't met.
type CreateProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductRequestValidationError) ErrorName() string {
	return "CreateProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductRequestValidationError{}

// Validate checks the field values on CreateProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateProductResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductResponseMultiError, or nil if none found.
func (m *CreateProductResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CreateProductResponseMultiError(errors)
	}

	return nil
}

// CreateProductResponseMultiError is an error wrapping multiple validation
// errors returned by CreateProductResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateProductResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductResponseMultiError) AllErrors() []error { return m }

// CreateProductResponseValidationError is the validation error returned by
// CreateProductResponse.Validate if the designated constraints aren't met.
type CreateProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductResponseValidationError) ErrorName() string {
	return "CreateProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductResponseValidationError{}

// Validate checks the field values on GetProductListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProductListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductListRequestMultiError, or nil if none found.
func (m *GetProductListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetProductListRequestMultiError(errors)
	}

	return nil
}

// GetProductListRequestMultiError is an error wrapping multiple validation
// errors returned by GetProductListRequest.ValidateAll() if the designated
// constraints aren't met.
type GetProductListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductListRequestMultiError) AllErrors() []error { return m }

// GetProductListRequestValidationError is the validation error returned by
// GetProductListRequest.Validate if the designated constraints aren't met.
type GetProductListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductListRequestValidationError) ErrorName() string {
	return "GetProductListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductListRequestValidationError{}

// Validate checks the field values on GetProductListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProductListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductListResponseMultiError, or nil if none found.
func (m *GetProductListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResult() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetProductListResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetProductListResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetProductListResponseValidationError{
					field:  fmt.Sprintf("Result[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetProductListResponseMultiError(errors)
	}

	return nil
}

// GetProductListResponseMultiError is an error wrapping multiple validation
// errors returned by GetProductListResponse.ValidateAll() if the designated
// constraints aren't met.
type GetProductListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductListResponseMultiError) AllErrors() []error { return m }

// GetProductListResponseValidationError is the validation error returned by
// GetProductListResponse.Validate if the designated constraints aren't met.
type GetProductListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductListResponseValidationError) ErrorName() string {
	return "GetProductListResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductListResponseValidationError{}

// Validate checks the field values on GetProductListByUuidsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProductListByUuidsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductListByUuidsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductListByUuidsRequestMultiError, or nil if none found.
func (m *GetProductListByUuidsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductListByUuidsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetProductListByUuidsRequestMultiError(errors)
	}

	return nil
}

// GetProductListByUuidsRequestMultiError is an error wrapping multiple
// validation errors returned by GetProductListByUuidsRequest.ValidateAll() if
// the designated constraints aren't met.
type GetProductListByUuidsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductListByUuidsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductListByUuidsRequestMultiError) AllErrors() []error { return m }

// GetProductListByUuidsRequestValidationError is the validation error returned
// by GetProductListByUuidsRequest.Validate if the designated constraints
// aren't met.
type GetProductListByUuidsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductListByUuidsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductListByUuidsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductListByUuidsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductListByUuidsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductListByUuidsRequestValidationError) ErrorName() string {
	return "GetProductListByUuidsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductListByUuidsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductListByUuidsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductListByUuidsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductListByUuidsRequestValidationError{}

// Validate checks the field values on GetProductListByUuidsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProductListByUuidsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductListByUuidsResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetProductListByUuidsResponseMultiError, or nil if none found.
func (m *GetProductListByUuidsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductListByUuidsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResult() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetProductListByUuidsResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetProductListByUuidsResponseValidationError{
						field:  fmt.Sprintf("Result[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetProductListByUuidsResponseValidationError{
					field:  fmt.Sprintf("Result[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetProductListByUuidsResponseMultiError(errors)
	}

	return nil
}

// GetProductListByUuidsResponseMultiError is an error wrapping multiple
// validation errors returned by GetProductListByUuidsResponse.ValidateAll()
// if the designated constraints aren't met.
type GetProductListByUuidsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductListByUuidsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductListByUuidsResponseMultiError) AllErrors() []error { return m }

// GetProductListByUuidsResponseValidationError is the validation error
// returned by GetProductListByUuidsResponse.Validate if the designated
// constraints aren't met.
type GetProductListByUuidsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductListByUuidsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductListByUuidsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductListByUuidsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductListByUuidsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductListByUuidsResponseValidationError) ErrorName() string {
	return "GetProductListByUuidsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductListByUuidsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductListByUuidsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductListByUuidsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductListByUuidsResponseValidationError{}

// Validate checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Product) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Product with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ProductMultiError, or nil if none found.
func (m *Product) ValidateAll() error {
	return m.validate(true)
}

func (m *Product) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Type

	// no validation rules for Weight

	// no validation rules for Price

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProductValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProductValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProductMultiError(errors)
	}

	return nil
}

// ProductMultiError is an error wrapping multiple validation errors returned
// by Product.ValidateAll() if the designated constraints aren't met.
type ProductMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProductMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProductMultiError) AllErrors() []error { return m }

// ProductValidationError is the validation error returned by Product.Validate
// if the designated constraints aren't met.
type ProductValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProductValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProductValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProductValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProductValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProductValidationError) ErrorName() string { return "ProductValidationError" }

// Error satisfies the builtin error interface
func (e ProductValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProduct.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProductValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProductValidationError{}
