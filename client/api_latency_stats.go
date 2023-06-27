/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
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


// LatencyStatsApiService LatencyStatsApi service
type LatencyStatsApiService service

type LatencyStatsApiGetDeviceWirelessLatencyStatsRequest struct {
	ctx context.Context
	ApiService *LatencyStatsApiService
	serial string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	fields *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r LatencyStatsApiGetDeviceWirelessLatencyStatsRequest) T0(t0 string) LatencyStatsApiGetDeviceWirelessLatencyStatsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r LatencyStatsApiGetDeviceWirelessLatencyStatsRequest) T1(t1 string) LatencyStatsApiGetDeviceWirelessLatencyStatsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r LatencyStatsApiGetDeviceWirelessLatencyStatsRequest) Timespan(timespan float32) LatencyStatsApiGetDeviceWirelessLatencyStatsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r LatencyStatsApiGetDeviceWirelessLatencyStatsRequest) Band(band string) LatencyStatsApiGetDeviceWirelessLatencyStatsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r LatencyStatsApiGetDeviceWirelessLatencyStatsRequest) Ssid(ssid int32) LatencyStatsApiGetDeviceWirelessLatencyStatsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r LatencyStatsApiGetDeviceWirelessLatencyStatsRequest) Vlan(vlan int32) LatencyStatsApiGetDeviceWirelessLatencyStatsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r LatencyStatsApiGetDeviceWirelessLatencyStatsRequest) ApTag(apTag string) LatencyStatsApiGetDeviceWirelessLatencyStatsRequest {
	r.apTag = &apTag
	return r
}

// Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
func (r LatencyStatsApiGetDeviceWirelessLatencyStatsRequest) Fields(fields string) LatencyStatsApiGetDeviceWirelessLatencyStatsRequest {
	r.fields = &fields
	return r
}

func (r LatencyStatsApiGetDeviceWirelessLatencyStatsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetDeviceWirelessLatencyStatsExecute(r)
}

/*
GetDeviceWirelessLatencyStats Aggregated latency info for a given AP on this network

Aggregated latency info for a given AP on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serial Serial
 @return LatencyStatsApiGetDeviceWirelessLatencyStatsRequest
*/
func (a *LatencyStatsApiService) GetDeviceWirelessLatencyStats(ctx context.Context, serial string) LatencyStatsApiGetDeviceWirelessLatencyStatsRequest {
	return LatencyStatsApiGetDeviceWirelessLatencyStatsRequest{
		ApiService: a,
		ctx: ctx,
		serial: serial,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LatencyStatsApiService) GetDeviceWirelessLatencyStatsExecute(r LatencyStatsApiGetDeviceWirelessLatencyStatsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyStatsApiService.GetDeviceWirelessLatencyStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/devices/{serial}/wireless/latencyStats"
	localVarPath = strings.Replace(localVarPath, "{"+"serial"+"}", url.PathEscape(parameterToString(r.serial, "")), -1)

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
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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

type LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest struct {
	ctx context.Context
	ApiService *LatencyStatsApiService
	networkId string
	clientId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	fields *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest) T0(t0 string) LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest) T1(t1 string) LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest) Timespan(timespan float32) LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest) Band(band string) LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest) Ssid(ssid int32) LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest) Vlan(vlan int32) LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest) ApTag(apTag string) LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest {
	r.apTag = &apTag
	return r
}

// Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
func (r LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest) Fields(fields string) LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest {
	r.fields = &fields
	return r
}

func (r LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessClientLatencyStatsExecute(r)
}

/*
GetNetworkWirelessClientLatencyStats Aggregated latency info for a given client on this network

Aggregated latency info for a given client on this network. Clients are identified by their MAC.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param clientId Client ID
 @return LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest
*/
func (a *LatencyStatsApiService) GetNetworkWirelessClientLatencyStats(ctx context.Context, networkId string, clientId string) LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest {
	return LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		clientId: clientId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LatencyStatsApiService) GetNetworkWirelessClientLatencyStatsExecute(r LatencyStatsApiGetNetworkWirelessClientLatencyStatsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyStatsApiService.GetNetworkWirelessClientLatencyStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/clients/{clientId}/latencyStats"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clientId"+"}", url.PathEscape(parameterToString(r.clientId, "")), -1)

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
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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

type LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest struct {
	ctx context.Context
	ApiService *LatencyStatsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	fields *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest) T0(t0 string) LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest) T1(t1 string) LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest) Timespan(timespan float32) LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest) Band(band string) LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest) Ssid(ssid int32) LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest) Vlan(vlan int32) LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest) ApTag(apTag string) LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.apTag = &apTag
	return r
}

// Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
func (r LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest) Fields(fields string) LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest {
	r.fields = &fields
	return r
}

func (r LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessClientsLatencyStatsExecute(r)
}

