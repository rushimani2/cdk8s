package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type IpamBlockProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// IPAMBlockSpec contains the specification for an IPAMBlock resource.
	Spec *IpamBlockSpec `field:"optional" json:"spec" yaml:"spec"`
}

