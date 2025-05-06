package crdprojectcalicoorg


// IPAMConfigSpec contains the specification for an IPAMConfig resource.
type IpamConfigSpec struct {
	AutoAllocateBlocks *bool `field:"required" json:"autoAllocateBlocks" yaml:"autoAllocateBlocks"`
	StrictAffinity *bool `field:"required" json:"strictAffinity" yaml:"strictAffinity"`
	// MaxBlocksPerHost, if non-zero, is the max number of blocks that can be affine to each host.
	MaxBlocksPerHost *float64 `field:"optional" json:"maxBlocksPerHost" yaml:"maxBlocksPerHost"`
}

