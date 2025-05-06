package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Felix Configuration contains the configuration for Felix.
type FelixConfigurationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// FelixConfigurationSpec contains the values of the Felix configuration.
	Spec *FelixConfigurationSpec `field:"optional" json:"spec" yaml:"spec"`
}

