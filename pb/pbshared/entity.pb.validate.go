// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: shared/entity.proto

package pbshared

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

// define the regex for a UUID once up-front
var _entity_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Entity with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Entity) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EntityId

	// no validation rules for EntityTable

	return nil
}

// EntityValidationError is the validation error returned by Entity.Validate if
// the designated constraints aren't met.
type EntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EntityValidationError) ErrorName() string { return "EntityValidationError" }

// Error satisfies the builtin error interface
func (e EntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EntityValidationError{}
