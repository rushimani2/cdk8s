package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type NetworkPolicySpecIngressSourceNotPorts interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecIngressSourceNotPorts
type jsiiProxy_NetworkPolicySpecIngressSourceNotPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecIngressSourceNotPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecIngressSourceNotPorts_FromNumber(value *float64) NetworkPolicySpecIngressSourceNotPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressSourceNotPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressSourceNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressSourceNotPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecIngressSourceNotPorts_FromString(value *string) NetworkPolicySpecIngressSourceNotPorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressSourceNotPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressSourceNotPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressSourceNotPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

