package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// NotProtocol is the negated version of the Protocol field.
type NetworkPolicySpecEgressNotProtocol interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecEgressNotProtocol
type jsiiProxy_NetworkPolicySpecEgressNotProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecEgressNotProtocol) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecEgressNotProtocol_FromNumber(value *float64) NetworkPolicySpecEgressNotProtocol {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressNotProtocol_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressNotProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressNotProtocol",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecEgressNotProtocol_FromString(value *string) NetworkPolicySpecEgressNotProtocol {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressNotProtocol_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressNotProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressNotProtocol",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

