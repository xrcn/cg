package builtin

import (
	"errors"

	"github.com/xrcn/cg/text/gregex"
)

// RuleNotRegex implements `not-regex` rule:
// Value should not match custom regular expression pattern.
//
// Format: not-regex:pattern
type RuleNotRegex struct{}

func init() {
	Register(RuleNotRegex{})
}

func (r RuleNotRegex) Name() string {
	return "not-regex"
}

func (r RuleNotRegex) Message() string {
	return "The {field} value `{value}` should not be in regex of: {pattern}"
}

func (r RuleNotRegex) Run(in RunInput) error {
	if gregex.IsMatchString(in.RulePattern, in.Value.String()) {
		return errors.New(in.Message)
	}
	return nil
}
