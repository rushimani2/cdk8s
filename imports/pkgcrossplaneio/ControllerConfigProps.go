package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A ControllerConfig applies settings to controllers like Provider pods. Deprecated: Use the [DeploymentRuntimeConfig](https://docs.crossplane.io/latest/concepts/providers#runtime-configuration) instead.
//
// Read the
// [Package Runtime Configuration](https://github.com/crossplane/crossplane/blob/11bbe13ea3604928cc4e24e8d0d18f3f5f7e847c/design/one-pager-package-runtime-config.md)
// design document for more details.
type ControllerConfigProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ControllerConfigSpec specifies the configuration for a packaged controller.
	//
	// Values provided will override package manager defaults. Labels and
	// annotations are passed to both the controller Deployment and ServiceAccount.
	Spec *ControllerConfigSpec `field:"optional" json:"spec" yaml:"spec"`
}

