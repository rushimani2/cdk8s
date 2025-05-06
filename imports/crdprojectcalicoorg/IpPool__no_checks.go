//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateIpPool_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateIpPool_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpPool_ManifestParameters(props *IpPoolProps) error {
	return nil
}

func validateIpPool_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewIpPoolParameters(scope constructs.Construct, id *string, props *IpPoolProps) error {
	return nil
}

