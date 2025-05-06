package crdprojectcalicoorg


// NetworkSetSpec contains the specification for a NetworkSet resource.
type NetworkSetSpec struct {
	// The list of IP networks that belong to this set.
	Nets *[]*string `field:"optional" json:"nets" yaml:"nets"`
}

