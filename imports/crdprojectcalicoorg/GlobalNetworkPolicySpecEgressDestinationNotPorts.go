package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type GlobalNetworkPolicySpecEgressDestinationNotPorts interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecEgressDestinationNotPorts
type jsiiProxy_GlobalNetworkPolicySpecEgressDestinationNotPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecEgressDestinationNotPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecEgressDestinationNotPorts_FromNumber(value *float64) GlobalNetworkPolicySpecEgressDestinationNotPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressDestinationNotPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressDestinationNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressDestinationNotPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecEgressDestinationNotPorts_FromString(value *string) GlobalNetworkPolicySpecEgressDestinationNotPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressDestinationNotPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressDestinationNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressDestinationNotPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

