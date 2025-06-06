package pkgcrossplaneio


// ephemeral represents a volume that is handled by a cluster storage driver.
//
// The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,
// and deleted when the pod is removed.
//
// Use this if:
// a) the volume is only needed while the pod runs,
// b) features of normal volumes like restoring from snapshot or capacity
// tracking are needed,
// c) the storage driver is specified through a storage class, and
// d) the storage driver supports dynamic volume provisioning through
// a PersistentVolumeClaim (see EphemeralVolumeSource for more
// information on the connection between this volume type
// and PersistentVolumeClaim).
//
// Use PersistentVolumeClaim or one of the vendor-specific
// APIs for volumes that persist for longer than the lifecycle
// of an individual pod.
//
// Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to
// be used that way - see the documentation of the driver for
// more information.
//
// A pod can use both types of ephemeral volumes and
// persistent volumes at the same time.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeral struct {
	// Will be used to create a stand-alone PVC to provision the volume.
	//
	// The pod in which this EphemeralVolumeSource is embedded will be the
	// owner of the PVC, i.e. the PVC will be deleted together with the
	// pod.  The name of the PVC will be `<pod name>-<volume name>` where
	// `<volume name>` is the name from the `PodSpec.Volumes` array
	// entry. Pod validation will reject the pod if the concatenated name
	// is not valid for a PVC (for example, too long).
	//
	// An existing PVC with that name that is not owned by the pod
	// will *not* be used for the pod to avoid using an unrelated
	// volume by mistake. Starting the pod is then blocked until
	// the unrelated PVC is removed. If such a pre-created PVC is
	// meant to be used by the pod, the PVC has to updated with an
	// owner reference to the pod once the pod exists. Normally
	// this should not be necessary, but it may be useful when
	// manually reconstructing a broken cluster.
	//
	// This field is read-only and no changes will be made by Kubernetes
	// to the PVC after it has been created.
	//
	// Required, must not be nil.
	VolumeClaimTemplate *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplate `field:"optional" json:"volumeClaimTemplate" yaml:"volumeClaimTemplate"`
}

