package apiextensionscrossplaneio


// DefaultCompositionUpdatePolicy is the policy used when updating composites after a new Composition Revision has been created if no policy has been specified on the composite.
type CompositeResourceDefinitionSpecDefaultCompositionUpdatePolicy string

const (
	// Automatic.
	CompositeResourceDefinitionSpecDefaultCompositionUpdatePolicy_AUTOMATIC CompositeResourceDefinitionSpecDefaultCompositionUpdatePolicy = "AUTOMATIC"
	// Manual.
	CompositeResourceDefinitionSpecDefaultCompositionUpdatePolicy_MANUAL CompositeResourceDefinitionSpecDefaultCompositionUpdatePolicy = "MANUAL"
)

