package builtin

import (
	"errors"

	"github.com/xrcn/cg/text/gregex"
)

// RuleEmail implements `email` rule:
// Email address.
//
// Format: email
type RuleEmail struct{}

func init() {
	Register(RuleEmail{})
}

func (r RuleEmail) Name() string {
	return "email"
}

func (r RuleEmail) Message() string {
	return "The {field} value `{value}` is not a valid email address"
}

func (r RuleEmail) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`^[a-zA-Z0-9_\-\.]+@[a-zA-Z0-9_\-]+(\.[a-zA-Z0-9_\-]+)+$`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
