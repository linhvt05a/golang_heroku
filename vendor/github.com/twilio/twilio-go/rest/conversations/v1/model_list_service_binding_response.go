/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListServiceBindingResponse struct for ListServiceBindingResponse
type ListServiceBindingResponse struct {
	Bindings []ConversationsV1ServiceBinding      `json:"bindings,omitempty"`
	Meta     ListConfigurationAddressResponseMeta `json:"meta,omitempty"`
}