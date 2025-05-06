package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type NetworkPolicySpecEgressSourceNotPorts interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecEgressSourceNotPorts
type jsiiProxy_NetworkPolicySpecEgressSourceNotPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecEgressSourceNotPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecEgressSourceNotPorts_FromNumber(value *float64) NetworkPolicySpecEgressSourceNotPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressSourceNotPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressSourceNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressSourceNotPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecEgressSourceNotPorts_FromString(value *string) NetworkPolicySpecEgressSourceNotPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressSourceNotPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressSourceNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressSourceNotPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

