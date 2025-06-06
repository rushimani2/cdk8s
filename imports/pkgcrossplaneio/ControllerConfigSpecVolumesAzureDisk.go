package pkgcrossplaneio


// azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
type ControllerConfigSpecVolumesAzureDisk struct {
	// diskName is the Name of the data disk in the blob storage.
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// diskURI is the URI of data disk in the blob storage.
	DiskUri *string `field:"required" json:"diskUri" yaml:"diskUri"`
	// cachingMode is the Host Caching mode: None, Read Only, Read Write.
	CachingMode *string `field:"optional" json:"cachingMode" yaml:"cachingMode"`
	// fsType is Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system.
	// Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set).
	//
	// defaults to shared.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// readOnly Defaults to false (read/write).
	//
	// ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// Default: false (read/write). ReadOnly here will force
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

