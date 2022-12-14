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

// InsightsV1Event struct for InsightsV1Event
type InsightsV1Event struct {
	Timestamp   *string      `json:"timestamp,omitempty"`
	CallSid     *string      `json:"call_sid,omitempty"`
	AccountSid  *string      `json:"account_sid,omitempty"`
	Edge        *string      `json:"edge,omitempty"`
	Group       *string      `json:"group,omitempty"`
	Level       *string      `json:"level,omitempty"`
	Name        *string      `json:"name,omitempty"`
	CarrierEdge *interface{} `json:"carrier_edge,omitempty"`
	SipEdge     *interface{} `json:"sip_edge,omitempty"`
	SdkEdge     *interface{} `json:"sdk_edge,omitempty"`
	ClientEdge  *interface{} `json:"client_edge,omitempty"`
}
