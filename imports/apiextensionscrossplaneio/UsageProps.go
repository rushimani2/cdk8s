package apiextensionscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A Usage defines a deletion blocking relationship between two resources.
//
// Usages prevent accidental deletion of a single resource or deletion of
// resources with dependent resources.
//
// Read the Crossplane documentation for
// [more information about Compositions](https://docs.crossplane.io/latest/concepts/usages).
type UsageProps struct {
	// UsageSpec defines the desired state of Usage.
	Spec *UsageSpec `field:"required" json:"spec" yaml:"spec"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

