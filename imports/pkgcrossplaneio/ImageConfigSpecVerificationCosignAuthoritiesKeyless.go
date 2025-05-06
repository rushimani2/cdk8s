package pkgcrossplaneio


// Keyless sets the configuration to verify the authority against a Fulcio instance.
type ImageConfigSpecVerificationCosignAuthoritiesKeyless struct {
	// Identities sets a list of identities.
	Identities *[]*ImageConfigSpecVerificationCosignAuthoritiesKeylessIdentities `field:"required" json:"identities" yaml:"identities"`
	// InsecureIgnoreSCT omits verifying if a certificate contains an embedded SCT.
	InsecureIgnoreSct *bool `field:"optional" json:"insecureIgnoreSct" yaml:"insecureIgnoreSct"`
}

