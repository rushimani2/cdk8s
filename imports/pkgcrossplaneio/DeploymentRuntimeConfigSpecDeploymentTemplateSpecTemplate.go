package pkgcrossplaneio


// Template describes the pods that will be created.
//
// The only allowed template.spec.restartPolicy value is "Always".
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplate struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior of the pod.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

