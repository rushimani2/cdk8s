//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateNetworkSet_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateNetworkSet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNetworkSet_ManifestParameters(props *NetworkSetProps) error {
	return nil
}

func validateNetworkSet_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewNetworkSetParameters(scope constructs.Construct, id *string, props *NetworkSetProps) error {
	return nil
}

