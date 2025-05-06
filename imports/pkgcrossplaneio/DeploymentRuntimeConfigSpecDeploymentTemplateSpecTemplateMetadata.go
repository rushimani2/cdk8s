package pkgcrossplaneio


// Standard object's metadata.
//
// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateMetadata struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Finalizers *[]*string `field:"optional" json:"finalizers" yaml:"finalizers"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

