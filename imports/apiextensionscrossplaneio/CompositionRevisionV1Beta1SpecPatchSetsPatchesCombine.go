package apiextensionscrossplaneio


// Combine is the patch configuration for a CombineFromComposite or CombineToComposite patch.
type CompositionRevisionV1Beta1SpecPatchSetsPatchesCombine struct {
	// Strategy defines the strategy to use to combine the input variable values.
	//
	// Currently only string is supported.
	Strategy CompositionRevisionV1Beta1SpecPatchSetsPatchesCombineStrategy `field:"required" json:"strategy" yaml:"strategy"`
	// Variables are the list of variables whose values will be retrieved and combined.
	Variables *[]*CompositionRevisionV1Beta1SpecPatchSetsPatchesCombineVariables `field:"required" json:"variables" yaml:"variables"`
	// String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.
	String *CompositionRevisionV1Beta1SpecPatchSetsPatchesCombineString `field:"optional" json:"string" yaml:"string"`
}

