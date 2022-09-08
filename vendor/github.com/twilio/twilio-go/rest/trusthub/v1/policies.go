/*
 * Twilio - Trusthub
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

// Fetch specific Policy Instance.
func (c *ApiService) FetchPolicies(Sid string) (*TrusthubV1Policies, error) {
	path := "/v1/Policies/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &TrusthubV1Policies{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListPolicies'
type ListPoliciesParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListPoliciesParams) SetPageSize(PageSize int) *ListPoliciesParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListPoliciesParams) SetLimit(Limit int) *ListPoliciesParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Policies records from the API. Request is executed immediately.
func (c *ApiService) PagePolicies(params *ListPoliciesParams, pageToken, pageNumber string) (*ListPoliciesResponse, error) {
	path := "/v1/Policies"

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ListPoliciesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Policies records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListPolicies(params *ListPoliciesParams) ([]TrusthubV1Policies, error) {
	response, errors := c.StreamPolicies(params)

	records := make([]TrusthubV1Policies, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Policies records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamPolicies(params *ListPoliciesParams) (chan TrusthubV1Policies, chan error) {
	if params == nil {
		params = &ListPoliciesParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan TrusthubV1Policies, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PagePolicies(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamPolicies(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamPolicies(response *ListPoliciesResponse, params *ListPoliciesParams, recordChannel chan TrusthubV1Policies, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Results
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListPoliciesResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListPoliciesResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListPoliciesResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPoliciesResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
