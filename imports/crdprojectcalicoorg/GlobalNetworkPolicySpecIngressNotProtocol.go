package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// NotProtocol is the negated version of the Protocol field.
type GlobalNetworkPolicySpecIngressNotProtocol interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecIngressNotProtocol
type jsiiProxy_GlobalNetworkPolicySpecIngressNotProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecIngressNotProtocol) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecIngressNotProtocol_FromNumber(value *float64) GlobalNetworkPolicySpecIngressNotProtocol {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressNotProtocol_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressNotProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressNotProtocol",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecIngressNotProtocol_FromString(value *string) GlobalNetworkPolicySpecIngressNotProtocol {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressNotProtocol_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressNotProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressNotProtocol",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

