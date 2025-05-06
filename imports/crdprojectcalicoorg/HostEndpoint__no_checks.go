//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateHostEndpoint_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateHostEndpoint_IsConstructParameters(x interface{}) error {
	return nil
}

func validateHostEndpoint_ManifestParameters(props *HostEndpointProps) error {
	return nil
}

func validateHostEndpoint_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewHostEndpointParameters(scope constructs.Construct, id *string, props *HostEndpointProps) error {
	return nil
}

