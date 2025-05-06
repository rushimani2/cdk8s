package crdprojectcalicoorg


// ServiceAccount enables and configures the service account controller.
//
// Enabled by default, set to nil to disable.
type KubeControllersConfigurationSpecControllersServiceAccount struct {
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore.
	//
	// [Default: 5m].
	ReconcilerPeriod *string `field:"optional" json:"reconcilerPeriod" yaml:"reconcilerPeriod"`
}

