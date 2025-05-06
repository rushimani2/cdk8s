package crdprojectcalicoorg


// Set source-destination-check on AWS EC2 instances.
//
// Accepted value must be one of "DoNothing", "Enable" or "Disable". [Default: DoNothing]
type FelixConfigurationSpecAwsSrcDstCheck string

const (
	// DoNothing.
	FelixConfigurationSpecAwsSrcDstCheck_DO_NOTHING FelixConfigurationSpecAwsSrcDstCheck = "DO_NOTHING"
	// Enable.
	FelixConfigurationSpecAwsSrcDstCheck_ENABLE FelixConfigurationSpecAwsSrcDstCheck = "ENABLE"
	// Disable.
	FelixConfigurationSpecAwsSrcDstCheck_DISABLE FelixConfigurationSpecAwsSrcDstCheck = "DISABLE"
)

