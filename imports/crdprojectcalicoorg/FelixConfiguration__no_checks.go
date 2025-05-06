//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateFelixConfiguration_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateFelixConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFelixConfiguration_ManifestParameters(props *FelixConfigurationProps) error {
	return nil
}

func validateFelixConfiguration_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewFelixConfigurationParameters(scope constructs.Construct, id *string, props *FelixConfigurationProps) error {
	return nil
}

