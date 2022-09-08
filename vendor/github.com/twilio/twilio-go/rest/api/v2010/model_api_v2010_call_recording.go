/*
 * Twilio - Api
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

	"github.com/twilio/twilio-go/client"
)

// ApiV2010CallRecording struct for ApiV2010CallRecording
type ApiV2010CallRecording struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to make the recording
	ApiVersion *string `json:"api_version,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid *string `json:"call_sid,omitempty"`
	// The number of channels in the final recording file
	Channels *int `json:"channels,omitempty"`
	// The Conference SID that identifies the conference associated with the recording
	ConferenceSid *string `json:"conference_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The length of the recording in seconds
	Duration *string `json:"duration,omitempty"`
	// How to decrypt the recording.
	EncryptionDetails *interface{} `json:"encryption_details,omitempty"`
	// More information about why the recording is missing, if status is `absent`.
	ErrorCode *int `json:"error_code,omitempty"`
	// The one-time cost of creating the recording.
	Price *float32 `json:"price,omitempty"`
	// The currency used in the price property.
	PriceUnit *string `json:"price_unit,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// How the recording was created
	Source *string `json:"source,omitempty"`
	// The start time of the recording, given in RFC 2822 format
	StartTime *string `json:"start_time,omitempty"`
	// The status of the recording
	Status *string `json:"status,omitempty"`
	// The recorded track. Can be: `inbound`, `outbound`, or `both`.
	Track *string `json:"track,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}

func (response *ApiV2010CallRecording) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AccountSid        *string      `json:"account_sid"`
		ApiVersion        *string      `json:"api_version"`
		CallSid           *string      `json:"call_sid"`
		Channels          *int         `json:"channels"`
		ConferenceSid     *string      `json:"conference_sid"`
		DateCreated       *string      `json:"date_created"`
		DateUpdated       *string      `json:"date_updated"`
		Duration          *string      `json:"duration"`
		EncryptionDetails *interface{} `json:"encryption_details"`
		ErrorCode         *int         `json:"error_code"`
		Price             *interface{} `json:"price"`
		PriceUnit         *string      `json:"price_unit"`
		Sid               *string      `json:"sid"`
		Source            *string      `json:"source"`
		StartTime         *string      `json:"start_time"`
		Status            *string      `json:"status"`
		Track             *string      `json:"track"`
		Uri               *string      `json:"uri"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = ApiV2010CallRecording{
		AccountSid:        raw.AccountSid,
		ApiVersion:        raw.ApiVersion,
		CallSid:           raw.CallSid,
		Channels:          raw.Channels,
		ConferenceSid:     raw.ConferenceSid,
		DateCreated:       raw.DateCreated,
		DateUpdated:       raw.DateUpdated,
		Duration:          raw.Duration,
		EncryptionDetails: raw.EncryptionDetails,
		ErrorCode:         raw.ErrorCode,
		PriceUnit:         raw.PriceUnit,
		Sid:               raw.Sid,
		Source:            raw.Source,
		StartTime:         raw.StartTime,
		Status:            raw.Status,
		Track:             raw.Track,
		Uri:               raw.Uri,
	}

	responsePrice, err := client.UnmarshalFloat32(raw.Price)
	if err != nil {
		return err
	}
	response.Price = responsePrice

	return
}