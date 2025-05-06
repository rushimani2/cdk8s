package apiextensionscrossplaneio


// Type sets the patching behaviour to be used.
//
// Each patch type may require
// its own fields to be set on the Patch object.
type CompositionRevisionV1Beta1SpecResourcesPatchesType string

const (
	// FromCompositeFieldPath.
	CompositionRevisionV1Beta1SpecResourcesPatchesType_FROM_COMPOSITE_FIELD_PATH CompositionRevisionV1Beta1SpecResourcesPatchesType = "FROM_COMPOSITE_FIELD_PATH"
	// PatchSet.
	CompositionRevisionV1Beta1SpecResourcesPatchesType_PATCH_SET CompositionRevisionV1Beta1SpecResourcesPatchesType = "PATCH_SET"
	// ToCompositeFieldPath.
	CompositionRevisionV1Beta1SpecResourcesPatchesType_TO_COMPOSITE_FIELD_PATH CompositionRevisionV1Beta1SpecResourcesPatchesType = "TO_COMPOSITE_FIELD_PATH"
	// CombineFromComposite.
	CompositionRevisionV1Beta1SpecResourcesPatchesType_COMBINE_FROM_COMPOSITE CompositionRevisionV1Beta1SpecResourcesPatchesType = "COMBINE_FROM_COMPOSITE"
	// CombineToComposite.
	CompositionRevisionV1Beta1SpecResourcesPatchesType_COMBINE_TO_COMPOSITE CompositionRevisionV1Beta1SpecResourcesPatchesType = "COMBINE_TO_COMPOSITE"
)

