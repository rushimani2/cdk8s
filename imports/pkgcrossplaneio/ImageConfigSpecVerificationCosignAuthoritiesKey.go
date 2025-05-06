package pkgcrossplaneio


// Key defines the type of key to validate the image.
type ImageConfigSpecVerificationCosignAuthoritiesKey struct {
	// HashAlgorithm always defaults to sha256 if the algorithm hasn't been explicitly set.
	HashAlgorithm *string `field:"required" json:"hashAlgorithm" yaml:"hashAlgorithm"`
	// SecretRef sets a reference to a secret with the key.
	SecretRef *ImageConfigSpecVerificationCosignAuthoritiesKeySecretRef `field:"required" json:"secretRef" yaml:"secretRef"`
}

