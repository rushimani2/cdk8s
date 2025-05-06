package pkgcrossplaneio


// Provider is the provider that should be used to verify the image.
type ImageConfigSpecVerificationProvider string

const (
	// Cosign.
	ImageConfigSpecVerificationProvider_COSIGN ImageConfigSpecVerificationProvider = "COSIGN"
)

