package apiextensionscrossplaneio


// Math is used to transform the input via mathematical operations such as multiplication.
type CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMath struct {
	// ClampMax makes sure that the value is not bigger than the given value.
	ClampMax *float64 `field:"optional" json:"clampMax" yaml:"clampMax"`
	// ClampMin makes sure that the value is not smaller than the given value.
	ClampMin *float64 `field:"optional" json:"clampMin" yaml:"clampMin"`
	// Multiply the value.
	Multiply *float64 `field:"optional" json:"multiply" yaml:"multiply"`
	// Type of the math transform to be run.
	Type CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMathType `field:"optional" json:"type" yaml:"type"`
}

