package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type NetworkPolicySpecIngressDestinationNotPorts interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecIngressDestinationNotPorts
type jsiiProxy_NetworkPolicySpecIngressDestinationNotPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecIngressDestinationNotPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecIngressDestinationNotPorts_FromNumber(value *float64) NetworkPolicySpecIngressDestinationNotPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressDestinationNotPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressDestinationNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressDestinationNotPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecIngressDestinationNotPorts_FromString(value *string) NetworkPolicySpecIngressDestinationNotPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressDestinationNotPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressDestinationNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressDestinationNotPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

