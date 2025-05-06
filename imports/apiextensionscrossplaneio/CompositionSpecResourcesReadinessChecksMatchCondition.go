package apiextensionscrossplaneio


// MatchCondition specifies the condition you'd like to match if you're using "MatchCondition" type.
type CompositionSpecResourcesReadinessChecksMatchCondition struct {
	// Status is the status of the condition you'd like to match.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Type indicates the type of condition you'd like to use.
	Type *string `field:"required" json:"type" yaml:"type"`
}

