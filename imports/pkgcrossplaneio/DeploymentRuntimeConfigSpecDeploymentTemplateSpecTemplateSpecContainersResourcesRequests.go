package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests interface {
	Value() interface{}
}

// The jsii proxy struct for DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests
type jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests struct {
	_ byte // padding
}

func (j *jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests_FromNumber(value *float64) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests_FromString(value *string) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

