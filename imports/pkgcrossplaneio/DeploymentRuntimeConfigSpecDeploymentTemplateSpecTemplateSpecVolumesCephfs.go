package pkgcrossplaneio


// cephFS represents a Ceph FS mount on the host that shares a pod's lifetime.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCephfs struct {
	// monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	Monitors *[]*string `field:"required" json:"monitors" yaml:"monitors"`
	// path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// readOnly is Optional: Defaults to false (read/write).
	//
	// ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	// Default: false (read/write). ReadOnly here will force
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	SecretFile *string `field:"optional" json:"secretFile" yaml:"secretFile"`
	// secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.
	//
	// More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
	SecretRef *DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCephfsSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
	// user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it.
	User *string `field:"optional" json:"user" yaml:"user"`
}

