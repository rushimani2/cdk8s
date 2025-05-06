package apiextensionscrossplaneio


// service is a reference to the service for this webhook. Either service or url must be specified.
//
// If the webhook is running within the cluster, then you should use `service`.
type CompositeResourceDefinitionSpecConversionWebhookClientConfigService struct {
	// name is the name of the service.
	//
	// Required.
	Name *string `field:"required" json:"name" yaml:"name"`
	// namespace is the namespace of the service.
	//
	// Required.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// path is an optional URL path at which the webhook will be contacted.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// port is an optional service port at which the webhook will be contacted.
	//
	// `port` should be a valid port number (1-65535, inclusive).
	// Defaults to 443 for backward compatibility.
	// Default: 443 for backward compatibility.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

