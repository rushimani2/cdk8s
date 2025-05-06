// apiextensionscrossplaneio
package apiextensionscrossplaneio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.CompositeResourceDefinition",
		reflect.TypeOf((*CompositeResourceDefinition)(nil)).Elem(),
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
			j := jsiiProxy_CompositeResourceDefinition{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionProps",
		reflect.TypeOf((*CompositeResourceDefinitionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpec",
		reflect.TypeOf((*CompositeResourceDefinitionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecClaimNames",
		reflect.TypeOf((*CompositeResourceDefinitionSpecClaimNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecConversion",
		reflect.TypeOf((*CompositeResourceDefinitionSpecConversion)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecConversionWebhook",
		reflect.TypeOf((*CompositeResourceDefinitionSpecConversionWebhook)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecConversionWebhookClientConfig",
		reflect.TypeOf((*CompositeResourceDefinitionSpecConversionWebhookClientConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecConversionWebhookClientConfigService",
		reflect.TypeOf((*CompositeResourceDefinitionSpecConversionWebhookClientConfigService)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecDefaultCompositeDeletePolicy",
		reflect.TypeOf((*CompositeResourceDefinitionSpecDefaultCompositeDeletePolicy)(nil)).Elem(),
		map[string]interface{}{
			"BACKGROUND": CompositeResourceDefinitionSpecDefaultCompositeDeletePolicy_BACKGROUND,
			"FOREGROUND": CompositeResourceDefinitionSpecDefaultCompositeDeletePolicy_FOREGROUND,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecDefaultCompositionRef",
		reflect.TypeOf((*CompositeResourceDefinitionSpecDefaultCompositionRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecDefaultCompositionUpdatePolicy",
		reflect.TypeOf((*CompositeResourceDefinitionSpecDefaultCompositionUpdatePolicy)(nil)).Elem(),
		map[string]interface{}{
			"AUTOMATIC": CompositeResourceDefinitionSpecDefaultCompositionUpdatePolicy_AUTOMATIC,
			"MANUAL": CompositeResourceDefinitionSpecDefaultCompositionUpdatePolicy_MANUAL,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecEnforcedCompositionRef",
		reflect.TypeOf((*CompositeResourceDefinitionSpecEnforcedCompositionRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecMetadata",
		reflect.TypeOf((*CompositeResourceDefinitionSpecMetadata)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecNames",
		reflect.TypeOf((*CompositeResourceDefinitionSpecNames)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecVersions",
		reflect.TypeOf((*CompositeResourceDefinitionSpecVersions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecVersionsAdditionalPrinterColumns",
		reflect.TypeOf((*CompositeResourceDefinitionSpecVersionsAdditionalPrinterColumns)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositeResourceDefinitionSpecVersionsSchema",
		reflect.TypeOf((*CompositeResourceDefinitionSpecVersionsSchema)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.Composition",
		reflect.TypeOf((*Composition)(nil)).Elem(),
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
			j := jsiiProxy_Composition{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionProps",
		reflect.TypeOf((*CompositionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.CompositionRevision",
		reflect.TypeOf((*CompositionRevision)(nil)).Elem(),
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
			j := jsiiProxy_CompositionRevision{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionProps",
		reflect.TypeOf((*CompositionRevisionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpec",
		reflect.TypeOf((*CompositionRevisionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecCompositeTypeRef",
		reflect.TypeOf((*CompositionRevisionSpecCompositeTypeRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecMode",
		reflect.TypeOf((*CompositionRevisionSpecMode)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCES": CompositionRevisionSpecMode_RESOURCES,
			"PIPELINE": CompositionRevisionSpecMode_PIPELINE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSets",
		reflect.TypeOf((*CompositionRevisionSpecPatchSets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatches",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesCombine",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesCombine)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesCombineStrategy",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesCombineStrategy)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionSpecPatchSetsPatchesCombineStrategy_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesCombineString",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesCombineString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesCombineVariables",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesCombineVariables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesPolicy",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath_OPTIONAL,
			"REQUIRED": CompositionRevisionSpecPatchSetsPatchesPolicyFromFieldPath_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesPolicyMergeOptions",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesPolicyMergeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransforms",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransforms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsConvert",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsConvert)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat_NONE,
			"QUANTITY": CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat_QUANTITY,
			"JSON": CompositionRevisionSpecPatchSetsPatchesTransformsConvertFormat_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_STRING,
			"INT": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_INT,
			"INT64": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_INT64,
			"BOOL": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_BOOL,
			"FLOAT64": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_FLOAT64,
			"OBJECT": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_OBJECT,
			"ARRAY": CompositionRevisionSpecPatchSetsPatchesTransformsConvertToType_ARRAY,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsMatch",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsMatch)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsMatchFallbackTo",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsMatchFallbackTo)(nil)).Elem(),
		map[string]interface{}{
			"VALUE": CompositionRevisionSpecPatchSetsPatchesTransformsMatchFallbackTo_VALUE,
			"INPUT": CompositionRevisionSpecPatchSetsPatchesTransformsMatchFallbackTo_INPUT,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsMatchPatterns",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsMatchPatterns)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsMatchPatternsType",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsMatchPatternsType)(nil)).Elem(),
		map[string]interface{}{
			"LITERAL": CompositionRevisionSpecPatchSetsPatchesTransformsMatchPatternsType_LITERAL,
			"REGEXP": CompositionRevisionSpecPatchSetsPatchesTransformsMatchPatternsType_REGEXP,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsMath",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsMath)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsMathType",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsMathType)(nil)).Elem(),
		map[string]interface{}{
			"MULTIPLY": CompositionRevisionSpecPatchSetsPatchesTransformsMathType_MULTIPLY,
			"CLAMP_MIN": CompositionRevisionSpecPatchSetsPatchesTransformsMathType_CLAMP_MIN,
			"CLAMP_MAX": CompositionRevisionSpecPatchSetsPatchesTransformsMathType_CLAMP_MAX,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsString",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsString)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert)(nil)).Elem(),
		map[string]interface{}{
			"TO_UPPER": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_UPPER,
			"TO_LOWER": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_LOWER,
			"TO_BASE64": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_BASE64,
			"FROM_BASE64": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_FROM_BASE64,
			"TO_JSON": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_JSON,
			"TO_SHA1": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_SHA1,
			"TO_SHA256": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_SHA256,
			"TO_SHA512": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_SHA512,
			"TO_ADLER32": CompositionRevisionSpecPatchSetsPatchesTransformsStringConvert_TO_ADLER32,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsStringJoin",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsStringJoin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsStringRegexp",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsStringRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsStringType",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsStringType)(nil)).Elem(),
		map[string]interface{}{
			"FORMAT": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_FORMAT,
			"CONVERT": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_CONVERT,
			"TRIM_PREFIX": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_TRIM_PREFIX,
			"TRIM_SUFFIX": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_TRIM_SUFFIX,
			"REGEXP": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_REGEXP,
			"JOIN": CompositionRevisionSpecPatchSetsPatchesTransformsStringType_JOIN,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesTransformsType",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesTransformsType)(nil)).Elem(),
		map[string]interface{}{
			"MAP": CompositionRevisionSpecPatchSetsPatchesTransformsType_MAP,
			"MATCH": CompositionRevisionSpecPatchSetsPatchesTransformsType_MATCH,
			"MATH": CompositionRevisionSpecPatchSetsPatchesTransformsType_MATH,
			"STRING": CompositionRevisionSpecPatchSetsPatchesTransformsType_STRING,
			"CONVERT": CompositionRevisionSpecPatchSetsPatchesTransformsType_CONVERT,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPatchSetsPatchesType",
		reflect.TypeOf((*CompositionRevisionSpecPatchSetsPatchesType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_COMPOSITE_FIELD_PATH": CompositionRevisionSpecPatchSetsPatchesType_FROM_COMPOSITE_FIELD_PATH,
			"PATCH_SET": CompositionRevisionSpecPatchSetsPatchesType_PATCH_SET,
			"TO_COMPOSITE_FIELD_PATH": CompositionRevisionSpecPatchSetsPatchesType_TO_COMPOSITE_FIELD_PATH,
			"COMBINE_FROM_COMPOSITE": CompositionRevisionSpecPatchSetsPatchesType_COMBINE_FROM_COMPOSITE,
			"COMBINE_TO_COMPOSITE": CompositionRevisionSpecPatchSetsPatchesType_COMBINE_TO_COMPOSITE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPipeline",
		reflect.TypeOf((*CompositionRevisionSpecPipeline)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPipelineCredentials",
		reflect.TypeOf((*CompositionRevisionSpecPipelineCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPipelineCredentialsSecretRef",
		reflect.TypeOf((*CompositionRevisionSpecPipelineCredentialsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecPipelineCredentialsSource",
		reflect.TypeOf((*CompositionRevisionSpecPipelineCredentialsSource)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompositionRevisionSpecPipelineCredentialsSource_NONE,
			"SECRET": CompositionRevisionSpecPipelineCredentialsSource_SECRET,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPipelineFunctionRef",
		reflect.TypeOf((*CompositionRevisionSpecPipelineFunctionRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecPublishConnectionDetailsWithStoreConfigRef",
		reflect.TypeOf((*CompositionRevisionSpecPublishConnectionDetailsWithStoreConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResources",
		reflect.TypeOf((*CompositionRevisionSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesConnectionDetails",
		reflect.TypeOf((*CompositionRevisionSpecResourcesConnectionDetails)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesConnectionDetailsType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesConnectionDetailsType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_CONNECTION_SECRET_KEY": CompositionRevisionSpecResourcesConnectionDetailsType_FROM_CONNECTION_SECRET_KEY,
			"FROM_FIELD_PATH": CompositionRevisionSpecResourcesConnectionDetailsType_FROM_FIELD_PATH,
			"FROM_VALUE": CompositionRevisionSpecResourcesConnectionDetailsType_FROM_VALUE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatches",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesCombine",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesCombine)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesCombineStrategy",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesCombineStrategy)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionSpecResourcesPatchesCombineStrategy_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesCombineString",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesCombineString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesCombineVariables",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesCombineVariables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesPolicy",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesPolicyFromFieldPath",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesPolicyFromFieldPath)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": CompositionRevisionSpecResourcesPatchesPolicyFromFieldPath_OPTIONAL,
			"REQUIRED": CompositionRevisionSpecResourcesPatchesPolicyFromFieldPath_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesPolicyMergeOptions",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesPolicyMergeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransforms",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransforms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsConvert",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsConvert)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsConvertFormat",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsConvertFormat)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompositionRevisionSpecResourcesPatchesTransformsConvertFormat_NONE,
			"QUANTITY": CompositionRevisionSpecResourcesPatchesTransformsConvertFormat_QUANTITY,
			"JSON": CompositionRevisionSpecResourcesPatchesTransformsConvertFormat_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsConvertToType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsConvertToType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_STRING,
			"INT": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_INT,
			"INT64": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_INT64,
			"BOOL": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_BOOL,
			"FLOAT64": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_FLOAT64,
			"OBJECT": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_OBJECT,
			"ARRAY": CompositionRevisionSpecResourcesPatchesTransformsConvertToType_ARRAY,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsMatch",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsMatch)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsMatchFallbackTo",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsMatchFallbackTo)(nil)).Elem(),
		map[string]interface{}{
			"VALUE": CompositionRevisionSpecResourcesPatchesTransformsMatchFallbackTo_VALUE,
			"INPUT": CompositionRevisionSpecResourcesPatchesTransformsMatchFallbackTo_INPUT,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsMatchPatterns",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsMatchPatterns)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsMatchPatternsType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsMatchPatternsType)(nil)).Elem(),
		map[string]interface{}{
			"LITERAL": CompositionRevisionSpecResourcesPatchesTransformsMatchPatternsType_LITERAL,
			"REGEXP": CompositionRevisionSpecResourcesPatchesTransformsMatchPatternsType_REGEXP,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsMath",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsMath)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsMathType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsMathType)(nil)).Elem(),
		map[string]interface{}{
			"MULTIPLY": CompositionRevisionSpecResourcesPatchesTransformsMathType_MULTIPLY,
			"CLAMP_MIN": CompositionRevisionSpecResourcesPatchesTransformsMathType_CLAMP_MIN,
			"CLAMP_MAX": CompositionRevisionSpecResourcesPatchesTransformsMathType_CLAMP_MAX,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsString",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsString)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsStringConvert",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsStringConvert)(nil)).Elem(),
		map[string]interface{}{
			"TO_UPPER": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_UPPER,
			"TO_LOWER": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_LOWER,
			"TO_BASE64": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_BASE64,
			"FROM_BASE64": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_FROM_BASE64,
			"TO_JSON": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_JSON,
			"TO_SHA1": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_SHA1,
			"TO_SHA256": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_SHA256,
			"TO_SHA512": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_SHA512,
			"TO_ADLER32": CompositionRevisionSpecResourcesPatchesTransformsStringConvert_TO_ADLER32,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsStringJoin",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsStringJoin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsStringRegexp",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsStringRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsStringType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsStringType)(nil)).Elem(),
		map[string]interface{}{
			"FORMAT": CompositionRevisionSpecResourcesPatchesTransformsStringType_FORMAT,
			"CONVERT": CompositionRevisionSpecResourcesPatchesTransformsStringType_CONVERT,
			"TRIM_PREFIX": CompositionRevisionSpecResourcesPatchesTransformsStringType_TRIM_PREFIX,
			"TRIM_SUFFIX": CompositionRevisionSpecResourcesPatchesTransformsStringType_TRIM_SUFFIX,
			"REGEXP": CompositionRevisionSpecResourcesPatchesTransformsStringType_REGEXP,
			"JOIN": CompositionRevisionSpecResourcesPatchesTransformsStringType_JOIN,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesTransformsType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesTransformsType)(nil)).Elem(),
		map[string]interface{}{
			"MAP": CompositionRevisionSpecResourcesPatchesTransformsType_MAP,
			"MATCH": CompositionRevisionSpecResourcesPatchesTransformsType_MATCH,
			"MATH": CompositionRevisionSpecResourcesPatchesTransformsType_MATH,
			"STRING": CompositionRevisionSpecResourcesPatchesTransformsType_STRING,
			"CONVERT": CompositionRevisionSpecResourcesPatchesTransformsType_CONVERT,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesPatchesType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesPatchesType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_COMPOSITE_FIELD_PATH": CompositionRevisionSpecResourcesPatchesType_FROM_COMPOSITE_FIELD_PATH,
			"PATCH_SET": CompositionRevisionSpecResourcesPatchesType_PATCH_SET,
			"TO_COMPOSITE_FIELD_PATH": CompositionRevisionSpecResourcesPatchesType_TO_COMPOSITE_FIELD_PATH,
			"COMBINE_FROM_COMPOSITE": CompositionRevisionSpecResourcesPatchesType_COMBINE_FROM_COMPOSITE,
			"COMBINE_TO_COMPOSITE": CompositionRevisionSpecResourcesPatchesType_COMBINE_TO_COMPOSITE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesReadinessChecks",
		reflect.TypeOf((*CompositionRevisionSpecResourcesReadinessChecks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesReadinessChecksMatchCondition",
		reflect.TypeOf((*CompositionRevisionSpecResourcesReadinessChecksMatchCondition)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionSpecResourcesReadinessChecksType",
		reflect.TypeOf((*CompositionRevisionSpecResourcesReadinessChecksType)(nil)).Elem(),
		map[string]interface{}{
			"MATCH_STRING": CompositionRevisionSpecResourcesReadinessChecksType_MATCH_STRING,
			"MATCH_INTEGER": CompositionRevisionSpecResourcesReadinessChecksType_MATCH_INTEGER,
			"NON_EMPTY": CompositionRevisionSpecResourcesReadinessChecksType_NON_EMPTY,
			"MATCH_CONDITION": CompositionRevisionSpecResourcesReadinessChecksType_MATCH_CONDITION,
			"MATCH_TRUE": CompositionRevisionSpecResourcesReadinessChecksType_MATCH_TRUE,
			"MATCH_FALSE": CompositionRevisionSpecResourcesReadinessChecksType_MATCH_FALSE,
			"NONE": CompositionRevisionSpecResourcesReadinessChecksType_NONE,
		},
	)
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1",
		reflect.TypeOf((*CompositionRevisionV1Beta1)(nil)).Elem(),
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
			j := jsiiProxy_CompositionRevisionV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1Props",
		reflect.TypeOf((*CompositionRevisionV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1Spec",
		reflect.TypeOf((*CompositionRevisionV1Beta1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecCompositeTypeRef",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecCompositeTypeRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecMode",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecMode)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCES": CompositionRevisionV1Beta1SpecMode_RESOURCES,
			"PIPELINE": CompositionRevisionV1Beta1SpecMode_PIPELINE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSets",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatches",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesCombine",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesCombine)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesCombineStrategy",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesCombineStrategy)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionV1Beta1SpecPatchSetsPatchesCombineStrategy_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesCombineString",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesCombineString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesCombineVariables",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesCombineVariables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesPolicy",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesPolicyFromFieldPath",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesPolicyFromFieldPath)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": CompositionRevisionV1Beta1SpecPatchSetsPatchesPolicyFromFieldPath_OPTIONAL,
			"REQUIRED": CompositionRevisionV1Beta1SpecPatchSetsPatchesPolicyFromFieldPath_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesPolicyMergeOptions",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesPolicyMergeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransforms",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransforms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvert",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvert)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat_NONE,
			"QUANTITY": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat_QUANTITY,
			"JSON": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertFormat_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertToType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertToType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertToType_STRING,
			"INT": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertToType_INT,
			"INT64": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertToType_INT64,
			"BOOL": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertToType_BOOL,
			"FLOAT64": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertToType_FLOAT64,
			"OBJECT": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertToType_OBJECT,
			"ARRAY": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsConvertToType_ARRAY,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatch",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatch)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchFallbackTo",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchFallbackTo)(nil)).Elem(),
		map[string]interface{}{
			"VALUE": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchFallbackTo_VALUE,
			"INPUT": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchFallbackTo_INPUT,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchPatterns",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchPatterns)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchPatternsType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchPatternsType)(nil)).Elem(),
		map[string]interface{}{
			"LITERAL": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchPatternsType_LITERAL,
			"REGEXP": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMatchPatternsType_REGEXP,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMath",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMath)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMathType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMathType)(nil)).Elem(),
		map[string]interface{}{
			"MULTIPLY": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMathType_MULTIPLY,
			"CLAMP_MIN": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMathType_CLAMP_MIN,
			"CLAMP_MAX": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsMathType_CLAMP_MAX,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsString",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsString)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert)(nil)).Elem(),
		map[string]interface{}{
			"TO_UPPER": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_UPPER,
			"TO_LOWER": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_LOWER,
			"TO_BASE64": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_BASE64,
			"FROM_BASE64": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_FROM_BASE64,
			"TO_JSON": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_JSON,
			"TO_SHA1": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_SHA1,
			"TO_SHA256": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_SHA256,
			"TO_SHA512": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_SHA512,
			"TO_ADLER32": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringConvert_TO_ADLER32,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringJoin",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringJoin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringRegexp",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringType)(nil)).Elem(),
		map[string]interface{}{
			"FORMAT": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringType_FORMAT,
			"CONVERT": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringType_CONVERT,
			"TRIM_PREFIX": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringType_TRIM_PREFIX,
			"TRIM_SUFFIX": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringType_TRIM_SUFFIX,
			"REGEXP": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringType_REGEXP,
			"JOIN": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsStringType_JOIN,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsType)(nil)).Elem(),
		map[string]interface{}{
			"MAP": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsType_MAP,
			"MATCH": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsType_MATCH,
			"MATH": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsType_MATH,
			"STRING": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsType_STRING,
			"CONVERT": CompositionRevisionV1Beta1SpecPatchSetsPatchesTransformsType_CONVERT,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPatchSetsPatchesType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPatchSetsPatchesType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_COMPOSITE_FIELD_PATH": CompositionRevisionV1Beta1SpecPatchSetsPatchesType_FROM_COMPOSITE_FIELD_PATH,
			"PATCH_SET": CompositionRevisionV1Beta1SpecPatchSetsPatchesType_PATCH_SET,
			"TO_COMPOSITE_FIELD_PATH": CompositionRevisionV1Beta1SpecPatchSetsPatchesType_TO_COMPOSITE_FIELD_PATH,
			"COMBINE_FROM_COMPOSITE": CompositionRevisionV1Beta1SpecPatchSetsPatchesType_COMBINE_FROM_COMPOSITE,
			"COMBINE_TO_COMPOSITE": CompositionRevisionV1Beta1SpecPatchSetsPatchesType_COMBINE_TO_COMPOSITE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPipeline",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPipeline)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPipelineCredentials",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPipelineCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPipelineCredentialsSecretRef",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPipelineCredentialsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPipelineCredentialsSource",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPipelineCredentialsSource)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompositionRevisionV1Beta1SpecPipelineCredentialsSource_NONE,
			"SECRET": CompositionRevisionV1Beta1SpecPipelineCredentialsSource_SECRET,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPipelineFunctionRef",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPipelineFunctionRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecPublishConnectionDetailsWithStoreConfigRef",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecPublishConnectionDetailsWithStoreConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResources",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesConnectionDetails",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesConnectionDetails)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_CONNECTION_SECRET_KEY": CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType_FROM_CONNECTION_SECRET_KEY,
			"FROM_FIELD_PATH": CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType_FROM_FIELD_PATH,
			"FROM_VALUE": CompositionRevisionV1Beta1SpecResourcesConnectionDetailsType_FROM_VALUE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatches",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesCombine",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesCombine)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesCombineStrategy",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesCombineStrategy)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionV1Beta1SpecResourcesPatchesCombineStrategy_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesCombineString",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesCombineString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesCombineVariables",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesCombineVariables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesPolicy",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesPolicyFromFieldPath",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesPolicyFromFieldPath)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": CompositionRevisionV1Beta1SpecResourcesPatchesPolicyFromFieldPath_OPTIONAL,
			"REQUIRED": CompositionRevisionV1Beta1SpecResourcesPatchesPolicyFromFieldPath_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesPolicyMergeOptions",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesPolicyMergeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransforms",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransforms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvert",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvert)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat_NONE,
			"QUANTITY": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat_QUANTITY,
			"JSON": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertFormat_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertToType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertToType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertToType_STRING,
			"INT": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertToType_INT,
			"INT64": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertToType_INT64,
			"BOOL": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertToType_BOOL,
			"FLOAT64": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertToType_FLOAT64,
			"OBJECT": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertToType_OBJECT,
			"ARRAY": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsConvertToType_ARRAY,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatch",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatch)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatchFallbackTo",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatchFallbackTo)(nil)).Elem(),
		map[string]interface{}{
			"VALUE": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatchFallbackTo_VALUE,
			"INPUT": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatchFallbackTo_INPUT,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatchPatterns",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatchPatterns)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatchPatternsType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatchPatternsType)(nil)).Elem(),
		map[string]interface{}{
			"LITERAL": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatchPatternsType_LITERAL,
			"REGEXP": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMatchPatternsType_REGEXP,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMath",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMath)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMathType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMathType)(nil)).Elem(),
		map[string]interface{}{
			"MULTIPLY": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMathType_MULTIPLY,
			"CLAMP_MIN": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMathType_CLAMP_MIN,
			"CLAMP_MAX": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsMathType_CLAMP_MAX,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsString",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsString)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert)(nil)).Elem(),
		map[string]interface{}{
			"TO_UPPER": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert_TO_UPPER,
			"TO_LOWER": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert_TO_LOWER,
			"TO_BASE64": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert_TO_BASE64,
			"FROM_BASE64": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert_FROM_BASE64,
			"TO_JSON": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert_TO_JSON,
			"TO_SHA1": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert_TO_SHA1,
			"TO_SHA256": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert_TO_SHA256,
			"TO_SHA512": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert_TO_SHA512,
			"TO_ADLER32": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringConvert_TO_ADLER32,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringJoin",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringJoin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringRegexp",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringType)(nil)).Elem(),
		map[string]interface{}{
			"FORMAT": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringType_FORMAT,
			"CONVERT": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringType_CONVERT,
			"TRIM_PREFIX": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringType_TRIM_PREFIX,
			"TRIM_SUFFIX": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringType_TRIM_SUFFIX,
			"REGEXP": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringType_REGEXP,
			"JOIN": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsStringType_JOIN,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesTransformsType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesTransformsType)(nil)).Elem(),
		map[string]interface{}{
			"MAP": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsType_MAP,
			"MATCH": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsType_MATCH,
			"MATH": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsType_MATH,
			"STRING": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsType_STRING,
			"CONVERT": CompositionRevisionV1Beta1SpecResourcesPatchesTransformsType_CONVERT,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesPatchesType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesPatchesType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_COMPOSITE_FIELD_PATH": CompositionRevisionV1Beta1SpecResourcesPatchesType_FROM_COMPOSITE_FIELD_PATH,
			"PATCH_SET": CompositionRevisionV1Beta1SpecResourcesPatchesType_PATCH_SET,
			"TO_COMPOSITE_FIELD_PATH": CompositionRevisionV1Beta1SpecResourcesPatchesType_TO_COMPOSITE_FIELD_PATH,
			"COMBINE_FROM_COMPOSITE": CompositionRevisionV1Beta1SpecResourcesPatchesType_COMBINE_FROM_COMPOSITE,
			"COMBINE_TO_COMPOSITE": CompositionRevisionV1Beta1SpecResourcesPatchesType_COMBINE_TO_COMPOSITE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesReadinessChecks",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesReadinessChecks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesReadinessChecksMatchCondition",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesReadinessChecksMatchCondition)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionRevisionV1Beta1SpecResourcesReadinessChecksType",
		reflect.TypeOf((*CompositionRevisionV1Beta1SpecResourcesReadinessChecksType)(nil)).Elem(),
		map[string]interface{}{
			"MATCH_STRING": CompositionRevisionV1Beta1SpecResourcesReadinessChecksType_MATCH_STRING,
			"MATCH_INTEGER": CompositionRevisionV1Beta1SpecResourcesReadinessChecksType_MATCH_INTEGER,
			"NON_EMPTY": CompositionRevisionV1Beta1SpecResourcesReadinessChecksType_NON_EMPTY,
			"MATCH_CONDITION": CompositionRevisionV1Beta1SpecResourcesReadinessChecksType_MATCH_CONDITION,
			"MATCH_TRUE": CompositionRevisionV1Beta1SpecResourcesReadinessChecksType_MATCH_TRUE,
			"MATCH_FALSE": CompositionRevisionV1Beta1SpecResourcesReadinessChecksType_MATCH_FALSE,
			"NONE": CompositionRevisionV1Beta1SpecResourcesReadinessChecksType_NONE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpec",
		reflect.TypeOf((*CompositionSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecCompositeTypeRef",
		reflect.TypeOf((*CompositionSpecCompositeTypeRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecMode",
		reflect.TypeOf((*CompositionSpecMode)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCES": CompositionSpecMode_RESOURCES,
			"PIPELINE": CompositionSpecMode_PIPELINE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSets",
		reflect.TypeOf((*CompositionSpecPatchSets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatches",
		reflect.TypeOf((*CompositionSpecPatchSetsPatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesCombine",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesCombine)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesCombineStrategy",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesCombineStrategy)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionSpecPatchSetsPatchesCombineStrategy_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesCombineString",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesCombineString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesCombineVariables",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesCombineVariables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesPolicy",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesPolicyFromFieldPath",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesPolicyFromFieldPath)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": CompositionSpecPatchSetsPatchesPolicyFromFieldPath_OPTIONAL,
			"REQUIRED": CompositionSpecPatchSetsPatchesPolicyFromFieldPath_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesPolicyMergeOptions",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesPolicyMergeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransforms",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransforms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsConvert",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsConvert)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsConvertFormat",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsConvertFormat)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompositionSpecPatchSetsPatchesTransformsConvertFormat_NONE,
			"QUANTITY": CompositionSpecPatchSetsPatchesTransformsConvertFormat_QUANTITY,
			"JSON": CompositionSpecPatchSetsPatchesTransformsConvertFormat_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsConvertToType",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsConvertToType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionSpecPatchSetsPatchesTransformsConvertToType_STRING,
			"INT": CompositionSpecPatchSetsPatchesTransformsConvertToType_INT,
			"INT64": CompositionSpecPatchSetsPatchesTransformsConvertToType_INT64,
			"BOOL": CompositionSpecPatchSetsPatchesTransformsConvertToType_BOOL,
			"FLOAT64": CompositionSpecPatchSetsPatchesTransformsConvertToType_FLOAT64,
			"OBJECT": CompositionSpecPatchSetsPatchesTransformsConvertToType_OBJECT,
			"ARRAY": CompositionSpecPatchSetsPatchesTransformsConvertToType_ARRAY,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsMatch",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsMatch)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsMatchFallbackTo",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsMatchFallbackTo)(nil)).Elem(),
		map[string]interface{}{
			"VALUE": CompositionSpecPatchSetsPatchesTransformsMatchFallbackTo_VALUE,
			"INPUT": CompositionSpecPatchSetsPatchesTransformsMatchFallbackTo_INPUT,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsMatchPatterns",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsMatchPatterns)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsMatchPatternsType",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsMatchPatternsType)(nil)).Elem(),
		map[string]interface{}{
			"LITERAL": CompositionSpecPatchSetsPatchesTransformsMatchPatternsType_LITERAL,
			"REGEXP": CompositionSpecPatchSetsPatchesTransformsMatchPatternsType_REGEXP,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsMath",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsMath)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsMathType",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsMathType)(nil)).Elem(),
		map[string]interface{}{
			"MULTIPLY": CompositionSpecPatchSetsPatchesTransformsMathType_MULTIPLY,
			"CLAMP_MIN": CompositionSpecPatchSetsPatchesTransformsMathType_CLAMP_MIN,
			"CLAMP_MAX": CompositionSpecPatchSetsPatchesTransformsMathType_CLAMP_MAX,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsString",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsString)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsStringConvert",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsStringConvert)(nil)).Elem(),
		map[string]interface{}{
			"TO_UPPER": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_UPPER,
			"TO_LOWER": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_LOWER,
			"TO_BASE64": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_BASE64,
			"FROM_BASE64": CompositionSpecPatchSetsPatchesTransformsStringConvert_FROM_BASE64,
			"TO_JSON": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_JSON,
			"TO_SHA1": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_SHA1,
			"TO_SHA256": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_SHA256,
			"TO_SHA512": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_SHA512,
			"TO_ADLER32": CompositionSpecPatchSetsPatchesTransformsStringConvert_TO_ADLER32,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsStringJoin",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsStringJoin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsStringRegexp",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsStringRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsStringType",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsStringType)(nil)).Elem(),
		map[string]interface{}{
			"FORMAT": CompositionSpecPatchSetsPatchesTransformsStringType_FORMAT,
			"CONVERT": CompositionSpecPatchSetsPatchesTransformsStringType_CONVERT,
			"TRIM_PREFIX": CompositionSpecPatchSetsPatchesTransformsStringType_TRIM_PREFIX,
			"TRIM_SUFFIX": CompositionSpecPatchSetsPatchesTransformsStringType_TRIM_SUFFIX,
			"REGEXP": CompositionSpecPatchSetsPatchesTransformsStringType_REGEXP,
			"JOIN": CompositionSpecPatchSetsPatchesTransformsStringType_JOIN,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesTransformsType",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesTransformsType)(nil)).Elem(),
		map[string]interface{}{
			"MAP": CompositionSpecPatchSetsPatchesTransformsType_MAP,
			"MATCH": CompositionSpecPatchSetsPatchesTransformsType_MATCH,
			"MATH": CompositionSpecPatchSetsPatchesTransformsType_MATH,
			"STRING": CompositionSpecPatchSetsPatchesTransformsType_STRING,
			"CONVERT": CompositionSpecPatchSetsPatchesTransformsType_CONVERT,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPatchSetsPatchesType",
		reflect.TypeOf((*CompositionSpecPatchSetsPatchesType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_COMPOSITE_FIELD_PATH": CompositionSpecPatchSetsPatchesType_FROM_COMPOSITE_FIELD_PATH,
			"PATCH_SET": CompositionSpecPatchSetsPatchesType_PATCH_SET,
			"TO_COMPOSITE_FIELD_PATH": CompositionSpecPatchSetsPatchesType_TO_COMPOSITE_FIELD_PATH,
			"COMBINE_FROM_COMPOSITE": CompositionSpecPatchSetsPatchesType_COMBINE_FROM_COMPOSITE,
			"COMBINE_TO_COMPOSITE": CompositionSpecPatchSetsPatchesType_COMBINE_TO_COMPOSITE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPipeline",
		reflect.TypeOf((*CompositionSpecPipeline)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPipelineCredentials",
		reflect.TypeOf((*CompositionSpecPipelineCredentials)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPipelineCredentialsSecretRef",
		reflect.TypeOf((*CompositionSpecPipelineCredentialsSecretRef)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecPipelineCredentialsSource",
		reflect.TypeOf((*CompositionSpecPipelineCredentialsSource)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompositionSpecPipelineCredentialsSource_NONE,
			"SECRET": CompositionSpecPipelineCredentialsSource_SECRET,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPipelineFunctionRef",
		reflect.TypeOf((*CompositionSpecPipelineFunctionRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecPublishConnectionDetailsWithStoreConfigRef",
		reflect.TypeOf((*CompositionSpecPublishConnectionDetailsWithStoreConfigRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResources",
		reflect.TypeOf((*CompositionSpecResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesConnectionDetails",
		reflect.TypeOf((*CompositionSpecResourcesConnectionDetails)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesConnectionDetailsType",
		reflect.TypeOf((*CompositionSpecResourcesConnectionDetailsType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_CONNECTION_SECRET_KEY": CompositionSpecResourcesConnectionDetailsType_FROM_CONNECTION_SECRET_KEY,
			"FROM_FIELD_PATH": CompositionSpecResourcesConnectionDetailsType_FROM_FIELD_PATH,
			"FROM_VALUE": CompositionSpecResourcesConnectionDetailsType_FROM_VALUE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatches",
		reflect.TypeOf((*CompositionSpecResourcesPatches)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesCombine",
		reflect.TypeOf((*CompositionSpecResourcesPatchesCombine)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesCombineStrategy",
		reflect.TypeOf((*CompositionSpecResourcesPatchesCombineStrategy)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionSpecResourcesPatchesCombineStrategy_STRING,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesCombineString",
		reflect.TypeOf((*CompositionSpecResourcesPatchesCombineString)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesCombineVariables",
		reflect.TypeOf((*CompositionSpecResourcesPatchesCombineVariables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesPolicy",
		reflect.TypeOf((*CompositionSpecResourcesPatchesPolicy)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesPolicyFromFieldPath",
		reflect.TypeOf((*CompositionSpecResourcesPatchesPolicyFromFieldPath)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL": CompositionSpecResourcesPatchesPolicyFromFieldPath_OPTIONAL,
			"REQUIRED": CompositionSpecResourcesPatchesPolicyFromFieldPath_REQUIRED,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesPolicyMergeOptions",
		reflect.TypeOf((*CompositionSpecResourcesPatchesPolicyMergeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransforms",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransforms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsConvert",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsConvert)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsConvertFormat",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsConvertFormat)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompositionSpecResourcesPatchesTransformsConvertFormat_NONE,
			"QUANTITY": CompositionSpecResourcesPatchesTransformsConvertFormat_QUANTITY,
			"JSON": CompositionSpecResourcesPatchesTransformsConvertFormat_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsConvertToType",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsConvertToType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": CompositionSpecResourcesPatchesTransformsConvertToType_STRING,
			"INT": CompositionSpecResourcesPatchesTransformsConvertToType_INT,
			"INT64": CompositionSpecResourcesPatchesTransformsConvertToType_INT64,
			"BOOL": CompositionSpecResourcesPatchesTransformsConvertToType_BOOL,
			"FLOAT64": CompositionSpecResourcesPatchesTransformsConvertToType_FLOAT64,
			"OBJECT": CompositionSpecResourcesPatchesTransformsConvertToType_OBJECT,
			"ARRAY": CompositionSpecResourcesPatchesTransformsConvertToType_ARRAY,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsMatch",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsMatch)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsMatchFallbackTo",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsMatchFallbackTo)(nil)).Elem(),
		map[string]interface{}{
			"VALUE": CompositionSpecResourcesPatchesTransformsMatchFallbackTo_VALUE,
			"INPUT": CompositionSpecResourcesPatchesTransformsMatchFallbackTo_INPUT,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsMatchPatterns",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsMatchPatterns)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsMatchPatternsType",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsMatchPatternsType)(nil)).Elem(),
		map[string]interface{}{
			"LITERAL": CompositionSpecResourcesPatchesTransformsMatchPatternsType_LITERAL,
			"REGEXP": CompositionSpecResourcesPatchesTransformsMatchPatternsType_REGEXP,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsMath",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsMath)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsMathType",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsMathType)(nil)).Elem(),
		map[string]interface{}{
			"MULTIPLY": CompositionSpecResourcesPatchesTransformsMathType_MULTIPLY,
			"CLAMP_MIN": CompositionSpecResourcesPatchesTransformsMathType_CLAMP_MIN,
			"CLAMP_MAX": CompositionSpecResourcesPatchesTransformsMathType_CLAMP_MAX,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsString",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsString)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsStringConvert",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsStringConvert)(nil)).Elem(),
		map[string]interface{}{
			"TO_UPPER": CompositionSpecResourcesPatchesTransformsStringConvert_TO_UPPER,
			"TO_LOWER": CompositionSpecResourcesPatchesTransformsStringConvert_TO_LOWER,
			"TO_BASE64": CompositionSpecResourcesPatchesTransformsStringConvert_TO_BASE64,
			"FROM_BASE64": CompositionSpecResourcesPatchesTransformsStringConvert_FROM_BASE64,
			"TO_JSON": CompositionSpecResourcesPatchesTransformsStringConvert_TO_JSON,
			"TO_SHA1": CompositionSpecResourcesPatchesTransformsStringConvert_TO_SHA1,
			"TO_SHA256": CompositionSpecResourcesPatchesTransformsStringConvert_TO_SHA256,
			"TO_SHA512": CompositionSpecResourcesPatchesTransformsStringConvert_TO_SHA512,
			"TO_ADLER32": CompositionSpecResourcesPatchesTransformsStringConvert_TO_ADLER32,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsStringJoin",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsStringJoin)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsStringRegexp",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsStringRegexp)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsStringType",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsStringType)(nil)).Elem(),
		map[string]interface{}{
			"FORMAT": CompositionSpecResourcesPatchesTransformsStringType_FORMAT,
			"CONVERT": CompositionSpecResourcesPatchesTransformsStringType_CONVERT,
			"TRIM_PREFIX": CompositionSpecResourcesPatchesTransformsStringType_TRIM_PREFIX,
			"TRIM_SUFFIX": CompositionSpecResourcesPatchesTransformsStringType_TRIM_SUFFIX,
			"REGEXP": CompositionSpecResourcesPatchesTransformsStringType_REGEXP,
			"JOIN": CompositionSpecResourcesPatchesTransformsStringType_JOIN,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesTransformsType",
		reflect.TypeOf((*CompositionSpecResourcesPatchesTransformsType)(nil)).Elem(),
		map[string]interface{}{
			"MAP": CompositionSpecResourcesPatchesTransformsType_MAP,
			"MATCH": CompositionSpecResourcesPatchesTransformsType_MATCH,
			"MATH": CompositionSpecResourcesPatchesTransformsType_MATH,
			"STRING": CompositionSpecResourcesPatchesTransformsType_STRING,
			"CONVERT": CompositionSpecResourcesPatchesTransformsType_CONVERT,
		},
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesPatchesType",
		reflect.TypeOf((*CompositionSpecResourcesPatchesType)(nil)).Elem(),
		map[string]interface{}{
			"FROM_COMPOSITE_FIELD_PATH": CompositionSpecResourcesPatchesType_FROM_COMPOSITE_FIELD_PATH,
			"PATCH_SET": CompositionSpecResourcesPatchesType_PATCH_SET,
			"TO_COMPOSITE_FIELD_PATH": CompositionSpecResourcesPatchesType_TO_COMPOSITE_FIELD_PATH,
			"COMBINE_FROM_COMPOSITE": CompositionSpecResourcesPatchesType_COMBINE_FROM_COMPOSITE,
			"COMBINE_TO_COMPOSITE": CompositionSpecResourcesPatchesType_COMBINE_TO_COMPOSITE,
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesReadinessChecks",
		reflect.TypeOf((*CompositionSpecResourcesReadinessChecks)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.CompositionSpecResourcesReadinessChecksMatchCondition",
		reflect.TypeOf((*CompositionSpecResourcesReadinessChecksMatchCondition)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"apiextensionscrossplaneio.CompositionSpecResourcesReadinessChecksType",
		reflect.TypeOf((*CompositionSpecResourcesReadinessChecksType)(nil)).Elem(),
		map[string]interface{}{
			"MATCH_STRING": CompositionSpecResourcesReadinessChecksType_MATCH_STRING,
			"MATCH_INTEGER": CompositionSpecResourcesReadinessChecksType_MATCH_INTEGER,
			"NON_EMPTY": CompositionSpecResourcesReadinessChecksType_NON_EMPTY,
			"MATCH_CONDITION": CompositionSpecResourcesReadinessChecksType_MATCH_CONDITION,
			"MATCH_TRUE": CompositionSpecResourcesReadinessChecksType_MATCH_TRUE,
			"MATCH_FALSE": CompositionSpecResourcesReadinessChecksType_MATCH_FALSE,
			"NONE": CompositionSpecResourcesReadinessChecksType_NONE,
		},
	)
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.EnvironmentConfig",
		reflect.TypeOf((*EnvironmentConfig)(nil)).Elem(),
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
			j := jsiiProxy_EnvironmentConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.EnvironmentConfigProps",
		reflect.TypeOf((*EnvironmentConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.EnvironmentConfigV1Beta1",
		reflect.TypeOf((*EnvironmentConfigV1Beta1)(nil)).Elem(),
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
			j := jsiiProxy_EnvironmentConfigV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.EnvironmentConfigV1Beta1Props",
		reflect.TypeOf((*EnvironmentConfigV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.Usage",
		reflect.TypeOf((*Usage)(nil)).Elem(),
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
			j := jsiiProxy_Usage{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageProps",
		reflect.TypeOf((*UsageProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageSpec",
		reflect.TypeOf((*UsageSpec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageSpecBy",
		reflect.TypeOf((*UsageSpecBy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageSpecByResourceRef",
		reflect.TypeOf((*UsageSpecByResourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageSpecByResourceSelector",
		reflect.TypeOf((*UsageSpecByResourceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageSpecOf",
		reflect.TypeOf((*UsageSpecOf)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageSpecOfResourceRef",
		reflect.TypeOf((*UsageSpecOfResourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageSpecOfResourceSelector",
		reflect.TypeOf((*UsageSpecOfResourceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"apiextensionscrossplaneio.UsageV1Beta1",
		reflect.TypeOf((*UsageV1Beta1)(nil)).Elem(),
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
			j := jsiiProxy_UsageV1Beta1{}
			_jsii_.InitJsiiProxy(&j.Type__cdk8sApiObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageV1Beta1Props",
		reflect.TypeOf((*UsageV1Beta1Props)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageV1Beta1Spec",
		reflect.TypeOf((*UsageV1Beta1Spec)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageV1Beta1SpecBy",
		reflect.TypeOf((*UsageV1Beta1SpecBy)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageV1Beta1SpecByResourceRef",
		reflect.TypeOf((*UsageV1Beta1SpecByResourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageV1Beta1SpecByResourceSelector",
		reflect.TypeOf((*UsageV1Beta1SpecByResourceSelector)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageV1Beta1SpecOf",
		reflect.TypeOf((*UsageV1Beta1SpecOf)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageV1Beta1SpecOfResourceRef",
		reflect.TypeOf((*UsageV1Beta1SpecOfResourceRef)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"apiextensionscrossplaneio.UsageV1Beta1SpecOfResourceSelector",
		reflect.TypeOf((*UsageV1Beta1SpecOfResourceSelector)(nil)).Elem(),
	)
}
