package secretscrossplaneio


// Plugin configures External secret store as a plugin.
type StoreConfigSpecPlugin struct {
	// ConfigRef contains store config reference info.
	ConfigRef *StoreConfigSpecPluginConfigRef `field:"optional" json:"configRef" yaml:"configRef"`
	// Endpoint is the endpoint of the gRPC server.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
}

