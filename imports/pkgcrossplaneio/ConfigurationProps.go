package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Configuration installs an OCI compatible Crossplane package, extending Crossplane with support for new kinds of CompositeResourceDefinitions and Compositions.
//
// Read the Crossplane documentation for
// [more information about Configuration packages](https://docs.crossplane.io/latest/concepts/packages).
type ConfigurationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ConfigurationSpec specifies details about a request to install a configuration to Crossplane.
	Spec *ConfigurationSpec `field:"optional" json:"spec" yaml:"spec"`
}

