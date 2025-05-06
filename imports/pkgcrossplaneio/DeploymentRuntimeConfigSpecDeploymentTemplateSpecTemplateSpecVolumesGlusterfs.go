package pkgcrossplaneio


// glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.
//
// More info: https://examples.k8s.io/volumes/glusterfs/README.md
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesGlusterfs struct {
	// endpoints is the endpoint name that details Glusterfs topology.
	//
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Endpoints *string `field:"required" json:"endpoints" yaml:"endpoints"`
	// path is the Glusterfs volume path.
	//
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	Path *string `field:"required" json:"path" yaml:"path"`
	// readOnly here will force the Glusterfs volume to be mounted with read-only permissions.
	//
	// Defaults to false.
	// More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

