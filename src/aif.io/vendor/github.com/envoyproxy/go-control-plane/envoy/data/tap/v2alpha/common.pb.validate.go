// Code generated by protoc-gen-validate
// source: envoy/data/tap/v2alpha/common.proto
// DO NOT EDIT!!!

package v2

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on Body with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Body) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Truncated

	switch m.BodyType.(type) {

	case *Body_AsBytes:
		// no validation rules for AsBytes

	case *Body_AsString:
		// no validation rules for AsString

	}

	return nil
}

// BodyValidationError is the validation error returned by Body.Validate if the
// designated constraints aren't met.
type BodyValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e BodyValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBody.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = BodyValidationError{}
