package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A ConfigurationRevision represents a revision of a Configuration. Crossplane creates new revisions when there are changes to a Configuration.
//
// Crossplane creates and manages ConfigurationRevision. Don't directly edit
// ConfigurationRevisions.
type ConfigurationRevisionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// PackageRevisionSpec specifies the desired state of a PackageRevision.
	Spec *ConfigurationRevisionSpec `field:"optional" json:"spec" yaml:"spec"`
}

