package pkgcrossplaneio

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Lock is the CRD type that tracks package dependencies.
type LockProps struct {
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	Packages *[]*LockPackages `field:"optional" json:"packages" yaml:"packages"`
}

