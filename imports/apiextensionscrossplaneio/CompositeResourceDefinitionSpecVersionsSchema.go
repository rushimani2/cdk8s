package apiextensionscrossplaneio


// Schema describes the schema used for validation, pruning, and defaulting of this version of the defined composite resource.
//
// Fields required by all
// composite resources will be injected into this schema automatically, and
// will override equivalently named fields in this schema. Omitting this
// schema results in a schema that contains only the fields required by all
// composite resources.
type CompositeResourceDefinitionSpecVersionsSchema struct {
	// OpenAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.
	OpenApiv3Schema interface{} `field:"optional" json:"openApiv3Schema" yaml:"openApiv3Schema"`
}

