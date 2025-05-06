package pkgcrossplaneio


// DeploymentRuntimeConfigSpec specifies the configuration for a packaged controller.
//
// Values provided will override package manager defaults. Labels and
// annotations are passed to both the controller Deployment and ServiceAccount.
type DeploymentRuntimeConfigSpec struct {
	// DeploymentTemplate is the template for the Deployment object.
	DeploymentTemplate *DeploymentRuntimeConfigSpecDeploymentTemplate `field:"optional" json:"deploymentTemplate" yaml:"deploymentTemplate"`
	// ServiceAccountTemplate is the template for the ServiceAccount object.
	ServiceAccountTemplate *DeploymentRuntimeConfigSpecServiceAccountTemplate `field:"optional" json:"serviceAccountTemplate" yaml:"serviceAccountTemplate"`
	// ServiceTemplate is the template for the Service object.
	ServiceTemplate *DeploymentRuntimeConfigSpecServiceTemplate `field:"optional" json:"serviceTemplate" yaml:"serviceTemplate"`
}

