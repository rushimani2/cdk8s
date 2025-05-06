package apiextensionscrossplaneio


// UsageSpec defines the desired state of Usage.
type UsageSpec struct {
	// Of is the resource that is "being used".
	Of *UsageSpecOf `field:"required" json:"of" yaml:"of"`
	// By is the resource that is "using the other resource".
	By *UsageSpecBy `field:"optional" json:"by" yaml:"by"`
	// Reason is the reason for blocking deletion of the resource.
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// ReplayDeletion will trigger a deletion on the used resource during the deletion of the usage itself, if it was attempted to be deleted at least once.
	ReplayDeletion *bool `field:"optional" json:"replayDeletion" yaml:"replayDeletion"`
}

