package crdprojectcalicoorg


// Destination contains the match criteria that apply to destination entity.
type NetworkPolicySpecEgressDestination struct {
	// NamespaceSelector is an optional field that contains a selector expression.
	//
	// Only traffic that originates from (or terminates at) endpoints within the selected namespaces will be matched. When both NamespaceSelector and another selector are defined on the same rule, then only workload endpoints that are matched by both selectors will be selected by the rule.
	// For NetworkPolicy, an empty NamespaceSelector implies that the Selector is limited to selecting only workload endpoints in the same namespace as the NetworkPolicy.
	// For NetworkPolicy, `global()` NamespaceSelector implies that the Selector is limited to selecting only GlobalNetworkSet or HostEndpoint.
	// For GlobalNetworkPolicy, an empty NamespaceSelector implies the Selector applies to workload endpoints across all namespaces.
	NamespaceSelector *string `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
	// Nets is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) IP addresses in any of the given subnets.
	Nets *[]*string `field:"optional" json:"nets" yaml:"nets"`
	// NotNets is the negated version of the Nets field.
	NotNets *[]*string `field:"optional" json:"notNets" yaml:"notNets"`
	// NotPorts is the negated version of the Ports field.
	//
	// Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to "TCP" or "UDP".
	NotPorts *[]NetworkPolicySpecEgressDestinationNotPorts `field:"optional" json:"notPorts" yaml:"notPorts"`
	// NotSelector is the negated version of the Selector field.
	//
	// See Selector field for subtleties with negated selectors.
	NotSelector *string `field:"optional" json:"notSelector" yaml:"notSelector"`
	// Ports is an optional field that restricts the rule to only apply to traffic that has a source (destination) port that matches one of these ranges/values.
	//
	// This value is a list of integers or strings that represent ranges of ports.
	// Since only some protocols have ports, if any ports are specified it requires the Protocol match in the Rule to be set to "TCP" or "UDP".
	Ports *[]NetworkPolicySpecEgressDestinationPorts `field:"optional" json:"ports" yaml:"ports"`
	// Selector is an optional field that contains a selector expression (see Policy for sample syntax).
	//
	// Only traffic that originates from (terminates at) endpoints matching the selector will be matched.
	// Note that: in addition to the negated version of the Selector (see NotSelector below), the selector expression syntax itself supports negation.  The two types of negation are subtly different. One negates the set of matched endpoints, the other negates the whole match:
	// Selector = "!has(my_label)" matches packets that are from other Calico-controlled 	endpoints that do not have the label "my_label".
	// NotSelector = "has(my_label)" matches packets that are not from Calico-controlled 	endpoints that do have the label "my_label".
	// The effect is that the latter will accept packets from non-Calico sources whereas the former is limited to packets from Calico-controlled endpoints.
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
	// ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.
	ServiceAccounts *NetworkPolicySpecEgressDestinationServiceAccounts `field:"optional" json:"serviceAccounts" yaml:"serviceAccounts"`
	// Services is an optional field that contains options for matching Kubernetes Services.
	//
	// If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.
	// Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.
	// Ports and NotPorts can only be specified with Services on ingress rules.
	Services *NetworkPolicySpecEgressDestinationServices `field:"optional" json:"services" yaml:"services"`
}

