package crdprojectcalicoorg

import (
	_init_ "example.com/cdk8s-demo/imports/crdprojectcalicoorg/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type FelixConfigurationSpecKubeNodePortRanges interface {
	Value() interface{}
}

// The jsii proxy struct for FelixConfigurationSpecKubeNodePortRanges
type jsiiProxy_FelixConfigurationSpecKubeNodePortRanges struct {
	_ byte // padding
}

func (j *jsiiProxy_FelixConfigurationSpecKubeNodePortRanges) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func FelixConfigurationSpecKubeNodePortRanges_FromNumber(value *float64) FelixConfigurationSpecKubeNodePortRanges {
	_init_.Initialize()

	if err := validateFelixConfigurationSpecKubeNodePortRanges_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns FelixConfigurationSpecKubeNodePortRanges

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.FelixConfigurationSpecKubeNodePortRanges",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func FelixConfigurationSpecKubeNodePortRanges_FromString(value *string) FelixConfigurationSpecKubeNodePortRanges {
	_init_.Initialize()

	if err := validateFelixConfigurationSpecKubeNodePortRanges_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns FelixConfigurationSpecKubeNodePortRanges

	_jsii_.StaticInvoke(
		"crdprojectcalicoorg.FelixConfigurationSpecKubeNodePortRanges",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

