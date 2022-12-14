/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
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

// ProxyV1ShortCode struct for ProxyV1ShortCode
type ProxyV1ShortCode struct {
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the resource's parent Service
	ServiceSid *string `json:"service_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The short code's number
	ShortCode *string `json:"short_code,omitempty"`
	// The ISO Country Code
	IsoCountry   *string                              `json:"iso_country,omitempty"`
	Capabilities *ProxyV1ServiceShortCodeCapabilities `json:"capabilities,omitempty"`
	// The absolute URL of the ShortCode resource
	Url *string `json:"url,omitempty"`
	// Whether the short code should be reserved for manual assignment to participants only
	IsReserved *bool `json:"is_reserved,omitempty"`
}
