package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type GlobalNetworkPolicySpecIngressDestinationNotPorts interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecIngressDestinationNotPorts
type jsiiProxy_GlobalNetworkPolicySpecIngressDestinationNotPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecIngressDestinationNotPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecIngressDestinationNotPorts_FromNumber(value *float64) GlobalNetworkPolicySpecIngressDestinationNotPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressDestinationNotPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressDestinationNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressDestinationNotPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecIngressDestinationNotPorts_FromString(value *string) GlobalNetworkPolicySpecIngressDestinationNotPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressDestinationNotPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressDestinationNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressDestinationNotPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

