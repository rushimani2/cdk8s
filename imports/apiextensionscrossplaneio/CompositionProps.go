package apiextensionscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Composition defines a collection of managed resources or functions that Crossplane uses to create and manage new composite resources.
//
// Read the Crossplane documentation for
// [more information about Compositions](https://docs.crossplane.io/latest/concepts/compositions).
type CompositionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// CompositionSpec specifies desired state of a composition.
	Spec *CompositionSpec `field:"optional" json:"spec" yaml:"spec"`
}

