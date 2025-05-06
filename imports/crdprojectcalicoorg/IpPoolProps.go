package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type IpPoolProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// IPPoolSpec contains the specification for an IPPool resource.
	Spec *IpPoolSpec `field:"optional" json:"spec" yaml:"spec"`
}

