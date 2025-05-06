package apiextensionscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A CompositeResourceDefinition defines the schema for a new custom Kubernetes API.
//
// Read the Crossplane documentation for
// [more information about CustomResourceDefinitions](https://docs.crossplane.io/latest/concepts/composite-resource-definitions).
type CompositeResourceDefinitionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// CompositeResourceDefinitionSpec specifies the desired state of the definition.
	Spec *CompositeResourceDefinitionSpec `field:"optional" json:"spec" yaml:"spec"`
}

