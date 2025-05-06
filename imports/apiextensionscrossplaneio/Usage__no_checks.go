//go:build no_runtime_type_checking

package apiextensionscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateUsage_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateUsage_IsConstructParameters(x interface{}) error {
	return nil
}

func validateUsage_ManifestParameters(props *UsageProps) error {
	return nil
}

func validateUsage_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewUsageParameters(scope constructs.Construct, id *string, props *UsageProps) error {
	return nil
}

