package v1

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/mjpitz/go-ga/client/v1/gatypes"
)

const baseURL = "https://www.google-analytics.com/collect"

// Encode accepts an analytics payload and encodes it to a URL
func Encode(payload *gatypes.Payload) (string, error) {
	values, err := Values(payload)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s?%s", baseURL, values.Encode()), nil
}

// Values accepts an analytics payload and encodes it to a query string values
func Values(payload *gatypes.Payload) (url.Values, error) {
	payload.Version = "1"

	if payload.TrackingID == "" {
		return nil, fmt.Errorf("a tracking id must be provided")
	}

	if payload.HitType == "" {
		return nil, fmt.Errorf("a hit type must be provided")
	}

	if payload.Users.ClientID == "" && payload.Users.UserID == "" {
		return nil, fmt.Errorf("either a client id or user id must be provided")
	}

	return query.Values(payload)
}

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
