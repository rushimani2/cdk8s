package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type BgpPeerProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// BGPPeerSpec contains the specification for a BGPPeer resource.
	Spec *BgpPeerSpec `field:"optional" json:"spec" yaml:"spec"`
}

