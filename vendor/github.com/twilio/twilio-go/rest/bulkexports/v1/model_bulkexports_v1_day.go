/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Bulkexports
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// BulkexportsV1Day struct for BulkexportsV1Day
type BulkexportsV1Day struct {
	// The date of the data in the file
	Day *string `json:"day,omitempty"`
	// Size of the file in bytes
	Size *int `json:"size,omitempty"`
	// The date when resource is created
	CreateDate *string `json:"create_date,omitempty"`
	// The friendly name specified when creating the job
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
}
