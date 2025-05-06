package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type NetworkPolicySpecIngressDestinationPorts interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecIngressDestinationPorts
type jsiiProxy_NetworkPolicySpecIngressDestinationPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecIngressDestinationPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecIngressDestinationPorts_FromNumber(value *float64) NetworkPolicySpecIngressDestinationPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressDestinationPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressDestinationPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressDestinationPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecIngressDestinationPorts_FromString(value *string) NetworkPolicySpecIngressDestinationPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressDestinationPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressDestinationPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressDestinationPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

