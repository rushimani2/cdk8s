package crdprojectcalicoorg


// A Rule encapsulates a set of match criteria and an action.
//
// Both selector-based security Policy and security Profiles reference rules - separated out as a list of rules for both ingress and egress packet matching.
// Each positive match criteria has a negated version, prefixed with "Not". All the match criteria within a rule must be satisfied for a packet to match. A single rule can contain the positive and negative version of a match and both must be satisfied for the rule to match.
type GlobalNetworkPolicySpecIngress struct {
	Action *string `field:"required" json:"action" yaml:"action"`
	// Destination contains the match criteria that apply to destination entity.
	Destination *GlobalNetworkPolicySpecIngressDestination `field:"optional" json:"destination" yaml:"destination"`
	// HTTP contains match criteria that apply to HTTP requests.
	Http *GlobalNetworkPolicySpecIngressHttp `field:"optional" json:"http" yaml:"http"`
	// ICMP is an optional field that restricts the rule to apply to a specific type and code of ICMP traffic.
	//
	// This should only be specified if the Protocol field is set to "ICMP" or "ICMPv6".
	Icmp *GlobalNetworkPolicySpecIngressIcmp `field:"optional" json:"icmp" yaml:"icmp"`
	// IPVersion is an optional field that restricts the rule to only match a specific IP version.
	IpVersion *float64 `field:"optional" json:"ipVersion" yaml:"ipVersion"`
	// Metadata contains additional information for this rule.
	Metadata *GlobalNetworkPolicySpecIngressMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// NotICMP is the negated version of the ICMP field.
	NotIcmp *GlobalNetworkPolicySpecIngressNotIcmp `field:"optional" json:"notIcmp" yaml:"notIcmp"`
	// NotProtocol is the negated version of the Protocol field.
	NotProtocol GlobalNetworkPolicySpecIngressNotProtocol `field:"optional" json:"notProtocol" yaml:"notProtocol"`
	// Protocol is an optional field that restricts the rule to only apply to traffic of a specific IP protocol.
	//
	// Required if any of the EntityRules contain Ports (because ports only apply to certain protocols).
	// Must be one of these string values: "TCP", "UDP", "ICMP", "ICMPv6", "SCTP", "UDPLite" or an integer in the range 1-255.
	Protocol GlobalNetworkPolicySpecIngressProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// Source contains the match criteria that apply to source entity.
	Source *GlobalNetworkPolicySpecIngressSource `field:"optional" json:"source" yaml:"source"`
}

