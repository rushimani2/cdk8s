package crdprojectcalicoorg


// HostEndpoint controls syncing nodes to host endpoints.
//
// Disabled by default, set to nil to disable.
type KubeControllersConfigurationSpecControllersNodeHostEndpoint struct {
	// AutoCreate enables automatic creation of host endpoints for every node.
	//
	// [Default: Disabled].
	AutoCreate *string `field:"optional" json:"autoCreate" yaml:"autoCreate"`
}

