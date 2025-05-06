//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateCalicoNodeStatus_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateCalicoNodeStatus_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCalicoNodeStatus_ManifestParameters(props *CalicoNodeStatusProps) error {
	return nil
}

func validateCalicoNodeStatus_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewCalicoNodeStatusParameters(scope constructs.Construct, id *string, props *CalicoNodeStatusProps) error {
	return nil
}

