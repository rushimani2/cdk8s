package crdprojectcalicoorg


// Services is an optional field that contains options for matching Kubernetes Services.
//
// If specified, only traffic that originates from or terminates at endpoints within the selected service(s) will be matched, and only to/from each endpoint's port.
// Services cannot be specified on the same rule as Selector, NotSelector, NamespaceSelector, Nets, NotNets or ServiceAccounts.
// Ports and NotPorts can only be specified with Services on ingress rules.
type GlobalNetworkPolicySpecEgressDestinationServices struct {
	// Name specifies the name of a Kubernetes Service to match.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Namespace specifies the namespace of the given Service.
	//
	// If left empty, the rule will match within this policy's namespace.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

