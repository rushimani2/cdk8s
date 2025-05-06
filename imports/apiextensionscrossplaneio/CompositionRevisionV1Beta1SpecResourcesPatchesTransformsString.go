package apiextensionscrossplaneio


// String is used to transform the input into a string or a different kind of string.
//
// Note that the input does not necessarily need to be a string.
type CompositionRevisionV1Beta1SpecResourcesPatchesTransformsString struct {
	// Optional conversion method to be specified.
	//
	// `ToUpper` and `ToLower` change the letter case of the input string.
	// `ToBase64` and `FromBase64` perform a base64 conversion based on the input string.
	// `ToJson` converts any input value into its raw JSON representation.
	// `ToSha1`, `ToSha256` and `ToSha512` generate a hash value based on the input
	// converted to JSON.
	// `ToAdler32` generate a addler32 hash based on the input string.
	Convert CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert `field:"optional" json:"convert" yaml:"convert"`
	// Format the input using a Go format string.
	//
	// See
	// https://golang.org/pkg/fmt/ for details.
	Fmt *string `field:"optional" json:"fmt" yaml:"fmt"`
	// Join defines parameters to join a slice of values to a string.
	Join *CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringJoin `field:"optional" json:"join" yaml:"join"`
	// Extract a match from the input using a regular expression.
	Regexp *CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringRegexp `field:"optional" json:"regexp" yaml:"regexp"`
	// Trim the prefix or suffix from the input.
	Trim *string `field:"optional" json:"trim" yaml:"trim"`
	// Type of the string transform to be run.
	Type CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringType `field:"optional" json:"type" yaml:"type"`
}

