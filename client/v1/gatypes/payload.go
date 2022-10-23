package gatypes

// https://developers.google.com/analytics/devguides/collection/protocol/v1/reference

// Payload represent a ping payload you can send to Google Analytics.
type Payload struct {
	// General
	Version    string `url:"v" json:"version,omitempty"`
	TrackingID string `url:"tid" json:"tracking_id,omitempty"`

	AnonymizeIP                       Boolean `url:"aip,omitempty" json:"anonymize_ip,omitempty"`
	DisableAdvertisingPersonalization Boolean `url:"npa,omitempty" json:"disable_ad_personalization,omitempty"`
	DataSource                        string  `url:"ds,omitempty" json:"data_source,omitempty"`
	QueueTime                         int64   `url:"qt,omitempty" json:"queue_time,omitempty"`
	CacheBuster                       string  `url:"z,omitempty" json:"cache_buster,omitempty"`

	HitType           string  `url:"t" json:"hit_type,omitempty"`
	NonInteractionHit Boolean `url:"ni,omitempty" json:"non_interaction_hit,omitempty"`

	// apply to all
	Users
	SessionControl
	Apps
	CustomDimensions StringList  `url:"cd,omitempty" json:"custom_dimensions,omitempty"`
	CustomMetrics    Float64List `url:"cm,omitempty" json:"custom_metrics,omitempty"`

	// one of
	Event
	Exception
	Item
	Social
	Timing
	Transaction
	Views

	// aply to Item and Transaction
	TransactionID string `url:"ti,omitempty" json:"transaction_id,omitempty"`
}

// Users encapsulates information about the user performing an action
type Users struct {
	ClientID string `url:"cid,omitempty" json:"client_id,omitempty"`
	UserID   string `url:"uid,omitempty" json:"user_id,omitempty"`
}

// SessionControl gives uses the ability to control session data
type SessionControl struct { // t=*
	SessionControl    string `url:"sc,omitempty" json:"session_control,omitempty"`
	OverrideIP        string `url:"uip,omitempty" json:"override_ip,omitempty"`
	OverrideUserAgent string `url:"ua,omitempty" json:"override_user_agent,omitempty"`
	OverrideGeography string `url:"geoid,omitempty" json:"override_geography,omitempty"`
}

// Apps encapsulates information about the application sending the event
type Apps struct { // t=*
	ApplicationName        string `url:"an,omitempty" json:"application_name,omitempty"`
	ApplicationID          string `url:"aid,omitempty" json:"application_id,omitempty"`
	ApplicationVersion     string `url:"av,omitempty" json:"application_version,omitempty"`
	ApplicationInstallerID string `url:"aiid,omitempty" json:"application_installer_id,omitempty"`
}
