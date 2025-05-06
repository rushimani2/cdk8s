package apiextensionscrossplaneio


// FunctionRef is a reference to the Composition Function this step should execute.
type CompositionRevisionV1Beta1SpecPipelineFunctionRef struct {
	// Name of the referenced Function.
	Name *string `field:"required" json:"name" yaml:"name"`
}

