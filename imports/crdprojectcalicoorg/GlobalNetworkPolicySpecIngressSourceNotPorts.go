package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type GlobalNetworkPolicySpecIngressSourceNotPorts interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecIngressSourceNotPorts
type jsiiProxy_GlobalNetworkPolicySpecIngressSourceNotPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecIngressSourceNotPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecIngressSourceNotPorts_FromNumber(value *float64) GlobalNetworkPolicySpecIngressSourceNotPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressSourceNotPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressSourceNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressSourceNotPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecIngressSourceNotPorts_FromString(value *string) GlobalNetworkPolicySpecIngressSourceNotPorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressSourceNotPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressSourceNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressSourceNotPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

