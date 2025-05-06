package apiextensionscrossplaneio


// Of is the resource that is "being used".
type UsageV1Beta1SpecOf struct {
	// API version of the referent.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Kind of the referent.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Reference to the resource.
	ResourceRef *UsageV1Beta1SpecOfResourceRef `field:"optional" json:"resourceRef" yaml:"resourceRef"`
	// Selector to the resource.
	//
	// This field will be ignored if ResourceRef is set.
	ResourceSelector *UsageV1Beta1SpecOfResourceSelector `field:"optional" json:"resourceSelector" yaml:"resourceSelector"`
}

