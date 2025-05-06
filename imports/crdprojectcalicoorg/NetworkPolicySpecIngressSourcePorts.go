package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type NetworkPolicySpecIngressSourcePorts interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecIngressSourcePorts
type jsiiProxy_NetworkPolicySpecIngressSourcePorts struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecIngressSourcePorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecIngressSourcePorts_FromNumber(value *float64) NetworkPolicySpecIngressSourcePorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressSourcePorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressSourcePorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressSourcePorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecIngressSourcePorts_FromString(value *string) NetworkPolicySpecIngressSourcePorts {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressSourcePorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressSourcePorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressSourcePorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

