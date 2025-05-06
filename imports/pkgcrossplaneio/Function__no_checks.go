//go:build no_runtime_type_checking

package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateFunction_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateFunction_IsConstructParameters(x interface{}) error {
	return nil
}

func validateFunction_ManifestParameters(props *FunctionProps) error {
	return nil
}

func validateFunction_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewFunctionParameters(scope constructs.Construct, id *string, props *FunctionProps) error {
	return nil
}

