// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsLbListenerInvalidProtocolRule checks the pattern is valid
type AwsLbListenerInvalidProtocolRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsLbListenerInvalidProtocolRule returns new rule with default attributes
func NewAwsLbListenerInvalidProtocolRule() *AwsLbListenerInvalidProtocolRule {
	return &AwsLbListenerInvalidProtocolRule{
		resourceType:  "aws_lb_listener",
		attributeName: "protocol",
		enum: []string{
			"HTTP",
			"HTTPS",
			"TCP",
			"TLS",
			"UDP",
			"TCP_UDP",
		},
	}
}

// Name returns the rule name
func (r *AwsLbListenerInvalidProtocolRule) Name() string {
	return "aws_lb_listener_invalid_protocol"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLbListenerInvalidProtocolRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsLbListenerInvalidProtocolRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsLbListenerInvalidProtocolRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLbListenerInvalidProtocolRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`protocol is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
