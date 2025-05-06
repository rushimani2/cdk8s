package crdprojectcalicoorg


// IPAMBlockSpec contains the specification for an IPAMBlock resource.
type IpamBlockSpec struct {
	// Array of allocations in-use within this block.
	//
	// nil entries mean the allocation is free. For non-nil entries at index i, the index is the ordinal of the allocation within this block and the value is the index of the associated attributes in the Attributes array.
	Allocations *[]*float64 `field:"required" json:"allocations" yaml:"allocations"`
	// Attributes is an array of arbitrary metadata associated with allocations in the block.
	//
	// To find attributes for a given allocation, use the value of the allocation's entry in the Allocations array as the index of the element in this array.
	Attributes *[]*IpamBlockSpecAttributes `field:"required" json:"attributes" yaml:"attributes"`
	// The block's CIDR.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// StrictAffinity on the IPAMBlock is deprecated and no longer used by the code.
	//
	// Use IPAMConfig StrictAffinity instead.
	StrictAffinity *bool `field:"required" json:"strictAffinity" yaml:"strictAffinity"`
	// Unallocated is an ordered list of allocations which are free in the block.
	Unallocated *[]*float64 `field:"required" json:"unallocated" yaml:"unallocated"`
	// Affinity of the block, if this block has one.
	//
	// If set, it will be of the form "host:<hostname>". If not set, this block is not affine to a host.
	Affinity *string `field:"optional" json:"affinity" yaml:"affinity"`
	// Deleted is an internal boolean used to workaround a limitation in the Kubernetes API whereby deletion will not return a conflict error if the block has been updated.
	//
	// It should not be set manually.
	Deleted *bool `field:"optional" json:"deleted" yaml:"deleted"`
	// We store a sequence number that is updated each time the block is written.
	//
	// Each allocation will also store the sequence number of the block at the time of its creation. When releasing an IP, passing the sequence number associated with the allocation allows us to protect against a race condition and ensure the IP hasn't been released and re-allocated since the release request.
	SequenceNumber *float64 `field:"optional" json:"sequenceNumber" yaml:"sequenceNumber"`
	// Map of allocated ordinal within the block to sequence number of the block at the time of allocation.
	//
	// Kubernetes does not allow numerical keys for maps, so the key is cast to a string.
	SequenceNumberForAllocation *map[string]*float64 `field:"optional" json:"sequenceNumberForAllocation" yaml:"sequenceNumberForAllocation"`
}

