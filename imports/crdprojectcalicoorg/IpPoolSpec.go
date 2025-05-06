package crdprojectcalicoorg


// IPPoolSpec contains the specification for an IPPool resource.
type IpPoolSpec struct {
	// The pool CIDR.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// AllowedUse controls what the IP pool will be used for.
	//
	// If not specified or empty, defaults to ["Tunnel", "Workload"] for back-compatibility.
	AllowedUses *[]*string `field:"optional" json:"allowedUses" yaml:"allowedUses"`
	// The block size to use for IP address assignments from this pool.
	//
	// Defaults to 26 for IPv4 and 122 for IPv6.
	// Default: 26 for IPv4 and 122 for IPv6.
	//
	BlockSize *float64 `field:"optional" json:"blockSize" yaml:"blockSize"`
	// Disable exporting routes from this IP Pool's CIDR over BGP.
	//
	// [Default: false].
	DisableBgpExport *bool `field:"optional" json:"disableBgpExport" yaml:"disableBgpExport"`
	// When disabled is true, Calico IPAM will not assign addresses from this pool.
	Disabled *bool `field:"optional" json:"disabled" yaml:"disabled"`
	// Deprecated: this field is only used for APIv1 backwards compatibility.
	//
	// Setting this field is not allowed, this field is for internal use only.
	Ipip *IpPoolSpecIpip `field:"optional" json:"ipip" yaml:"ipip"`
	// Contains configuration for IPIP tunneling for this pool.
	//
	// If not specified, then this is defaulted to "Never" (i.e. IPIP tunneling is disabled).
	IpipMode *string `field:"optional" json:"ipipMode" yaml:"ipipMode"`
	// Deprecated: this field is only used for APIv1 backwards compatibility.
	//
	// Setting this field is not allowed, this field is for internal use only.
	NatOutgoing *bool `field:"optional" json:"natOutgoing" yaml:"natOutgoing"`
	// Allows IPPool to allocate for a specific node by label selector.
	NodeSelector *string `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// Contains configuration for VXLAN tunneling for this pool.
	//
	// If not specified, then this is defaulted to "Never" (i.e. VXLAN tunneling is disabled).
	VxlanMode *string `field:"optional" json:"vxlanMode" yaml:"vxlanMode"`
}

