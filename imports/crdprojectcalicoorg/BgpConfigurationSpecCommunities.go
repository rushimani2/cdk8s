package crdprojectcalicoorg


// Community contains standard or large community value and its name.
type BgpConfigurationSpecCommunities struct {
	// Name given to community value.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value must be of format `aa:nn` or `aa:nn:mm`.
	//
	// For standard community use `aa:nn` format, where `aa` and `nn` are 16 bit number. For large community use `aa:nn:mm` format, where `aa`, `nn` and `mm` are 32 bit number. Where, `aa` is an AS Number, `nn` and `mm` are per-AS identifier.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

