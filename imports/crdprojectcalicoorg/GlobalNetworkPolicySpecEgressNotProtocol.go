package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// NotProtocol is the negated version of the Protocol field.
type GlobalNetworkPolicySpecEgressNotProtocol interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecEgressNotProtocol
type jsiiProxy_GlobalNetworkPolicySpecEgressNotProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecEgressNotProtocol) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecEgressNotProtocol_FromNumber(value *float64) GlobalNetworkPolicySpecEgressNotProtocol {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressNotProtocol_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressNotProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressNotProtocol",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecEgressNotProtocol_FromString(value *string) GlobalNetworkPolicySpecEgressNotProtocol {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressNotProtocol_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressNotProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressNotProtocol",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

