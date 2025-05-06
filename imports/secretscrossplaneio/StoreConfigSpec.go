package secretscrossplaneio


// A StoreConfigSpec defines the desired state of a StoreConfig.
type StoreConfigSpec struct {
	// DefaultScope used for scoping secrets for "cluster-scoped" resources.
	//
	// If store type is "Kubernetes", this would mean the default namespace to
	// store connection secrets for cluster scoped resources.
	// In case of "Vault", this would be used as the default parent path.
	// Typically, should be set as Crossplane installation namespace.
	DefaultScope *string `field:"required" json:"defaultScope" yaml:"defaultScope"`
	// Kubernetes configures a Kubernetes secret store.
	//
	// If the "type" is "Kubernetes" but no config provided, in cluster config
	// will be used.
	Kubernetes *StoreConfigSpecKubernetes `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// Plugin configures External secret store as a plugin.
	Plugin *StoreConfigSpecPlugin `field:"optional" json:"plugin" yaml:"plugin"`
	// Type configures which secret store to be used.
	//
	// Only the configuration
	// block for this store will be used and others will be ignored if provided.
	// Default is Kubernetes.
	// Default: Kubernetes.
	//
	Type StoreConfigSpecType `field:"optional" json:"type" yaml:"type"`
}

