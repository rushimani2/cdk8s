package apiextensionscrossplaneio


// FromFieldPath specifies how to patch from a field path.
//
// The default is
// 'Optional', which means the patch will be a no-op if the specified
// fromFieldPath does not exist. Use 'Required' if the patch should fail if
// the specified path does not exist.
type CompositionRevisionV1Beta1SpecResourcesPatchesPolicyFromFieldPath string

const (
	// Optional.
	CompositionRevisionV1Beta1SpecResourcesPatchesPolicyFromFieldPath_OPTIONAL CompositionRevisionV1Beta1SpecResourcesPatchesPolicyFromFieldPath = "OPTIONAL"
	// Required.
	CompositionRevisionV1Beta1SpecResourcesPatchesPolicyFromFieldPath_REQUIRED CompositionRevisionV1Beta1SpecResourcesPatchesPolicyFromFieldPath = "REQUIRED"
)

