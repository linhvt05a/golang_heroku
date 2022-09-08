/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// AutopilotV1TaskStatistics struct for AutopilotV1TaskStatistics
type AutopilotV1TaskStatistics struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Assistant that is the parent of the Task associated with the resource
	AssistantSid *string `json:"assistant_sid,omitempty"`
	// The total number of Fields associated with the Task
	FieldsCount *int `json:"fields_count,omitempty"`
	// The total number of Samples associated with the Task
	SamplesCount *int `json:"samples_count,omitempty"`
	// The SID of the Task for which the statistics were collected
	TaskSid *string `json:"task_sid,omitempty"`
	// The absolute URL of the TaskStatistics resource
	Url *string `json:"url,omitempty"`
}
