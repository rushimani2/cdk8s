package crdprojectcalicoorg


// Selects a key of a secret in the node pod's namespace.
type BgpConfigurationSpecNodeMeshPasswordSecretKeyRef struct {
	// The key of the secret to select from.
	//
	// Must be a valid secret key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the Secret or its key must be defined.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

