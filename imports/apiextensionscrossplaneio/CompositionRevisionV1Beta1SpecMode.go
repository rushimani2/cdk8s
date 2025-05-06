package apiextensionscrossplaneio


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
type CompositionRevisionV1Beta1SpecMode string

const (
	// Resources.
	CompositionRevisionV1Beta1SpecMode_RESOURCES CompositionRevisionV1Beta1SpecMode = "RESOURCES"
	// Pipeline.
	CompositionRevisionV1Beta1SpecMode_PIPELINE CompositionRevisionV1Beta1SpecMode = "PIPELINE"
)

