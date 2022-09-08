/*
 * Twilio - Fax
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"

	"github.com/twilio/twilio-go/client"
)

// FaxV1Fax struct for FaxV1Fax
type FaxV1Fax struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to transmit the fax
	ApiVersion *string `json:"api_version,omitempty"`
	// The ISO 8601 formatted date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 formatted date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The direction of the fax
	Direction *string `json:"direction,omitempty"`
	// The time it took to transmit the fax
	Duration *int `json:"duration,omitempty"`
	// The number the fax was sent from
	From *string `json:"from,omitempty"`
	// The URLs of the fax's related resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the FaxMedia resource that is associated with the Fax
	MediaSid *string `json:"media_sid,omitempty"`
	// The Twilio-hosted URL that can be used to download fax media
	MediaUrl *string `json:"media_url,omitempty"`
	// The number of pages contained in the fax document
	NumPages *int `json:"num_pages,omitempty"`
	// The fax transmission price
	Price *float32 `json:"price,omitempty"`
	// The ISO 4217 currency used for billing
	PriceUnit *string `json:"price_unit,omitempty"`
	// The quality of the fax
	Quality *string `json:"quality,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The status of the fax
	Status *string `json:"status,omitempty"`
	// The phone number that received the fax
	To *string `json:"to,omitempty"`
	// The absolute URL of the fax resource
	Url *string `json:"url,omitempty"`
}

func (response *FaxV1Fax) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AccountSid  *string                 `json:"account_sid"`
		ApiVersion  *string                 `json:"api_version"`
		DateCreated *time.Time              `json:"date_created"`
		DateUpdated *time.Time              `json:"date_updated"`
		Direction   *string                 `json:"direction"`
		Duration    *int                    `json:"duration"`
		From        *string                 `json:"from"`
		Links       *map[string]interface{} `json:"links"`
		MediaSid    *string                 `json:"media_sid"`
		MediaUrl    *string                 `json:"media_url"`
		NumPages    *int                    `json:"num_pages"`
		Price       *interface{}            `json:"price"`
		PriceUnit   *string                 `json:"price_unit"`
		Quality     *string                 `json:"quality"`
		Sid         *string                 `json:"sid"`
		Status      *string                 `json:"status"`
		To          *string                 `json:"to"`
		Url         *string                 `json:"url"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = FaxV1Fax{
		AccountSid:  raw.AccountSid,
		ApiVersion:  raw.ApiVersion,
		DateCreated: raw.DateCreated,
		DateUpdated: raw.DateUpdated,
		Direction:   raw.Direction,
		Duration:    raw.Duration,
		From:        raw.From,
		Links:       raw.Links,
		MediaSid:    raw.MediaSid,
		MediaUrl:    raw.MediaUrl,
		NumPages:    raw.NumPages,
		PriceUnit:   raw.PriceUnit,
		Quality:     raw.Quality,
		Sid:         raw.Sid,
		Status:      raw.Status,
		To:          raw.To,
		Url:         raw.Url,
	}

	responsePrice, err := client.UnmarshalFloat32(raw.Price)
	if err != nil {
		return err
	}
	response.Price = responsePrice

	return
}
