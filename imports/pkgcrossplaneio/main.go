// pkgcrossplaneio
package pkgcrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"pkgcrossplaneio.Configuration",
		reflect.TypeOf((*Configuration)(nil)).Elem(),
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
			j := jsiiProxy_Configuration{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationProps",
		reflect.TypeOf((*ConfigurationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ConfigurationRevision",
		reflect.TypeOf((*ConfigurationRevision)(nil)).Elem(),
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
			j := jsiiProxy_ConfigurationRevision{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationRevisionProps",
		reflect.TypeOf((*ConfigurationRevisionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationRevisionSpec",
		reflect.TypeOf((*ConfigurationRevisionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationRevisionSpecPackagePullSecrets",
		reflect.TypeOf((*ConfigurationRevisionSpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationSpec",
		reflect.TypeOf((*ConfigurationSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ConfigurationSpecPackagePullSecrets",
		reflect.TypeOf((*ConfigurationSpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfig",
		reflect.TypeOf((*ControllerConfig)(nil)).Elem(),
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
			j := jsiiProxy_ControllerConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigProps",
		reflect.TypeOf((*ControllerConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpec",
		reflect.TypeOf((*ControllerConfigSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinity",
		reflect.TypeOf((*ControllerConfigSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinity",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*ControllerConfigSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinity",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinity",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnv",
		reflect.TypeOf((*ControllerConfigSpecEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvFrom",
		reflect.TypeOf((*ControllerConfigSpecEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvFromConfigMapRef",
		reflect.TypeOf((*ControllerConfigSpecEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvFromSecretRef",
		reflect.TypeOf((*ControllerConfigSpecEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFrom",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromFieldRef",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromResourceFieldRef",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecEnvValueFromSecretKeyRef",
		reflect.TypeOf((*ControllerConfigSpecEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecImagePullSecrets",
		reflect.TypeOf((*ControllerConfigSpecImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecMetadata",
		reflect.TypeOf((*ControllerConfigSpecMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContext",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContextAppArmorProfile",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContextAppArmorProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContextSeLinuxOptions",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContextSeccompProfile",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContextSysctls",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPodSecurityContextWindowsOptions",
		reflect.TypeOf((*ControllerConfigSpecPodSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecPorts",
		reflect.TypeOf((*ControllerConfigSpecPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecResources",
		reflect.TypeOf((*ControllerConfigSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecResourcesClaims",
		reflect.TypeOf((*ControllerConfigSpecResourcesClaims)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecResourcesLimits",
		reflect.TypeOf((*ControllerConfigSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecResourcesRequests",
		reflect.TypeOf((*ControllerConfigSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContext",
		reflect.TypeOf((*ControllerConfigSpecSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContextAppArmorProfile",
		reflect.TypeOf((*ControllerConfigSpecSecurityContextAppArmorProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContextCapabilities",
		reflect.TypeOf((*ControllerConfigSpecSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContextSeLinuxOptions",
		reflect.TypeOf((*ControllerConfigSpecSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContextSeccompProfile",
		reflect.TypeOf((*ControllerConfigSpecSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecSecurityContextWindowsOptions",
		reflect.TypeOf((*ControllerConfigSpecSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecTolerations",
		reflect.TypeOf((*ControllerConfigSpecTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumeMounts",
		reflect.TypeOf((*ControllerConfigSpecVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumes",
		reflect.TypeOf((*ControllerConfigSpecVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesAwsElasticBlockStore",
		reflect.TypeOf((*ControllerConfigSpecVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesAzureDisk",
		reflect.TypeOf((*ControllerConfigSpecVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesAzureFile",
		reflect.TypeOf((*ControllerConfigSpecVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesCephfs",
		reflect.TypeOf((*ControllerConfigSpecVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesCephfsSecretRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesCinder",
		reflect.TypeOf((*ControllerConfigSpecVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesCinderSecretRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesConfigMap",
		reflect.TypeOf((*ControllerConfigSpecVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesConfigMapItems",
		reflect.TypeOf((*ControllerConfigSpecVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesCsi",
		reflect.TypeOf((*ControllerConfigSpecVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesDownwardApi",
		reflect.TypeOf((*ControllerConfigSpecVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesDownwardApiItems",
		reflect.TypeOf((*ControllerConfigSpecVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEmptyDir",
		reflect.TypeOf((*ControllerConfigSpecVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*ControllerConfigSpecVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeral",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateMetadata",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesFc",
		reflect.TypeOf((*ControllerConfigSpecVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesFlexVolume",
		reflect.TypeOf((*ControllerConfigSpecVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesFlocker",
		reflect.TypeOf((*ControllerConfigSpecVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesGcePersistentDisk",
		reflect.TypeOf((*ControllerConfigSpecVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesGitRepo",
		reflect.TypeOf((*ControllerConfigSpecVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesGlusterfs",
		reflect.TypeOf((*ControllerConfigSpecVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesHostPath",
		reflect.TypeOf((*ControllerConfigSpecVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesImage",
		reflect.TypeOf((*ControllerConfigSpecVolumesImage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesIscsi",
		reflect.TypeOf((*ControllerConfigSpecVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesIscsiSecretRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesNfs",
		reflect.TypeOf((*ControllerConfigSpecVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesPersistentVolumeClaim",
		reflect.TypeOf((*ControllerConfigSpecVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesPhotonPersistentDisk",
		reflect.TypeOf((*ControllerConfigSpecVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesPortworxVolume",
		reflect.TypeOf((*ControllerConfigSpecVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjected",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSources",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesClusterTrustBundle",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesClusterTrustBundle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesClusterTrustBundleLabelSelector",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesClusterTrustBundleLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesClusterTrustBundleLabelSelectorMatchExpressions",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesClusterTrustBundleLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_ControllerConfigSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesSecret",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*ControllerConfigSpecVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesQuobyte",
		reflect.TypeOf((*ControllerConfigSpecVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesRbd",
		reflect.TypeOf((*ControllerConfigSpecVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesRbdSecretRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesScaleIo",
		reflect.TypeOf((*ControllerConfigSpecVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesScaleIoSecretRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesSecret",
		reflect.TypeOf((*ControllerConfigSpecVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesSecretItems",
		reflect.TypeOf((*ControllerConfigSpecVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesStorageos",
		reflect.TypeOf((*ControllerConfigSpecVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesStorageosSecretRef",
		reflect.TypeOf((*ControllerConfigSpecVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ControllerConfigSpecVolumesVsphereVolume",
		reflect.TypeOf((*ControllerConfigSpecVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfig",
		reflect.TypeOf((*DeploymentRuntimeConfig)(nil)).Elem(),
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
			j := jsiiProxy_DeploymentRuntimeConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigProps",
		reflect.TypeOf((*DeploymentRuntimeConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplate",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateMetadata",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategy",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdate",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxSurge",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxSurge)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxSurge{}
		},
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecStrategyRollingUpdateMaxUnavailable{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplate",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateMetadata",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinity",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinity",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreference)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTerms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsMatchFields)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinity",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinity",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTerm)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecution)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionNamespaceSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainers",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnv",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvFrom",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvFromConfigMapRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvFromSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFrom",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromFieldRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecycle",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStart",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartSleep",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStop",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopSleep",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbe",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeGrpc",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersPorts",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbe",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeGrpc",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResizePolicy",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResizePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResources",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesClaims",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesClaims)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesLimits",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContext",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContextAppArmorProfile",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContextAppArmorProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContextCapabilities",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbe",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeGrpc",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersVolumeDevices",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersVolumeMounts",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecDnsConfig",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecDnsConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecDnsConfigOptions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecDnsConfigOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainers",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnv",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvFrom",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvFromConfigMapRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvFromSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFrom",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromFieldRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecycle",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStart",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartSleep",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStop",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopSleep",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbe",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeGrpc",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersPorts",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbe",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeGrpc",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResizePolicy",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResizePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResources",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesClaims",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesClaims)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesLimits",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesRequests",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContext",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContextAppArmorProfile",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContextAppArmorProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContextCapabilities",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbe",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeGrpc",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersVolumeDevices",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersVolumeMounts",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecHostAliases",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecHostAliases)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecImagePullSecrets",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecImagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainers",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainers)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnv",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnv)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvFrom",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvFromConfigMapRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvFromConfigMapRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvFromSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvFromSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFrom",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFrom)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromConfigMapKeyRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromConfigMapKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromFieldRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromResourceFieldRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromResourceFieldRefDivisor",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromSecretKeyRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersEnvValueFromSecretKeyRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecycle",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecycle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStart",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStart)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartSleep",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePostStartTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStop",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStop)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopSleep",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopSleep)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLifecyclePreStopTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbe",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeGrpc",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersLivenessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersPorts",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersPorts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbe",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeGrpc",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersReadinessProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResizePolicy",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResizePolicy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResources",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResourcesClaims",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResourcesClaims)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResourcesLimits",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResourcesRequests",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContext",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContextAppArmorProfile",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContextAppArmorProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContextCapabilities",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContextCapabilities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContextSeLinuxOptions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContextSeccompProfile",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContextWindowsOptions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbe",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbe)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeExec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeExec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeGrpc",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeGrpc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeHttpGet",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeHttpGet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeHttpGetHttpHeaders",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeHttpGetHttpHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeHttpGetPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeHttpGetPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeHttpGetPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeTcpSocket",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeTcpSocket)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeTcpSocketPort",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeTcpSocketPort)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersStartupProbeTcpSocketPort{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersVolumeDevices",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersVolumeDevices)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersVolumeMounts",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecInitContainersVolumeMounts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecOs",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecOs)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecOverhead",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecOverhead)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecOverhead{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecReadinessGates",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecReadinessGates)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecResourceClaims",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecResourceClaims)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSchedulingGates",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSchedulingGates)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContext",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContextAppArmorProfile",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContextAppArmorProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContextSeLinuxOptions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContextSeLinuxOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContextSeccompProfile",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContextSeccompProfile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContextSysctls",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContextSysctls)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContextWindowsOptions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecSecurityContextWindowsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecTolerations",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecTolerations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecTopologySpreadConstraints",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecTopologySpreadConstraints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecTopologySpreadConstraintsLabelSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecTopologySpreadConstraintsLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecTopologySpreadConstraintsLabelSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecTopologySpreadConstraintsLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumes",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesAwsElasticBlockStore",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesAwsElasticBlockStore)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesAzureDisk",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesAzureDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesAzureFile",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesAzureFile)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCephfs",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCephfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCephfsSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCephfsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCinder",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCinder)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCinderSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCinderSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesConfigMap",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesConfigMapItems",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCsi",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCsiNodePublishSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesCsiNodePublishSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApi",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApiItems",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApiItemsFieldRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEmptyDir",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEmptyDir)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEmptyDirSizeLimit",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEmptyDirSizeLimit)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEmptyDirSizeLimit{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeral",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeral)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplate",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateMetadata",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpec",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecDataSourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResources",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesLimits{}
		},
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecResourcesRequests{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesEphemeralVolumeClaimTemplateSpecSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesFc",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesFc)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesFlexVolume",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesFlexVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesFlexVolumeSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesFlexVolumeSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesFlocker",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesFlocker)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesGcePersistentDisk",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesGcePersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesGitRepo",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesGitRepo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesGlusterfs",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesGlusterfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesHostPath",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesHostPath)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesImage",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesImage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesIscsi",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesIscsi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesIscsiSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesIscsiSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesNfs",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesNfs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesPersistentVolumeClaim",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesPersistentVolumeClaim)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesPhotonPersistentDisk",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesPhotonPersistentDisk)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesPortworxVolume",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesPortworxVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjected",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSources",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundle",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundle)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundleLabelSelector",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundleLabelSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundleLabelSelectorMatchExpressions",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesClusterTrustBundleLabelSelectorMatchExpressions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesConfigMap",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesConfigMap)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesConfigMapItems",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesConfigMapItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApi",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApi)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApiItems",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApiItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsFieldRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesDownwardApiItemsResourceFieldRefDivisor{}
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesSecret",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesSecretItems",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesServiceAccountToken",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesProjectedSourcesServiceAccountToken)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesQuobyte",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesQuobyte)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesRbd",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesRbd)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesRbdSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesRbdSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesScaleIo",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesScaleIo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesScaleIoSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesScaleIoSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesSecret",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesSecret)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesSecretItems",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesSecretItems)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesStorageos",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesStorageos)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesStorageosSecretRef",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesStorageosSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesVsphereVolume",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecVolumesVsphereVolume)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecServiceAccountTemplate",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecServiceAccountTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecServiceAccountTemplateMetadata",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecServiceAccountTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecServiceTemplate",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecServiceTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.DeploymentRuntimeConfigSpecServiceTemplateMetadata",
		reflect.TypeOf((*DeploymentRuntimeConfigSpecServiceTemplateMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.Function",
		reflect.TypeOf((*Function)(nil)).Elem(),
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
			j := jsiiProxy_Function{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionProps",
		reflect.TypeOf((*FunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.FunctionRevision",
		reflect.TypeOf((*FunctionRevision)(nil)).Elem(),
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
			j := jsiiProxy_FunctionRevision{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionRevisionProps",
		reflect.TypeOf((*FunctionRevisionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionRevisionSpec",
		reflect.TypeOf((*FunctionRevisionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionRevisionSpecControllerConfigRef",
		reflect.TypeOf((*FunctionRevisionSpecControllerConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionRevisionSpecPackagePullSecrets",
		reflect.TypeOf((*FunctionRevisionSpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionRevisionSpecRuntimeConfigRef",
		reflect.TypeOf((*FunctionRevisionSpecRuntimeConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.FunctionRevisionV1Beta1",
		reflect.TypeOf((*FunctionRevisionV1Beta1)(nil)).Elem(),
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
			j := jsiiProxy_FunctionRevisionV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionRevisionV1Beta1Props",
		reflect.TypeOf((*FunctionRevisionV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionRevisionV1Beta1Spec",
		reflect.TypeOf((*FunctionRevisionV1Beta1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionRevisionV1Beta1SpecControllerConfigRef",
		reflect.TypeOf((*FunctionRevisionV1Beta1SpecControllerConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionRevisionV1Beta1SpecPackagePullSecrets",
		reflect.TypeOf((*FunctionRevisionV1Beta1SpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionRevisionV1Beta1SpecRuntimeConfigRef",
		reflect.TypeOf((*FunctionRevisionV1Beta1SpecRuntimeConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionSpec",
		reflect.TypeOf((*FunctionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionSpecControllerConfigRef",
		reflect.TypeOf((*FunctionSpecControllerConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionSpecPackagePullSecrets",
		reflect.TypeOf((*FunctionSpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionSpecRuntimeConfigRef",
		reflect.TypeOf((*FunctionSpecRuntimeConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.FunctionV1Beta1",
		reflect.TypeOf((*FunctionV1Beta1)(nil)).Elem(),
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
			j := jsiiProxy_FunctionV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionV1Beta1Props",
		reflect.TypeOf((*FunctionV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionV1Beta1Spec",
		reflect.TypeOf((*FunctionV1Beta1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionV1Beta1SpecControllerConfigRef",
		reflect.TypeOf((*FunctionV1Beta1SpecControllerConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionV1Beta1SpecPackagePullSecrets",
		reflect.TypeOf((*FunctionV1Beta1SpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.FunctionV1Beta1SpecRuntimeConfigRef",
		reflect.TypeOf((*FunctionV1Beta1SpecRuntimeConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ImageConfig",
		reflect.TypeOf((*ImageConfig)(nil)).Elem(),
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
			j := jsiiProxy_ImageConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigProps",
		reflect.TypeOf((*ImageConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpec",
		reflect.TypeOf((*ImageConfigSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecMatchImages",
		reflect.TypeOf((*ImageConfigSpecMatchImages)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"pkgcrossplaneio.ImageConfigSpecMatchImagesType",
		reflect.TypeOf((*ImageConfigSpecMatchImagesType)(nil)).Elem(),
		map[string]interface{}{
			"PREFIX": ImageConfigSpecMatchImagesType_PREFIX,
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecRegistry",
		reflect.TypeOf((*ImageConfigSpecRegistry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecRegistryAuthentication",
		reflect.TypeOf((*ImageConfigSpecRegistryAuthentication)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecRegistryAuthenticationPullSecretRef",
		reflect.TypeOf((*ImageConfigSpecRegistryAuthenticationPullSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecVerification",
		reflect.TypeOf((*ImageConfigSpecVerification)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecVerificationCosign",
		reflect.TypeOf((*ImageConfigSpecVerificationCosign)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecVerificationCosignAuthorities",
		reflect.TypeOf((*ImageConfigSpecVerificationCosignAuthorities)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecVerificationCosignAuthoritiesAttestations",
		reflect.TypeOf((*ImageConfigSpecVerificationCosignAuthoritiesAttestations)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecVerificationCosignAuthoritiesKey",
		reflect.TypeOf((*ImageConfigSpecVerificationCosignAuthoritiesKey)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecVerificationCosignAuthoritiesKeySecretRef",
		reflect.TypeOf((*ImageConfigSpecVerificationCosignAuthoritiesKeySecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecVerificationCosignAuthoritiesKeyless",
		reflect.TypeOf((*ImageConfigSpecVerificationCosignAuthoritiesKeyless)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ImageConfigSpecVerificationCosignAuthoritiesKeylessIdentities",
		reflect.TypeOf((*ImageConfigSpecVerificationCosignAuthoritiesKeylessIdentities)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"pkgcrossplaneio.ImageConfigSpecVerificationProvider",
		reflect.TypeOf((*ImageConfigSpecVerificationProvider)(nil)).Elem(),
		map[string]interface{}{
			"COSIGN": ImageConfigSpecVerificationProvider_COSIGN,
		},
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.Lock",
		reflect.TypeOf((*Lock)(nil)).Elem(),
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
			j := jsiiProxy_Lock{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.LockPackages",
		reflect.TypeOf((*LockPackages)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.LockPackagesDependencies",
		reflect.TypeOf((*LockPackagesDependencies)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"pkgcrossplaneio.LockPackagesDependenciesType",
		reflect.TypeOf((*LockPackagesDependenciesType)(nil)).Elem(),
		map[string]interface{}{
			"CONFIGURATION": LockPackagesDependenciesType_CONFIGURATION,
			"PROVIDER": LockPackagesDependenciesType_PROVIDER,
			"FUNCTION": LockPackagesDependenciesType_FUNCTION,
		},
	)
	_jsii_.RegisterEnum(
		"pkgcrossplaneio.LockPackagesType",
		reflect.TypeOf((*LockPackagesType)(nil)).Elem(),
		map[string]interface{}{
			"CONFIGURATION": LockPackagesType_CONFIGURATION,
			"PROVIDER": LockPackagesType_PROVIDER,
			"FUNCTION": LockPackagesType_FUNCTION,
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.LockProps",
		reflect.TypeOf((*LockProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.Provider",
		reflect.TypeOf((*Provider)(nil)).Elem(),
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
			j := jsiiProxy_Provider{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderProps",
		reflect.TypeOf((*ProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"pkgcrossplaneio.ProviderRevision",
		reflect.TypeOf((*ProviderRevision)(nil)).Elem(),
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
			j := jsiiProxy_ProviderRevision{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderRevisionProps",
		reflect.TypeOf((*ProviderRevisionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderRevisionSpec",
		reflect.TypeOf((*ProviderRevisionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderRevisionSpecControllerConfigRef",
		reflect.TypeOf((*ProviderRevisionSpecControllerConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderRevisionSpecPackagePullSecrets",
		reflect.TypeOf((*ProviderRevisionSpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderRevisionSpecRuntimeConfigRef",
		reflect.TypeOf((*ProviderRevisionSpecRuntimeConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderSpec",
		reflect.TypeOf((*ProviderSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderSpecControllerConfigRef",
		reflect.TypeOf((*ProviderSpecControllerConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderSpecPackagePullSecrets",
		reflect.TypeOf((*ProviderSpecPackagePullSecrets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"pkgcrossplaneio.ProviderSpecRuntimeConfigRef",
		reflect.TypeOf((*ProviderSpecRuntimeConfigRef)(nil)).Elem(),
	)
}
