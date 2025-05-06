package apiextensionscrossplaneio


// Combine is the patch configuration for a CombineFromComposite or CombineToComposite patch.
type CompositionRevisionV1Beta1SpecResourcesPatchesCombine struct {
	// Strategy defines the strategy to use to combine the input variable values.
	//
	// Currently only string is supported.
	Strategy CompositionRevisionV1Beta1SpecResourcesPatchesCombineStrategy `field:"required" json:"strategy" yaml:"strategy"`
	// Variables are the list of variables whose values will be retrieved and combined.
	Variables *[]*CompositionRevisionV1Beta1SpecResourcesPatchesCombineVariables `field:"required" json:"variables" yaml:"variables"`
	// String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.
	String *CompositionRevisionV1Beta1SpecResourcesPatchesCombineString `field:"optional" json:"string" yaml:"string"`
}

