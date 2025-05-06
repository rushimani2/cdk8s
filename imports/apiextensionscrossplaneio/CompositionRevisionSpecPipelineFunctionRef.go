package apiextensionscrossplaneio


// FunctionRef is a reference to the Composition Function this step should execute.
type CompositionRevisionSpecPipelineFunctionRef struct {
	// Name of the referenced Function.
	Name *string `field:"required" json:"name" yaml:"name"`
}

