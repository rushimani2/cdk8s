package crdprojectcalicoorg


// IPReservationSpec contains the specification for an IPReservation resource.
type IpReservationSpec struct {
	// ReservedCIDRs is a list of CIDRs and/or IP addresses that Calico IPAM will exclude from new allocations.
	ReservedCidRs *[]*string `field:"optional" json:"reservedCidRs" yaml:"reservedCidRs"`
}

