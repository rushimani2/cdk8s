//go:build no_runtime_type_checking

package metapkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateProviderV1Alpha1_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateProviderV1Alpha1_IsConstructParameters(x interface{}) error {
	return nil
}

func validateProviderV1Alpha1_ManifestParameters(props *ProviderV1Alpha1Props) error {
	return nil
}

func validateProviderV1Alpha1_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewProviderV1Alpha1Parameters(scope constructs.Construct, id *string, props *ProviderV1Alpha1Props) error {
	return nil
}

