//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateClusterInformation_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateClusterInformation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterInformation_ManifestParameters(props *ClusterInformationProps) error {
	return nil
}

func validateClusterInformation_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewClusterInformationParameters(scope constructs.Construct, id *string, props *ClusterInformationProps) error {
	return nil
}

