/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"net/url"
)

// Optional parameters for the method 'CreateGooddata'
type CreateGooddataParams struct {
	// The Token HTTP request header
	Token *string `json:"Token,omitempty"`
}

func (params *CreateGooddataParams) SetToken(Token string) *CreateGooddataParams {
	params.Token = &Token
	return params
}

// To create a GoodData Session id to access GoodData dashboards
func (c *ApiService) CreateGooddata(params *CreateGooddataParams) (*FlexV1Gooddata, error) {
	path := "/v1/Accounts/GoodData"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Token != nil {
		headers["Token"] = *params.Token
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Gooddata{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
