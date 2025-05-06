package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type IpamHandleProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// IPAMHandleSpec contains the specification for an IPAMHandle resource.
	Spec *IpamHandleSpec `field:"optional" json:"spec" yaml:"spec"`
}

