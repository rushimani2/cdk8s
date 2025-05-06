package pkgcrossplaneio


// Deprecated.
//
// TCPSocket is NOT supported as a LifecycleHandler and kept
// for the backward compatibility. There are no validation of this field and
// lifecycle hooks will fail in runtime when tcp handler is specified.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartTcpSocket struct {
	// Number or name of the port to access on the container.
	//
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	Port DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartTcpSocketPort `field:"required" json:"port" yaml:"port"`
	// Optional: Host name to connect to, defaults to the pod IP.
	Host *string `field:"optional" json:"host" yaml:"host"`
}

