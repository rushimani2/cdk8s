package crdprojectcalicoorg


// ProtoPort is combination of protocol, port, and CIDR.
//
// Protocol and port must be specified.
type FelixConfigurationSpecFailsafeInboundHostPorts struct {
	Port *float64 `field:"required" json:"port" yaml:"port"`
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	Net *string `field:"optional" json:"net" yaml:"net"`
}

