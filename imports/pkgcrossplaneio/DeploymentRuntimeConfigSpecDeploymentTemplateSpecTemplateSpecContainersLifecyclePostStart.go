package pkgcrossplaneio


// PostStart is called immediately after a container is created.
//
// If the handler fails,
// the container is terminated and restarted according to its restart policy.
// Other management of the container blocks until the hook completes.
// More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStart struct {
	// Exec specifies the action to take.
	Exec *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartExec `field:"optional" json:"exec" yaml:"exec"`
	// HTTPGet specifies the http request to perform.
	HttpGet *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartHttpGet `field:"optional" json:"httpGet" yaml:"httpGet"`
	// Sleep represents the duration that the container should sleep before being terminated.
	Sleep *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartSleep `field:"optional" json:"sleep" yaml:"sleep"`
	// Deprecated.
	//
	// TCPSocket is NOT supported as a LifecycleHandler and kept
	// for the backward compatibility. There are no validation of this field and
	// lifecycle hooks will fail in runtime when tcp handler is specified.
	TcpSocket *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartTcpSocket `field:"optional" json:"tcpSocket" yaml:"tcpSocket"`
}

