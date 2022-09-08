/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// WirelessV1Sim struct for WirelessV1Sim
type WirelessV1Sim struct {
	// The SID of the Account to which the Sim resource belongs
	AccountSid *string `json:"account_sid,omitempty"`
	// The HTTP method we use to call commands_callback_url
	CommandsCallbackMethod *string `json:"commands_callback_method,omitempty"`
	// The URL we call when the SIM originates a machine-to-machine Command
	CommandsCallbackUrl *string `json:"commands_callback_url,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the Sim resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Deprecated
	EId *string `json:"e_id,omitempty"`
	// The string that you assigned to describe the Sim resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The ICCID associated with the SIM
	Iccid *string `json:"iccid,omitempty"`
	// Deprecated
	IpAddress *string `json:"ip_address,omitempty"`
	// The URLs of related subresources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the RatePlan resource to which the Sim resource is assigned.
	RatePlanSid *string `json:"rate_plan_sid,omitempty"`
	// The connectivity reset status of the SIM
	ResetStatus *string `json:"reset_status,omitempty"`
	// The unique string that identifies the Sim resource
	Sid *string `json:"sid,omitempty"`
	// Deprecated
	SmsFallbackMethod *string `json:"sms_fallback_method,omitempty"`
	// Deprecated
	SmsFallbackUrl *string `json:"sms_fallback_url,omitempty"`
	// Deprecated
	SmsMethod *string `json:"sms_method,omitempty"`
	// Deprecated
	SmsUrl *string `json:"sms_url,omitempty"`
	// The status of the Sim resource
	Status *string `json:"status,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName *string `json:"unique_name,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
	// Deprecated. The HTTP method we use to call voice_fallback_url
	VoiceFallbackMethod *string `json:"voice_fallback_method,omitempty"`
	// Deprecated. The URL we call when an error occurs while retrieving or executing the TwiML requested from voice_url
	VoiceFallbackUrl *string `json:"voice_fallback_url,omitempty"`
	// Deprecated. The HTTP method we use to call voice_url
	VoiceMethod *string `json:"voice_method,omitempty"`
	// Deprecated. The URL we call when the SIM-connected device makes a voice call
	VoiceUrl *string `json:"voice_url,omitempty"`
}