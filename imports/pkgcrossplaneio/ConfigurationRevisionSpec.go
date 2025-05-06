package pkgcrossplaneio


// PackageRevisionSpec specifies the desired state of a PackageRevision.
type ConfigurationRevisionSpec struct {
	// DesiredState of the PackageRevision.
	//
	// Can be either Active or Inactive.
	DesiredState *string `field:"required" json:"desiredState" yaml:"desiredState"`
	// Package image used by install Pod to extract package contents.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Revision number.
	//
	// Indicates when the revision will be garbage collected
	// based on the parent's RevisionHistoryLimit.
	Revision *float64 `field:"required" json:"revision" yaml:"revision"`
	// Map of string keys and values that can be used to organize and categorize (scope and select) objects.
	//
	// May match selectors of replication controllers
	// and services.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	CommonLabels *map[string]*string `field:"optional" json:"commonLabels" yaml:"commonLabels"`
	// IgnoreCrossplaneConstraints indicates to the package manager whether to honor Crossplane version constrains specified by the package.
	//
	// Default is false.
	// Default: false.
	//
	IgnoreCrossplaneConstraints *bool `field:"optional" json:"ignoreCrossplaneConstraints" yaml:"ignoreCrossplaneConstraints"`
	// PackagePullPolicy defines the pull policy for the package.
	//
	// It is also
	// applied to any images pulled for the package, such as a provider's
	// controller image.
	// Default is IfNotPresent.
	// Default: IfNotPresent.
	//
	PackagePullPolicy *string `field:"optional" json:"packagePullPolicy" yaml:"packagePullPolicy"`
	// PackagePullSecrets are named secrets in the same namespace that can be used to fetch packages from private registries.
	//
	// They are also applied to
	// any images pulled for the package, such as a provider's controller image.
	PackagePullSecrets *[]*ConfigurationRevisionSpecPackagePullSecrets `field:"optional" json:"packagePullSecrets" yaml:"packagePullSecrets"`
	// SkipDependencyResolution indicates to the package manager whether to skip resolving dependencies for a package.
	//
	// Setting this value to true may have
	// unintended consequences.
	// Default is false.
	// Default: false.
	//
	SkipDependencyResolution *bool `field:"optional" json:"skipDependencyResolution" yaml:"skipDependencyResolution"`
}

