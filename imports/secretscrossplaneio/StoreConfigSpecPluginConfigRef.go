package secretscrossplaneio


// ConfigRef contains store config reference info.
type StoreConfigSpecPluginConfigRef struct {
	// APIVersion of the referenced config.
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Kind of the referenced config.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name of the referenced config.
	Name *string `field:"required" json:"name" yaml:"name"`
}

