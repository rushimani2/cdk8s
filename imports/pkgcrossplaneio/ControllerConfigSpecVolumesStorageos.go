package pkgcrossplaneio


// storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.
type ControllerConfigSpecVolumesStorageos struct {
	// fsType is the filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// readOnly defaults to false (read/write).
	//
	// ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secretRef specifies the secret to use for obtaining the StorageOS API credentials.
	//
	// If not specified, default values will be attempted.
	SecretRef *ControllerConfigSpecVolumesStorageosSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	// volumeName is the human-readable name of the StorageOS volume.
	//
	// Volume
	// names are only unique within a namespace.
	VolumeName *string `field:"optional" json:"volumeName" yaml:"volumeName"`
	// volumeNamespace specifies the scope of the volume within StorageOS.
	//
	// If no
	// namespace is specified then the Pod's namespace will be used.  This allows the
	// Kubernetes name scoping to be mirrored within StorageOS for tighter integration.
	// Set VolumeName to any name to override the default behaviour.
	// Set to "default" if you are not using namespaces within StorageOS.
	// Namespaces that do not pre-exist within StorageOS will be created.
	VolumeNamespace *string `field:"optional" json:"volumeNamespace" yaml:"volumeNamespace"`
}

