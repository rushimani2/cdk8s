package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type KubeControllersConfigurationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// KubeControllersConfigurationSpec contains the values of the Kubernetes controllers configuration.
	Spec *KubeControllersConfigurationSpec `field:"optional" json:"spec" yaml:"spec"`
}

