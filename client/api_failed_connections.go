/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// FailedConnectionsApiService FailedConnectionsApi service
type FailedConnectionsApiService service

type FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest struct {
	ctx context.Context
	ApiService *FailedConnectionsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	serial *string
	clientId *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) T0(t0 string) FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) T1(t1 string) FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) Timespan(timespan float32) FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) Band(band string) FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) Ssid(ssid int32) FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) Vlan(vlan int32) FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) ApTag(apTag string) FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest {
	r.apTag = &apTag
	return r
}

// Filter by AP
func (r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) Serial(serial string) FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest {
	r.serial = &serial
	return r
}

// Filter by client MAC
func (r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) ClientId(clientId string) FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest {
	r.clientId = &clientId
	return r
}

func (r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) Execute() ([]GetNetworkWirelessFailedConnections200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessFailedConnectionsExecute(r)
}

/*
GetNetworkWirelessFailedConnections List of all failed client connection events on this network in a given time range

List of all failed client connection events on this network in a given time range

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId
 @return FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest
*/
func (a *FailedConnectionsApiService) GetNetworkWirelessFailedConnections(ctx context.Context, networkId string) FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest {
	return FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkWirelessFailedConnections200ResponseInner
func (a *FailedConnectionsApiService) GetNetworkWirelessFailedConnectionsExecute(r FailedConnectionsApiGetNetworkWirelessFailedConnectionsRequest) ([]GetNetworkWirelessFailedConnections200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkWirelessFailedConnections200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FailedConnectionsApiService.GetNetworkWirelessFailedConnections")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/failedConnections"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	if r.band != nil {
		localVarQueryParams.Add("band", parameterToString(*r.band, ""))
	}
	if r.ssid != nil {
		localVarQueryParams.Add("ssid", parameterToString(*r.ssid, ""))
	}
	if r.vlan != nil {
		localVarQueryParams.Add("vlan", parameterToString(*r.vlan, ""))
	}
	if r.apTag != nil {
		localVarQueryParams.Add("apTag", parameterToString(*r.apTag, ""))
	}
	if r.serial != nil {
		localVarQueryParams.Add("serial", parameterToString(*r.serial, ""))
	}
	if r.clientId != nil {
		localVarQueryParams.Add("clientId", parameterToString(*r.clientId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
