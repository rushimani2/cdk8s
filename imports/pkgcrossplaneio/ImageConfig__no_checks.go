//go:build no_runtime_type_checking

package pkgcrossplaneio

// Building without runtime type checking enabled, so all the below just return nil

func validateImageConfig_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateImageConfig_IsConstructParameters(x interface{}) error {
	return nil
}

func validateImageConfig_ManifestParameters(props *ImageConfigProps) error {
	return nil
}

func validateImageConfig_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewImageConfigParameters(scope constructs.Construct, id *string, props *ImageConfigProps) error {
	return nil
}

