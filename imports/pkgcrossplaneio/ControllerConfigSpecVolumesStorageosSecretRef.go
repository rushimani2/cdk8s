package pkgcrossplaneio


// secretRef specifies the secret to use for obtaining the StorageOS API credentials.
//
// If not specified, default values will be attempted.
type ControllerConfigSpecVolumesStorageosSecretRef struct {
	// Name of the referent.
	//
	// This field is effectively required, but due to backwards compatibility is
	// allowed to be empty. Instances of this type with an empty value here are
	// almost certainly wrong.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	Name *string `field:"optional" json:"name" yaml:"name"`
}

