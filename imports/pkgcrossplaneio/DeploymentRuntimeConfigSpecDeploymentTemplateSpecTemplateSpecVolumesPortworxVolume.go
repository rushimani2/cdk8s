package pkgcrossplaneio


// portworxVolume represents a portworx volume attached and mounted on kubelets host machine.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesPortworxVolume struct {
	// volumeID uniquely identifies a Portworx volume.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system.
	//
	// Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// readOnly defaults to false (read/write).
	//
	// ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

