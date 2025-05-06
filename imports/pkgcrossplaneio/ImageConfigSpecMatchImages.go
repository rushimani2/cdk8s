package pkgcrossplaneio


// ImageMatch defines a rule for matching image.
type ImageConfigSpecMatchImages struct {
	// Prefix is the prefix that should be matched.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Type is the type of match.
	Type ImageConfigSpecMatchImagesType `field:"optional" json:"type" yaml:"type"`
}

