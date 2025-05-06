package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// NotProtocol is the negated version of the Protocol field.
type NetworkPolicySpecIngressNotProtocol interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecIngressNotProtocol
type jsiiProxy_NetworkPolicySpecIngressNotProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecIngressNotProtocol) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecIngressNotProtocol_FromNumber(value *float64) NetworkPolicySpecIngressNotProtocol {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressNotProtocol_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressNotProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressNotProtocol",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecIngressNotProtocol_FromString(value *string) NetworkPolicySpecIngressNotProtocol {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressNotProtocol_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressNotProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressNotProtocol",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

