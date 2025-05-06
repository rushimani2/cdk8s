package pkgcrossplaneio


// Sleep represents the duration that the container should sleep before being terminated.
type DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopSleep struct {
	// Seconds is the number of seconds to sleep.
	Seconds *float64 `field:"required" json:"seconds" yaml:"seconds"`
}

