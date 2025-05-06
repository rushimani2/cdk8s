package pkgcrossplaneio


// DeploymentTemplate is the template for the Deployment object.
type DeploymentRuntimeConfigSpecDeploymentTemplate struct {
	// Metadata contains the configurable metadata fields for the Deployment.
	Metadata *DeploymentRuntimeConfigSpecDeploymentTemplateMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Spec contains the configurable spec fields for the Deployment object.
	Spec *DeploymentRuntimeConfigSpecDeploymentTemplateSpec `field:"optional" json:"spec" yaml:"spec"`
}

