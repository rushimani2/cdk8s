package pkgcrossplaneio


// LockPackage is a package that is in the lock.
type LockPackages struct {
	// Dependencies are the list of dependencies of this package.
	//
	// The order of
	// the dependencies will dictate the order in which they are resolved.
	Dependencies *[]*LockPackagesDependencies `field:"required" json:"dependencies" yaml:"dependencies"`
	// Name corresponds to the name of the package revision for this package.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Source is the OCI image name without a tag or digest.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Version is the tag or digest of the OCI image.
	Version *string `field:"required" json:"version" yaml:"version"`
	// APIVersion of the package.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Kind of the package (not the kind of the package revision).
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Type is the type of package.
	//
	// Deprecated: Specify an apiVersion and kind instead.
	Type LockPackagesType `field:"optional" json:"type" yaml:"type"`
}

