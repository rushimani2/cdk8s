package pkgcrossplaneio


// The deployment strategy to use to replace existing pods with new ones.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategy struct {
	// Rolling update config params.
	//
	// Present only if DeploymentStrategyType =
	// RollingUpdate.
	RollingUpdate *DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdate `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of deployment.
	//
	// Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	// Default: RollingUpdate.
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

