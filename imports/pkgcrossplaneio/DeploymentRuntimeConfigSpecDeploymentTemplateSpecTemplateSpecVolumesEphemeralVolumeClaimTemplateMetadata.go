package pkgcrossplaneio


// May contain labels and annotations that will be copied into the PVC when creating it.
//
// No other fields are allowed and will be rejected during
// validation.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateMetadata struct {
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	Finalizers *[]*string `field:"optional" json:"finalizers" yaml:"finalizers"`
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	Name *string `field:"optional" json:"name" yaml:"name"`
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

