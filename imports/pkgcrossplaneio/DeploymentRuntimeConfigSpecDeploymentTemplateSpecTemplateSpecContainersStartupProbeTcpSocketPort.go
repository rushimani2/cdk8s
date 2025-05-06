package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Number or name of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort interface {
	Value() interface{}
}

// The jsii proxy struct for DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort
type jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort struct {
	_ byte // padding
}

func (j *jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort_FromNumber(value *float64) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort_FromString(value *string) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

