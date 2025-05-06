package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type GlobalNetworkPolicySpecIngressDestinationPorts interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecIngressDestinationPorts
type jsiiProxy_GlobalNetworkPolicySpecIngressDestinationPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecIngressDestinationPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecIngressDestinationPorts_FromNumber(value *float64) GlobalNetworkPolicySpecIngressDestinationPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressDestinationPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressDestinationPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressDestinationPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecIngressDestinationPorts_FromString(value *string) GlobalNetworkPolicySpecIngressDestinationPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressDestinationPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressDestinationPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressDestinationPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

