package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type NetworkPolicySpecEgressDestinationNotPorts interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecEgressDestinationNotPorts
type jsiiProxy_NetworkPolicySpecEgressDestinationNotPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecEgressDestinationNotPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecEgressDestinationNotPorts_FromNumber(value *float64) NetworkPolicySpecEgressDestinationNotPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressDestinationNotPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressDestinationNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressDestinationNotPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecEgressDestinationNotPorts_FromString(value *string) NetworkPolicySpecEgressDestinationNotPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressDestinationNotPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressDestinationNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressDestinationNotPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

