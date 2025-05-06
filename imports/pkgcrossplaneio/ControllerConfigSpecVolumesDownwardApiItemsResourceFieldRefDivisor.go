package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specifies the output format of the exposed resources, defaults to "1".
type ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor interface {
	Value() interface{}
}

// The jsii proxy struct for ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor
type jsiiProxy_ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor struct {
	_ byte // padding
}

func (j *jsiiProxy_ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor_FromNumber(value *float64) ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor_FromString(value *string) ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor {
	_init_.Initialize()

	if err := validateControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

