package pkgcrossplaneio


// ControllerConfigRef references a ControllerConfig resource that will be used to configure the packaged controller Deployment.
//
// Deprecated: Use RuntimeConfigReference instead.
type FunctionV1Beta1SpecControllerConfigRef struct {
	// Name of the ControllerConfig.
	Name *string `field:"required" json:"name" yaml:"name"`
}

