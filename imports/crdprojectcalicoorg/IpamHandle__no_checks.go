//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateIpamHandle_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateIpamHandle_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpamHandle_ManifestParameters(props *IpamHandleProps) error {
	return nil
}

func validateIpamHandle_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewIpamHandleParameters(scope constructs.Construct, id *string, props *IpamHandleProps) error {
	return nil
}

