//go:build no_runtime_type_checking

package apiextensionscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateEnvironmentConfig_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateEnvironmentConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEnvironmentConfig_ManifestParameters(props *EnvironmentConfigProps) error {
	return nil
}

func validateEnvironmentConfig_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewEnvironmentConfigParameters(scope constructs.Construct, id *string, props *EnvironmentConfigProps) error {
	return nil
}

