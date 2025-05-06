package crdprojectcalicoorg


// Policy enables and configures the policy controller.
//
// Enabled by default, set to nil to disable.
type KubeControllersConfigurationSpecControllersPolicy struct {
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore.
	//
	// [Default: 5m].
	ReconcilerPeriod *string `field:"optional" json:"reconcilerPeriod" yaml:"reconcilerPeriod"`
}

