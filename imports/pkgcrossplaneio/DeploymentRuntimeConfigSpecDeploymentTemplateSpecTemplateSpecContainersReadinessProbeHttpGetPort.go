package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Name or number of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort
type jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort_FromNumber(value *float64) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort_FromString(value *string) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

