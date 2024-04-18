package aoscxgo

import (
	"errors"
)

type Firmware struct {
	BootedImage      string                 `json:"booted_image"`
	CurrentVersion   string                 `json:"current_version"`
	DefaultImage     string                 `json:"default_image"`
	PrimaryVersion   string                 `json:"primary_version"`
	SecondaryVersion string                 `json:"secondary_version"`
	FirmwareDetails  map[string]interface{} `json:"details"`
	materialized     bool
}

func (f *Firmware) Get(c *Client) error {
	base_uri := "firmware"

	url := "https://" + c.Hostname + "/rest/" + c.Version + "/" + base_uri

	res, body := get(c.Transport, c.Cookie, url)

	if res.Status != "200 OK" {
		f.materialized = false
		return &RequestError{
			StatusCode: res.Status,
			Err:        errors.New("retrieval error"),
		}
	}

	for key, value := range body {
		switch key {
		case "booted_image":
			if strVal, ok := value.(string); ok {
				f.BootedImage = strVal
			}
		case "current_version":
			if strVal, ok := value.(string); ok {
				f.CurrentVersion = strVal
			}
		case "default_image":
			if strVal, ok := value.(string); ok {
				f.DefaultImage = strVal
			}
		case "primary_version":
			if strVal, ok := value.(string); ok {
				f.PrimaryVersion = strVal
			}
		case "secondary_version":
			if strVal, ok := value.(string); ok {
				f.SecondaryVersion = strVal
			}
		}
	}
	return nil
}
