package apiextensionscrossplaneio


// DefaultCompositeDeletePolicy is the policy used when deleting the Composite that is associated with the Claim if no policy has been specified.
type CompositeResourceDefinitionSpecDefaultCompositeDeletePolicy string

const (
	// Background.
	CompositeResourceDefinitionSpecDefaultCompositeDeletePolicy_BACKGROUND CompositeResourceDefinitionSpecDefaultCompositeDeletePolicy = "BACKGROUND"
	// Foreground.
	CompositeResourceDefinitionSpecDefaultCompositeDeletePolicy_FOREGROUND CompositeResourceDefinitionSpecDefaultCompositeDeletePolicy = "FOREGROUND"
)

