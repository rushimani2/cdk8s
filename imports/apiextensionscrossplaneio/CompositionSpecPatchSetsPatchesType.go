package apiextensionscrossplaneio


// Type sets the patching behaviour to be used.
//
// Each patch type may require
// its own fields to be set on the Patch object.
type CompositionSpecPatchSetsPatchesType string

const (
	// FromCompositeFieldPath.
	CompositionSpecPatchSetsPatchesType_FROM_COMPOSITE_FIELD_PATH CompositionSpecPatchSetsPatchesType = "FROM_COMPOSITE_FIELD_PATH"
	// PatchSet.
	CompositionSpecPatchSetsPatchesType_PATCH_SET CompositionSpecPatchSetsPatchesType = "PATCH_SET"
	// ToCompositeFieldPath.
	CompositionSpecPatchSetsPatchesType_TO_COMPOSITE_FIELD_PATH CompositionSpecPatchSetsPatchesType = "TO_COMPOSITE_FIELD_PATH"
	// CombineFromComposite.
	CompositionSpecPatchSetsPatchesType_COMBINE_FROM_COMPOSITE CompositionSpecPatchSetsPatchesType = "COMBINE_FROM_COMPOSITE"
	// CombineToComposite.
	CompositionSpecPatchSetsPatchesType_COMBINE_TO_COMPOSITE CompositionSpecPatchSetsPatchesType = "COMBINE_TO_COMPOSITE"
)

