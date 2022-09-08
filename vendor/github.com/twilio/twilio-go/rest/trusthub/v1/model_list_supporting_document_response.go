/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSupportingDocumentResponse struct for ListSupportingDocumentResponse
type ListSupportingDocumentResponse struct {
	Meta    ListCustomerProfileResponseMeta `json:"meta,omitempty"`
	Results []TrusthubV1SupportingDocument  `json:"results,omitempty"`
}
