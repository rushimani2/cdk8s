package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// A FunctionRevision represents a revision of a Function. Crossplane creates new revisions when there are changes to the Function.
//
// Crossplane creates and manages FunctionRevisions. Don't directly edit
// FunctionRevisions.
type FunctionRevisionProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// FunctionRevisionSpec specifies configuration for a FunctionRevision.
	Spec *FunctionRevisionSpec `field:"optional" json:"spec" yaml:"spec"`
}

