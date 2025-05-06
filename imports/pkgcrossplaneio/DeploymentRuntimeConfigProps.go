package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// The DeploymentRuntimeConfig provides settings for the Kubernetes Deployment of a Provider or composition function package.
//
// Read the Crossplane documentation for
// [more information about DeploymentRuntimeConfigs](https://docs.crossplane.io/latest/concepts/providers/#runtime-configuration).
type DeploymentRuntimeConfigProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// DeploymentRuntimeConfigSpec specifies the configuration for a packaged controller.
	//
	// Values provided will override package manager defaults. Labels and
	// annotations are passed to both the controller Deployment and ServiceAccount.
	Spec *DeploymentRuntimeConfigSpec `field:"optional" json:"spec" yaml:"spec"`
}

