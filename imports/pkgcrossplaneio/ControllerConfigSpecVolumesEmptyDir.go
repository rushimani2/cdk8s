package pkgcrossplaneio


// emptyDir represents a temporary directory that shares a pod's lifetime.
//
// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
type ControllerConfigSpecVolumesEmptyDir struct {
	// medium represents what type of storage medium should back this directory.
	//
	// The default is "" which means to use the node's default medium.
	// Must be an empty string (default) or Memory.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	// sizeLimit is the total amount of local storage required for this EmptyDir volume.
	//
	// The size limit is also applicable for memory medium.
	// The maximum usage on memory medium EmptyDir would be the minimum value between
	// the SizeLimit specified here and the sum of memory limits of all containers in a pod.
	// The default is nil which means that the limit is undefined.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
	SizeLimit ControllerConfigSpecVolumesEmptyDirSizeLimit `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

