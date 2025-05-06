package apiextensionscrossplaneio


// Selector to the resource.
//
// This field will be ignored if ResourceRef is set.
type UsageV1Beta1SpecByResourceSelector struct {
	// MatchControllerRef ensures an object with the same controller reference as the selecting object is selected.
	MatchControllerRef *bool `field:"optional" json:"matchControllerRef" yaml:"matchControllerRef"`
	// MatchLabels ensures an object with matching labels is selected.
	MatchLabels *map[string]*string `field:"optional" json:"matchLabels" yaml:"matchLabels"`
}

