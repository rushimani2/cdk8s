package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type BlockAffinityProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// BlockAffinitySpec contains the specification for a BlockAffinity resource.
	Spec *BlockAffinitySpec `field:"optional" json:"spec" yaml:"spec"`
}

