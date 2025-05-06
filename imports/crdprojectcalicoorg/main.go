// crdprojectcalicoorg
package crdprojectcalicoorg

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.BgpConfiguration",
		reflect.TypeOf((*BgpConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BgpConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpConfigurationProps",
		reflect.TypeOf((*BgpConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpConfigurationSpec",
		reflect.TypeOf((*BgpConfigurationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpConfigurationSpecCommunities",
		reflect.TypeOf((*BgpConfigurationSpecCommunities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpConfigurationSpecNodeMeshPassword",
		reflect.TypeOf((*BgpConfigurationSpecNodeMeshPassword)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpConfigurationSpecNodeMeshPasswordSecretKeyRef",
		reflect.TypeOf((*BgpConfigurationSpecNodeMeshPasswordSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpConfigurationSpecPrefixAdvertisements",
		reflect.TypeOf((*BgpConfigurationSpecPrefixAdvertisements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpConfigurationSpecServiceClusterIPs",
		reflect.TypeOf((*BgpConfigurationSpecServiceClusterIPs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpConfigurationSpecServiceExternalIPs",
		reflect.TypeOf((*BgpConfigurationSpecServiceExternalIPs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpConfigurationSpecServiceLoadBalancerIPs",
		reflect.TypeOf((*BgpConfigurationSpecServiceLoadBalancerIPs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.BgpPeer",
		reflect.TypeOf((*BgpPeer)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BgpPeer{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpPeerProps",
		reflect.TypeOf((*BgpPeerProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpPeerSpec",
		reflect.TypeOf((*BgpPeerSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpPeerSpecPassword",
		reflect.TypeOf((*BgpPeerSpecPassword)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BgpPeerSpecPasswordSecretKeyRef",
		reflect.TypeOf((*BgpPeerSpecPasswordSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.BlockAffinity",
		reflect.TypeOf((*BlockAffinity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BlockAffinity{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BlockAffinityProps",
		reflect.TypeOf((*BlockAffinityProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.BlockAffinitySpec",
		reflect.TypeOf((*BlockAffinitySpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.CalicoNodeStatus",
		reflect.TypeOf((*CalicoNodeStatus)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CalicoNodeStatus{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.CalicoNodeStatusProps",
		reflect.TypeOf((*CalicoNodeStatusProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.CalicoNodeStatusSpec",
		reflect.TypeOf((*CalicoNodeStatusSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.ClusterInformation",
		reflect.TypeOf((*ClusterInformation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ClusterInformation{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.ClusterInformationProps",
		reflect.TypeOf((*ClusterInformationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.ClusterInformationSpec",
		reflect.TypeOf((*ClusterInformationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.FelixConfiguration",
		reflect.TypeOf((*FelixConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_FelixConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.FelixConfigurationProps",
		reflect.TypeOf((*FelixConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.FelixConfigurationSpec",
		reflect.TypeOf((*FelixConfigurationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"crdprojectcalicoorg.FelixConfigurationSpecAwsSrcDstCheck",
		reflect.TypeOf((*FelixConfigurationSpecAwsSrcDstCheck)(nil)).Elem(),
		map[string]interface{}{
			"DO_NOTHING": FelixConfigurationSpecAwsSrcDstCheck_DO_NOTHING,
			"ENABLE": FelixConfigurationSpecAwsSrcDstCheck_ENABLE,
			"DISABLE": FelixConfigurationSpecAwsSrcDstCheck_DISABLE,
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.FelixConfigurationSpecBpfPsnatPorts",
		reflect.TypeOf((*FelixConfigurationSpecBpfPsnatPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_FelixConfigurationSpecBpfPsnatPorts{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.FelixConfigurationSpecFailsafeInboundHostPorts",
		reflect.TypeOf((*FelixConfigurationSpecFailsafeInboundHostPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.FelixConfigurationSpecFailsafeOutboundHostPorts",
		reflect.TypeOf((*FelixConfigurationSpecFailsafeOutboundHostPorts)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"crdprojectcalicoorg.FelixConfigurationSpecFloatingIPs",
		reflect.TypeOf((*FelixConfigurationSpecFloatingIPs)(nil)).Elem(),
		map[string]interface{}{
			"ENABLED": FelixConfigurationSpecFloatingIPs_ENABLED,
			"DISABLED": FelixConfigurationSpecFloatingIPs_DISABLED,
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.FelixConfigurationSpecKubeNodePortRanges",
		reflect.TypeOf((*FelixConfigurationSpecKubeNodePortRanges)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_FelixConfigurationSpecKubeNodePortRanges{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.FelixConfigurationSpecNatPortRange",
		reflect.TypeOf((*FelixConfigurationSpecNatPortRange)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_FelixConfigurationSpecNatPortRange{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.FelixConfigurationSpecRouteTableRange",
		reflect.TypeOf((*FelixConfigurationSpecRouteTableRange)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.FelixConfigurationSpecRouteTableRanges",
		reflect.TypeOf((*FelixConfigurationSpecRouteTableRanges)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicy",
		reflect.TypeOf((*GlobalNetworkPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GlobalNetworkPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicyProps",
		reflect.TypeOf((*GlobalNetworkPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpec",
		reflect.TypeOf((*GlobalNetworkPolicySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgress",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressDestination",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressDestination)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressDestinationNotPorts",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressDestinationNotPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecEgressDestinationNotPorts{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressDestinationPorts",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressDestinationPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecEgressDestinationPorts{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressDestinationServiceAccounts",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressDestinationServiceAccounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressDestinationServices",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressDestinationServices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressHttp",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressHttpPaths",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressHttpPaths)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressIcmp",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressIcmp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressMetadata",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressNotIcmp",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressNotIcmp)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressNotProtocol",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressNotProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecEgressNotProtocol{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressProtocol",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecEgressProtocol{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressSource",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressSourceNotPorts",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressSourceNotPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecEgressSourceNotPorts{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressSourcePorts",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressSourcePorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecEgressSourcePorts{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressSourceServiceAccounts",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressSourceServiceAccounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecEgressSourceServices",
		reflect.TypeOf((*GlobalNetworkPolicySpecEgressSourceServices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngress",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressDestination",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressDestination)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressDestinationNotPorts",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressDestinationNotPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecIngressDestinationNotPorts{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressDestinationPorts",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressDestinationPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecIngressDestinationPorts{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressDestinationServiceAccounts",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressDestinationServiceAccounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressDestinationServices",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressDestinationServices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressHttp",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressHttpPaths",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressHttpPaths)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressIcmp",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressIcmp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressMetadata",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressNotIcmp",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressNotIcmp)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressNotProtocol",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressNotProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecIngressNotProtocol{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressProtocol",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecIngressProtocol{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressSource",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressSourceNotPorts",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressSourceNotPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecIngressSourceNotPorts{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressSourcePorts",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressSourcePorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_GlobalNetworkPolicySpecIngressSourcePorts{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressSourceServiceAccounts",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressSourceServiceAccounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkPolicySpecIngressSourceServices",
		reflect.TypeOf((*GlobalNetworkPolicySpecIngressSourceServices)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.GlobalNetworkSet",
		reflect.TypeOf((*GlobalNetworkSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GlobalNetworkSet{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkSetProps",
		reflect.TypeOf((*GlobalNetworkSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.GlobalNetworkSetSpec",
		reflect.TypeOf((*GlobalNetworkSetSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.HostEndpoint",
		reflect.TypeOf((*HostEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_HostEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.HostEndpointProps",
		reflect.TypeOf((*HostEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.HostEndpointSpec",
		reflect.TypeOf((*HostEndpointSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.HostEndpointSpecPorts",
		reflect.TypeOf((*HostEndpointSpecPorts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.HostEndpointSpecPortsProtocol",
		reflect.TypeOf((*HostEndpointSpecPortsProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_HostEndpointSpecPortsProtocol{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.IpPool",
		reflect.TypeOf((*IpPool)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpPool{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpPoolProps",
		reflect.TypeOf((*IpPoolProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpPoolSpec",
		reflect.TypeOf((*IpPoolSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpPoolSpecIpip",
		reflect.TypeOf((*IpPoolSpecIpip)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.IpReservation",
		reflect.TypeOf((*IpReservation)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpReservation{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpReservationProps",
		reflect.TypeOf((*IpReservationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpReservationSpec",
		reflect.TypeOf((*IpReservationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.IpamBlock",
		reflect.TypeOf((*IpamBlock)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpamBlock{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpamBlockProps",
		reflect.TypeOf((*IpamBlockProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpamBlockSpec",
		reflect.TypeOf((*IpamBlockSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpamBlockSpecAttributes",
		reflect.TypeOf((*IpamBlockSpecAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.IpamConfig",
		reflect.TypeOf((*IpamConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpamConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpamConfigProps",
		reflect.TypeOf((*IpamConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpamConfigSpec",
		reflect.TypeOf((*IpamConfigSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.IpamHandle",
		reflect.TypeOf((*IpamHandle)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_IpamHandle{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpamHandleProps",
		reflect.TypeOf((*IpamHandleProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.IpamHandleSpec",
		reflect.TypeOf((*IpamHandleSpec)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.KubeControllersConfiguration",
		reflect.TypeOf((*KubeControllersConfiguration)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_KubeControllersConfiguration{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.KubeControllersConfigurationProps",
		reflect.TypeOf((*KubeControllersConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.KubeControllersConfigurationSpec",
		reflect.TypeOf((*KubeControllersConfigurationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.KubeControllersConfigurationSpecControllers",
		reflect.TypeOf((*KubeControllersConfigurationSpecControllers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.KubeControllersConfigurationSpecControllersNamespace",
		reflect.TypeOf((*KubeControllersConfigurationSpecControllersNamespace)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.KubeControllersConfigurationSpecControllersNode",
		reflect.TypeOf((*KubeControllersConfigurationSpecControllersNode)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.KubeControllersConfigurationSpecControllersNodeHostEndpoint",
		reflect.TypeOf((*KubeControllersConfigurationSpecControllersNodeHostEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.KubeControllersConfigurationSpecControllersPolicy",
		reflect.TypeOf((*KubeControllersConfigurationSpecControllersPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.KubeControllersConfigurationSpecControllersServiceAccount",
		reflect.TypeOf((*KubeControllersConfigurationSpecControllersServiceAccount)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.KubeControllersConfigurationSpecControllersWorkloadEndpoint",
		reflect.TypeOf((*KubeControllersConfigurationSpecControllersWorkloadEndpoint)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicy",
		reflect.TypeOf((*NetworkPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicyProps",
		reflect.TypeOf((*NetworkPolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpec",
		reflect.TypeOf((*NetworkPolicySpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgress",
		reflect.TypeOf((*NetworkPolicySpecEgress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressDestination",
		reflect.TypeOf((*NetworkPolicySpecEgressDestination)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecEgressDestinationNotPorts",
		reflect.TypeOf((*NetworkPolicySpecEgressDestinationNotPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecEgressDestinationNotPorts{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecEgressDestinationPorts",
		reflect.TypeOf((*NetworkPolicySpecEgressDestinationPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecEgressDestinationPorts{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressDestinationServiceAccounts",
		reflect.TypeOf((*NetworkPolicySpecEgressDestinationServiceAccounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressDestinationServices",
		reflect.TypeOf((*NetworkPolicySpecEgressDestinationServices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressHttp",
		reflect.TypeOf((*NetworkPolicySpecEgressHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressHttpPaths",
		reflect.TypeOf((*NetworkPolicySpecEgressHttpPaths)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressIcmp",
		reflect.TypeOf((*NetworkPolicySpecEgressIcmp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressMetadata",
		reflect.TypeOf((*NetworkPolicySpecEgressMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressNotIcmp",
		reflect.TypeOf((*NetworkPolicySpecEgressNotIcmp)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecEgressNotProtocol",
		reflect.TypeOf((*NetworkPolicySpecEgressNotProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecEgressNotProtocol{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecEgressProtocol",
		reflect.TypeOf((*NetworkPolicySpecEgressProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecEgressProtocol{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressSource",
		reflect.TypeOf((*NetworkPolicySpecEgressSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecEgressSourceNotPorts",
		reflect.TypeOf((*NetworkPolicySpecEgressSourceNotPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecEgressSourceNotPorts{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecEgressSourcePorts",
		reflect.TypeOf((*NetworkPolicySpecEgressSourcePorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecEgressSourcePorts{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressSourceServiceAccounts",
		reflect.TypeOf((*NetworkPolicySpecEgressSourceServiceAccounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecEgressSourceServices",
		reflect.TypeOf((*NetworkPolicySpecEgressSourceServices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngress",
		reflect.TypeOf((*NetworkPolicySpecIngress)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressDestination",
		reflect.TypeOf((*NetworkPolicySpecIngressDestination)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecIngressDestinationNotPorts",
		reflect.TypeOf((*NetworkPolicySpecIngressDestinationNotPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecIngressDestinationNotPorts{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecIngressDestinationPorts",
		reflect.TypeOf((*NetworkPolicySpecIngressDestinationPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecIngressDestinationPorts{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressDestinationServiceAccounts",
		reflect.TypeOf((*NetworkPolicySpecIngressDestinationServiceAccounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressDestinationServices",
		reflect.TypeOf((*NetworkPolicySpecIngressDestinationServices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressHttp",
		reflect.TypeOf((*NetworkPolicySpecIngressHttp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressHttpPaths",
		reflect.TypeOf((*NetworkPolicySpecIngressHttpPaths)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressIcmp",
		reflect.TypeOf((*NetworkPolicySpecIngressIcmp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressMetadata",
		reflect.TypeOf((*NetworkPolicySpecIngressMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressNotIcmp",
		reflect.TypeOf((*NetworkPolicySpecIngressNotIcmp)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecIngressNotProtocol",
		reflect.TypeOf((*NetworkPolicySpecIngressNotProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecIngressNotProtocol{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecIngressProtocol",
		reflect.TypeOf((*NetworkPolicySpecIngressProtocol)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecIngressProtocol{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressSource",
		reflect.TypeOf((*NetworkPolicySpecIngressSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecIngressSourceNotPorts",
		reflect.TypeOf((*NetworkPolicySpecIngressSourceNotPorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecIngressSourceNotPorts{}
		},
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkPolicySpecIngressSourcePorts",
		reflect.TypeOf((*NetworkPolicySpecIngressSourcePorts)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_NetworkPolicySpecIngressSourcePorts{}
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressSourceServiceAccounts",
		reflect.TypeOf((*NetworkPolicySpecIngressSourceServiceAccounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkPolicySpecIngressSourceServices",
		reflect.TypeOf((*NetworkPolicySpecIngressSourceServices)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"crdprojectcalicoorg.NetworkSet",
		reflect.TypeOf((*NetworkSet)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addJsonPatch", GoMethod: "AddJsonPatch"},
			_jsii_.MemberProperty{JsiiProperty: "apiGroup", GoGetter: "ApiGroup"},
			_jsii_.MemberProperty{JsiiProperty: "apiVersion", GoGetter: "ApiVersion"},
			_jsii_.MemberProperty{JsiiProperty: "chart", GoGetter: "Chart"},
			_jsii_.MemberProperty{JsiiProperty: "kind", GoGetter: "Kind"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toJson", GoMethod: "ToJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NetworkSet{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkSetProps",
		reflect.TypeOf((*NetworkSetProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"crdprojectcalicoorg.NetworkSetSpec",
		reflect.TypeOf((*NetworkSetSpec)(nil)).Elem(),
	)
}
