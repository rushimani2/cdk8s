package pkgcrossplaneio


// LabelSelector is used to find matching pods.
//
// Pods that match this label selector are counted to determine the number of pods
// in their corresponding topology domain.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecTopologySpreadConstraintsLabelSelector struct {
	// matchExpressions is a list of label selector requirements.
	//
	// The requirements are ANDed.
	MatchExpressions *[]*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecTopologySpreadConstraintsLabelSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// matchLabels is a map of {key,value} pairs.
	//
	// A single {key,value} in the matchLabels
	// map is equivalent to an element of matchExpressions, whose key field is "key", the
	// operator is "In", and the values array contains only "value". The requirements are ANDed.
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

