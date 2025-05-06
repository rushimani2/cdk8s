package crdprojectcalicoorg


// IPAMHandleSpec contains the specification for an IPAMHandle resource.
type IpamHandleSpec struct {
	Block *map[string]*float64 `field:"required" json:"block" yaml:"block"`
	HandleId *string `field:"required" json:"handleId" yaml:"handleId"`
	Deleted *bool `field:"optional" json:"deleted" yaml:"deleted"`
}

