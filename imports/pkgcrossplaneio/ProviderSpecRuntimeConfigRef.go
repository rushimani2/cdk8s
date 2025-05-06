package pkgcrossplaneio


// RuntimeConfigRef references a RuntimeConfig resource that will be used to configure the package runtime.
type ProviderSpecRuntimeConfigRef struct {
	// Name of the RuntimeConfig.
	Name *string `field:"required" json:"name" yaml:"name"`
	// API version of the referent.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Kind of the referent.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

