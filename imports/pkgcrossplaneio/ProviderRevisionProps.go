package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A ProviderRevision represents a revision of a Provider. Crossplane creates new revisions when there are changes to a Provider.
//
// Crossplane creates and manages ProviderRevisions. Don't directly edit
// ProviderRevisions.
type ProviderRevisionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// ProviderRevisionSpec specifies configuration for a ProviderRevision.
	Spec *ProviderRevisionSpec `field:"optional" json:"spec" yaml:"spec"`
}

