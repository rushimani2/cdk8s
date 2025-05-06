package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The maximum number of pods that can be unavailable during the update.
//
// Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
// Absolute number is calculated from percentage by rounding down.
// This can not be 0 if MaxSurge is 0.
// Defaults to 25%.
// Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods
// immediately when the rolling update starts. Once new pods are ready, old ReplicaSet
// can be scaled down further, followed by scaling up the new ReplicaSet, ensuring
// that the total number of pods available at all times during the update is at
// least 70% of desired pods.
// Default: 25%.
//
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable interface {
	Value() interface{}
}

// The jsii proxy struct for DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable
type jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable struct {
	_ byte // padding
}

func (j *jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable_FromNumber(value *float64) DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable_FromString(value *string) DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

