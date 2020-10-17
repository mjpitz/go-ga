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
	Users          *Users          `url:"-"`
	SessionControl *SessionControl `url:"-"`
	Apps           *Apps           `url:"-"`

	// one of
	Event       *Event       `url:"-"`
	Exception   *Exception   `url:"-"`
	Item        *Item        `url:"-"`
	Social      *Social      `url:"-"`
	Timing      *Timing      `url:"-"`
	Transaction *Transaction `url:"-"`
	PageView    *PageView    `url:"-"`
	ScreenView  *ScreenView  `url:"-"`
}

type Users struct {
	ClientID string `url:"cid"`
	UserID   string `url:"uid"`
}

type SessionControl struct { // t=*
	SessionControl    string `url:"sc"`
	OverrideIP        string `url:"uip"`
	OverrideUserAgent string `url:"ua"`
	OverrideGeography string `url:"geoid"`
}

type Apps struct { // t=*
	ApplicationName        string `url:"an"`
	ApplicationID          string `url:"aid"`
	ApplicationVersion     string `url:"av"`
	ApplicationInstallerID string `url:"aiid"`
}
