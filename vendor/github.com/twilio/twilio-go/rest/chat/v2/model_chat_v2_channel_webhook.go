/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Chat
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// ChatV2ChannelWebhook struct for ChatV2ChannelWebhook
type ChatV2ChannelWebhook struct {
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Service that the Channel Webhook resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the Channel the Channel Webhook resource belongs to
	ChannelSid *string `json:"channel_sid,omitempty"`
	// The type of webhook
	Type *string `json:"type,omitempty"`
	// The absolute URL of the Channel Webhook resource
	Url *string `json:"url,omitempty"`
	// The JSON string that describes the configuration object for the channel webhook
	Configuration *interface{} `json:"configuration,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}
