package v1

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"

	"github.com/mjpitz/go-ga/client/v1/gatypes"
)

// NewClient creates a client to emit information to Google Analytics
func NewClient(trackingID string) *Client {
	return &Client{
		baseURL:    "https://www.google-analytics.com/collect",
		trackingID: trackingID,
	}
}

// Client encapsulates communication with the Google Analytics API.
type Client struct {
	baseURL    string
	trackingID string
}

func (c *Client) validateAndEncode(payload *gatypes.Payload) (url.Values, error) {
	payload.Version = "1"
	payload.TrackingID = c.trackingID

	if payload.HitType == "" {
		return nil, fmt.Errorf("a hit type must be provided")
	}

	if payload.Users == nil || (payload.Users.ClientID == "" && payload.Users.UserID == "") {
		return nil, fmt.Errorf("either a client id or user id must be provided")
	}

	return query.Values(payload)
}

// GetURL returns a string URL for the given payload that can be sent via a GET request.
func (c *Client) GetURL(payload *gatypes.Payload) (string, error) {
	values, err := c.validateAndEncode(payload)
	if err != nil {
		return "", err
	}

	getURL := fmt.Sprintf("%s?%s", c.baseURL, values.Encode())
	return getURL, nil
}

// SendPost will send the provided payload as a POST request.
func (c *Client) SendPost(payload *gatypes.Payload) error {
	values, err := c.validateAndEncode(payload)
	if err != nil {
		return err
	}

	resp, err := http.PostForm(c.baseURL, values)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
