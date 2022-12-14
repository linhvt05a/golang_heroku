/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
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

// VerifyV2Verification struct for VerifyV2Verification
type VerifyV2Verification struct {
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The phone number or email being verified
	To      *string `json:"to,omitempty"`
	Channel *string `json:"channel,omitempty"`
	// The status of the verification resource
	Status *string `json:"status,omitempty"`
	// Whether the verification was successful
	Valid *bool `json:"valid,omitempty"`
	// Information about the phone number being verified
	Lookup *interface{} `json:"lookup,omitempty"`
	// The amount of the associated PSD2 compliant transaction.
	Amount *string `json:"amount,omitempty"`
	// The payee of the associated PSD2 compliant transaction
	Payee *string `json:"payee,omitempty"`
	// An array of verification attempt objects.
	SendCodeAttempts *[]interface{} `json:"send_code_attempts,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The set of fields used for a silent network auth (`sna`) verification
	Sna *interface{} `json:"sna,omitempty"`
	// The absolute URL of the Verification resource
	Url *string `json:"url,omitempty"`
}
