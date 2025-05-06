package crdprojectcalicoorg


// BGPConfigurationSpec contains the values of the BGP configuration.
type BgpConfigurationSpec struct {
	// ASNumber is the default AS number used by a node.
	//
	// [Default: 64512].
	AsNumber *float64 `field:"optional" json:"asNumber" yaml:"asNumber"`
	// BindMode indicates whether to listen for BGP connections on all addresses (None) or only on the node's canonical IP address Node.Spec.BGP.IPvXAddress (NodeIP). Default behaviour is to listen for BGP connections on all addresses.
	BindMode *string `field:"optional" json:"bindMode" yaml:"bindMode"`
	// Communities is a list of BGP community values and their arbitrary names for tagging routes.
	Communities *[]*BgpConfigurationSpecCommunities `field:"optional" json:"communities" yaml:"communities"`
	// ListenPort is the port where BGP protocol should listen.
	//
	// Defaults to 179.
	// Default: 179.
	//
	ListenPort *float64 `field:"optional" json:"listenPort" yaml:"listenPort"`
	// LogSeverityScreen is the log severity above which logs are sent to the stdout.
	//
	// [Default: INFO].
	LogSeverityScreen *string `field:"optional" json:"logSeverityScreen" yaml:"logSeverityScreen"`
	// Time to allow for software restart for node-to-mesh peerings.
	//
	// When specified, this is configured as the graceful restart timeout.  When not specified, the BIRD default of 120s is used. This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled
	NodeMeshMaxRestartTime *string `field:"optional" json:"nodeMeshMaxRestartTime" yaml:"nodeMeshMaxRestartTime"`
	// Optional BGP password for full node-to-mesh peerings.
	//
	// This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled.
	NodeMeshPassword *BgpConfigurationSpecNodeMeshPassword `field:"optional" json:"nodeMeshPassword" yaml:"nodeMeshPassword"`
	// NodeToNodeMeshEnabled sets whether full node to node BGP mesh is enabled.
	//
	// [Default: true].
	NodeToNodeMeshEnabled *bool `field:"optional" json:"nodeToNodeMeshEnabled" yaml:"nodeToNodeMeshEnabled"`
	// PrefixAdvertisements contains per-prefix advertisement configuration.
	PrefixAdvertisements *[]*BgpConfigurationSpecPrefixAdvertisements `field:"optional" json:"prefixAdvertisements" yaml:"prefixAdvertisements"`
	// ServiceClusterIPs are the CIDR blocks from which service cluster IPs are allocated.
	//
	// If specified, Calico will advertise these blocks, as well as any cluster IPs within them.
	ServiceClusterIPs *[]*BgpConfigurationSpecServiceClusterIPs `field:"optional" json:"serviceClusterIPs" yaml:"serviceClusterIPs"`
	// ServiceExternalIPs are the CIDR blocks for Kubernetes Service External IPs.
	//
	// Kubernetes Service ExternalIPs will only be advertised if they are within one of these blocks.
	ServiceExternalIPs *[]*BgpConfigurationSpecServiceExternalIPs `field:"optional" json:"serviceExternalIPs" yaml:"serviceExternalIPs"`
	// ServiceLoadBalancerIPs are the CIDR blocks for Kubernetes Service LoadBalancer IPs.
	//
	// Kubernetes Service status.LoadBalancer.Ingress IPs will only be advertised if they are within one of these blocks.
	ServiceLoadBalancerIPs *[]*BgpConfigurationSpecServiceLoadBalancerIPs `field:"optional" json:"serviceLoadBalancerIPs" yaml:"serviceLoadBalancerIPs"`
}

