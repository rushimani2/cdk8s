package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type GlobalNetworkPolicySpecEgressDestinationPorts interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecEgressDestinationPorts
type jsiiProxy_GlobalNetworkPolicySpecEgressDestinationPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecEgressDestinationPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecEgressDestinationPorts_FromNumber(value *float64) GlobalNetworkPolicySpecEgressDestinationPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressDestinationPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressDestinationPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressDestinationPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecEgressDestinationPorts_FromString(value *string) GlobalNetworkPolicySpecEgressDestinationPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressDestinationPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressDestinationPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressDestinationPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

