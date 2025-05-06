package crdprojectcalicoorg


// GlobalNetworkSetSpec contains the specification for a NetworkSet resource.
type GlobalNetworkSetSpec struct {
	// The list of IP networks that belong to this set.
	Nets *[]*string `field:"optional" json:"nets" yaml:"nets"`
}

