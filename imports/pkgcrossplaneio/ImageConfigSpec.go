package pkgcrossplaneio


// ImageConfigSpec contains the configuration for matching images.
type ImageConfigSpec struct {
	// MatchImages is a list of image matching rules that should be satisfied.
	MatchImages *[]*ImageConfigSpecMatchImages `field:"required" json:"matchImages" yaml:"matchImages"`
	// Registry is the configuration for the registry.
	Registry *ImageConfigSpecRegistry `field:"optional" json:"registry" yaml:"registry"`
	// Verification contains the configuration for verifying the image.
	Verification *ImageConfigSpecVerification `field:"optional" json:"verification" yaml:"verification"`
}

