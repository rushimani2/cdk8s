package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type HostEndpointProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// HostEndpointSpec contains the specification for a HostEndpoint resource.
	Spec *HostEndpointSpec `field:"optional" json:"spec" yaml:"spec"`
}

