package pkgcrossplaneio


// Registry is the configuration for the registry.
type ImageConfigSpecRegistry struct {
	// Authentication is the authentication information for the registry.
	Authentication *ImageConfigSpecRegistryAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
}

