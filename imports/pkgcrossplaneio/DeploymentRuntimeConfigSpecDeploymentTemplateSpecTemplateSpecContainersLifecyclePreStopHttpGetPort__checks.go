//go:build !no_runtime_type_checking

package pkgcrossplaneio

import (
	"fmt"
)

func validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopHttpGetPort_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecContainersLifecyclePreStopHttpGetPort_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

