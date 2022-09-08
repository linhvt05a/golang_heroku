/*
 * Twilio - Api
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

// Optional parameters for the method 'ListUsageRecordLastMonth'
type ListUsageRecordLastMonthParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the UsageRecord resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The [usage category](https://www.twilio.com/docs/usage/api/usage-record#usage-categories) of the UsageRecord resources to read. Only UsageRecord resources in the specified category are retrieved.
	Category *string `json:"Category,omitempty"`
	// Only include usage that has occurred on or after this date. Specify the date in GMT and format as `YYYY-MM-DD`. You can also specify offsets from the current date, such as: `-30days`, which will set the start date to be 30 days before the current date.
	StartDate *string `json:"StartDate,omitempty"`
	// Only include usage that occurred on or before this date. Specify the date in GMT and format as `YYYY-MM-DD`.  You can also specify offsets from the current date, such as: `+30days`, which will set the end date to 30 days from the current date.
	EndDate *string `json:"EndDate,omitempty"`
	// Whether to include usage from the master account and all its subaccounts. Can be: `true` (the default) to include usage from the master account and all subaccounts or `false` to retrieve usage from only the specified account.
	IncludeSubaccounts *bool `json:"IncludeSubaccounts,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUsageRecordLastMonthParams) SetPathAccountSid(PathAccountSid string) *ListUsageRecordLastMonthParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListUsageRecordLastMonthParams) SetCategory(Category string) *ListUsageRecordLastMonthParams {
	params.Category = &Category
	return params
}
func (params *ListUsageRecordLastMonthParams) SetStartDate(StartDate string) *ListUsageRecordLastMonthParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListUsageRecordLastMonthParams) SetEndDate(EndDate string) *ListUsageRecordLastMonthParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListUsageRecordLastMonthParams) SetIncludeSubaccounts(IncludeSubaccounts bool) *ListUsageRecordLastMonthParams {
	params.IncludeSubaccounts = &IncludeSubaccounts
	return params
}
func (params *ListUsageRecordLastMonthParams) SetPageSize(PageSize int) *ListUsageRecordLastMonthParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUsageRecordLastMonthParams) SetLimit(Limit int) *ListUsageRecordLastMonthParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UsageRecordLastMonth records from the API. Request is executed immediately.
func (c *ApiService) PageUsageRecordLastMonth(params *ListUsageRecordLastMonthParams, pageToken, pageNumber string) (*ListUsageRecordLastMonthResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/Usage/Records/LastMonth.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Category != nil {
		data.Set("Category", *params.Category)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint(*params.StartDate))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint(*params.EndDate))
	}
	if params != nil && params.IncludeSubaccounts != nil {
		data.Set("IncludeSubaccounts", fmt.Sprint(*params.IncludeSubaccounts))
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

	ps := &ListUsageRecordLastMonthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UsageRecordLastMonth records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUsageRecordLastMonth(params *ListUsageRecordLastMonthParams) ([]ApiV2010UsageRecordLastMonth, error) {
	response, errors := c.StreamUsageRecordLastMonth(params)

	records := make([]ApiV2010UsageRecordLastMonth, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams UsageRecordLastMonth records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUsageRecordLastMonth(params *ListUsageRecordLastMonthParams) (chan ApiV2010UsageRecordLastMonth, chan error) {
	if params == nil {
		params = &ListUsageRecordLastMonthParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ApiV2010UsageRecordLastMonth, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageUsageRecordLastMonth(params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamUsageRecordLastMonth(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamUsageRecordLastMonth(response *ListUsageRecordLastMonthResponse, params *ListUsageRecordLastMonthParams, recordChannel chan ApiV2010UsageRecordLastMonth, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.UsageRecords
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListUsageRecordLastMonthResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListUsageRecordLastMonthResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListUsageRecordLastMonthResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageRecordLastMonthResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
