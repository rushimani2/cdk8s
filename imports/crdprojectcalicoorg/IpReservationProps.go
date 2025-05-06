package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type IpReservationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// IPReservationSpec contains the specification for an IPReservation resource.
	Spec *IpReservationSpec `field:"optional" json:"spec" yaml:"spec"`
}

