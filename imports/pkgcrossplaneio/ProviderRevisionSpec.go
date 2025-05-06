package pkgcrossplaneio


// ProviderRevisionSpec specifies configuration for a ProviderRevision.
type ProviderRevisionSpec struct {
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
	// ControllerConfigRef references a ControllerConfig resource that will be used to configure the packaged controller Deployment.
	//
	// Deprecated: Use RuntimeConfigReference instead.
	ControllerConfigRef *ProviderRevisionSpecControllerConfigRef `field:"optional" json:"controllerConfigRef" yaml:"controllerConfigRef"`
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
	PackagePullSecrets *[]*ProviderRevisionSpecPackagePullSecrets `field:"optional" json:"packagePullSecrets" yaml:"packagePullSecrets"`
	// RuntimeConfigRef references a RuntimeConfig resource that will be used to configure the package runtime.
	RuntimeConfigRef *ProviderRevisionSpecRuntimeConfigRef `field:"optional" json:"runtimeConfigRef" yaml:"runtimeConfigRef"`
	// SkipDependencyResolution indicates to the package manager whether to skip resolving dependencies for a package.
	//
	// Setting this value to true may have
	// unintended consequences.
	// Default is false.
	// Default: false.
	//
	SkipDependencyResolution *bool `field:"optional" json:"skipDependencyResolution" yaml:"skipDependencyResolution"`
	// TLSClientSecretName is the name of the TLS Secret that stores client certificates of the Provider.
	TlsClientSecretName *string `field:"optional" json:"tlsClientSecretName" yaml:"tlsClientSecretName"`
	// TLSServerSecretName is the name of the TLS Secret that stores server certificates of the Provider.
	TlsServerSecretName *string `field:"optional" json:"tlsServerSecretName" yaml:"tlsServerSecretName"`
}

