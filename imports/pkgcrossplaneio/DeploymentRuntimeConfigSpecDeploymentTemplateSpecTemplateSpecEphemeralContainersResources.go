package pkgcrossplaneio


// Resources are not allowed for ephemeral containers.
//
// Ephemeral containers use spare resources
// already allocated to the pod.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResources struct {
	// Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.
	//
	// This is an alpha field and requires enabling the
	// DynamicResourceAllocation feature gate.
	//
	// This field is immutable. It can only be set for containers.
	Claims *[]*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesClaims `field:"optional" json:"claims" yaml:"claims"`
	// Limits describes the maximum amount of compute resources allowed.
	//
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits *map[string]DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	// Requests describes the minimum amount of compute resources required.
	//
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
	// otherwise to an implementation-defined value. Requests cannot exceed Limits.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests *map[string]DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

