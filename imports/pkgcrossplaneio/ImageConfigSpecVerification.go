package pkgcrossplaneio


// Verification contains the configuration for verifying the image.
type ImageConfigSpecVerification struct {
	// Provider is the provider that should be used to verify the image.
	Provider ImageConfigSpecVerificationProvider `field:"required" json:"provider" yaml:"provider"`
	// Cosign is the configuration for verifying the image using cosign.
	Cosign *ImageConfigSpecVerificationCosign `field:"optional" json:"cosign" yaml:"cosign"`
}

