package apiextensionscrossplaneio


// PublishConnectionDetailsWithStoreConfig specifies the secret store config with which the connection details of composite resources dynamically provisioned using this composition will be published.
//
// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
// unless the relevant Crossplane feature flag is enabled, and may be
// changed or removed without notice.
type CompositionRevisionSpecPublishConnectionDetailsWithStoreConfigRef struct {
	// Name of the referenced StoreConfig.
	Name *string `field:"required" json:"name" yaml:"name"`
}

