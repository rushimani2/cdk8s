package apiextensionscrossplaneio


// Type sets the connection detail fetching behaviour to be used.
//
// Each
// connection detail type may require its own fields to be set on the
// ConnectionDetail object. If the type is omitted Crossplane will attempt
// to infer it based on which other fields were specified. If multiple
// fields are specified the order of precedence is:
// 1. FromValue
// 2. FromConnectionSecretKey
// 3. FromFieldPath
type CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType string

const (
	// FromConnectionSecretKey.
	CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType_FROM_CONNECTION_SECRET_KEY CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType = "FROM_CONNECTION_SECRET_KEY"
	// FromFieldPath.
	CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType_FROM_FIELD_PATH CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType = "FROM_FIELD_PATH"
	// FromValue.
	CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType_FROM_VALUE CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType = "FROM_VALUE"
)

