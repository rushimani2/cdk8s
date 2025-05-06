package apiextensionscrossplaneio


// Optional conversion method to be specified.
//
// `ToUpper` and `ToLower` change the letter case of the input string.
// `ToBase64` and `FromBase64` perform a base64 conversion based on the input string.
// `ToJson` converts any input value into its raw JSON representation.
// `ToSha1`, `ToSha256` and `ToSha512` generate a hash value based on the input
// converted to JSON.
// `ToAdler32` generate a addler32 hash based on the input string.
type CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert string

const (
	// ToUpper.
	CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_UPPER CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert = "TO_UPPER"
	// ToLower.
	CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_LOWER CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert = "TO_LOWER"
	// ToBase64.
	CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_BASE64 CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert = "TO_BASE64"
	// FromBase64.
	CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_FROM_BASE64 CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert = "FROM_BASE64"
	// ToJson.
	CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_JSON CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert = "TO_JSON"
	// ToSha1.
	CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_SHA1 CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert = "TO_SHA1"
	// ToSha256.
	CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_SHA256 CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert = "TO_SHA256"
	// ToSha512.
	CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_SHA512 CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert = "TO_SHA512"
	// ToAdler32.
	CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_ADLER32 CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert = "TO_ADLER32"
)

