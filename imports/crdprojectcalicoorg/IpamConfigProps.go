package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type IpamConfigProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// IPAMConfigSpec contains the specification for an IPAMConfig resource.
	Spec *IpamConfigSpec `field:"optional" json:"spec" yaml:"spec"`
}

