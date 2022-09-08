/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEntityResponse struct for ListEntityResponse
type ListEntityResponse struct {
	Entities []VerifyV2Entity                    `json:"entities,omitempty"`
	Meta     ListVerificationAttemptResponseMeta `json:"meta,omitempty"`
}
