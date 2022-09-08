/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// StudioV2TestUser struct for StudioV2TestUser
type StudioV2TestUser struct {
	// Unique identifier of the flow.
	Sid *string `json:"sid,omitempty"`
	// List of test user identities that can test draft versions of the flow.
	TestUsers *[]string `json:"test_users,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
