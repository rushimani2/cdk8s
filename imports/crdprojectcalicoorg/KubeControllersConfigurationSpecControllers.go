package crdprojectcalicoorg


// Controllers enables and configures individual Kubernetes controllers.
type KubeControllersConfigurationSpecControllers struct {
	// Namespace enables and configures the namespace controller.
	//
	// Enabled by default, set to nil to disable.
	Namespace *KubeControllersConfigurationSpecControllersNamespace `field:"optional" json:"namespace" yaml:"namespace"`
	// Node enables and configures the node controller.
	//
	// Enabled by default, set to nil to disable.
	Node *KubeControllersConfigurationSpecControllersNode `field:"optional" json:"node" yaml:"node"`
	// Policy enables and configures the policy controller.
	//
	// Enabled by default, set to nil to disable.
	Policy *KubeControllersConfigurationSpecControllersPolicy `field:"optional" json:"policy" yaml:"policy"`
	// ServiceAccount enables and configures the service account controller.
	//
	// Enabled by default, set to nil to disable.
	ServiceAccount *KubeControllersConfigurationSpecControllersServiceAccount `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// WorkloadEndpoint enables and configures the workload endpoint controller.
	//
	// Enabled by default, set to nil to disable.
	WorkloadEndpoint *KubeControllersConfigurationSpecControllersWorkloadEndpoint `field:"optional" json:"workloadEndpoint" yaml:"workloadEndpoint"`
}

