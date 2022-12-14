/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// SupersimV1SettingsUpdate struct for SupersimV1SettingsUpdate
type SupersimV1SettingsUpdate struct {
	// The unique identifier of this Settings Update
	Sid *string `json:"sid,omitempty"`
	// The ICCID associated with the SIM
	Iccid *string `json:"iccid,omitempty"`
	// The SID of the Super SIM to which this Settings Update was applied
	SimSid *string `json:"sim_sid,omitempty"`
	Status *string `json:"status,omitempty"`
	// Array containing the different Settings Packages that will be applied to the SIM after the update completes
	Packages *[]interface{} `json:"packages,omitempty"`
	// The time when the update successfully completed and the new settings were applied to the SIM
	DateCompleted *time.Time `json:"date_completed,omitempty"`
	// The date this Settings Update was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Settings Update was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
}
