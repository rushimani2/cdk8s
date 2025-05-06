package pkgcrossplaneio


// Cosign is the configuration for verifying the image using cosign.
type ImageConfigSpecVerificationCosign struct {
	// Authorities defines the rules for discovering and validating signatures.
	Authorities *[]*ImageConfigSpecVerificationCosignAuthorities `field:"required" json:"authorities" yaml:"authorities"`
}

