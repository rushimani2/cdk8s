package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol.
//
// Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).
// Must be one of these string values: "TCP", "UDP", "ICMP", "ICMPv6", "SCTP", "UDPLite" or an integer in the range 1-255.
type NetworkPolicySpecEgressProtocol interface {
	Value() interface{}
}

// The jsii proxy struct for NetworkPolicySpecEgressProtocol
type jsiiProxy_NetworkPolicySpecEgressProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkPolicySpecEgressProtocol) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func NetworkPolicySpecEgressProtocol_FromNumber(value *float64) NetworkPolicySpecEgressProtocol {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressProtocol_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressProtocol",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkPolicySpecEgressProtocol_FromString(value *string) NetworkPolicySpecEgressProtocol {
	_init_.Initialize()

	if err := validateNetworkPolicySpecEgressProtocol_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkPolicySpecEgressProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.NetworkPolicySpecEgressProtocol",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

