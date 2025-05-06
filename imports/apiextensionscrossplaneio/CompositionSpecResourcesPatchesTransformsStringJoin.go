package apiextensionscrossplaneio


// Join defines parameters to join a slice of values to a string.
type CompositionSpecResourcesPatchesTransformsStringJoin struct {
	// Separator defines the character that should separate the values from each other in the joined string.
	Separator *string `field:"required" json:"separator" yaml:"separator"`
}

