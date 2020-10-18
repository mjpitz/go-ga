package gatypes

// https://developers.google.com/analytics/devguides/collection/protocol/v1/reference

// Payload represent a ping payload you can send to Google Analytics.
type Payload struct {
	// General

	Version    string `url:"v"`
	TrackingID string `url:"tid"`

	AnonymizeIP                       Boolean `url:"aip,omitempty"`
	DisableAdvertisingPersonalization Boolean `url:"npa,omitempty"`
	DataSource                        string  `url:"ds,omitempty"`
	QueueTime                         int64   `url:"qt,omitempty"`
	CacheBuster                       string  `url:"z,omitempty"`

	HitType           string  `url:"t"`
	NonInteractionHit Boolean `url:"ni,omitempty"`

	// apply to all
	Users
	SessionControl
	Apps

	// one of
	Event
	Exception
	Item
	Social
	Timing
	Transaction
	PageView
	ScreenView
}

// Users encapsulates information about the user performing an action
type Users struct {
	ClientID string `url:"cid,omitempty"`
	UserID   string `url:"uid,omitempty"`
}

// SessionControl gives uses the ability to control session data
type SessionControl struct { // t=*
	SessionControl    string `url:"sc,omitempty"`
	OverrideIP        string `url:"uip,omitempty"`
	OverrideUserAgent string `url:"ua,omitempty"`
	OverrideGeography string `url:"geoid,omitempty"`
}

// Apps encapsulates information about the application sending the event
type Apps struct { // t=*
	ApplicationName        string `url:"an,omitempty"`
	ApplicationID          string `url:"aid,omitempty"`
	ApplicationVersion     string `url:"av,omitempty"`
	ApplicationInstallerID string `url:"aiid,omitempty"`
}
