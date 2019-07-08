// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsLaunchConfigurationInvalidTypeRule checks the pattern is valid
type AwsLaunchConfigurationInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsLaunchConfigurationInvalidTypeRule returns new rule with default attributes
func NewAwsLaunchConfigurationInvalidTypeRule() *AwsLaunchConfigurationInvalidTypeRule {
	return &AwsLaunchConfigurationInvalidTypeRule{
		resourceType:  "aws_launch_configuration",
		attributeName: "instance_type",
		enum: []string{
			"t1.micro",
			"t2.nano",
			"t2.micro",
			"t2.small",
			"t2.medium",
			"t2.large",
			"t2.xlarge",
			"t2.2xlarge",
			"t3.nano",
			"t3.micro",
			"t3.small",
			"t3.medium",
			"t3.large",
			"t3.xlarge",
			"t3.2xlarge",
			"t3a.nano",
			"t3a.micro",
			"t3a.small",
			"t3a.medium",
			"t3a.large",
			"t3a.xlarge",
			"t3a.2xlarge",
			"m1.small",
			"m1.medium",
			"m1.large",
			"m1.xlarge",
			"m3.medium",
			"m3.large",
			"m3.xlarge",
			"m3.2xlarge",
			"m4.large",
			"m4.xlarge",
			"m4.2xlarge",
			"m4.4xlarge",
			"m4.10xlarge",
			"m4.16xlarge",
			"m2.xlarge",
			"m2.2xlarge",
			"m2.4xlarge",
			"cr1.8xlarge",
			"r3.large",
			"r3.xlarge",
			"r3.2xlarge",
			"r3.4xlarge",
			"r3.8xlarge",
			"r4.large",
			"r4.xlarge",
			"r4.2xlarge",
			"r4.4xlarge",
			"r4.8xlarge",
			"r4.16xlarge",
			"r5.large",
			"r5.xlarge",
			"r5.2xlarge",
			"r5.4xlarge",
			"r5.8xlarge",
			"r5.12xlarge",
			"r5.16xlarge",
			"r5.24xlarge",
			"r5.metal",
			"r5a.large",
			"r5a.xlarge",
			"r5a.2xlarge",
			"r5a.4xlarge",
			"r5a.8xlarge",
			"r5a.12xlarge",
			"r5a.16xlarge",
			"r5a.24xlarge",
			"r5d.large",
			"r5d.xlarge",
			"r5d.2xlarge",
			"r5d.4xlarge",
			"r5d.8xlarge",
			"r5d.12xlarge",
			"r5d.16xlarge",
			"r5d.24xlarge",
			"r5d.metal",
			"r5ad.large",
			"r5ad.xlarge",
			"r5ad.2xlarge",
			"r5ad.4xlarge",
			"r5ad.8xlarge",
			"r5ad.12xlarge",
			"r5ad.16xlarge",
			"r5ad.24xlarge",
			"x1.16xlarge",
			"x1.32xlarge",
			"x1e.xlarge",
			"x1e.2xlarge",
			"x1e.4xlarge",
			"x1e.8xlarge",
			"x1e.16xlarge",
			"x1e.32xlarge",
			"i2.xlarge",
			"i2.2xlarge",
			"i2.4xlarge",
			"i2.8xlarge",
			"i3.large",
			"i3.xlarge",
			"i3.2xlarge",
			"i3.4xlarge",
			"i3.8xlarge",
			"i3.16xlarge",
			"i3.metal",
			"i3en.large",
			"i3en.xlarge",
			"i3en.2xlarge",
			"i3en.3xlarge",
			"i3en.6xlarge",
			"i3en.12xlarge",
			"i3en.24xlarge",
			"hi1.4xlarge",
			"hs1.8xlarge",
			"c1.medium",
			"c1.xlarge",
			"c3.large",
			"c3.xlarge",
			"c3.2xlarge",
			"c3.4xlarge",
			"c3.8xlarge",
			"c4.large",
			"c4.xlarge",
			"c4.2xlarge",
			"c4.4xlarge",
			"c4.8xlarge",
			"c5.large",
			"c5.xlarge",
			"c5.2xlarge",
			"c5.4xlarge",
			"c5.9xlarge",
			"c5.12xlarge",
			"c5.18xlarge",
			"c5.24xlarge",
			"c5.metal",
			"c5d.large",
			"c5d.xlarge",
			"c5d.2xlarge",
			"c5d.4xlarge",
			"c5d.9xlarge",
			"c5d.18xlarge",
			"c5n.large",
			"c5n.xlarge",
			"c5n.2xlarge",
			"c5n.4xlarge",
			"c5n.9xlarge",
			"c5n.18xlarge",
			"cc1.4xlarge",
			"cc2.8xlarge",
			"g2.2xlarge",
			"g2.8xlarge",
			"g3.4xlarge",
			"g3.8xlarge",
			"g3.16xlarge",
			"g3s.xlarge",
			"cg1.4xlarge",
			"p2.xlarge",
			"p2.8xlarge",
			"p2.16xlarge",
			"p3.2xlarge",
			"p3.8xlarge",
			"p3.16xlarge",
			"p3dn.24xlarge",
			"d2.xlarge",
			"d2.2xlarge",
			"d2.4xlarge",
			"d2.8xlarge",
			"f1.2xlarge",
			"f1.4xlarge",
			"f1.16xlarge",
			"m5.large",
			"m5.xlarge",
			"m5.2xlarge",
			"m5.4xlarge",
			"m5.8xlarge",
			"m5.12xlarge",
			"m5.16xlarge",
			"m5.24xlarge",
			"m5.metal",
			"m5a.large",
			"m5a.xlarge",
			"m5a.2xlarge",
			"m5a.4xlarge",
			"m5a.8xlarge",
			"m5a.12xlarge",
			"m5a.16xlarge",
			"m5a.24xlarge",
			"m5d.large",
			"m5d.xlarge",
			"m5d.2xlarge",
			"m5d.4xlarge",
			"m5d.8xlarge",
			"m5d.12xlarge",
			"m5d.16xlarge",
			"m5d.24xlarge",
			"m5d.metal",
			"m5ad.large",
			"m5ad.xlarge",
			"m5ad.2xlarge",
			"m5ad.4xlarge",
			"m5ad.8xlarge",
			"m5ad.12xlarge",
			"m5ad.16xlarge",
			"m5ad.24xlarge",
			"h1.2xlarge",
			"h1.4xlarge",
			"h1.8xlarge",
			"h1.16xlarge",
			"z1d.large",
			"z1d.xlarge",
			"z1d.2xlarge",
			"z1d.3xlarge",
			"z1d.6xlarge",
			"z1d.12xlarge",
			"z1d.metal",
			"u-6tb1.metal",
			"u-9tb1.metal",
			"u-12tb1.metal",
			"a1.medium",
			"a1.large",
			"a1.xlarge",
			"a1.2xlarge",
			"a1.4xlarge",
		},
	}
}

// Name returns the rule name
func (r *AwsLaunchConfigurationInvalidTypeRule) Name() string {
	return "aws_launch_configuration_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLaunchConfigurationInvalidTypeRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsLaunchConfigurationInvalidTypeRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsLaunchConfigurationInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLaunchConfigurationInvalidTypeRule) Check(runner *tflint.Runner) error {
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
					`instance_type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
