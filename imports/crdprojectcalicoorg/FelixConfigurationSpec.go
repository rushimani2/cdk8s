package crdprojectcalicoorg


// FelixConfigurationSpec contains the values of the Felix configuration.
type FelixConfigurationSpec struct {
	// AllowIPIPPacketsFromWorkloads controls whether Felix will add a rule to drop IPIP encapsulated traffic from workloads [Default: false].
	AllowIpipPacketsFromWorkloads *bool `field:"optional" json:"allowIpipPacketsFromWorkloads" yaml:"allowIpipPacketsFromWorkloads"`
	// AllowVXLANPacketsFromWorkloads controls whether Felix will add a rule to drop VXLAN encapsulated traffic from workloads [Default: false].
	AllowVxlanPacketsFromWorkloads *bool `field:"optional" json:"allowVxlanPacketsFromWorkloads" yaml:"allowVxlanPacketsFromWorkloads"`
	// Set source-destination-check on AWS EC2 instances.
	//
	// Accepted value must be one of "DoNothing", "Enable" or "Disable". [Default: DoNothing]
	AwsSrcDstCheck FelixConfigurationSpecAwsSrcDstCheck `field:"optional" json:"awsSrcDstCheck" yaml:"awsSrcDstCheck"`
	// BPFConnectTimeLoadBalancingEnabled when in BPF mode, controls whether Felix installs the connection-time load balancer.
	//
	// The connect-time load balancer is required for the host to be able to reach Kubernetes services and it improves the performance of pod-to-service connections.  The only reason to disable it is for debugging purposes.  [Default: true]
	BpfConnectTimeLoadBalancingEnabled *bool `field:"optional" json:"bpfConnectTimeLoadBalancingEnabled" yaml:"bpfConnectTimeLoadBalancingEnabled"`
	// BPFDataIfacePattern is a regular expression that controls which interfaces Felix should attach BPF programs to in order to catch traffic to/from the network.
	//
	// This needs to match the interfaces that Calico workload traffic flows over as well as any interfaces that handle incoming traffic to nodeports and services from outside the cluster.  It should not match the workload interfaces (usually named cali...).
	BpfDataIfacePattern *string `field:"optional" json:"bpfDataIfacePattern" yaml:"bpfDataIfacePattern"`
	// BPFDisableUnprivileged, if enabled, Felix sets the kernel.unprivileged_bpf_disabled sysctl to disable unprivileged use of BPF.  This ensures that unprivileged users cannot access Calico's BPF maps and cannot insert their own BPF programs to interfere with Calico's. [Default: true].
	BpfDisableUnprivileged *bool `field:"optional" json:"bpfDisableUnprivileged" yaml:"bpfDisableUnprivileged"`
	// BPFEnabled, if enabled Felix will use the BPF dataplane.
	//
	// [Default: false].
	BpfEnabled *bool `field:"optional" json:"bpfEnabled" yaml:"bpfEnabled"`
	// BPFEnforceRPF enforce strict RPF on all interfaces with BPF programs regardless of what is the per-interfaces or global setting.
	//
	// Possible values are Disabled or Strict. [Default: Strict]
	BpfEnforceRpf *string `field:"optional" json:"bpfEnforceRpf" yaml:"bpfEnforceRpf"`
	// BPFExternalServiceMode in BPF mode, controls how connections from outside the cluster to services (node ports and cluster IPs) are forwarded to remote workloads.
	//
	// If set to "Tunnel" then both request and response traffic is tunneled to the remote node.  If set to "DSR", the request traffic is tunneled but the response traffic is sent directly from the remote node.  In "DSR" mode, the remote node appears to use the IP of the ingress node; this requires a permissive L2 network.  [Default: Tunnel]
	BpfExternalServiceMode *string `field:"optional" json:"bpfExternalServiceMode" yaml:"bpfExternalServiceMode"`
	// BPFExtToServiceConnmark in BPF mode, control a 32bit mark that is set on connections from an external client to a local service.
	//
	// This mark allows us to control how packets of that connection are routed within the host and how is routing interpreted by RPF check. [Default: 0]
	BpfExtToServiceConnmark *float64 `field:"optional" json:"bpfExtToServiceConnmark" yaml:"bpfExtToServiceConnmark"`
	// BPFHostConntrackBypass Controls whether to bypass Linux conntrack in BPF mode for workloads and services.
	//
	// [Default: true - bypass Linux conntrack].
	BpfHostConntrackBypass *bool `field:"optional" json:"bpfHostConntrackBypass" yaml:"bpfHostConntrackBypass"`
	// BPFKubeProxyEndpointSlicesEnabled in BPF mode, controls whether Felix's embedded kube-proxy accepts EndpointSlices or not.
	BpfKubeProxyEndpointSlicesEnabled *bool `field:"optional" json:"bpfKubeProxyEndpointSlicesEnabled" yaml:"bpfKubeProxyEndpointSlicesEnabled"`
	// BPFKubeProxyIptablesCleanupEnabled, if enabled in BPF mode, Felix will proactively clean up the upstream Kubernetes kube-proxy's iptables chains.
	//
	// Should only be enabled if kube-proxy is not running.  [Default: true]
	BpfKubeProxyIptablesCleanupEnabled *bool `field:"optional" json:"bpfKubeProxyIptablesCleanupEnabled" yaml:"bpfKubeProxyIptablesCleanupEnabled"`
	// BPFKubeProxyMinSyncPeriod, in BPF mode, controls the minimum time between updates to the dataplane for Felix's embedded kube-proxy.
	//
	// Lower values give reduced set-up latency.  Higher values reduce Felix CPU usage by batching up more work.  [Default: 1s]
	BpfKubeProxyMinSyncPeriod *string `field:"optional" json:"bpfKubeProxyMinSyncPeriod" yaml:"bpfKubeProxyMinSyncPeriod"`
	// BPFLogLevel controls the log level of the BPF programs when in BPF dataplane mode.
	//
	// One of "Off", "Info", or "Debug".  The logs are emitted to the BPF trace pipe, accessible with the command `tc exec bpf debug`. [Default: Off].
	BpfLogLevel *string `field:"optional" json:"bpfLogLevel" yaml:"bpfLogLevel"`
	// BPFMapSizeConntrack sets the size for the conntrack map.
	//
	// This map must be large enough to hold an entry for each active connection.  Warning: changing the size of the conntrack map can cause disruption.
	BpfMapSizeConntrack *float64 `field:"optional" json:"bpfMapSizeConntrack" yaml:"bpfMapSizeConntrack"`
	// BPFMapSizeIfState sets the size for ifstate map.
	//
	// The ifstate map must be large enough to hold an entry for each device (host + workloads) on a host.
	BpfMapSizeIfState *float64 `field:"optional" json:"bpfMapSizeIfState" yaml:"bpfMapSizeIfState"`
	// BPFMapSizeIPSets sets the size for ipsets map.
	//
	// The IP sets map must be large enough to hold an entry for each endpoint matched by every selector in the source/destination matches in network policy.  Selectors such as "all()" can result in large numbers of entries (one entry per endpoint in that case).
	BpfMapSizeIpSets *float64 `field:"optional" json:"bpfMapSizeIpSets" yaml:"bpfMapSizeIpSets"`
	BpfMapSizeNatAffinity *float64 `field:"optional" json:"bpfMapSizeNatAffinity" yaml:"bpfMapSizeNatAffinity"`
	// BPFMapSizeNATBackend sets the size for nat back end map.
	//
	// This is the total number of endpoints. This is mostly more than the size of the number of services.
	BpfMapSizeNatBackend *float64 `field:"optional" json:"bpfMapSizeNatBackend" yaml:"bpfMapSizeNatBackend"`
	// BPFMapSizeNATFrontend sets the size for nat front end map.
	//
	// FrontendMap should be large enough to hold an entry for each nodeport, external IP and each port in each service.
	BpfMapSizeNatFrontend *float64 `field:"optional" json:"bpfMapSizeNatFrontend" yaml:"bpfMapSizeNatFrontend"`
	// BPFMapSizeRoute sets the size for the routes map.
	//
	// The routes map should be large enough to hold one entry per workload and a handful of entries per host (enough to cover its own IPs and tunnel IPs).
	BpfMapSizeRoute *float64 `field:"optional" json:"bpfMapSizeRoute" yaml:"bpfMapSizeRoute"`
	// BPFPolicyDebugEnabled when true, Felix records detailed information about the BPF policy programs, which can be examined with the calico-bpf command-line tool.
	BpfPolicyDebugEnabled *bool `field:"optional" json:"bpfPolicyDebugEnabled" yaml:"bpfPolicyDebugEnabled"`
	// BPFPSNATPorts sets the range from which we randomly pick a port if there is a source port collision.
	//
	// This should be within the ephemeral range as defined by RFC 6056 (1024–65535) and preferably outside the  ephemeral ranges used by common operating systems. Linux uses 32768–60999, while others mostly use the IANA defined range 49152–65535. It is not necessarily a problem if this range overlaps with the operating systems. Both ends of the range are inclusive. [Default: 20000:29999]
	BpfPsnatPorts FelixConfigurationSpecBpfPsnatPorts `field:"optional" json:"bpfPsnatPorts" yaml:"bpfPsnatPorts"`
	// ChainInsertMode controls whether Felix hooks the kernel's top-level iptables chains by inserting a rule at the top of the chain or by appending a rule at the bottom.
	//
	// insert is the safe default since it prevents Calico's rules from being bypassed. If you switch to append mode, be sure that the other rules in the chains signal acceptance by falling through to the Calico rules, otherwise the Calico policy will be bypassed. [Default: insert]
	ChainInsertMode *string `field:"optional" json:"chainInsertMode" yaml:"chainInsertMode"`
	// DataplaneDriver filename of the external dataplane driver to use.
	//
	// Only used if UseInternalDataplaneDriver is set to false.
	DataplaneDriver *string `field:"optional" json:"dataplaneDriver" yaml:"dataplaneDriver"`
	// DataplaneWatchdogTimeout is the readiness/liveness timeout used for Felix's (internal) dataplane driver.
	//
	// Increase this value if you experience spurious non-ready or non-live events when Felix is under heavy load. Decrease the value to get felix to report non-live or non-ready more quickly. [Default: 90s]
	DataplaneWatchdogTimeout *string `field:"optional" json:"dataplaneWatchdogTimeout" yaml:"dataplaneWatchdogTimeout"`
	DebugDisableLogDropping *bool `field:"optional" json:"debugDisableLogDropping" yaml:"debugDisableLogDropping"`
	DebugMemoryProfilePath *string `field:"optional" json:"debugMemoryProfilePath" yaml:"debugMemoryProfilePath"`
	DebugSimulateCalcGraphHangAfter *string `field:"optional" json:"debugSimulateCalcGraphHangAfter" yaml:"debugSimulateCalcGraphHangAfter"`
	DebugSimulateDataplaneHangAfter *string `field:"optional" json:"debugSimulateDataplaneHangAfter" yaml:"debugSimulateDataplaneHangAfter"`
	// DefaultEndpointToHostAction controls what happens to traffic that goes from a workload endpoint to the host itself (after the traffic hits the endpoint egress policy).
	//
	// By default Calico blocks traffic from workload endpoints to the host itself with an iptables "DROP" action. If you want to allow some or all traffic from endpoint to host, set this parameter to RETURN or ACCEPT. Use RETURN if you have your own rules in the iptables "INPUT" chain; Calico will insert its rules at the top of that chain, then "RETURN" packets to the "INPUT" chain once it has completed processing workload endpoint egress policy. Use ACCEPT to unconditionally accept packets from workloads after processing workload endpoint egress policy. [Default: Drop]
	DefaultEndpointToHostAction *string `field:"optional" json:"defaultEndpointToHostAction" yaml:"defaultEndpointToHostAction"`
	// This defines the route protocol added to programmed device routes, by default this will be RTPROT_BOOT when left blank.
	DeviceRouteProtocol *float64 `field:"optional" json:"deviceRouteProtocol" yaml:"deviceRouteProtocol"`
	// This is the IPv4 source address to use on programmed device routes.
	//
	// By default the source address is left blank, leaving the kernel to choose the source address used.
	DeviceRouteSourceAddress *string `field:"optional" json:"deviceRouteSourceAddress" yaml:"deviceRouteSourceAddress"`
	// This is the IPv6 source address to use on programmed device routes.
	//
	// By default the source address is left blank, leaving the kernel to choose the source address used.
	DeviceRouteSourceAddressIPv6 *string `field:"optional" json:"deviceRouteSourceAddressIPv6" yaml:"deviceRouteSourceAddressIPv6"`
	DisableConntrackInvalidCheck *bool `field:"optional" json:"disableConntrackInvalidCheck" yaml:"disableConntrackInvalidCheck"`
	EndpointReportingDelay *string `field:"optional" json:"endpointReportingDelay" yaml:"endpointReportingDelay"`
	EndpointReportingEnabled *bool `field:"optional" json:"endpointReportingEnabled" yaml:"endpointReportingEnabled"`
	// ExternalNodesCIDRList is a list of CIDR's of external-non-calico-nodes which may source tunnel traffic and have the tunneled traffic be accepted at calico nodes.
	ExternalNodesList *[]*string `field:"optional" json:"externalNodesList" yaml:"externalNodesList"`
	// FailsafeInboundHostPorts is a list of UDP/TCP ports and CIDRs that Felix will allow incoming traffic to host endpoints on irrespective of the security policy.
	//
	// This is useful to avoid accidentally cutting off a host with incorrect configuration. For back-compatibility, if the protocol is not specified, it defaults to "tcp". If a CIDR is not specified, it will allow traffic from all addresses. To disable all inbound host ports, use the value none. The default value allows ssh access and DHCP. [Default: tcp:22, udp:68, tcp:179, tcp:2379, tcp:2380, tcp:6443, tcp:6666, tcp:6667]
	FailsafeInboundHostPorts *[]*FelixConfigurationSpecFailsafeInboundHostPorts `field:"optional" json:"failsafeInboundHostPorts" yaml:"failsafeInboundHostPorts"`
	// FailsafeOutboundHostPorts is a list of UDP/TCP ports and CIDRs that Felix will allow outgoing traffic from host endpoints to irrespective of the security policy.
	//
	// This is useful to avoid accidentally cutting off a host with incorrect configuration. For back-compatibility, if the protocol is not specified, it defaults to "tcp". If a CIDR is not specified, it will allow traffic from all addresses. To disable all outbound host ports, use the value none. The default value opens etcd's standard ports to ensure that Felix does not get cut off from etcd as well as allowing DHCP and DNS. [Default: tcp:179, tcp:2379, tcp:2380, tcp:6443, tcp:6666, tcp:6667, udp:53, udp:67]
	FailsafeOutboundHostPorts *[]*FelixConfigurationSpecFailsafeOutboundHostPorts `field:"optional" json:"failsafeOutboundHostPorts" yaml:"failsafeOutboundHostPorts"`
	// FeatureDetectOverride is used to override the feature detection.
	//
	// Values are specified in a comma separated list with no spaces, example; "SNATFullyRandom=true,MASQFullyRandom=false,RestoreSupportsLock=". "true" or "false" will force the feature, empty or omitted values are auto-detected.
	FeatureDetectOverride *string `field:"optional" json:"featureDetectOverride" yaml:"featureDetectOverride"`
	// FloatingIPs configures whether or not Felix will program floating IP addresses.
	FloatingIPs FelixConfigurationSpecFloatingIPs `field:"optional" json:"floatingIPs" yaml:"floatingIPs"`
	// GenericXDPEnabled enables Generic XDP so network cards that don't support XDP offload or driver modes can use XDP.
	//
	// This is not recommended since it doesn't provide better performance than iptables. [Default: false]
	GenericXdpEnabled *bool `field:"optional" json:"genericXdpEnabled" yaml:"genericXdpEnabled"`
	HealthEnabled *bool `field:"optional" json:"healthEnabled" yaml:"healthEnabled"`
	HealthHost *string `field:"optional" json:"healthHost" yaml:"healthHost"`
	HealthPort *float64 `field:"optional" json:"healthPort" yaml:"healthPort"`
	// InterfaceExclude is a comma-separated list of interfaces that Felix should exclude when monitoring for host endpoints.
	//
	// The default value ensures that Felix ignores Kubernetes' IPVS dummy interface, which is used internally by kube-proxy. If you want to exclude multiple interface names using a single value, the list supports regular expressions. For regular expressions you must wrap the value with '/'. For example having values '/^kube/,veth1' will exclude all interfaces that begin with 'kube' and also the interface 'veth1'. [Default: kube-ipvs0]
	InterfaceExclude *string `field:"optional" json:"interfaceExclude" yaml:"interfaceExclude"`
	// InterfacePrefix is the interface name prefix that identifies workload endpoints and so distinguishes them from host endpoint interfaces.
	//
	// Note: in environments other than bare metal, the orchestrators configure this appropriately. For example our Kubernetes and Docker integrations set the 'cali' value, and our OpenStack integration sets the 'tap' value. [Default: cali]
	InterfacePrefix *string `field:"optional" json:"interfacePrefix" yaml:"interfacePrefix"`
	// InterfaceRefreshInterval is the period at which Felix rescans local interfaces to verify their state.
	//
	// The rescan can be disabled by setting the interval to 0.
	InterfaceRefreshInterval *string `field:"optional" json:"interfaceRefreshInterval" yaml:"interfaceRefreshInterval"`
	// IPIPEnabled overrides whether Felix should configure an IPIP interface on the host.
	//
	// Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]
	IpipEnabled *bool `field:"optional" json:"ipipEnabled" yaml:"ipipEnabled"`
	// IPIPMTU is the MTU to set on the tunnel device.
	//
	// See Configuring MTU [Default: 1440].
	IpipMtu *float64 `field:"optional" json:"ipipMtu" yaml:"ipipMtu"`
	// IpsetsRefreshInterval is the period at which Felix re-checks all iptables state to ensure that no other process has accidentally broken Calico's rules.
	//
	// Set to 0 to disable iptables refresh. [Default: 90s]
	IpsetsRefreshInterval *string `field:"optional" json:"ipsetsRefreshInterval" yaml:"ipsetsRefreshInterval"`
	// IptablesBackend specifies which backend of iptables will be used.
	//
	// The default is legacy.
	IptablesBackend *string `field:"optional" json:"iptablesBackend" yaml:"iptablesBackend"`
	IptablesFilterAllowAction *string `field:"optional" json:"iptablesFilterAllowAction" yaml:"iptablesFilterAllowAction"`
	// IptablesLockFilePath is the location of the iptables lock file.
	//
	// You may need to change this if the lock file is not in its standard location (for example if you have mapped it into Felix's container at a different path). [Default: /run/xtables.lock]
	IptablesLockFilePath *string `field:"optional" json:"iptablesLockFilePath" yaml:"iptablesLockFilePath"`
	// IptablesLockProbeInterval is the time that Felix will wait between attempts to acquire the iptables lock if it is not available.
	//
	// Lower values make Felix more responsive when the lock is contended, but use more CPU. [Default: 50ms]
	IptablesLockProbeInterval *string `field:"optional" json:"iptablesLockProbeInterval" yaml:"iptablesLockProbeInterval"`
	// IptablesLockTimeout is the time that Felix will wait for the iptables lock, or 0, to disable.
	//
	// To use this feature, Felix must share the iptables lock file with all other processes that also take the lock. When running Felix inside a container, this requires the /run directory of the host to be mounted into the calico/node or calico/felix container. [Default: 0s disabled]
	IptablesLockTimeout *string `field:"optional" json:"iptablesLockTimeout" yaml:"iptablesLockTimeout"`
	IptablesMangleAllowAction *string `field:"optional" json:"iptablesMangleAllowAction" yaml:"iptablesMangleAllowAction"`
	// IptablesMarkMask is the mask that Felix selects its IPTables Mark bits from.
	//
	// Should be a 32 bit hexadecimal number with at least 8 bits set, none of which clash with any other mark bits in use on the system. [Default: 0xff000000]
	IptablesMarkMask *float64 `field:"optional" json:"iptablesMarkMask" yaml:"iptablesMarkMask"`
	IptablesNatOutgoingInterfaceFilter *string `field:"optional" json:"iptablesNatOutgoingInterfaceFilter" yaml:"iptablesNatOutgoingInterfaceFilter"`
	// IptablesPostWriteCheckInterval is the period after Felix has done a write to the dataplane that it schedules an extra read back in order to check the write was not clobbered by another process.
	//
	// This should only occur if another application on the system doesn't respect the iptables lock. [Default: 1s]
	IptablesPostWriteCheckInterval *string `field:"optional" json:"iptablesPostWriteCheckInterval" yaml:"iptablesPostWriteCheckInterval"`
	// IptablesRefreshInterval is the period at which Felix re-checks the IP sets in the dataplane to ensure that no other process has accidentally broken Calico's rules.
	//
	// Set to 0 to disable IP sets refresh. Note: the default for this value is lower than the other refresh intervals as a workaround for a Linux kernel bug that was fixed in kernel version 4.11. If you are using v4.11 or greater you may want to set this to, a higher value to reduce Felix CPU usage. [Default: 10s]
	IptablesRefreshInterval *string `field:"optional" json:"iptablesRefreshInterval" yaml:"iptablesRefreshInterval"`
	// IPv6Support controls whether Felix enables support for IPv6 (if supported by the in-use dataplane).
	Ipv6Support *bool `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
	// KubeNodePortRanges holds list of port ranges used for service node ports.
	//
	// Only used if felix detects kube-proxy running in ipvs mode. Felix uses these ranges to separate host and workload traffic. [Default: 30000:32767].
	KubeNodePortRanges *[]FelixConfigurationSpecKubeNodePortRanges `field:"optional" json:"kubeNodePortRanges" yaml:"kubeNodePortRanges"`
	// LogDebugFilenameRegex controls which source code files have their Debug log output included in the logs.
	//
	// Only logs from files with names that match the given regular expression are included.  The filter only applies to Debug level logs.
	LogDebugFilenameRegex *string `field:"optional" json:"logDebugFilenameRegex" yaml:"logDebugFilenameRegex"`
	// LogFilePath is the full path to the Felix log.
	//
	// Set to none to disable file logging. [Default: /var/log/calico/felix.log]
	LogFilePath *string `field:"optional" json:"logFilePath" yaml:"logFilePath"`
	// LogPrefix is the log prefix that Felix uses when rendering LOG rules.
	//
	// [Default: calico-packet].
	LogPrefix *string `field:"optional" json:"logPrefix" yaml:"logPrefix"`
	// LogSeverityFile is the log severity above which logs are sent to the log file.
	//
	// [Default: Info].
	LogSeverityFile *string `field:"optional" json:"logSeverityFile" yaml:"logSeverityFile"`
	// LogSeverityScreen is the log severity above which logs are sent to the stdout.
	//
	// [Default: Info].
	LogSeverityScreen *string `field:"optional" json:"logSeverityScreen" yaml:"logSeverityScreen"`
	// LogSeveritySys is the log severity above which logs are sent to the syslog.
	//
	// Set to None for no logging to syslog. [Default: Info]
	LogSeveritySys *string `field:"optional" json:"logSeveritySys" yaml:"logSeveritySys"`
	MaxIpsetSize *float64 `field:"optional" json:"maxIpsetSize" yaml:"maxIpsetSize"`
	// MetadataAddr is the IP address or domain name of the server that can answer VM queries for cloud-init metadata.
	//
	// In OpenStack, this corresponds to the machine running nova-api (or in Ubuntu, nova-api-metadata). A value of none (case insensitive) means that Felix should not set up any NAT rule for the metadata path. [Default: 127.0.0.1]
	MetadataAddr *string `field:"optional" json:"metadataAddr" yaml:"metadataAddr"`
	// MetadataPort is the port of the metadata server.
	//
	// This, combined with global.MetadataAddr (if not 'None'), is used to set up a NAT rule, from 169.254.169.254:80 to MetadataAddr:MetadataPort. In most cases this should not need to be changed [Default: 8775].
	MetadataPort *float64 `field:"optional" json:"metadataPort" yaml:"metadataPort"`
	// MTUIfacePattern is a regular expression that controls which interfaces Felix should scan in order to calculate the host's MTU.
	//
	// This should not match workload interfaces (usually named cali...).
	MtuIfacePattern *string `field:"optional" json:"mtuIfacePattern" yaml:"mtuIfacePattern"`
	// NATOutgoingAddress specifies an address to use when performing source NAT for traffic in a natOutgoing pool that is leaving the network.
	//
	// By default the address used is an address on the interface the traffic is leaving on (ie it uses the iptables MASQUERADE target).
	NatOutgoingAddress *string `field:"optional" json:"natOutgoingAddress" yaml:"natOutgoingAddress"`
	// NATPortRange specifies the range of ports that is used for port mapping when doing outgoing NAT.
	//
	// When unset the default behavior of the network stack is used.
	NatPortRange FelixConfigurationSpecNatPortRange `field:"optional" json:"natPortRange" yaml:"natPortRange"`
	NetlinkTimeout *string `field:"optional" json:"netlinkTimeout" yaml:"netlinkTimeout"`
	// OpenstackRegion is the name of the region that a particular Felix belongs to.
	//
	// In a multi-region Calico/OpenStack deployment, this must be configured somehow for each Felix (here in the datamodel, or in felix.cfg or the environment on each compute node), and must match the [calico] openstack_region value configured in neutron.conf on each node. [Default: Empty]
	OpenstackRegion *string `field:"optional" json:"openstackRegion" yaml:"openstackRegion"`
	// PolicySyncPathPrefix is used to by Felix to communicate policy changes to external services, like Application layer policy.
	//
	// [Default: Empty].
	PolicySyncPathPrefix *string `field:"optional" json:"policySyncPathPrefix" yaml:"policySyncPathPrefix"`
	// PrometheusGoMetricsEnabled disables Go runtime metrics collection, which the Prometheus client does by default, when set to false.
	//
	// This reduces the number of metrics reported, reducing Prometheus load. [Default: true]
	PrometheusGoMetricsEnabled *bool `field:"optional" json:"prometheusGoMetricsEnabled" yaml:"prometheusGoMetricsEnabled"`
	// PrometheusMetricsEnabled enables the Prometheus metrics server in Felix if set to true.
	//
	// [Default: false].
	PrometheusMetricsEnabled *bool `field:"optional" json:"prometheusMetricsEnabled" yaml:"prometheusMetricsEnabled"`
	// PrometheusMetricsHost is the host that the Prometheus metrics server should bind to.
	//
	// [Default: empty].
	PrometheusMetricsHost *string `field:"optional" json:"prometheusMetricsHost" yaml:"prometheusMetricsHost"`
	// PrometheusMetricsPort is the TCP port that the Prometheus metrics server should bind to.
	//
	// [Default: 9091].
	PrometheusMetricsPort *float64 `field:"optional" json:"prometheusMetricsPort" yaml:"prometheusMetricsPort"`
	// PrometheusProcessMetricsEnabled disables process metrics collection, which the Prometheus client does by default, when set to false.
	//
	// This reduces the number of metrics reported, reducing Prometheus load. [Default: true]
	PrometheusProcessMetricsEnabled *bool `field:"optional" json:"prometheusProcessMetricsEnabled" yaml:"prometheusProcessMetricsEnabled"`
	// PrometheusWireGuardMetricsEnabled disables wireguard metrics collection, which the Prometheus client does by default, when set to false.
	//
	// This reduces the number of metrics reported, reducing Prometheus load. [Default: true]
	PrometheusWireGuardMetricsEnabled *bool `field:"optional" json:"prometheusWireGuardMetricsEnabled" yaml:"prometheusWireGuardMetricsEnabled"`
	// Whether or not to remove device routes that have not been programmed by Felix.
	//
	// Disabling this will allow external applications to also add device routes. This is enabled by default which means we will remove externally added routes.
	RemoveExternalRoutes *bool `field:"optional" json:"removeExternalRoutes" yaml:"removeExternalRoutes"`
	// ReportingInterval is the interval at which Felix reports its status into the datastore or 0 to disable.
	//
	// Must be non-zero in OpenStack deployments. [Default: 30s]
	ReportingInterval *string `field:"optional" json:"reportingInterval" yaml:"reportingInterval"`
	// ReportingTTL is the time-to-live setting for process-wide status reports.
	//
	// [Default: 90s].
	ReportingTtl *string `field:"optional" json:"reportingTtl" yaml:"reportingTtl"`
	// RouteRefreshInterval is the period at which Felix re-checks the routes in the dataplane to ensure that no other process has accidentally broken Calico's rules.
	//
	// Set to 0 to disable route refresh. [Default: 90s]
	RouteRefreshInterval *string `field:"optional" json:"routeRefreshInterval" yaml:"routeRefreshInterval"`
	// RouteSource configures where Felix gets its routing information.
	//
	// - WorkloadIPs: use workload endpoints to construct routes. - CalicoIPAM: the default - use IPAM data to construct routes.
	RouteSource *string `field:"optional" json:"routeSource" yaml:"routeSource"`
	// RouteSyncDisabled will disable all operations performed on the route table.
	//
	// Set to true to run in network-policy mode only.
	RouteSyncDisabled *bool `field:"optional" json:"routeSyncDisabled" yaml:"routeSyncDisabled"`
	// Deprecated in favor of RouteTableRanges.
	//
	// Calico programs additional Linux route tables for various purposes. RouteTableRange specifies the indices of the route tables that Calico should use.
	RouteTableRange *FelixConfigurationSpecRouteTableRange `field:"optional" json:"routeTableRange" yaml:"routeTableRange"`
	// Calico programs additional Linux route tables for various purposes.
	//
	// RouteTableRanges specifies a set of table index ranges that Calico should use. Deprecates`RouteTableRange`, overrides `RouteTableRange`.
	RouteTableRanges *[]*FelixConfigurationSpecRouteTableRanges `field:"optional" json:"routeTableRanges" yaml:"routeTableRanges"`
	// When service IP advertisement is enabled, prevent routing loops to service IPs that are not in use, by dropping or rejecting packets that do not get DNAT'd by kube-proxy.
	//
	// Unless set to "Disabled", in which case such routing loops continue to be allowed. [Default: Drop]
	ServiceLoopPrevention *string `field:"optional" json:"serviceLoopPrevention" yaml:"serviceLoopPrevention"`
	// SidecarAccelerationEnabled enables experimental sidecar acceleration [Default: false].
	SidecarAccelerationEnabled *bool `field:"optional" json:"sidecarAccelerationEnabled" yaml:"sidecarAccelerationEnabled"`
	// UsageReportingEnabled reports anonymous Calico version number and cluster size to projectcalico.org. Logs warnings returned by the usage server. For example, if a significant security vulnerability has been discovered in the version of Calico being used. [Default: true].
	UsageReportingEnabled *bool `field:"optional" json:"usageReportingEnabled" yaml:"usageReportingEnabled"`
	// UsageReportingInitialDelay controls the minimum delay before Felix makes a report.
	//
	// [Default: 300s].
	UsageReportingInitialDelay *string `field:"optional" json:"usageReportingInitialDelay" yaml:"usageReportingInitialDelay"`
	// UsageReportingInterval controls the interval at which Felix makes reports.
	//
	// [Default: 86400s].
	UsageReportingInterval *string `field:"optional" json:"usageReportingInterval" yaml:"usageReportingInterval"`
	// UseInternalDataplaneDriver, if true, Felix will use its internal dataplane programming logic.
	//
	// If false, it will launch an external dataplane driver and communicate with it over protobuf.
	UseInternalDataplaneDriver *bool `field:"optional" json:"useInternalDataplaneDriver" yaml:"useInternalDataplaneDriver"`
	// VXLANEnabled overrides whether Felix should create the VXLAN tunnel device for IPv4 VXLAN networking.
	//
	// Optional as Felix determines this based on the existing IP pools. [Default: nil (unset)]
	VxlanEnabled *bool `field:"optional" json:"vxlanEnabled" yaml:"vxlanEnabled"`
	// VXLANMTU is the MTU to set on the IPv4 VXLAN tunnel device.
	//
	// See Configuring MTU [Default: 1410].
	VxlanMtu *float64 `field:"optional" json:"vxlanMtu" yaml:"vxlanMtu"`
	// VXLANMTUV6 is the MTU to set on the IPv6 VXLAN tunnel device.
	//
	// See Configuring MTU [Default: 1390].
	VxlanMtuv6 *float64 `field:"optional" json:"vxlanMtuv6" yaml:"vxlanMtuv6"`
	VxlanPort *float64 `field:"optional" json:"vxlanPort" yaml:"vxlanPort"`
	VxlanVni *float64 `field:"optional" json:"vxlanVni" yaml:"vxlanVni"`
	// WireguardEnabled controls whether Wireguard is enabled for IPv4 (encapsulating IPv4 traffic over an IPv4 underlay network).
	//
	// [Default: false].
	WireguardEnabled *bool `field:"optional" json:"wireguardEnabled" yaml:"wireguardEnabled"`
	// WireguardEnabledV6 controls whether Wireguard is enabled for IPv6 (encapsulating IPv6 traffic over an IPv6 underlay network).
	//
	// [Default: false].
	WireguardEnabledV6 *bool `field:"optional" json:"wireguardEnabledV6" yaml:"wireguardEnabledV6"`
	// WireguardHostEncryptionEnabled controls whether Wireguard host-to-host encryption is enabled.
	//
	// [Default: false].
	WireguardHostEncryptionEnabled *bool `field:"optional" json:"wireguardHostEncryptionEnabled" yaml:"wireguardHostEncryptionEnabled"`
	// WireguardInterfaceName specifies the name to use for the IPv4 Wireguard interface.
	//
	// [Default: wireguard.cali]
	WireguardInterfaceName *string `field:"optional" json:"wireguardInterfaceName" yaml:"wireguardInterfaceName"`
	// WireguardInterfaceNameV6 specifies the name to use for the IPv6 Wireguard interface.
	//
	// [Default: wg-v6.cali]
	WireguardInterfaceNameV6 *string `field:"optional" json:"wireguardInterfaceNameV6" yaml:"wireguardInterfaceNameV6"`
	// WireguardKeepAlive controls Wireguard PersistentKeepalive option.
	//
	// Set 0 to disable. [Default: 0]
	WireguardKeepAlive *string `field:"optional" json:"wireguardKeepAlive" yaml:"wireguardKeepAlive"`
	// WireguardListeningPort controls the listening port used by IPv4 Wireguard.
	//
	// [Default: 51820].
	WireguardListeningPort *float64 `field:"optional" json:"wireguardListeningPort" yaml:"wireguardListeningPort"`
	// WireguardListeningPortV6 controls the listening port used by IPv6 Wireguard.
	//
	// [Default: 51821].
	WireguardListeningPortV6 *float64 `field:"optional" json:"wireguardListeningPortV6" yaml:"wireguardListeningPortV6"`
	// WireguardMTU controls the MTU on the IPv4 Wireguard interface.
	//
	// See Configuring MTU [Default: 1440].
	WireguardMtu *float64 `field:"optional" json:"wireguardMtu" yaml:"wireguardMtu"`
	// WireguardMTUV6 controls the MTU on the IPv6 Wireguard interface.
	//
	// See Configuring MTU [Default: 1420].
	WireguardMtuv6 *float64 `field:"optional" json:"wireguardMtuv6" yaml:"wireguardMtuv6"`
	// WireguardRoutingRulePriority controls the priority value to use for the Wireguard routing rule.
	//
	// [Default: 99].
	WireguardRoutingRulePriority *float64 `field:"optional" json:"wireguardRoutingRulePriority" yaml:"wireguardRoutingRulePriority"`
	// WorkloadSourceSpoofing controls whether pods can use the allowedSourcePrefixes annotation to send traffic with a source IP address that is not theirs.
	//
	// This is disabled by default. When set to "Any", pods can request any prefix.
	WorkloadSourceSpoofing *string `field:"optional" json:"workloadSourceSpoofing" yaml:"workloadSourceSpoofing"`
	// XDPEnabled enables XDP acceleration for suitable untracked incoming deny rules.
	//
	// [Default: true].
	XdpEnabled *bool `field:"optional" json:"xdpEnabled" yaml:"xdpEnabled"`
	// XDPRefreshInterval is the period at which Felix re-checks all XDP state to ensure that no other process has accidentally broken Calico's BPF maps or attached programs.
	//
	// Set to 0 to disable XDP refresh. [Default: 90s]
	XdpRefreshInterval *string `field:"optional" json:"xdpRefreshInterval" yaml:"xdpRefreshInterval"`
}

