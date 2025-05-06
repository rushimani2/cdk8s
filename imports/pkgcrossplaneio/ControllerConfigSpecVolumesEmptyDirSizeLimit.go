package pkgcrossplaneio

import (
	_init_ "example.com/cdk8s-demo/imports/pkgcrossplaneio/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// sizeLimit is the total amount of local storage required for this EmptyDir volume.
//
// The size limit is also applicable for memory medium.
// The maximum usage on memory medium EmptyDir would be the minimum value between
// the SizeLimit specified here and the sum of memory limits of all containers in a pod.
// The default is nil which means that the limit is undefined.
// More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
type ControllerConfigSpecVolumesEmptyDirSizeLimit interface {
	Value() interface{}
}

// The jsii proxy struct for ControllerConfigSpecVolumesEmptyDirSizeLimit
type jsiiProxy_ControllerConfigSpecVolumesEmptyDirSizeLimit struct {
	_ byte // padding
}

func (j *jsiiProxy_ControllerConfigSpecVolumesEmptyDirSizeLimit) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func ControllerConfigSpecVolumesEmptyDirSizeLimit_FromNumber(value *float64) ControllerConfigSpecVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateControllerConfigSpecVolumesEmptyDirSizeLimit_FromNumberParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEmptyDirSizeLimit",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ControllerConfigSpecVolumesEmptyDirSizeLimit_FromString(value *string) ControllerConfigSpecVolumesEmptyDirSizeLimit {
	_init_.Initialize()

	if err := validateControllerConfigSpecVolumesEmptyDirSizeLimit_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns ControllerConfigSpecVolumesEmptyDirSizeLimit

	_jsii_.StaticInvoke(
		"pkgcrossplaneio.ControllerConfigSpecVolumesEmptyDirSizeLimit",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

