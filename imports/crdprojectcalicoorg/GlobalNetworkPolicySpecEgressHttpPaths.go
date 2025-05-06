package crdprojectcalicoorg


// HTTPPath specifies an HTTP path to match.
//
// It may be either of the form: exact: <path>: which matches the path exactly or prefix: <path-prefix>: which matches the path prefix.
type GlobalNetworkPolicySpecEgressHttpPaths struct {
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

