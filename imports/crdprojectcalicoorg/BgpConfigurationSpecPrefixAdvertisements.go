package crdprojectcalicoorg


// PrefixAdvertisement configures advertisement properties for the specified CIDR.
type BgpConfigurationSpecPrefixAdvertisements struct {
	// CIDR for which properties should be advertised.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Communities can be list of either community names already defined in `Specs.Communities` or community value of format `aa:nn` or `aa:nn:mm`. For standard community use `aa:nn` format, where `aa` and `nn` are 16 bit number. For large community use `aa:nn:mm` format, where `aa`, `nn` and `mm` are 32 bit number. Where,`aa` is an AS Number, `nn` and `mm` are per-AS identifier.
	Communities *[]*string `field:"optional" json:"communities" yaml:"communities"`
}

