package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type GlobalNetworkPolicySpecIngressSourcePorts interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecIngressSourcePorts
type jsiiProxy_GlobalNetworkPolicySpecIngressSourcePorts struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecIngressSourcePorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecIngressSourcePorts_FromNumber(value *float64) GlobalNetworkPolicySpecIngressSourcePorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressSourcePorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressSourcePorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressSourcePorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecIngressSourcePorts_FromString(value *string) GlobalNetworkPolicySpecIngressSourcePorts {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressSourcePorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressSourcePorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressSourcePorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

