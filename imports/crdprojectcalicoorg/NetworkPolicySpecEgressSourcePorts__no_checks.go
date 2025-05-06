//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateNetworkPolicySpecEgressSourcePorts_FromNumberParameters(value *float64) error {
	return nil
}

func validateNetworkPolicySpecEgressSourcePorts_FromStringParameters(value *string) error {
	return nil
}

