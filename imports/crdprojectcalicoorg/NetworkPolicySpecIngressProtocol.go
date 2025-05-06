package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol.
//
// Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).
// Must be one of these string values: "TCP", "UDP", "ICMP", "ICMPv6", "SCTP", "UDPLite" or an integer in the range 1-255.
type NetworkPolicySpecIngressProtocol interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecIngressProtocol
type jsiiProxy_NetworkPolicySpecIngressProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecIngressProtocol) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecIngressProtocol_FromNumber(value *float64) NetworkPolicySpecIngressProtocol {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressProtocol_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressProtocol",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecIngressProtocol_FromString(value *string) NetworkPolicySpecIngressProtocol {
	_init_.Initialize()

	if err := validateNetworkPolicySpecIngressProtocol_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecIngressProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecIngressProtocol",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

