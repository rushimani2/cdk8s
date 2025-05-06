package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type NetworkPolicySpecEgressSourcePorts interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecEgressSourcePorts
type jsiiProxy_NetworkPolicySpecEgressSourcePorts struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecEgressSourcePorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecEgressSourcePorts_FromNumber(value *float64) NetworkPolicySpecEgressSourcePorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressSourcePorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressSourcePorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressSourcePorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecEgressSourcePorts_FromString(value *string) NetworkPolicySpecEgressSourcePorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressSourcePorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressSourcePorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressSourcePorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

