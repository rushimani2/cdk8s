package pkgcrossplaneio


// PodResourceClaim references exactly one ResourceClaim, either directly or by naming a ResourceClaimTemplate which is then turned into a ResourceClaim for the pod.
//
// It adds a name to it that uniquely identifies the ResourceClaim inside the Pod.
// Containers that need access to the ResourceClaim reference it with this name.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecResourceClaims struct {
	// Name uniquely identifies this resource claim inside the pod.
	//
	// This must be a DNS_LABEL.
	Name *string `field:"required" json:"name" yaml:"name"`
	// ResourceClaimName is the name of a ResourceClaim object in the same namespace as this pod.
	//
	// Exactly one of ResourceClaimName and ResourceClaimTemplateName must
	// be set.
	ResourceClaimName *string `field:"optional" json:"resourceClaimName" yaml:"resourceClaimName"`
	// ResourceClaimTemplateName is the name of a ResourceClaimTemplate object in the same namespace as this pod.
	//
	// The template will be used to create a new ResourceClaim, which will
	// be bound to this pod. When this pod is deleted, the ResourceClaim
	// will also be deleted. The pod name and resource name, along with a
	// generated component, will be used to form a unique name for the
	// ResourceClaim, which will be recorded in pod.status.resourceClaimStatuses.
	//
	// This field is immutable and no changes will be made to the
	// corresponding ResourceClaim by the control plane after creating the
	// ResourceClaim.
	//
	// Exactly one of ResourceClaimName and ResourceClaimTemplateName must
	// be set.
	ResourceClaimTemplateName *string `field:"optional" json:"resourceClaimTemplateName" yaml:"resourceClaimTemplateName"`
}

