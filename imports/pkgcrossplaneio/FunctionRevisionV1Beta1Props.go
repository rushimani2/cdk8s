package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A FunctionRevision represents a revision of a Function. Crossplane creates new revisions when there are changes to the Function.
//
// Crossplane creates and manages FunctionRevisions. Don't directly edit
// FunctionRevisions.
type FunctionRevisionV1Beta1Props struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// FunctionRevisionSpec specifies configuration for a FunctionRevision.
	Spec *FunctionRevisionV1Beta1Spec `field:"optional" json:"spec" yaml:"spec"`
}

