//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateBgpPeer_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateBgpPeer_IsConstructParameters(x interface{}) error {
	return nil
}

func validateBgpPeer_ManifestParameters(props *BgpPeerProps) error {
	return nil
}

func validateBgpPeer_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewBgpPeerParameters(scope constructs.Construct, id *string, props *BgpPeerProps) error {
	return nil
}

