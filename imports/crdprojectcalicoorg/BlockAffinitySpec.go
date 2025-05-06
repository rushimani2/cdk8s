package crdprojectcalicoorg


// BlockAffinitySpec contains the specification for a BlockAffinity resource.
type BlockAffinitySpec struct {
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Deleted indicates that this block affinity is being deleted.
	//
	// This field is a string for compatibility with older releases that mistakenly treat this field as a string.
	Deleted *string `field:"required" json:"deleted" yaml:"deleted"`
	Node *string `field:"required" json:"node" yaml:"node"`
	State *string `field:"required" json:"state" yaml:"state"`
}

