package pkgcrossplaneio


// SecretRef sets a reference to a secret with the key.
type ImageConfigSpecVerificationCosignAuthoritiesKeySecretRef struct {
	// The key to select.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the secret.
	Name *string `field:"required" json:"name" yaml:"name"`
}

