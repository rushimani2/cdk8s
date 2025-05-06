package pkgcrossplaneio


// A node selector term, associated with the corresponding weight.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference struct {
	// A list of node selector requirements by node's labels.
	MatchExpressions *[]*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// A list of node selector requirements by node's fields.
	MatchFields *[]*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields `field:"optional" json:"matchFields" yaml:"matchFields"`
}

