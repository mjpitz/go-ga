package gatypes

// TrafficSources encapsulates information about the system
type TrafficSources struct { // t=(pageview,screenview)
	DocumentReferrer   string `url:"dr,omitempty" json:"document_referrer,omitempty"`
	CampaignName       string `url:"cn,omitempty" json:"campaign_name,omitempty"`
	CampaignSource     string `url:"cs,omitempty" json:"campaign_source,omitempty"`
	CampaignMedium     string `url:"cm,omitempty" json:"campaign_medium,omitempty"`
	CampaignKeyword    string `url:"ck,omitempty" json:"campaign_keyword,omitempty"`
	CampaignContent    string `url:"cc,omitempty" json:"campaign_content,omitempty"`
	CampaignID         string `url:"ci,omitempty" json:"campaign_id,omitempty"`
	GoogleAdsID        string `url:"gclid,omitempty" json:"google_ads_id,omitempty"`
	GoogleDisplayAdsID string `url:"dclid,omitempty" json:"google_display_ads_id,omitempty"`
}

// SystemInformation encapsulates information about the system
type SystemInformation struct { // t=(pageview,screenview)
	ScreenResolution string `url:"sr,omitempty" json:"screen_resolution,omitempty"`
	ViewPortSize     string `url:"vp,omitempty" json:"view_port_size,omitempty"`
	DocumentEncoding string `url:"de,omitempty" json:"document_encoding,omitempty"`
	ScreenColors     string `url:"sd,omitempty" json:"screen_colors,omitempty"`
	UserLanguage     string `url:"ul,omitempty" json:"user_language,omitempty"`
	JavaEnabled      bool   `url:"je,omitempty" json:"java_enabled,omitempty"`
	FlashVersion     string `url:"fl,omitempty" json:"flash_version,omitempty"`
}

// ContentInformation encapsulates information about the page
type ContentInformation struct { // t=(pageview,screenview)
	DocumentLocationURL string     `url:"dl,omitempty" json:"document_location,omitempty"`
	DocumentHostName    string     `url:"dh,omitempty" json:"document_host_name,omitempty"`
	DocumentPath        string     `url:"dp,omitempty" json:"document_path,omitempty"`
	DocumentTitle       string     `url:"dt,omitempty" json:"document_title,omitempty"`
	ScreenName          string     `url:"cd,omitempty" json:"screen_name,omitempty"`
	ContentGroup        StringList `url:"cg,omitempty" json:"content_group,omitempty"`
	LinkID              string     `url:"linkid,omitempty" json:"link_id,omitempty"`
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
