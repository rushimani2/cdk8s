package apiextensionscrossplaneio


// Type specifies how the pattern matches the input.
//
// * `literal` - the pattern value has to exactly match (case sensitive) the
// input string. This is the default.
//
// * `regexp` - the pattern treated as a regular expression against
// which the input string is tested. Crossplane will throw an error if the
// key is not a valid regexp.
type CompositionSpecResourcesPatchesTransformsMatchPatternsType string

const (
	// literal.
	CompositionSpecResourcesPatchesTransformsMatchPatternsType_LITERAL CompositionSpecResourcesPatchesTransformsMatchPatternsType = "LITERAL"
	// regexp.
	CompositionSpecResourcesPatchesTransformsMatchPatternsType_REGEXP CompositionSpecResourcesPatchesTransformsMatchPatternsType = "REGEXP"
)

