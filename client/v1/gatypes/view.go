package gatypes

// TrafficSources encapsulates information about the system
type TrafficSources struct { // t=(pageview,screenview)
	DocumentReferrer   string `url:"dr,omitempty"`
	CampaignName       string `url:"cn,omitempty"`
	CampaignSource     string `url:"cs,omitempty"`
	CampaignMedium     string `url:"cm,omitempty"`
	CampaignKeyword    string `url:"ck,omitempty"`
	CampaignContent    string `url:"cc,omitempty"`
	CampaignID         string `url:"ci,omitempty"`
	GoogleAdsID        string `url:"gclid,omitempty"`
	GoogleDisplayAdsID string `url:"dclid,omitempty"`
}

// SystemInformation encapsulates information about the system
type SystemInformation struct { // t=(pageview,screenview)
	ScreenResolution string `url:"sr,omitempty"`
	ViewPortSize     string `url:"vp,omitempty"`
	DocumentEncoding string `url:"de,omitempty"`
	ScreenColors     string `url:"sd,omitempty"`
	UserLanguage     string `url:"ul,omitempty"`
	JavaEnabled      bool   `url:"je,omitempty"`
	FlashVersion     string `url:"fl,omitempty"`
}

// ContentInformation encapsulates information about the page
type ContentInformation struct { // t=(pageview,screenview)
	DocumentLocationURL string `url:"dl,omitempty"`
	DocumentHostName    string `url:"dh,omitempty"`
	DocumentPath        string `url:"dp,omitempty"`
	DocumentTitle       string `url:"dt,omitempty"`
	ScreenName          string `url:"cd,omitempty"`
	// ContentGroup []string `url:"cg[i]"`
	LinkID string `url:"linkid,omitempty"`
}

// PageView encapsulates metrics for a pageview event
type PageView struct {
	TrafficSources
	SystemInformation
	ContentInformation
}

// ScreenView encapsulates metrics for a screenview event
type ScreenView struct {
	TrafficSources
	SystemInformation
	ContentInformation
}
