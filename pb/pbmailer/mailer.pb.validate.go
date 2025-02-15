// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: mailer.proto

package pbmailer

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
var _mailer_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ForgotPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ForgotPasswordRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		return ForgotPasswordRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
	}

	// no validation rules for Name

	for idx, item := range m.GetTokenizedUsers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ForgotPasswordRequestValidationError{
					field:  fmt.Sprintf("TokenizedUsers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

func (m *ForgotPasswordRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *ForgotPasswordRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// ForgotPasswordRequestValidationError is the validation error returned by
// ForgotPasswordRequest.Validate if the designated constraints aren't met.
type ForgotPasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForgotPasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForgotPasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForgotPasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForgotPasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForgotPasswordRequestValidationError) ErrorName() string {
	return "ForgotPasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ForgotPasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForgotPasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForgotPasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForgotPasswordRequestValidationError{}

// Validate checks the field values on TokenizedUser with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TokenizedUser) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := _TokenizedUser_Role_InLookup[m.GetRole()]; !ok {
		return TokenizedUserValidationError{
			field:  "Role",
			reason: "value must be in list [service-provider consumer merchant]",
		}
	}

	// no validation rules for Token

	return nil
}

// TokenizedUserValidationError is the validation error returned by
// TokenizedUser.Validate if the designated constraints aren't met.
type TokenizedUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TokenizedUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TokenizedUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TokenizedUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TokenizedUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TokenizedUserValidationError) ErrorName() string { return "TokenizedUserValidationError" }

// Error satisfies the builtin error interface
func (e TokenizedUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTokenizedUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TokenizedUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TokenizedUserValidationError{}

var _TokenizedUser_Role_InLookup = map[string]struct{}{
	"service-provider": {},
	"consumer":         {},
	"merchant":         {},
}
