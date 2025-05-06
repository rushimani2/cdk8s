package crdprojectcalicoorg


// FloatingIPs configures whether or not Felix will program floating IP addresses.
type FelixConfigurationSpecFloatingIPs string

const (
	// Enabled.
	FelixConfigurationSpecFloatingIPs_ENABLED FelixConfigurationSpecFloatingIPs = "ENABLED"
	// Disabled.
	FelixConfigurationSpecFloatingIPs_DISABLED FelixConfigurationSpecFloatingIPs = "DISABLED"
)

