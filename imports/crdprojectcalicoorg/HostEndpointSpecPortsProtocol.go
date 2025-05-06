package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type HostEndpointSpecPortsProtocol interface {
	Value() interface{}
}

// The jsii proxy struct for HostEndpointSpecPortsProtocol
type jsiiProxy_HostEndpointSpecPortsProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_HostEndpointSpecPortsProtocol) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func HostEndpointSpecPortsProtocol_FromNumber(value *float64) HostEndpointSpecPortsProtocol {
	_init_.Initialize()

	if err := validateHostEndpointSpecPortsProtocol_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns HostEndpointSpecPortsProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.HostEndpointSpecPortsProtocol",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HostEndpointSpecPortsProtocol_FromString(value *string) HostEndpointSpecPortsProtocol {
	_init_.Initialize()

	if err := validateHostEndpointSpecPortsProtocol_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns HostEndpointSpecPortsProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.HostEndpointSpecPortsProtocol",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

