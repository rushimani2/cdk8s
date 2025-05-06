//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateBlockAffinity_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateBlockAffinity_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBlockAffinity_ManifestParameters(props *BlockAffinityProps) error {
	return nil
}

func validateBlockAffinity_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewBlockAffinityParameters(scope constructs.Construct, id *string, props *BlockAffinityProps) error {
	return nil
}

