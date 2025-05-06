package pkgcrossplaneio


// Identity may contain the issuer and/or the subject found in the transparency log.
//
// Issuer/Subject uses a strict match, while IssuerRegExp and SubjectRegExp
// apply a regexp for matching.
type ImageConfigSpecVerificationCosignAuthoritiesKeylessIdentities struct {
	// Issuer defines the issuer for this identity.
	Issuer *string `field:"optional" json:"issuer" yaml:"issuer"`
	// IssuerRegExp specifies a regular expression to match the issuer for this identity.
	//
	// This has precedence over the Issuer field.
	IssuerRegExp *string `field:"optional" json:"issuerRegExp" yaml:"issuerRegExp"`
	// Subject defines the subject for this identity.
	Subject *string `field:"optional" json:"subject" yaml:"subject"`
	// SubjectRegExp specifies a regular expression to match the subject for this identity.
	//
	// This has precedence over the Subject field.
	SubjectRegExp *string `field:"optional" json:"subjectRegExp" yaml:"subjectRegExp"`
}

