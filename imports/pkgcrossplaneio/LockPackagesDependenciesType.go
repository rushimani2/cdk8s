package pkgcrossplaneio


// Type is the type of package.
//
// Can be either Configuration or Provider.
// Deprecated: Specify an apiVersion and kind instead.
type LockPackagesDependenciesType string

const (
	// Configuration.
	LockPackagesDependenciesType_CONFIGURATION LockPackagesDependenciesType = "CONFIGURATION"
	// Provider.
	LockPackagesDependenciesType_PROVIDER LockPackagesDependenciesType = "PROVIDER"
	// Function.
	LockPackagesDependenciesType_FUNCTION LockPackagesDependenciesType = "FUNCTION"
)

