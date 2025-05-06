package pkgcrossplaneio


// Attestation defines the type of attestation to validate and optionally apply a policy decision to it.
//
// Authority block is used to verify the
// specified attestation types, and if Policy is specified, then it's applied
// only after the validation of the Attestation signature has been verified.
type ImageConfigSpecVerificationCosignAuthoritiesAttestations struct {
	// Name of the attestation.
	Name *string `field:"required" json:"name" yaml:"name"`
	// PredicateType defines which predicate type to verify.
	//
	// Matches cosign
	// verify-attestation options.
	PredicateType *string `field:"required" json:"predicateType" yaml:"predicateType"`
}

