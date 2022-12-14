/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// InsightsV1AccountSettings struct for InsightsV1AccountSettings
type InsightsV1AccountSettings struct {
	AccountSid       *string `json:"account_sid,omitempty"`
	AdvancedFeatures *bool   `json:"advanced_features,omitempty"`
	VoiceTrace       *bool   `json:"voice_trace,omitempty"`
	Url              *string `json:"url,omitempty"`
}
