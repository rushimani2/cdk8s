package apiextensionscrossplaneio


// The expected input format.
//
// * `quantity` - parses the input as a K8s [`resource.Quantity`](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity).
// Only used during `string -> float64` conversions.
// * `json` - parses the input as a JSON string.
// Only used during `string -> object` or `string -> list` conversions.
//
// If this property is null, the default conversion is applied.
type CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat string

const (
	// none.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat_NONE CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat = "NONE"
	// quantity.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat_QUANTITY CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat = "QUANTITY"
	// json.
	CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat_JSON CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat = "JSON"
)

