package apiextensionscrossplaneio


// The expected input format.
//
// * `quantity` - parses the input as a K8s [`resource.Quantity`](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity).
// Only used during `string -> float64` conversions.
// * `json` - parses the input as a JSON string.
// Only used during `string -> object` or `string -> list` conversions.
//
// If this property is null, the default conversion is applied.
type CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat string

const (
	// none.
	CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat_NONE CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat = "NONE"
	// quantity.
	CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat_QUANTITY CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat = "QUANTITY"
	// json.
	CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat_JSON CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat = "JSON"
)

