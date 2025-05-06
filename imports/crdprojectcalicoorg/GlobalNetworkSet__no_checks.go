//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateGlobalNetworkSet_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateGlobalNetworkSet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGlobalNetworkSet_ManifestParameters(props *GlobalNetworkSetProps) error {
	return nil
}

func validateGlobalNetworkSet_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewGlobalNetworkSetParameters(scope constructs.Construct, id *string, props *GlobalNetworkSetProps) error {
	return nil
}

