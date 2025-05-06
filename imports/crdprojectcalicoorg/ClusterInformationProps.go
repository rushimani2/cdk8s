package crdprojectcalicoorg

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// ClusterInformation contains the cluster specific information.
type ClusterInformationProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ClusterInformationSpec contains the values of describing the cluster.
	Spec *ClusterInformationSpec `field:"optional" json:"spec" yaml:"spec"`
}

