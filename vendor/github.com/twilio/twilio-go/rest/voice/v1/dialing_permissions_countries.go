/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Retrieve voice dialing country permissions identified by the given ISO country code
func (c *ApiService) FetchDialingPermissionsCountry(IsoCode string) (*VoiceV1DialingPermissionsCountryInstance, error) {
	path := "/v1/DialingPermissions/Countries/{IsoCode}"
	path = strings.Replace(path, "{"+"IsoCode"+"}", IsoCode, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1DialingPermissionsCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListDialingPermissionsCountry'
type ListDialingPermissionsCountryParams struct {
	// Filter to retrieve the country permissions by specifying the [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)
	IsoCode *string `json:"IsoCode,omitempty"`
	// Filter to retrieve the country permissions by specifying the continent
	Continent *string `json:"Continent,omitempty"`
	// Filter the results by specified [country codes](https://www.itu.int/itudoc/itu-t/ob-lists/icc/e164_763.html)
	CountryCode *string `json:"CountryCode,omitempty"`
	// Filter to retrieve the country permissions with dialing to low-risk numbers enabled. Can be: `true` or `false`.
	LowRiskNumbersEnabled *bool `json:"LowRiskNumbersEnabled,omitempty"`
	// Filter to retrieve the country permissions with dialing to high-risk special service numbers enabled. Can be: `true` or `false`
	HighRiskSpecialNumbersEnabled *bool `json:"HighRiskSpecialNumbersEnabled,omitempty"`
	// Filter to retrieve the country permissions with dialing to high-risk [toll fraud](https://www.twilio.com/learn/voice-and-video/toll-fraud) numbers enabled. Can be: `true` or `false`.
	HighRiskTollfraudNumbersEnabled *bool `json:"HighRiskTollfraudNumbersEnabled,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListDialingPermissionsCountryParams) SetIsoCode(IsoCode string) *ListDialingPermissionsCountryParams {
	params.IsoCode = &IsoCode
	return params
}
func (params *ListDialingPermissionsCountryParams) SetContinent(Continent string) *ListDialingPermissionsCountryParams {
	params.Continent = &Continent
	return params
}
func (params *ListDialingPermissionsCountryParams) SetCountryCode(CountryCode string) *ListDialingPermissionsCountryParams {
	params.CountryCode = &CountryCode
	return params
}
func (params *ListDialingPermissionsCountryParams) SetLowRiskNumbersEnabled(LowRiskNumbersEnabled bool) *ListDialingPermissionsCountryParams {
	params.LowRiskNumbersEnabled = &LowRiskNumbersEnabled
	return params
}
func (params *ListDialingPermissionsCountryParams) SetHighRiskSpecialNumbersEnabled(HighRiskSpecialNumbersEnabled bool) *ListDialingPermissionsCountryParams {
	params.HighRiskSpecialNumbersEnabled = &HighRiskSpecialNumbersEnabled
	return params
}
func (params *ListDialingPermissionsCountryParams) SetHighRiskTollfraudNumbersEnabled(HighRiskTollfraudNumbersEnabled bool) *ListDialingPermissionsCountryParams {
	params.HighRiskTollfraudNumbersEnabled = &HighRiskTollfraudNumbersEnabled
	return params
}
func (params *ListDialingPermissionsCountryParams) SetPageSize(PageSize int) *ListDialingPermissionsCountryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListDialingPermissionsCountryParams) SetLimit(Limit int) *ListDialingPermissionsCountryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of DialingPermissionsCountry records from the API. Request is executed immediately.
func (c *ApiService) PageDialingPermissionsCountry(params *ListDialingPermissionsCountryParams, pageToken, pageNumber string) (*ListDialingPermissionsCountryResponse, error) {
	path := "/v1/DialingPermissions/Countries"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IsoCode != nil {
		data.Set("IsoCode", *params.IsoCode)
	}
	if params != nil && params.Continent != nil {
		data.Set("Continent", *params.Continent)
	}
	if params != nil && params.CountryCode != nil {
		data.Set("CountryCode", *params.CountryCode)
	}
	if params != nil && params.LowRiskNumbersEnabled != nil {
		data.Set("LowRiskNumbersEnabled", fmt.Sprint(*params.LowRiskNumbersEnabled))
	}
	if params != nil && params.HighRiskSpecialNumbersEnabled != nil {
		data.Set("HighRiskSpecialNumbersEnabled", fmt.Sprint(*params.HighRiskSpecialNumbersEnabled))
	}
	if params != nil && params.HighRiskTollfraudNumbersEnabled != nil {
		data.Set("HighRiskTollfraudNumbersEnabled", fmt.Sprint(*params.HighRiskTollfraudNumbersEnabled))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDialingPermissionsCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists DialingPermissionsCountry records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListDialingPermissionsCountry(params *ListDialingPermissionsCountryParams) ([]VoiceV1DialingPermissionsCountry, error) {
	response, errors := c.StreamDialingPermissionsCountry(params)

	records := make([]VoiceV1DialingPermissionsCountry, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams DialingPermissionsCountry records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamDialingPermissionsCountry(params *ListDialingPermissionsCountryParams) (chan VoiceV1DialingPermissionsCountry, chan error) {
	if params == nil {
		params = &ListDialingPermissionsCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan VoiceV1DialingPermissionsCountry, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageDialingPermissionsCountry(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamDialingPermissionsCountry(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamDialingPermissionsCountry(response *ListDialingPermissionsCountryResponse, params *ListDialingPermissionsCountryParams, recordChannel chan VoiceV1DialingPermissionsCountry, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Content
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListDialingPermissionsCountryResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListDialingPermissionsCountryResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListDialingPermissionsCountryResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListDialingPermissionsCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}