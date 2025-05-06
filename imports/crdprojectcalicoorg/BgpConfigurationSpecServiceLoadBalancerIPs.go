package crdprojectcalicoorg


// ServiceLoadBalancerIPBlock represents a single allowed LoadBalancer IP CIDR block.
type BgpConfigurationSpecServiceLoadBalancerIPs struct {
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

