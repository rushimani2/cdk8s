//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateBgpConfiguration_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateBgpConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBgpConfiguration_ManifestParameters(props *BgpConfigurationProps) error {
	return nil
}

func validateBgpConfiguration_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewBgpConfigurationParameters(scope constructs.Construct, id *string, props *BgpConfigurationProps) error {
	return nil
}

