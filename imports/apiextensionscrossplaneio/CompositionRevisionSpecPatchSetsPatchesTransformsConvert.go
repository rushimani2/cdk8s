package apiextensionscrossplaneio


// Convert is used to cast the input into the given output type.
type CompositionRevisionSpecPatchSetsPatchesTransformsConvert struct {
	// ToType is the type of the output of this transform.
	ToType CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType `field:"required" json:"toType" yaml:"toType"`
	// The expected input format.
	//
	// * `quantity` - parses the input as a K8s [`resource.Quantity`](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity).
	// Only used during `string -> float64` conversions.
	// * `json` - parses the input as a JSON string.
	// Only used during `string -> object` or `string -> list` conversions.
	//
	// If this property is null, the default conversion is applied.
	Format CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat `field:"optional" json:"format" yaml:"format"`
}

