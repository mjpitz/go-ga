package v1

import (
	"net/http"

	"github.com/mjpitz/go-ga/client/v1/gatypes"
)

// NewClient creates a client to emit information to Google Analytics
func NewClient(trackingID, userAgent string) *Client {
	return &Client{
		baseURL:    baseURL,
		trackingID: trackingID,
		userAgent:  userAgent,
	}
}

// Client encapsulates communication with the Google Analytics API.
type Client struct {
	baseURL    string
	trackingID string
	userAgent  string
}

// SendGet will send the provided payload as a GET request.
func (c *Client) SendGet(payload *gatypes.Payload) error {
	payload.TrackingID = c.trackingID
	payload.SessionControl.OverrideUserAgent = c.userAgent

	getURL, err := Encode(payload)
	if err != nil {
		return err
	}

	resp, err := http.Get(getURL)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}

// SendPost will send the provided payload as a POST request.
func (c *Client) SendPost(payload *gatypes.Payload) error {
	payload.TrackingID = c.trackingID
	payload.SessionControl.OverrideUserAgent = c.userAgent

	values, err := Values(payload)
	if err != nil {
		return err
	}

	resp, err := http.PostForm(c.baseURL, values)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
