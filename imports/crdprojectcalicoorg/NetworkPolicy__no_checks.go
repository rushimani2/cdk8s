//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateNetworkPolicy_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateNetworkPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNetworkPolicy_ManifestParameters(props *NetworkPolicyProps) error {
	return nil
}

func validateNetworkPolicy_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewNetworkPolicyParameters(scope constructs.Construct, id *string, props *NetworkPolicyProps) error {
	return nil
}

