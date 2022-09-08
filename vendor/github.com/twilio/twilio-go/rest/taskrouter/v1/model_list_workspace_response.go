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

// ListWorkspaceResponse struct for ListWorkspaceResponse
type ListWorkspaceResponse struct {
	Meta       ListWorkspaceResponseMeta `json:"meta,omitempty"`
	Workspaces []TaskrouterV1Workspace   `json:"workspaces,omitempty"`
}
