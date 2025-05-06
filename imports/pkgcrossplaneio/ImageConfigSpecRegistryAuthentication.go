package pkgcrossplaneio


// Authentication is the authentication information for the registry.
type ImageConfigSpecRegistryAuthentication struct {
	// PullSecretRef is a reference to a secret that contains the credentials for the registry.
	PullSecretRef *ImageConfigSpecRegistryAuthenticationPullSecretRef `field:"required" json:"pullSecretRef" yaml:"pullSecretRef"`
}

