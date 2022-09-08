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

// ConversationsV1ServiceWebhookConfiguration struct for ConversationsV1ServiceWebhookConfiguration
type ConversationsV1ServiceWebhookConfiguration struct {
	// The unique ID of the Account responsible for this service.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique ID of the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource) this conversation belongs to.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The list of events that your configured webhook targets will receive. Events not configured here will not fire.
	Filters *[]string `json:"filters,omitempty"`
	// The HTTP method to be used when sending a webhook request
	Method *string `json:"method,omitempty"`
	// The absolute url the post-event webhook request should be sent to.
	PostWebhookUrl *string `json:"post_webhook_url,omitempty"`
	// The absolute url the pre-event webhook request should be sent to.
	PreWebhookUrl *string `json:"pre_webhook_url,omitempty"`
	// An absolute URL for this webhook.
	Url *string `json:"url,omitempty"`
}
