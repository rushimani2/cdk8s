package crdprojectcalicoorg


// ServiceClusterIPBlock represents a single allowed ClusterIP CIDR block.
type BgpConfigurationSpecServiceClusterIPs struct {
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

