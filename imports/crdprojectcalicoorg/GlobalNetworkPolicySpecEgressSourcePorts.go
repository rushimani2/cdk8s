package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type GlobalNetworkPolicySpecEgressSourcePorts interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecEgressSourcePorts
type jsiiProxy_GlobalNetworkPolicySpecEgressSourcePorts struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecEgressSourcePorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecEgressSourcePorts_FromNumber(value *float64) GlobalNetworkPolicySpecEgressSourcePorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressSourcePorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressSourcePorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressSourcePorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecEgressSourcePorts_FromString(value *string) GlobalNetworkPolicySpecEgressSourcePorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressSourcePorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressSourcePorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressSourcePorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

