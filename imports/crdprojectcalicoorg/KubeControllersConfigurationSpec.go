package crdprojectcalicoorg


// KubeControllersConfigurationSpec contains the values of the Kubernetes controllers configuration.
type KubeControllersConfigurationSpec struct {
	// Controllers enables and configures individual Kubernetes controllers.
	Controllers *KubeControllersConfigurationSpecControllers `field:"required" json:"controllers" yaml:"controllers"`
	// DebugProfilePort configures the port to serve memory and cpu profiles on.
	//
	// If not specified, profiling is disabled.
	DebugProfilePort *float64 `field:"optional" json:"debugProfilePort" yaml:"debugProfilePort"`
	// EtcdV3CompactionPeriod is the period between etcdv3 compaction requests.
	//
	// Set to 0 to disable. [Default: 10m]
	EtcdV3CompactionPeriod *string `field:"optional" json:"etcdV3CompactionPeriod" yaml:"etcdV3CompactionPeriod"`
	// HealthChecks enables or disables support for health checks [Default: Enabled].
	HealthChecks *string `field:"optional" json:"healthChecks" yaml:"healthChecks"`
	// LogSeverityScreen is the log severity above which logs are sent to the stdout.
	//
	// [Default: Info].
	LogSeverityScreen *string `field:"optional" json:"logSeverityScreen" yaml:"logSeverityScreen"`
	// PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to.
	//
	// Set to 0 to disable. [Default: 9094]
	PrometheusMetricsPort *float64 `field:"optional" json:"prometheusMetricsPort" yaml:"prometheusMetricsPort"`
}

