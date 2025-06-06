package pkgcrossplaneio


// If specified, the pod's scheduling constraints.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinity struct {
	// Describes node affinity scheduling rules for the pod.
	NodeAffinity *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinity `field:"optional" json:"nodeAffinity" yaml:"nodeAffinity"`
	// Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).
	PodAffinity *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinity `field:"optional" json:"podAffinity" yaml:"podAffinity"`
	// Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).
	PodAntiAffinity *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinity `field:"optional" json:"podAntiAffinity" yaml:"podAntiAffinity"`
}

