//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateIpamConfig_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateIpamConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpamConfig_ManifestParameters(props *IpamConfigProps) error {
	return nil
}

func validateIpamConfig_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewIpamConfigParameters(scope constructs.Construct, id *string, props *IpamConfigProps) error {
	return nil
}

