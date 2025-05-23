package crdprojectcalicoorg


// HostEndpointSpec contains the specification for a HostEndpoint resource.
type HostEndpointSpec struct {
	// The expected IP addresses (IPv4 and IPv6) of the endpoint.
	//
	// If "InterfaceName" is not present, Calico will look for an interface matching any of the IPs in the list and apply policy to that. Note: 	When using the selector match criteria in an ingress or egress security Policy 	or Profile, Calico converts the selector into a set of IP addresses. For host 	endpoints, the ExpectedIPs field is used for that purpose. (If only the interface 	name is specified, Calico does not learn the IPs of the interface for use in match 	criteria.)
	ExpectedIPs *[]*string `field:"optional" json:"expectedIPs" yaml:"expectedIPs"`
	// Either "*", or the name of a specific Linux interface to apply policy to;
	//
	// or empty.  "*" indicates that this HostEndpoint governs all traffic to, from or through the default network namespace of the host named by the "Node" field; entering and leaving that namespace via any interface, including those from/to non-host-networked local workloads.
	// If InterfaceName is not "*", this HostEndpoint only governs traffic that enters or leaves the host through the specific interface named by InterfaceName, or - when InterfaceName is empty - through the specific interface that has one of the IPs in ExpectedIPs. Therefore, when InterfaceName is empty, at least one expected IP must be specified.  Only external interfaces (such as "eth0") are supported here; it isn't possible for a HostEndpoint to protect traffic through a specific local workload interface.
	// Note: Only some kinds of policy are implemented for "*" HostEndpoints; initially just pre-DNAT policy.  Please check Calico documentation for the latest position.
	InterfaceName *string `field:"optional" json:"interfaceName" yaml:"interfaceName"`
	// The node name identifying the Calico node instance.
	Node *string `field:"optional" json:"node" yaml:"node"`
	// Ports contains the endpoint's named ports, which may be referenced in security policy rules.
	Ports *[]*HostEndpointSpecPorts `field:"optional" json:"ports" yaml:"ports"`
	// A list of identifiers of security Profile objects that apply to this endpoint.
	//
	// Each profile is applied in the order that they appear in this list.  Profile rules are applied after the selector-based security policy.
	Profiles *[]*string `field:"optional" json:"profiles" yaml:"profiles"`
}

