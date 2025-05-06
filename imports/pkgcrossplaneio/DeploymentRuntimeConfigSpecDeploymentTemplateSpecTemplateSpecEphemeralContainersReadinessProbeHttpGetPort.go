package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Name or number of the port to access on the container.
//
// Number must be in the range 1 to 65535.
// Name must be an IANA_SVC_NAME.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort interface {
	Value() interface{}
}

// The jsii proxy struct for DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort
type jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort struct {
	_ byte // padding
}

func (j *jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort_FromNumber(value *float64) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort_FromString(value *string) DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort {
	_init_.Initialize()

	if err := validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

