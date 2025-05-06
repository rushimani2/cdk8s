package pkgcrossplaneio


// Metadata contains the configurable metadata fields for the Service.
type DeploymentRuntimeConfigSpecServiceTemplateMetadata struct {
	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects.
	// More info: http:https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// Labels will be merged with internal labels
	// used by crossplane, and labels with a crossplane.io key might be
	// overwritten.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Name is the name of the object.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

