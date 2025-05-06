package pkgcrossplaneio


// Label selector for pods.
//
// Existing ReplicaSets whose pods are
// selected by this will be the ones affected by this deployment.
// It must match the pod template's labels.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecSelector struct {
	// matchExpressions is a list of label selector requirements.
	//
	// The requirements are ANDed.
	MatchExpressions *[]*DeploymentRuntimeConfigSpecDeploymentTemplateSpecSelectorMatchExpressions `field:"optional" json:"matchExpressions" yaml:"matchExpressions"`
	// matchLabels is a map of {key,value} pairs.
	//
	// A single {key,value} in the matchLabels
	// map is equivalent to an element of matchExpressions, whose key field is "key", the
	// operator is "In", and the values array contains only "value". The requirements are ANDed.
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

