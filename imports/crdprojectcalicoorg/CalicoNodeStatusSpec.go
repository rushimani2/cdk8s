package crdprojectcalicoorg


// CalicoNodeStatusSpec contains the specification for a CalicoNodeStatus resource.
type CalicoNodeStatusSpec struct {
	// Classes declares the types of information to monitor for this calico/node, and allows for selective status reporting about certain subsets of information.
	Classes *[]*string `field:"optional" json:"classes" yaml:"classes"`
	// The node name identifies the Calico node instance for node status.
	Node *string `field:"optional" json:"node" yaml:"node"`
	// UpdatePeriodSeconds is the period at which CalicoNodeStatus should be updated.
	//
	// Set to 0 to disable CalicoNodeStatus refresh. Maximum update period is one day.
	UpdatePeriodSeconds *float64 `field:"optional" json:"updatePeriodSeconds" yaml:"updatePeriodSeconds"`
}

