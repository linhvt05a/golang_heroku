/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTaskReservationResponse struct for ListTaskReservationResponse
type ListTaskReservationResponse struct {
	Meta         ListWorkspaceResponseMeta     `json:"meta,omitempty"`
	Reservations []TaskrouterV1TaskReservation `json:"reservations,omitempty"`
}
