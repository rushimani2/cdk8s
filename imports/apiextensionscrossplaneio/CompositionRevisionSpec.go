package apiextensionscrossplaneio


// CompositionRevisionSpec specifies the desired state of the composition revision.
type CompositionRevisionSpec struct {
	// CompositeTypeRef specifies the type of composite resource that this composition is compatible with.
	CompositeTypeRef *CompositionRevisionSpecCompositeTypeRef `field:"required" json:"compositeTypeRef" yaml:"compositeTypeRef"`
	// Revision number. Newer revisions have larger numbers.
	//
	// This number can change. When a Composition transitions from state A
	// -> B -> A there will be only two CompositionRevisions. Crossplane will
	// edit the original CompositionRevision to change its revision number from
	// 0 to 2.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
	// Mode controls what type or "mode" of Composition will be used.
	//
	// "Pipeline" indicates that a Composition specifies a pipeline of
	// Composition Functions, each of which is responsible for producing
	// composed resources that Crossplane should create or update.
	//
	// "Resources" indicates that a Composition uses what is commonly referred
	// to as "Patch & Transform" or P&T composition. This mode of Composition
	// uses an array of resources, each a template for a composed resource.
	//
	// All Compositions should use Pipeline mode. Resources mode is deprecated.
	// Resources mode won't be removed in Crossplane 1.x, and will remain the
	// default to avoid breaking legacy Compositions. However, it's no longer
	// accepting new features, and only accepting security related bug fixes.
	Mode CompositionRevisionSpecMode `field:"optional" json:"mode" yaml:"mode"`
	// PatchSets define a named set of patches that may be included by any resource in this Composition.
	//
	// PatchSets cannot themselves refer to other
	// PatchSets.
	//
	// PatchSets are only used by the "Resources" mode of Composition. They
	// are ignored by other modes.
	//
	// Deprecated: Use Composition Functions instead.
	PatchSets *[]*CompositionRevisionSpecPatchSets `field:"optional" json:"patchSets" yaml:"patchSets"`
	// Pipeline is a list of composition function steps that will be used when a composite resource referring to this composition is created.
	//
	// One of
	// resources and pipeline must be specified - you cannot specify both.
	//
	// The Pipeline is only used by the "Pipeline" mode of Composition. It is
	// ignored by other modes.
	Pipeline *[]*CompositionRevisionSpecPipeline `field:"optional" json:"pipeline" yaml:"pipeline"`
	// PublishConnectionDetailsWithStoreConfig specifies the secret store config with which the connection details of composite resources dynamically provisioned using this composition will be published.
	//
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	PublishConnectionDetailsWithStoreConfigRef *CompositionRevisionSpecPublishConnectionDetailsWithStoreConfigRef `field:"optional" json:"publishConnectionDetailsWithStoreConfigRef" yaml:"publishConnectionDetailsWithStoreConfigRef"`
	// Resources is a list of resource templates that will be used when a composite resource referring to this composition is created.
	//
	// Resources are only used by the "Resources" mode of Composition. They are
	// ignored by other modes.
	//
	// Deprecated: Use Composition Functions instead.
	Resources *[]*CompositionRevisionSpecResources `field:"optional" json:"resources" yaml:"resources"`
	// WriteConnectionSecretsToNamespace specifies the namespace in which the connection secrets of composite resource dynamically provisioned using this composition will be created.
	//
	// This field is planned to be replaced in a future release in favor of
	// PublishConnectionDetailsWithStoreConfigRef. Currently, both could be
	// set independently and connection details would be published to both
	// without affecting each other as long as related fields at MR level
	// specified.
	WriteConnectionSecretsToNamespace *string `field:"optional" json:"writeConnectionSecretsToNamespace" yaml:"writeConnectionSecretsToNamespace"`
}

