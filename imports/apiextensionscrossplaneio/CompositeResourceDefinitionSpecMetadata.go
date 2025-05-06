package apiextensionscrossplaneio


// Metadata specifies the desired metadata for the defined composite resource and claim CRD's.
type CompositeResourceDefinitionSpecMetadata struct {
	// Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata.
	//
	// They are not
	// queryable and should be preserved when modifying objects.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels
	// and services.
	// These labels are added to the composite resource and claim CRD's in addition
	// to any labels defined by `CompositionResourceDefinition` `metadata.labels`.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

