package apiextensionscrossplaneio


// ReadinessCheck is used to indicate how to tell whether a resource is ready for consumption.
type CompositionSpecResourcesReadinessChecks struct {
	// Type indicates the type of probe you'd like to use.
	Type CompositionSpecResourcesReadinessChecksType `field:"required" json:"type" yaml:"type"`
	// FieldPath shows the path of the field whose value will be used.
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
	// MatchCondition specifies the condition you'd like to match if you're using "MatchCondition" type.
	MatchCondition *CompositionSpecResourcesReadinessChecksMatchCondition `field:"optional" json:"matchCondition" yaml:"matchCondition"`
	// MatchInt is the value you'd like to match if you're using "MatchInt" type.
	MatchInteger *float64 `field:"optional" json:"matchInteger" yaml:"matchInteger"`
	// MatchString is the value you'd like to match if you're using "MatchString" type.
	MatchString *string `field:"optional" json:"matchString" yaml:"matchString"`
}

