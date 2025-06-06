package apiextensionscrossplaneio


// The expected input format.
//
// * `quantity` - parses the input as a K8s [`resource.Quantity`](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity).
// Only used during `string -> float64` conversions.
// * `json` - parses the input as a JSON string.
// Only used during `string -> object` or `string -> list` conversions.
//
// If this property is null, the default conversion is applied.
type CompositionSpecResourcesPatchesTransformsConvertFormat string

const (
	// none.
	CompositionSpecResourcesPatchesTransformsConvertFormat_NONE CompositionSpecResourcesPatchesTransformsConvertFormat = "NONE"
	// quantity.
	CompositionSpecResourcesPatchesTransformsConvertFormat_QUANTITY CompositionSpecResourcesPatchesTransformsConvertFormat = "QUANTITY"
	// json.
	CompositionSpecResourcesPatchesTransformsConvertFormat_JSON CompositionSpecResourcesPatchesTransformsConvertFormat = "JSON"
)

