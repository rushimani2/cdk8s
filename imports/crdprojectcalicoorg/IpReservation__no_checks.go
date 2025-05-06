//go:build no_runtime_type_checking

package crdprojectcalicoorg

// Building without runtime type checking enabled, so all the below just return nil

func validateIpReservation_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateIpReservation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateIpReservation_ManifestParameters(props *IpReservationProps) error {
	return nil
}

func validateIpReservation_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewIpReservationParameters(scope constructs.Construct, id *string, props *IpReservationProps) error {
	return nil
}

