package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// NATPortRange specifies the range of ports that is used for port mapping when doing outgoing NAT.
//
// When unset the default behavior of the network stack is used.
type FelixConfigurationSpecNatPortRange interface {
	Value() interface{}
}

// The jsii proxy struct for FelixConfigurationSpecNatPortRange
type jsiiProxy_FelixConfigurationSpecNatPortRange struct {
	_ byte // padding
}

func (j *jsiiProxy_FelixConfigurationSpecNatPortRange) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func FelixConfigurationSpecNatPortRange_FromNumber(value *float64) FelixConfigurationSpecNatPortRange {
	_init_.Initialize()

	if err := validateFelixConfigurationSpecNatPortRange_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns FelixConfigurationSpecNatPortRange

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.FelixConfigurationSpecNatPortRange",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func FelixConfigurationSpecNatPortRange_FromString(value *string) FelixConfigurationSpecNatPortRange {
	_init_.Initialize()

	if err := validateFelixConfigurationSpecNatPortRange_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns FelixConfigurationSpecNatPortRange

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.FelixConfigurationSpecNatPortRange",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

