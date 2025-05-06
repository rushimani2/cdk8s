package crdprojectcalicoorg


// ServiceExternalIPBlock represents a single allowed External IP CIDR block.
type BgpConfigurationSpecServiceExternalIPs struct {
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

