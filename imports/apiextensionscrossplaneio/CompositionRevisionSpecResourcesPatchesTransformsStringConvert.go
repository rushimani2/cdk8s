package apiextensionscrossplaneio


// Optional conversion method to be specified.
//
// `ToUpper` and `ToLower` change the letter case of the input string.
// `ToBase64` and `FromBase64` perform a base64 conversion based on the input string.
// `ToJson` converts any input value into its raw JSON representation.
// `ToSha1`, `ToSha256` and `ToSha512` generate a hash value based on the input
// converted to JSON.
// `ToAdler32` generate a addler32 hash based on the input string.
type CompositionRevisionSpecResourcesPatchesTransformsStringConvert string

const (
	// ToUpper.
	CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_UPPER CompositionRevisionSpecResourcesPatchesTransformsStringConvert = "TO_UPPER"
	// ToLower.
	CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_LOWER CompositionRevisionSpecResourcesPatchesTransformsStringConvert = "TO_LOWER"
	// ToBase64.
	CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_BASE64 CompositionRevisionSpecResourcesPatchesTransformsStringConvert = "TO_BASE64"
	// FromBase64.
	CompositionRevisionSpecResourcesPatchesTransformsStringConvert_FROM_BASE64 CompositionRevisionSpecResourcesPatchesTransformsStringConvert = "FROM_BASE64"
	// ToJson.
	CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_JSON CompositionRevisionSpecResourcesPatchesTransformsStringConvert = "TO_JSON"
	// ToSha1.
	CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_SHA1 CompositionRevisionSpecResourcesPatchesTransformsStringConvert = "TO_SHA1"
	// ToSha256.
	CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_SHA256 CompositionRevisionSpecResourcesPatchesTransformsStringConvert = "TO_SHA256"
	// ToSha512.
	CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_SHA512 CompositionRevisionSpecResourcesPatchesTransformsStringConvert = "TO_SHA512"
	// ToAdler32.
	CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_ADLER32 CompositionRevisionSpecResourcesPatchesTransformsStringConvert = "TO_ADLER32"
)

