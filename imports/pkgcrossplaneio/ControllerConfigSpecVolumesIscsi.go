package pkgcrossplaneio


// iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod.
//
// More info: https://examples.k8s.io/volumes/iscsi/README.md
type ControllerConfigSpecVolumesIscsi struct {
	// iqn is the target iSCSI Qualified Name.
	Iqn *string `field:"required" json:"iqn" yaml:"iqn"`
	// lun represents iSCSI Target Lun number.
	Lun *float64 `field:"required" json:"lun" yaml:"lun"`
	// targetPortal is iSCSI Target Portal.
	//
	// The Portal is either an IP or ip_addr:port if the port
	// is other than default (typically TCP ports 860 and 3260).
	TargetPortal *string `field:"required" json:"targetPortal" yaml:"targetPortal"`
	// chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication.
	ChapAuthDiscovery *bool `field:"optional" json:"chapAuthDiscovery" yaml:"chapAuthDiscovery"`
	// chapAuthSession defines whether support iSCSI Session CHAP authentication.
	ChapAuthSession *bool `field:"optional" json:"chapAuthSession" yaml:"chapAuthSession"`
	// fsType is the filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
	// More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// initiatorName is the custom iSCSI Initiator Name.
	//
	// If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface
	// <target portal>:<volume name> will be created for the connection.
	InitiatorName *string `field:"optional" json:"initiatorName" yaml:"initiatorName"`
	// iscsiInterface is the interface Name that uses an iSCSI transport.
	//
	// Defaults to 'default' (tcp).
	// Default: default' (tcp).
	//
	IscsiInterface *string `field:"optional" json:"iscsiInterface" yaml:"iscsiInterface"`
	// portals is the iSCSI Target Portal List.
	//
	// The portal is either an IP or ip_addr:port if the port
	// is other than default (typically TCP ports 860 and 3260).
	Portals *[]*string `field:"optional" json:"portals" yaml:"portals"`
	// readOnly here will force the ReadOnly setting in VolumeMounts.
	//
	// Defaults to false.
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// secretRef is the CHAP Secret for iSCSI target and initiator authentication.
	SecretRef *ControllerConfigSpecVolumesIscsiSecretRef `field:"optional" json:"secretRef" yaml:"secretRef"`
}

