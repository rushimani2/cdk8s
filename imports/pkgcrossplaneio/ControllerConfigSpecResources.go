package pkgcrossplaneio


// Compute Resources required by this container.
//
// Cannot be updated.
// More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/
type ControllerConfigSpecResources struct {
	// Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.
	//
	// This is an alpha field and requires enabling the
	// DynamicResourceAllocation feature gate.
	//
	// This field is immutable. It can only be set for containers.
	Claims *[]*ControllerConfigSpecResourcesClaims `field:"optional" json:"claims" yaml:"claims"`
	// Limits describes the maximum amount of compute resources allowed.
	//
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Limits *map[string]ControllerConfigSpecResourcesLimits `field:"optional" json:"limits" yaml:"limits"`
	// Requests describes the minimum amount of compute resources required.
	//
	// If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,
	// otherwise to an implementation-defined value. Requests cannot exceed Limits.
	// More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Requests *map[string]ControllerConfigSpecResourcesRequests `field:"optional" json:"requests" yaml:"requests"`
}

