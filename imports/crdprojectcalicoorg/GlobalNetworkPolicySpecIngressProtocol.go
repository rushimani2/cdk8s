package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol.
//
// Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).
// Must be one of these string values: "TCP", "UDP", "ICMP", "ICMPv6", "SCTP", "UDPLite" or an integer in the range 1-255.
type GlobalNetworkPolicySpecIngressProtocol interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecIngressProtocol
type jsiiProxy_GlobalNetworkPolicySpecIngressProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecIngressProtocol) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecIngressProtocol_FromNumber(value *float64) GlobalNetworkPolicySpecIngressProtocol {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressProtocol_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressProtocol",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecIngressProtocol_FromString(value *string) GlobalNetworkPolicySpecIngressProtocol {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecIngressProtocol_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecIngressProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressProtocol",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

