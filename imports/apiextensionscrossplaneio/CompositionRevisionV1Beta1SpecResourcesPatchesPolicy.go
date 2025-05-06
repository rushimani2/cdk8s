package apiextensionscrossplaneio


// Policy configures the specifics of patching behaviour.
type CompositionRevisionV1Beta1SpecResourcesPatchesPolicy struct {
	// FromFieldPath specifies how to patch from a field path.
	//
	// The default is
	// 'Optional', which means the patch will be a no-op if the specified
	// fromFieldPath does not exist. Use 'Required' if the patch should fail if
	// the specified path does not exist.
	FromFieldPath CompositionRevisionV1Beta1SpecResourcesPatchesPolicyFromFieldPath `field:"optional" json:"fromFieldPath" yaml:"fromFieldPath"`
	// MergeOptions Specifies merge options on a field path.
	MergeOptions *CompositionRevisionV1Beta1SpecResourcesPatchesPolicyMergeOptions `field:"optional" json:"mergeOptions" yaml:"mergeOptions"`
}

