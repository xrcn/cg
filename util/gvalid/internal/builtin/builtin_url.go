package builtin

import (
	"errors"

	"github.com/xrcn/cg/text/gregex"
)

// RuleUrl implements `url` rule:
// URL.
//
// Format: url
type RuleUrl struct{}

func init() {
	Register(RuleUrl{})
}

func (r RuleUrl) Name() string {
	return "url"
}

func (r RuleUrl) Message() string {
	return "The {field} value `{value}` is not a valid URL address"
}

func (r RuleUrl) Run(in RunInput) error {
	ok := gregex.IsMatchString(
		`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`,
		in.Value.String(),
	)
	if ok {
		return nil
	}
	return errors.New(in.Message)
}
