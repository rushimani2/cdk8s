package crdprojectcalicoorg


// Optional BGP password for full node-to-mesh peerings.
//
// This field can only be set on the default BGPConfiguration instance and requires that NodeMesh is enabled.
type BgpConfigurationSpecNodeMeshPassword struct {
	// Selects a key of a secret in the node pod's namespace.
	SecretKeyRef *BgpConfigurationSpecNodeMeshPasswordSecretKeyRef `field:"optional" json:"secretKeyRef" yaml:"secretKeyRef"`
}

