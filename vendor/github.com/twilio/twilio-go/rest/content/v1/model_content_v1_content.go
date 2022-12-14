/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Content
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

// ContentV1Content struct for ContentV1Content
type ContentV1Content struct {
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// A string name used to describe the Content resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Two-letter language code identifying the language the Content resource is in.
	Language *string `json:"language,omitempty"`
	// Defines the default placeholder values for variables included in the Content resource
	Variables *interface{} `json:"variables,omitempty"`
	// The Content types (e.g. twilio/text) for this Content resource
	Types *interface{} `json:"types,omitempty"`
	// The URL of the resource, relative to `https://content.twilio.com`
	Url *string `json:"url,omitempty"`
	// A list of links related to the Content resource
	Links *map[string]interface{} `json:"links,omitempty"`
}
