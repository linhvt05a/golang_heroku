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

// ApiV2010CallNotification struct for ApiV2010CallNotification
type ApiV2010CallNotification struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The API version used to create the Call Notification resource
	ApiVersion *string `json:"api_version,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid *string `json:"call_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// A unique error code corresponding to the notification
	ErrorCode *string `json:"error_code,omitempty"`
	// An integer log level
	Log *string `json:"log,omitempty"`
	// The date the notification was generated
	MessageDate *string `json:"message_date,omitempty"`
	// The text of the notification
	MessageText *string `json:"message_text,omitempty"`
	// A URL for more information about the error code
	MoreInfo *string `json:"more_info,omitempty"`
	// HTTP method used with the request url
	RequestMethod *string `json:"request_method,omitempty"`
	// URL of the resource that generated the notification
	RequestUrl *string `json:"request_url,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}