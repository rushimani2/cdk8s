package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// GlobalNetworkSet contains a set of arbitrary IP sub-networks/CIDRs that share labels to allow rules to refer to them via selectors.
//
// The labels of GlobalNetworkSet are not namespaced.
type GlobalNetworkSetProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// GlobalNetworkSetSpec contains the specification for a NetworkSet resource.
	Spec *GlobalNetworkSetSpec `field:"optional" json:"spec" yaml:"spec"`
}

