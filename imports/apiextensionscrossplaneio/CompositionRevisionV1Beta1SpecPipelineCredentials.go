package apiextensionscrossplaneio


// FunctionCredentials are optional credentials that a Composition Function needs to run.
type CompositionRevisionV1Beta1SpecPipelineCredentials struct {
	// Name of this set of credentials.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Source of the function credentials.
	Source CompositionRevisionV1Beta1SpecPipelineCredentialsSource `field:"required" json:"source" yaml:"source"`
	// A SecretRef is a reference to a secret containing credentials that should be supplied to the function.
	SecretRef *CompositionRevisionV1Beta1SpecPipelineCredentialsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

