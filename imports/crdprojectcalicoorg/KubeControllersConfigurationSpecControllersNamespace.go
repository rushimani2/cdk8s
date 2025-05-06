package crdprojectcalicoorg


// Namespace enables and configures the namespace controller.
//
// Enabled by default, set to nil to disable.
type KubeControllersConfigurationSpecControllersNamespace struct {
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore.
	//
	// [Default: 5m].
	ReconcilerPeriod *string `field:"optional" json:"reconcilerPeriod" yaml:"reconcilerPeriod"`
}

