/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ConversationsV1Configuration struct for ConversationsV1Configuration
type ConversationsV1Configuration struct {
	// The SID of the Account responsible for this configuration.
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the default Conversation Service that every new conversation is associated with.
	DefaultChatServiceSid *string `json:"default_chat_service_sid,omitempty"`
	// The SID of the default Messaging Service that every new conversation is associated with.
	DefaultMessagingServiceSid *string `json:"default_messaging_service_sid,omitempty"`
	// Default ISO8601 duration when conversation will be switched to `inactive` state.
	DefaultInactiveTimer *string `json:"default_inactive_timer,omitempty"`
	// Default ISO8601 duration when conversation will be switched to `closed` state.
	DefaultClosedTimer *string `json:"default_closed_timer,omitempty"`
	// An absolute URL for this global configuration.
	Url *string `json:"url,omitempty"`
	// Absolute URLs to access the webhook and default service configurations.
	Links *map[string]interface{} `json:"links,omitempty"`
}
