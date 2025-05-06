package crdprojectcalicoorg


// Node enables and configures the node controller.
//
// Enabled by default, set to nil to disable.
type KubeControllersConfigurationSpecControllersNode struct {
	// HostEndpoint controls syncing nodes to host endpoints.
	//
	// Disabled by default, set to nil to disable.
	HostEndpoint *KubeControllersConfigurationSpecControllersNodeHostEndpoint `field:"optional" json:"hostEndpoint" yaml:"hostEndpoint"`
	// LeakGracePeriod is the period used by the controller to determine if an IP address has been leaked.
	//
	// Set to 0 to disable IP garbage collection. [Default: 15m]
	LeakGracePeriod *string `field:"optional" json:"leakGracePeriod" yaml:"leakGracePeriod"`
	// ReconcilerPeriod is the period to perform reconciliation with the Calico datastore.
	//
	// [Default: 5m].
	ReconcilerPeriod *string `field:"optional" json:"reconcilerPeriod" yaml:"reconcilerPeriod"`
	// SyncLabels controls whether to copy Kubernetes node labels to Calico nodes.
	//
	// [Default: Enabled].
	SyncLabels *string `field:"optional" json:"syncLabels" yaml:"syncLabels"`
}

