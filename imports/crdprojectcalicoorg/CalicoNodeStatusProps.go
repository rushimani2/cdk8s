package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type CalicoNodeStatusProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// CalicoNodeStatusSpec contains the specification for a CalicoNodeStatus resource.
	Spec *CalicoNodeStatusSpec `field:"optional" json:"spec" yaml:"spec"`
}

