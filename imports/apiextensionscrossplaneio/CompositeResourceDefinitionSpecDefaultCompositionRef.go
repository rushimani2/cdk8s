package apiextensionscrossplaneio


// DefaultCompositionRef refers to the Composition resource that will be used in case no composition selector is given.
type CompositeResourceDefinitionSpecDefaultCompositionRef struct {
	// Name of the Composition.
	Name *string `field:"required" json:"name" yaml:"name"`
}

