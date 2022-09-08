/*
 * Twilio - Ip_messaging
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

// IpMessagingV1Message struct for IpMessagingV1Message
type IpMessagingV1Message struct {
	AccountSid  *string    `json:"account_sid,omitempty"`
	Attributes  *string    `json:"attributes,omitempty"`
	Body        *string    `json:"body,omitempty"`
	ChannelSid  *string    `json:"channel_sid,omitempty"`
	DateCreated *time.Time `json:"date_created,omitempty"`
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	From        *string    `json:"from,omitempty"`
	Index       *int       `json:"index,omitempty"`
	ServiceSid  *string    `json:"service_sid,omitempty"`
	Sid         *string    `json:"sid,omitempty"`
	To          *string    `json:"to,omitempty"`
	Url         *string    `json:"url,omitempty"`
	WasEdited   *bool      `json:"was_edited,omitempty"`
}
