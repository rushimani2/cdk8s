//go:build no_runtime_type_checking

package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateFunctionRevision_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateFunctionRevision_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFunctionRevision_ManifestParameters(props *FunctionRevisionProps) error {
	return nil
}

func validateFunctionRevision_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewFunctionRevisionParameters(scope constructs.Construct, id *string, props *FunctionRevisionProps) error {
	return nil
}

