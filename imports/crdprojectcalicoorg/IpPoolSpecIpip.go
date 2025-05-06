package crdprojectcalicoorg


// Deprecated: this field is only used for APIv1 backwards compatibility.
//
// Setting this field is not allowed, this field is for internal use only.
type IpPoolSpecIpip struct {
	// When enabled is true, ipip tunneling will be used to deliver packets to destinations within this pool.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The IPIP mode.
	//
	// This can be one of "always" or "cross-subnet".  A mode of "always" will also use IPIP tunneling for routing to destination IP addresses within this pool.  A mode of "cross-subnet" will only use IPIP tunneling when the destination node is on a different subnet to the originating node.  The default value (if not specified) is "always".
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