/*
GetNetworkWirelessClientsLatencyStats Aggregated latency info for this network, grouped by clients

Aggregated latency info for this network, grouped by clients

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest
*/
func (a *LatencyStatsApiService) GetNetworkWirelessClientsLatencyStats(ctx context.Context, networkId string) LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest {
	return LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *LatencyStatsApiService) GetNetworkWirelessClientsLatencyStatsExecute(r LatencyStatsApiGetNetworkWirelessClientsLatencyStatsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyStatsApiService.GetNetworkWirelessClientsLatencyStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/clients/latencyStats"
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
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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

type LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest struct {
	ctx context.Context
	ApiService *LatencyStatsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	fields *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest) T0(t0 string) LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest) T1(t1 string) LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest) Timespan(timespan float32) LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest) Band(band string) LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest) Ssid(ssid int32) LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest) Vlan(vlan int32) LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest) ApTag(apTag string) LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.apTag = &apTag
	return r
}

// Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
func (r LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest) Fields(fields string) LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest {
	r.fields = &fields
	return r
}

func (r LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest) Execute() ([]map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessDevicesLatencyStatsExecute(r)
}

/*
GetNetworkWirelessDevicesLatencyStats Aggregated latency info for this network, grouped by node

Aggregated latency info for this network, grouped by node

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest
*/
func (a *LatencyStatsApiService) GetNetworkWirelessDevicesLatencyStats(ctx context.Context, networkId string) LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest {
	return LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []map[string]interface{}
func (a *LatencyStatsApiService) GetNetworkWirelessDevicesLatencyStatsExecute(r LatencyStatsApiGetNetworkWirelessDevicesLatencyStatsRequest) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyStatsApiService.GetNetworkWirelessDevicesLatencyStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/devices/latencyStats"
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
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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

type LatencyStatsApiGetNetworkWirelessLatencyStatsRequest struct {
	ctx context.Context
	ApiService *LatencyStatsApiService
	networkId string
	t0 *string
	t1 *string
	timespan *float32
	band *string
	ssid *int32
	vlan *int32
	apTag *string
	fields *string
}

// The beginning of the timespan for the data. The maximum lookback period is 180 days from today.
func (r LatencyStatsApiGetNetworkWirelessLatencyStatsRequest) T0(t0 string) LatencyStatsApiGetNetworkWirelessLatencyStatsRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 7 days after t0.
func (r LatencyStatsApiGetNetworkWirelessLatencyStatsRequest) T1(t1 string) LatencyStatsApiGetNetworkWirelessLatencyStatsRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.
func (r LatencyStatsApiGetNetworkWirelessLatencyStatsRequest) Timespan(timespan float32) LatencyStatsApiGetNetworkWirelessLatencyStatsRequest {
	r.timespan = &timespan
	return r
}

// Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information.
func (r LatencyStatsApiGetNetworkWirelessLatencyStatsRequest) Band(band string) LatencyStatsApiGetNetworkWirelessLatencyStatsRequest {
	r.band = &band
	return r
}

// Filter results by SSID
func (r LatencyStatsApiGetNetworkWirelessLatencyStatsRequest) Ssid(ssid int32) LatencyStatsApiGetNetworkWirelessLatencyStatsRequest {
	r.ssid = &ssid
	return r
}

// Filter results by VLAN
func (r LatencyStatsApiGetNetworkWirelessLatencyStatsRequest) Vlan(vlan int32) LatencyStatsApiGetNetworkWirelessLatencyStatsRequest {
	r.vlan = &vlan
	return r
}

// Filter results by AP Tag
func (r LatencyStatsApiGetNetworkWirelessLatencyStatsRequest) ApTag(apTag string) LatencyStatsApiGetNetworkWirelessLatencyStatsRequest {
	r.apTag = &apTag
	return r
}

// Partial selection: If present, this call will return only the selected fields of [\&quot;rawDistribution\&quot;, \&quot;avg\&quot;]. All fields will be returned by default. Selected fields must be entered as a comma separated string.
func (r LatencyStatsApiGetNetworkWirelessLatencyStatsRequest) Fields(fields string) LatencyStatsApiGetNetworkWirelessLatencyStatsRequest {
	r.fields = &fields
	return r
}

func (r LatencyStatsApiGetNetworkWirelessLatencyStatsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessLatencyStatsExecute(r)
}

/*
GetNetworkWirelessLatencyStats Aggregated latency info for this network

Aggregated latency info for this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return LatencyStatsApiGetNetworkWirelessLatencyStatsRequest
*/
func (a *LatencyStatsApiService) GetNetworkWirelessLatencyStats(ctx context.Context, networkId string) LatencyStatsApiGetNetworkWirelessLatencyStatsRequest {
	return LatencyStatsApiGetNetworkWirelessLatencyStatsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LatencyStatsApiService) GetNetworkWirelessLatencyStatsExecute(r LatencyStatsApiGetNetworkWirelessLatencyStatsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LatencyStatsApiService.GetNetworkWirelessLatencyStats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/latencyStats"
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
	if r.fields != nil {
		localVarQueryParams.Add("fields", parameterToString(*r.fields, ""))
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
