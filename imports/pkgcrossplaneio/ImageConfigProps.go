package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// The ImageConfig resource is used to configure settings for package images.
type ImageConfigProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ImageConfigSpec contains the configuration for matching images.
	Spec *ImageConfigSpec `field:"optional" json:"spec" yaml:"spec"`
}

