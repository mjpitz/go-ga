package gatypes

// Social encapsulates fields on a social event
type Social struct { // t=social
	SocialNetwork      string `url:"sn,omitempty" json:"social_network,omitempty"`
	SocialAction       string `url:"sa,omitempty" json:"social_action,omitempty"`
	SocialActionTarget string `url:"st,omitempty" json:"social_action_target,omitempty"`
}
