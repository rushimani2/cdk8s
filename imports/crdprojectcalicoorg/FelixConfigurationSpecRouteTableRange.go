package crdprojectcalicoorg


// Deprecated in favor of RouteTableRanges.
//
// Calico programs additional Linux route tables for various purposes. RouteTableRange specifies the indices of the route tables that Calico should use.
type FelixConfigurationSpecRouteTableRange struct {
	Max *float64 `field:"required" json:"max" yaml:"max"`
	Min *float64 `field:"required" json:"min" yaml:"min"`
}

