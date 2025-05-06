package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type GlobalNetworkPolicyProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Spec *GlobalNetworkPolicySpec `field:"optional" json:"spec" yaml:"spec"`
}

