package pkgcrossplaneio


// azureFile represents an Azure File Service mount on the host and bind mount to the pod.
type ControllerConfigSpecVolumesAzureFile struct {
	// secretName is the  name of secret that contains Azure Storage Account Name and Key.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// shareName is the azure share Name.
	ShareName *string `field:"required" json:"shareName" yaml:"shareName"`
	// readOnly defaults to false (read/write).
	//
	// ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

