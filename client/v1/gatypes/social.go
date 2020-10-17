package gatypes

type Social struct { // t=social
	SocialNetwork      string `url:"sn,omitempty"`
	SocialAction       string `url:"sa,omitempty"`
	SocialActionTarget string `url:"st,omitempty"`
}
