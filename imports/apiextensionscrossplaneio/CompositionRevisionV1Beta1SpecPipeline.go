package apiextensionscrossplaneio


// A PipelineStep in a Composition Function pipeline.
type CompositionRevisionV1Beta1SpecPipeline struct {
	// FunctionRef is a reference to the Composition Function this step should execute.
	FunctionRef *CompositionRevisionV1Beta1SpecPipelineFunctionRef `field:"required" json:"functionRef" yaml:"functionRef"`
	// Step name.
	//
	// Must be unique within its Pipeline.
	Step *string `field:"required" json:"step" yaml:"step"`
	// Credentials are optional credentials that the Composition Function needs.
	Credentials *[]*CompositionRevisionV1Beta1SpecPipelineCredentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Input is an optional, arbitrary Kubernetes resource (i.e. a resource with an apiVersion and kind) that will be passed to the Composition Function as the 'input' of its RunFunctionRequest.
	Input interface{} `field:"optional" json:"input" yaml:"input"`
}

