package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// NetworkSet is the Namespaced-equivalent of the GlobalNetworkSet.
type NetworkSetProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// NetworkSetSpec contains the specification for a NetworkSet resource.
	Spec *NetworkSetSpec `field:"optional" json:"spec" yaml:"spec"`
}

