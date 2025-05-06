package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// BPFPSNATPorts sets the range from which we randomly pick a port if there is a source port collision.
//
// This should be within the ephemeral range as defined by RFC 6056 (1024–65535) and preferably outside the  ephemeral ranges used by common operating systems. Linux uses 32768–60999, while others mostly use the IANA defined range 49152–65535. It is not necessarily a problem if this range overlaps with the operating systems. Both ends of the range are inclusive. [Default: 20000:29999]
type FelixConfigurationSpecBpfPsnatPorts interface {
	Value() interface{}
}

// The jsii proxy struct for FelixConfigurationSpecBpfPsnatPorts
type jsiiProxy_FelixConfigurationSpecBpfPsnatPorts struct {
	_ byte // padding
}

func (j *jsiiProxy_FelixConfigurationSpecBpfPsnatPorts) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func FelixConfigurationSpecBpfPsnatPorts_FromNumber(value *float64) FelixConfigurationSpecBpfPsnatPorts {
	_init_.Initialize()

	if err := validateFelixConfigurationSpecBpfPsnatPorts_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns FelixConfigurationSpecBpfPsnatPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.FelixConfigurationSpecBpfPsnatPorts",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func FelixConfigurationSpecBpfPsnatPorts_FromString(value *string) FelixConfigurationSpecBpfPsnatPorts {
	_init_.Initialize()

	if err := validateFelixConfigurationSpecBpfPsnatPorts_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns FelixConfigurationSpecBpfPsnatPorts

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.FelixConfigurationSpecBpfPsnatPorts",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

