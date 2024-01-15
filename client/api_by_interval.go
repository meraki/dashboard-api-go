/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// ByIntervalApiService ByIntervalApi service
type ByIntervalApiService service

type ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest struct {
	ctx context.Context
	ApiService *ByIntervalApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
	interval *int32
	version *int32
	operationIds *[]string
	sourceIps *[]string
	adminIds *[]string
	userAgent *string
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) T0(t0 string) ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) T1(t1 string) ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 31 days. If interval is provided, the timespan will be autocalculated.
func (r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Timespan(timespan float32) ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.timespan = &timespan
	return r
}

// The time interval in seconds for returned data. The valid intervals are: 120, 3600, 14400, 21600. The default is 21600. Interval is calculated if time params are provided.
func (r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Interval(interval int32) ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.interval = &interval
	return r
}

// Filter by API version of the endpoint. Allowable values are: [0, 1]
func (r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Version(version int32) ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.version = &version
	return r
}

// Filter by operation ID of the endpoint
func (r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) OperationIds(operationIds []string) ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.operationIds = &operationIds
	return r
}

// Filter by source IP that made the API request
func (r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) SourceIps(sourceIps []string) ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.sourceIps = &sourceIps
	return r
}

// Filter by admin ID of user that made the API request
func (r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) AdminIds(adminIds []string) ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.adminIds = &adminIds
	return r
}

// Filter by user agent string for API request. This will filter by a complete or partial match.
func (r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) UserAgent(userAgent string) ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	r.userAgent = &userAgent
	return r
}

func (r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) Execute() ([]GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationApiRequestsOverviewResponseCodesByIntervalExecute(r)
}

/*
GetOrganizationApiRequestsOverviewResponseCodesByInterval Tracks organizations' API requests by response code across a given time period

Tracks organizations' API requests by response code across a given time period

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest
*/
func (a *ByIntervalApiService) GetOrganizationApiRequestsOverviewResponseCodesByInterval(ctx context.Context, organizationId string) ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest {
	return ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner
func (a *ByIntervalApiService) GetOrganizationApiRequestsOverviewResponseCodesByIntervalExecute(r ByIntervalApiGetOrganizationApiRequestsOverviewResponseCodesByIntervalRequest) ([]GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationApiRequestsOverviewResponseCodesByInterval200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByIntervalApiService.GetOrganizationApiRequestsOverviewResponseCodesByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/apiRequests/overview/responseCodes/byInterval"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
	}
	if r.version != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
	}
	if r.operationIds != nil {
		t := *r.operationIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "operationIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "operationIds", t, "multi")
		}
	}
	if r.sourceIps != nil {
		t := *r.sourceIps
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "sourceIps", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "sourceIps", t, "multi")
		}
	}
	if r.adminIds != nil {
		t := *r.adminIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "adminIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "adminIds", t, "multi")
		}
	}
	if r.userAgent != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "userAgent", r.userAgent, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest struct {
	ctx context.Context
	ApiService *ByIntervalApiService
	organizationId string
	networkIds *[]string
	serials *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	interval *int32
}

// Filter results by network.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) NetworkIds(networkIds []string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest {
	r.networkIds = &networkIds
	return r
}

// Filter results by device.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) Serials(serials []string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest {
	r.serials = &serials
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) PerPage(perPage int32) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) StartingAfter(startingAfter string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) EndingBefore(endingBefore string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) T0(t0 string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) T1(t1 string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) Timespan(timespan float32) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest {
	r.timespan = &timespan
	return r
}

// The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) Interval(interval int32) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest {
	r.interval = &interval
	return r
}

func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) Execute() ([]GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalExecute(r)
}

/*
GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval Get a time-series of average channel utilization for all bands, segmented by device.

Get a time-series of average channel utilization for all bands, segmented by device.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest
*/
func (a *ByIntervalApiService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval(ctx context.Context, organizationId string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest {
	return ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner
func (a *ByIntervalApiService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalExecute(r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByIntervalRequest) ([]GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByIntervalApiService.GetOrganizationWirelessDevicesChannelUtilizationHistoryByDeviceByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/devices/channelUtilization/history/byDevice/byInterval"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", t, "multi")
		}
	}
	if r.serials != nil {
		t := *r.serials
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "serials", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "serials", t, "multi")
		}
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingAfter", r.startingAfter, "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endingBefore", r.endingBefore, "")
	}
	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest struct {
	ctx context.Context
	ApiService *ByIntervalApiService
	organizationId string
	networkIds *[]string
	serials *[]string
	perPage *int32
	startingAfter *string
	endingBefore *string
	t0 *string
	t1 *string
	timespan *float32
	interval *int32
}

// Filter results by network.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) NetworkIds(networkIds []string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.networkIds = &networkIds
	return r
}

// Filter results by device.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) Serials(serials []string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.serials = &serials
	return r
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) PerPage(perPage int32) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) StartingAfter(startingAfter string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) EndingBefore(endingBefore string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.endingBefore = &endingBefore
	return r
}

// The beginning of the timespan for the data. The maximum lookback period is 31 days from today.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) T0(t0 string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 31 days after t0.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) T1(t1 string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) Timespan(timespan float32) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.timespan = &timespan
	return r
}

// The time interval in seconds for returned data. The valid intervals are: 300, 600, 3600, 7200, 14400, 21600. The default is 3600.
func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) Interval(interval int32) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	r.interval = &interval
	return r
}

func (r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) Execute() ([]GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalExecute(r)
}

/*
GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval Get a time-series of average channel utilization for all bands

Get a time-series of average channel utilization for all bands

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest
*/
func (a *ByIntervalApiService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval(ctx context.Context, organizationId string) ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest {
	return ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner
func (a *ByIntervalApiService) GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalExecute(r ByIntervalApiGetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByIntervalRequest) ([]GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ByIntervalApiService.GetOrganizationWirelessDevicesChannelUtilizationHistoryByNetworkByInterval")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/wireless/devices/channelUtilization/history/byNetwork/byInterval"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "networkIds", t, "multi")
		}
	}
	if r.serials != nil {
		t := *r.serials
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "serials", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "serials", t, "multi")
		}
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perPage", r.perPage, "")
	}
	if r.startingAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startingAfter", r.startingAfter, "")
	}
	if r.endingBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endingBefore", r.endingBefore, "")
	}
	if r.t0 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t0", r.t0, "")
	}
	if r.t1 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "t1", r.t1, "")
	}
	if r.timespan != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timespan", r.timespan, "")
	}
	if r.interval != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "interval", r.interval, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
