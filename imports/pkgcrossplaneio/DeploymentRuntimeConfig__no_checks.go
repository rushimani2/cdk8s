//go:build no_runtime_type_checking

package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateDeploymentRuntimeConfig_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateDeploymentRuntimeConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDeploymentRuntimeConfig_ManifestParameters(props *DeploymentRuntimeConfigProps) error {
	return nil
}

func validateDeploymentRuntimeConfig_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewDeploymentRuntimeConfigParameters(scope constructs.Construct, id *string, props *DeploymentRuntimeConfigProps) error {
	return nil
}

