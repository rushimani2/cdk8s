//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateGlobalNetworkPolicy_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateGlobalNetworkPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGlobalNetworkPolicy_ManifestParameters(props *GlobalNetworkPolicyProps) error {
	return nil
}

func validateGlobalNetworkPolicy_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewGlobalNetworkPolicyParameters(scope constructs.Construct, id *string, props *GlobalNetworkPolicyProps) error {
	return nil
}

