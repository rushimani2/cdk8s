package crdprojectcalicoorg


// WorkloadEndpoint enables and configures the workload endpoint controller.
//
// Enabled by default, set to nil to disable.
type KubeControllersConfigurationSpecControllersWorkloadEndpoint struct {
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore.
	//
	// [Default: 5m].
	ReconcilerPeriod *string `field:"optional" json:"reconcilerPeriod" yaml:"reconcilerPeriod"`
}

