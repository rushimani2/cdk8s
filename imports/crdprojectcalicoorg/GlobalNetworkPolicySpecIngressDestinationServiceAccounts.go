package crdprojectcalicoorg


// ServiceAccounts is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a matching service account.
type GlobalNetworkPolicySpecIngressDestinationServiceAccounts struct {
	// Names is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account whose name is in the list.
	Names *[]*string `field:"optional" json:"names" yaml:"names"`
	// Selector is an optional field that restricts the rule to only apply to traffic that originates from (or terminates at) a pod running as a service account that matches the given label selector.
	//
	// If both Names and Selector are specified then they are AND'ed.
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
}

