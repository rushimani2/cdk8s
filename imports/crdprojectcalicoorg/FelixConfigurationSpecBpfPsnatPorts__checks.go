//go:build !no_runtime_type_checking

package crdprojectcalicoorg

import (
	"fmt"
)

func validateFelixConfigurationSpecBpfPsnatPorts_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateFelixConfigurationSpecBpfPsnatPorts_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

