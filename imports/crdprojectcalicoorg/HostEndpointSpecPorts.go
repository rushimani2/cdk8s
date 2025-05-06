package crdprojectcalicoorg


type HostEndpointSpecPorts struct {
	Name *string `field:"required" json:"name" yaml:"name"`
	Port *float64 `field:"required" json:"port" yaml:"port"`
	Protocol HostEndpointSpecPortsProtocol `field:"required" json:"protocol" yaml:"protocol"`
}

