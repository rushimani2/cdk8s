package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests
type jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumber(value *float64) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromString(value *string) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

