package crdprojectcalicoorg


// Metadata contains additional information for this rule.
type NetworkPolicySpecEgressMetadata struct {
	// Annotations is a set of key value pairs that give extra information about the rule.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
}

