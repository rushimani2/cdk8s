package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type GlobalNetworkPolicySpecEgressSourceNotPorts interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecEgressSourceNotPorts
type jsiiProxy_GlobalNetworkPolicySpecEgressSourceNotPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecEgressSourceNotPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecEgressSourceNotPorts_FromNumber(value *float64) GlobalNetworkPolicySpecEgressSourceNotPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressSourceNotPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressSourceNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressSourceNotPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecEgressSourceNotPorts_FromString(value *string) GlobalNetworkPolicySpecEgressSourceNotPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressSourceNotPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressSourceNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressSourceNotPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

