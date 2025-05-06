package pkgcrossplaneio


// A Dependency is a dependency of a package in the lock.
type LockPackagesDependencies struct {
	// Constraints is a valid semver range or a digest, which will be used to select a valid dependency version.
	Constraints *string `field:"required" json:"constraints" yaml:"constraints"`
	// Package is the OCI image name without a tag or digest.
	Package *string `field:"required" json:"package" yaml:"package"`
	// APIVersion of the package.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// Kind of the package (not the kind of the package revision).
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Type is the type of package.
	//
	// Can be either Configuration or Provider.
	// Deprecated: Specify an apiVersion and kind instead.
	Type LockPackagesDependenciesType `field:"optional" json:"type" yaml:"type"`
}

