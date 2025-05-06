//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeControllersConfiguration_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeControllersConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeControllersConfiguration_ManifestParameters(props *KubeControllersConfigurationProps) error {
	return nil
}

func validateKubeControllersConfiguration_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeControllersConfigurationParameters(scope constructs.Construct, id *string, props *KubeControllersConfigurationProps) error {
	return nil
}

