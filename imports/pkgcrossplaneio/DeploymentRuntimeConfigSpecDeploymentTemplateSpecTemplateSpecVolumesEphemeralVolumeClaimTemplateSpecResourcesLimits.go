package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits
type jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumber(value *float64) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromString(value *string) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

