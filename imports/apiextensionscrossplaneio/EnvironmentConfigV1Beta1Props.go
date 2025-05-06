package apiextensionscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// An EnvironmentConfig contains user-defined unstructured values for use in a Composition.
//
// Read the Crossplane documentation for
// [more information about EnvironmentConfigs](https://docs.crossplane.io/latest/concepts/environment-configs).
type EnvironmentConfigV1Beta1Props struct {
	// The data of this EnvironmentConfig.
	//
	// This may contain any kind of structure that can be serialized into JSON.
	Data *map[string]interface{} `field:"optional" json:"data" yaml:"data"`
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

