package crdprojectcalicoorg


// ICMP is an optional field that restricts the rule to apply to a specific type and code of ICMP traffic.
//
// This should only be specified if the Protocol field is set to "ICMP" or "ICMPv6".
type GlobalNetworkPolicySpecIngressIcmp struct {
	// Match on a specific ICMP code.
	//
	// If specified, the Type value must also be specified. This is a technical limitation imposed by the kernel's iptables firewall, which Calico uses to enforce the rule.
	Code *float64 `field:"optional" json:"code" yaml:"code"`
	// Match on a specific ICMP type.
	//
	// For example a value of 8 refers to ICMP Echo Request (i.e. pings).
	Type *float64 `field:"optional" json:"type" yaml:"type"`
}

