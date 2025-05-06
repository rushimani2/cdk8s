package apiextensionscrossplaneio


// The expected input format.
//
// * `quantity` - parses the input as a K8s [`resource.Quantity`](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity).
// Only used during `string -> float64` conversions.
// * `json` - parses the input as a JSON string.
// Only used during `string -> object` or `string -> list` conversions.
//
// If this property is null, the default conversion is applied.
type CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat string

const (
	// none.
	CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat_NONE CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat = "NONE"
	// quantity.
	CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat_QUANTITY CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat = "QUANTITY"
	// json.
	CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat_JSON CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat = "JSON"
)

