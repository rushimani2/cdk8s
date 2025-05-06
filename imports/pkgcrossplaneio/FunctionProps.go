package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Function installs an OCI compatible Crossplane package, extending Crossplane with support for a new kind of composition function.
//
// Read the Crossplane documentation for
// [more information about Functions](https://docs.crossplane.io/latest/concepts/composition-functions).
type FunctionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// FunctionSpec specifies the configuration of a Function.
	Spec *FunctionSpec `field:"optional" json:"spec" yaml:"spec"`
}

