package apiextensionscrossplaneio


// Type sets the patching behaviour to be used.
//
// Each patch type may require
// its own fields to be set on the Patch object.
type CompositionRevisionV1Beta1SpecPatchSetsPatchesType string

const (
	// FromCompositeFieldPath.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesType_FROM_COMPOSITE_FIELD_PATH CompositionRevisionV1Beta1SpecPatchSetsPatchesType = "FROM_COMPOSITE_FIELD_PATH"
	// PatchSet.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesType_PATCH_SET CompositionRevisionV1Beta1SpecPatchSetsPatchesType = "PATCH_SET"
	// ToCompositeFieldPath.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesType_TO_COMPOSITE_FIELD_PATH CompositionRevisionV1Beta1SpecPatchSetsPatchesType = "TO_COMPOSITE_FIELD_PATH"
	// CombineFromComposite.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesType_COMBINE_FROM_COMPOSITE CompositionRevisionV1Beta1SpecPatchSetsPatchesType = "COMBINE_FROM_COMPOSITE"
	// CombineToComposite.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesType_COMBINE_TO_COMPOSITE CompositionRevisionV1Beta1SpecPatchSetsPatchesType = "COMBINE_TO_COMPOSITE"
)

