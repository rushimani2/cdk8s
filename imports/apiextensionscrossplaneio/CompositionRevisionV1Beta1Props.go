package apiextensionscrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A CompositionRevision represents a revision of a Composition. Crossplane creates new revisions when there are changes to the Composition.
//
// Crossplane creates and manages CompositionRevisions. Don't directly edit
// CompositionRevisions.
type CompositionRevisionV1Beta1Props struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// CompositionRevisionSpec specifies the desired state of the composition revision.
	Spec *CompositionRevisionV1Beta1Spec `field:"optional" json:"spec" yaml:"spec"`
}

