package apiextensionscrossplaneio


// Match is a more complex version of Map that matches a list of patterns.
type CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatch struct {
	// Determines to what value the transform should fallback if no pattern matches.
	FallbackTo CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchFallbackTo `field:"optional" json:"fallbackTo" yaml:"fallbackTo"`
	// The fallback value that should be returned by the transform if now pattern matches.
	FallbackValue interface{} `field:"optional" json:"fallbackValue" yaml:"fallbackValue"`
	// The patterns that should be tested against the input string.
	//
	// Patterns are tested in order. The value of the first match is used as
	// result of this transform.
	Patterns *[]*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchPatterns `field:"optional" json:"patterns" yaml:"patterns"`
}

