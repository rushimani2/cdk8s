package apiextensionscrossplaneio


// A SecretRef is a reference to a secret containing credentials that should be supplied to the function.
type CompositionRevisionV1Beta1SpecPipelineCredentialsSecretRef struct {
	// Name of the secret.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Namespace of the secret.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

