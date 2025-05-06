package pkgcrossplaneio


// Source for the environment variable's value.
//
// Cannot be used if value is not empty.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFrom struct {
	// Selects a key of a ConfigMap.
	ConfigMapKeyRef *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromConfigMapKeyRef `field:"optional" json:"configMapKeyRef" yaml:"configMapKeyRef"`
	// Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
	FieldRef *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromFieldRef `field:"optional" json:"fieldRef" yaml:"fieldRef"`
	// Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	ResourceFieldRef *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRef `field:"optional" json:"resourceFieldRef" yaml:"resourceFieldRef"`
	// Selects a key of a secret in the pod's namespace.
	SecretKeyRef *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

