package apiextensionscrossplaneio


// The expected input format.
//
// * `quantity` - parses the input as a K8s [`resource.Quantity`](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity).
// Only used during `string -> float64` conversions.
// * `json` - parses the input as a JSON string.
// Only used during `string -> object` or `string -> list` conversions.
//
// If this property is null, the default conversion is applied.
type CompositionRevisionSpecResourcesPatchesTransformsConvertFormat string

const (
	// none.
	CompositionRevisionSpecResourcesPatchesTransformsConvertFormat_NONE CompositionRevisionSpecResourcesPatchesTransformsConvertFormat = "NONE"
	// quantity.
	CompositionRevisionSpecResourcesPatchesTransformsConvertFormat_QUANTITY CompositionRevisionSpecResourcesPatchesTransformsConvertFormat = "QUANTITY"
	// json.
	CompositionRevisionSpecResourcesPatchesTransformsConvertFormat_JSON CompositionRevisionSpecResourcesPatchesTransformsConvertFormat = "JSON"
)

