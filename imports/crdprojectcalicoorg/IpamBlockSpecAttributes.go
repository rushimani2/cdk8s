package crdprojectcalicoorg


type IpamBlockSpecAttributes struct {
	HandleId *string `field:"optional" json:"handleId" yaml:"handleId"`
	Secondary *map[string]*string `field:"optional" json:"secondary" yaml:"secondary"`
}

