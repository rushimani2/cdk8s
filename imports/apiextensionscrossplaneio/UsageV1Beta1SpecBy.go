package apiextensionscrossplaneio


// By is the resource that is "using the other resource".
type UsageV1Beta1SpecBy struct {
	// API version of the referent.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Kind of the referent.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Reference to the resource.
	ResourceRef *UsageV1Beta1SpecByResourceRef `field:"optional" json:"resourceRef" yaml:"resourceRef"`
	// Selector to the resource.
	//
	// This field will be ignored if ResourceRef is set.
	ResourceSelector *UsageV1Beta1SpecByResourceSelector `field:"optional" json:"resourceSelector" yaml:"resourceSelector"`
}

