package pkgcrossplaneio


// CosignAuthority defines the rules for discovering and validating signatures.
type ImageConfigSpecVerificationCosignAuthorities struct {
	// Name is the name for this authority.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Attestations is a list of individual attestations for this authority, once the signature for this authority has been verified.
	Attestations *[]*ImageConfigSpecVerificationCosignAuthoritiesAttestations `field:"optional" json:"attestations" yaml:"attestations"`
	// Key defines the type of key to validate the image.
	Key *ImageConfigSpecVerificationCosignAuthoritiesKey `field:"optional" json:"key" yaml:"key"`
	// Keyless sets the configuration to verify the authority against a Fulcio instance.
	Keyless *ImageConfigSpecVerificationCosignAuthoritiesKeyless `field:"optional" json:"keyless" yaml:"keyless"`
}

