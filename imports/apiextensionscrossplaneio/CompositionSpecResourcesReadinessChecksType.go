package apiextensionscrossplaneio


// Type indicates the type of probe you'd like to use.
type CompositionSpecResourcesReadinessChecksType string

const (
	// MatchString.
	CompositionSpecResourcesReadinessChecksType_MATCH_STRING CompositionSpecResourcesReadinessChecksType = "MATCH_STRING"
	// MatchInteger.
	CompositionSpecResourcesReadinessChecksType_MATCH_INTEGER CompositionSpecResourcesReadinessChecksType = "MATCH_INTEGER"
	// NonEmpty.
	CompositionSpecResourcesReadinessChecksType_NON_EMPTY CompositionSpecResourcesReadinessChecksType = "NON_EMPTY"
	// MatchCondition.
	CompositionSpecResourcesReadinessChecksType_MATCH_CONDITION CompositionSpecResourcesReadinessChecksType = "MATCH_CONDITION"
	// MatchTrue.
	CompositionSpecResourcesReadinessChecksType_MATCH_TRUE CompositionSpecResourcesReadinessChecksType = "MATCH_TRUE"
	// MatchFalse.
	CompositionSpecResourcesReadinessChecksType_MATCH_FALSE CompositionSpecResourcesReadinessChecksType = "MATCH_FALSE"
	// None.
	CompositionSpecResourcesReadinessChecksType_NONE CompositionSpecResourcesReadinessChecksType = "NONE"
)

