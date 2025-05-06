package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type NetworkPolicySpecEgressDestinationPorts interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecEgressDestinationPorts
type jsiiProxy_NetworkPolicySpecEgressDestinationPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecEgressDestinationPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecEgressDestinationPorts_FromNumber(value *float64) NetworkPolicySpecEgressDestinationPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressDestinationPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressDestinationPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressDestinationPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecEgressDestinationPorts_FromString(value *string) NetworkPolicySpecEgressDestinationPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressDestinationPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressDestinationPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressDestinationPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

