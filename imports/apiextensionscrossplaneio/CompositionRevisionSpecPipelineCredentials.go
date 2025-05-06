package apiextensionscrossplaneio


// FunctionCredentials are optional credentials that a Composition Function needs to run.
type CompositionRevisionSpecPipelineCredentials struct {
	// Name of this set of credentials.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Source of the function credentials.
	Source CompositionRevisionSpecPipelineCredentialsSource `field:"required" json:"source" yaml:"source"`
	// A SecretRef is a reference to a secret containing credentials that should be supplied to the function.
	SecretRef *CompositionRevisionSpecPipelineCredentialsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

