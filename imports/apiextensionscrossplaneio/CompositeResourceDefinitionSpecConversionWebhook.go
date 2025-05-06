package apiextensionscrossplaneio


// webhook describes how to call the conversion webhook.
//
// Required when `strategy` is set to `"Webhook"`.
type CompositeResourceDefinitionSpecConversionWebhook struct {
	// conversionReviewVersions is an ordered list of preferred `ConversionReview` versions the Webhook expects.
	//
	// The API server will use the first version in
	// the list which it supports. If none of the versions specified in this list
	// are supported by API server, conversion will fail for the custom resource.
	// If a persisted Webhook configuration specifies allowed versions and does not
	// include any versions known to the API Server, calls to the webhook will fail.
	ConversionReviewVersions *[]*string `field:"required" json:"conversionReviewVersions" yaml:"conversionReviewVersions"`
	// clientConfig is the instructions for how to call the webhook if strategy is `Webhook`.
	ClientConfig *CompositeResourceDefinitionSpecConversionWebhookClientConfig `field:"optional" json:"clientConfig" yaml:"clientConfig"`
}

