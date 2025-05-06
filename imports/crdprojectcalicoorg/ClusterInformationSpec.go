package crdprojectcalicoorg


// ClusterInformationSpec contains the values of describing the cluster.
type ClusterInformationSpec struct {
	// CalicoVersion is the version of Calico that the cluster is running.
	CalicoVersion *string `field:"optional" json:"calicoVersion" yaml:"calicoVersion"`
	// ClusterGUID is the GUID of the cluster.
	ClusterGuid *string `field:"optional" json:"clusterGuid" yaml:"clusterGuid"`
	// ClusterType describes the type of the cluster.
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// DatastoreReady is used during significant datastore migrations to signal to components such as Felix that it should wait before accessing the datastore.
	DatastoreReady *bool `field:"optional" json:"datastoreReady" yaml:"datastoreReady"`
	// Variant declares which variant of Calico should be active.
	Variant *string `field:"optional" json:"variant" yaml:"variant"`
}

