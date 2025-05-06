package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol.
//
// Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).
// Must be one of these string values: "TCP", "UDP", "ICMP", "ICMPv6", "SCTP", "UDPLite" or an integer in the range 1-255.
type GlobalNetworkPolicySpecEgressProtocol interface {
	Value() interface{}
}

// The jsii proxy struct for GlobalNetworkPolicySpecEgressProtocol
type jsiiProxy_GlobalNetworkPolicySpecEgressProtocol struct {
	_ byte // padding
}

func (j *jsiiProxy_GlobalNetworkPolicySpecEgressProtocol) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GlobalNetworkPolicySpecEgressProtocol_FromNumber(value *float64) GlobalNetworkPolicySpecEgressProtocol {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressProtocol_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressProtocol",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func GlobalNetworkPolicySpecEgressProtocol_FromString(value *string) GlobalNetworkPolicySpecEgressProtocol {
	_init_.Initialize()

	if err := validateGlobalNetworkPolicySpecEgressProtocol_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns GlobalNetworkPolicySpecEgressProtocol

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressProtocol",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

