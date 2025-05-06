//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateIpamBlock_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateIpamBlock_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpamBlock_ManifestParameters(props *IpamBlockProps) error {
	return nil
}

func validateIpamBlock_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewIpamBlockParameters(scope constructs.Construct, id *string, props *IpamBlockProps) error {
	return nil
}

