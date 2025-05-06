package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Provider installs an OCI compatible Crossplane package, extending Crossplane with support for new kinds of managed resources.
//
// Read the Crossplane documentation for
// [more information about Providers](https://docs.crossplane.io/latest/concepts/providers).
type ProviderProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ProviderSpec specifies details about a request to install a provider to Crossplane.
	Spec *ProviderSpec `field:"optional" json:"spec" yaml:"spec"`
}

