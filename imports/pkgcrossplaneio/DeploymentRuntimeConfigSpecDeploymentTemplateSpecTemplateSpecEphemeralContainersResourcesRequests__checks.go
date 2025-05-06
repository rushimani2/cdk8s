//go:build !no_runtime_type_checking

package pkgcrossplaneio

import (
	"fmt"
)

func validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesRequests_FromNumberParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateDeploymentRuntimeConfigSpecDeploymentTemplateSpecTemplateSpecEphemeralContainersResourcesRequests_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

