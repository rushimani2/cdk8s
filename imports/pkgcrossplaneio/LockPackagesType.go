package pkgcrossplaneio


// Type is the type of package.
//
// Deprecated: Specify an apiVersion and kind instead.
type LockPackagesType string

const (
	// Configuration.
	LockPackagesType_CONFIGURATION LockPackagesType = "CONFIGURATION"
	// Provider.
	LockPackagesType_PROVIDER LockPackagesType = "PROVIDER"
	// Function.
	LockPackagesType_FUNCTION LockPackagesType = "FUNCTION"
)

