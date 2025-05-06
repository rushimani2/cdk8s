package pkgcrossplaneio


// cinder represents a cinder volume attached and mounted on kubelets host machine.
//
// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
type ControllerConfigSpecVolumesCinder struct {
	// volumeID used to identify the volume in cinder.
	//
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// fsType is the filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// readOnly defaults to false (read/write).
	//
	// ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.
	SecretRef *ControllerConfigSpecVolumesCinderSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

