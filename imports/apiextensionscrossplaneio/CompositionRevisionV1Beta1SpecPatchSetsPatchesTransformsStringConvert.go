package apiextensionscrossplaneio


// Optional conversion method to be specified.
//
// `ToUpper` and `ToLower` change the letter case of the input string.
// `ToBase64` and `FromBase64` perform a base64 conversion based on the input string.
// `ToJson` converts any input value into its raw JSON representation.
// `ToSha1`, `ToSha256` and `ToSha512` generate a hash value based on the input
// converted to JSON.
// `ToAdler32` generate a addler32 hash based on the input string.
type CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert string

const (
	// ToUpper.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_UPPER CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert = "TO_UPPER"
	// ToLower.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_LOWER CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert = "TO_LOWER"
	// ToBase64.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_BASE64 CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert = "TO_BASE64"
	// FromBase64.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_FROM_BASE64 CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert = "FROM_BASE64"
	// ToJson.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_JSON CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert = "TO_JSON"
	// ToSha1.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_SHA1 CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert = "TO_SHA1"
	// ToSha256.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_SHA256 CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert = "TO_SHA256"
	// ToSha512.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_SHA512 CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert = "TO_SHA512"
	// ToAdler32.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_ADLER32 CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert = "TO_ADLER32"
)

