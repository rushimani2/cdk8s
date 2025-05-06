package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// BGPConfiguration contains the configuration for any BGP routing.
type BgpConfigurationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// BGPConfigurationSpec contains the values of the BGP configuration.
	Spec *BgpConfigurationSpec `field:"optional" json:"spec" yaml:"spec"`
}

