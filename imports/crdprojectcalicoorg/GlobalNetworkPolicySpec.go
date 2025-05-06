package crdprojectcalicoorg


type GlobalNetworkPolicySpec struct {
	// ApplyOnForward indicates to apply the rules in this policy on forward traffic.
	ApplyOnForward *bool `field:"optional" json:"applyOnForward" yaml:"applyOnForward"`
	// DoNotTrack indicates whether packets matched by the rules in this policy should go through the data plane's connection tracking, such as Linux conntrack.
	//
	// If True, the rules in this policy are applied before any data plane connection tracking, and packets allowed by this policy are marked as not to be tracked.
	DoNotTrack *bool `field:"optional" json:"doNotTrack" yaml:"doNotTrack"`
	// The ordered set of egress rules.
	//
	// Each rule contains a set of packet match criteria and a corresponding action to apply.
	Egress *[]*GlobalNetworkPolicySpecEgress `field:"optional" json:"egress" yaml:"egress"`
	// The ordered set of ingress rules.
	//
	// Each rule contains a set of packet match criteria and a corresponding action to apply.
	Ingress *[]*GlobalNetworkPolicySpecIngress `field:"optional" json:"ingress" yaml:"ingress"`
	// NamespaceSelector is an optional field for an expression used to select a pod based on namespaces.
	NamespaceSelector *string `field:"optional" json:"namespaceSelector" yaml:"namespaceSelector"`
	// Order is an optional field that specifies the order in which the policy is applied.
	//
	// Policies with higher "order" are applied after those with lower order.  If the order is omitted, it may be considered to be "infinite" - i.e. the policy will be applied last.  Policies with identical order will be applied in alphanumerical order based on the Policy "Name".
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// PreDNAT indicates to apply the rules in this policy before any DNAT.
	PreDnat *bool `field:"optional" json:"preDnat" yaml:"preDnat"`
	// The selector is an expression used to pick pick out the endpoints that the policy should be applied to.
	//
	// Selector expressions follow this syntax:
	// label == "string_literal"  ->  comparison, e.g. my_label == "foo bar" 	label != "string_literal"   ->  not equal; also matches if label is not present 	label in { "a", "b", "c", ... }  ->  true if the value of label X is one of "a", "b", "c" 	label not in { "a", "b", "c", ... }  ->  true if the value of label X is not one of "a", "b", "c" 	has(label_name)  -> True if that label is present 	! expr -> negation of expr 	expr && expr  -> Short-circuit and 	expr || expr  -> Short-circuit or 	( expr ) -> parens for grouping 	all() or the empty selector -> matches all endpoints.
	// Label names are allowed to contain alphanumerics, -, _ and /. String literals are more permissive but they do not support escape characters.
	// Examples (with made-up labels):
	// type == "webserver" && deployment == "prod" 	type in {"frontend", "backend"} 	deployment != "dev" 	! has(label_name)
	Selector *string `field:"optional" json:"selector" yaml:"selector"`
	// ServiceAccountSelector is an optional field for an expression used to select a pod based on service accounts.
	ServiceAccountSelector *string `field:"optional" json:"serviceAccountSelector" yaml:"serviceAccountSelector"`
	// Types indicates whether this policy applies to ingress, or to egress, or to both.
	//
	// When not explicitly specified (and so the value on creation is empty or nil), Calico defaults Types according to what Ingress and Egress rules are present in the policy.  The default is:
	// - [ PolicyTypeIngress ], if there are no Egress rules (including the case where there are   also no Ingress rules)
	// - [ PolicyTypeEgress ], if there are Egress rules but no Ingress rules
	// - [ PolicyTypeIngress, PolicyTypeEgress ], if there are both Ingress and Egress rules.
	// When the policy is read back again, Types will always be one of these values, never empty or nil.
	Types *[]*string `field:"optional" json:"types" yaml:"types"`
}

