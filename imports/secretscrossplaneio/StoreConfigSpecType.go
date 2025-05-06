package secretscrossplaneio


// Type configures which secret store to be used.
//
// Only the configuration
// block for this store will be used and others will be ignored if provided.
// Default is Kubernetes.
// Default: Kubernetes.
//
type StoreConfigSpecType string

const (
	// Kubernetes.
	StoreConfigSpecType_KUBERNETES StoreConfigSpecType = "KUBERNETES"
	// Vault.
	StoreConfigSpecType_VAULT StoreConfigSpecType = "VAULT"
	// Plugin.
	StoreConfigSpecType_PLUGIN StoreConfigSpecType = "PLUGIN"
)

