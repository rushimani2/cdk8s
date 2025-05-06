//go:build no_runtime_type_checking

package apiextensionscrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateEnvironmentConfigV1Beta1_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateEnvironmentConfigV1Beta1_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEnvironmentConfigV1Beta1_ManifestParameters(props *EnvironmentConfigV1Beta1Props) error {
	return nil
}

func validateEnvironmentConfigV1Beta1_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewEnvironmentConfigV1Beta1Parameters(scope constructs.Construct, id *string, props *EnvironmentConfigV1Beta1Props) error {
	return nil
}

