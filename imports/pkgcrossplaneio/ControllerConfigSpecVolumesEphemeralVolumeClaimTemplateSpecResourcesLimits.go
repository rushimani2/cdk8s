package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits interface {
	Value() interface{}
}

// The jsii proxy struct for ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits
type jsiiProxy_ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits struct {
	_ byte // padding
}

func (j *jsiiProxy_ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumber(value *float64) ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromString(value *string) ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits {
	_init_.Initialize()

	if err := validateControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

