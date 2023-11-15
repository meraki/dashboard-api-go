/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
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
)


// FloorPlansApiService FloorPlansApi service
type FloorPlansApiService service

type FloorPlansApiCreateNetworkFloorPlanRequest struct {
	ctx context.Context
	ApiService *FloorPlansApiService
	networkId string
	createNetworkFloorPlanRequest *CreateNetworkFloorPlanRequest
}

func (r FloorPlansApiCreateNetworkFloorPlanRequest) CreateNetworkFloorPlanRequest(createNetworkFloorPlanRequest CreateNetworkFloorPlanRequest) FloorPlansApiCreateNetworkFloorPlanRequest {
	r.createNetworkFloorPlanRequest = &createNetworkFloorPlanRequest
	return r
}

func (r FloorPlansApiCreateNetworkFloorPlanRequest) Execute() (*GetNetworkFloorPlans200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateNetworkFloorPlanExecute(r)
}

/*
CreateNetworkFloorPlan Upload a floor plan

Upload a floor plan

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return FloorPlansApiCreateNetworkFloorPlanRequest
*/
func (a *FloorPlansApiService) CreateNetworkFloorPlan(ctx context.Context, networkId string) FloorPlansApiCreateNetworkFloorPlanRequest {
	return FloorPlansApiCreateNetworkFloorPlanRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkFloorPlans200ResponseInner
func (a *FloorPlansApiService) CreateNetworkFloorPlanExecute(r FloorPlansApiCreateNetworkFloorPlanRequest) (*GetNetworkFloorPlans200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkFloorPlans200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FloorPlansApiService.CreateNetworkFloorPlan")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/floorPlans"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkFloorPlanRequest == nil {
		return localVarReturnValue, nil, reportError("createNetworkFloorPlanRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.createNetworkFloorPlanRequest
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

type FloorPlansApiDeleteNetworkFloorPlanRequest struct {
	ctx context.Context
	ApiService *FloorPlansApiService
	networkId string
	floorPlanId string
}

func (r FloorPlansApiDeleteNetworkFloorPlanRequest) Execute() (*GetNetworkFloorPlans200ResponseInner, *http.Response, error) {
	return r.ApiService.DeleteNetworkFloorPlanExecute(r)
}

/*
DeleteNetworkFloorPlan Destroy a floor plan

Destroy a floor plan

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param floorPlanId Floor plan ID
 @return FloorPlansApiDeleteNetworkFloorPlanRequest
*/
func (a *FloorPlansApiService) DeleteNetworkFloorPlan(ctx context.Context, networkId string, floorPlanId string) FloorPlansApiDeleteNetworkFloorPlanRequest {
	return FloorPlansApiDeleteNetworkFloorPlanRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		floorPlanId: floorPlanId,
	}
}

// Execute executes the request
//  @return GetNetworkFloorPlans200ResponseInner
func (a *FloorPlansApiService) DeleteNetworkFloorPlanExecute(r FloorPlansApiDeleteNetworkFloorPlanRequest) (*GetNetworkFloorPlans200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkFloorPlans200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FloorPlansApiService.DeleteNetworkFloorPlan")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/floorPlans/{floorPlanId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"floorPlanId"+"}", url.PathEscape(parameterValueToString(r.floorPlanId, "floorPlanId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type FloorPlansApiGetNetworkFloorPlanRequest struct {
	ctx context.Context
	ApiService *FloorPlansApiService
	networkId string
	floorPlanId string
}

func (r FloorPlansApiGetNetworkFloorPlanRequest) Execute() (*GetNetworkFloorPlans200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkFloorPlanExecute(r)
}

/*
GetNetworkFloorPlan Find a floor plan by ID

Find a floor plan by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param floorPlanId Floor plan ID
 @return FloorPlansApiGetNetworkFloorPlanRequest
*/
func (a *FloorPlansApiService) GetNetworkFloorPlan(ctx context.Context, networkId string, floorPlanId string) FloorPlansApiGetNetworkFloorPlanRequest {
	return FloorPlansApiGetNetworkFloorPlanRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		floorPlanId: floorPlanId,
	}
}

// Execute executes the request
//  @return GetNetworkFloorPlans200ResponseInner
func (a *FloorPlansApiService) GetNetworkFloorPlanExecute(r FloorPlansApiGetNetworkFloorPlanRequest) (*GetNetworkFloorPlans200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkFloorPlans200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FloorPlansApiService.GetNetworkFloorPlan")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/floorPlans/{floorPlanId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"floorPlanId"+"}", url.PathEscape(parameterValueToString(r.floorPlanId, "floorPlanId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type FloorPlansApiGetNetworkFloorPlansRequest struct {
	ctx context.Context
	ApiService *FloorPlansApiService
	networkId string
}

func (r FloorPlansApiGetNetworkFloorPlansRequest) Execute() ([]GetNetworkFloorPlans200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkFloorPlansExecute(r)
}

/*
GetNetworkFloorPlans List the floor plans that belong to your network

List the floor plans that belong to your network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return FloorPlansApiGetNetworkFloorPlansRequest
*/
func (a *FloorPlansApiService) GetNetworkFloorPlans(ctx context.Context, networkId string) FloorPlansApiGetNetworkFloorPlansRequest {
	return FloorPlansApiGetNetworkFloorPlansRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkFloorPlans200ResponseInner
func (a *FloorPlansApiService) GetNetworkFloorPlansExecute(r FloorPlansApiGetNetworkFloorPlansRequest) ([]GetNetworkFloorPlans200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkFloorPlans200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FloorPlansApiService.GetNetworkFloorPlans")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/floorPlans"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type FloorPlansApiUpdateNetworkFloorPlanRequest struct {
	ctx context.Context
	ApiService *FloorPlansApiService
	networkId string
	floorPlanId string
	updateNetworkFloorPlanRequest *UpdateNetworkFloorPlanRequest
}

func (r FloorPlansApiUpdateNetworkFloorPlanRequest) UpdateNetworkFloorPlanRequest(updateNetworkFloorPlanRequest UpdateNetworkFloorPlanRequest) FloorPlansApiUpdateNetworkFloorPlanRequest {
	r.updateNetworkFloorPlanRequest = &updateNetworkFloorPlanRequest
	return r
}

func (r FloorPlansApiUpdateNetworkFloorPlanRequest) Execute() (*GetNetworkFloorPlans200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateNetworkFloorPlanExecute(r)
}

/*
UpdateNetworkFloorPlan Update a floor plan's geolocation and other meta data

Update a floor plan's geolocation and other meta data

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param floorPlanId Floor plan ID
 @return FloorPlansApiUpdateNetworkFloorPlanRequest
*/
func (a *FloorPlansApiService) UpdateNetworkFloorPlan(ctx context.Context, networkId string, floorPlanId string) FloorPlansApiUpdateNetworkFloorPlanRequest {
	return FloorPlansApiUpdateNetworkFloorPlanRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		floorPlanId: floorPlanId,
	}
}

// Execute executes the request
//  @return GetNetworkFloorPlans200ResponseInner
func (a *FloorPlansApiService) UpdateNetworkFloorPlanExecute(r FloorPlansApiUpdateNetworkFloorPlanRequest) (*GetNetworkFloorPlans200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkFloorPlans200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FloorPlansApiService.UpdateNetworkFloorPlan")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/floorPlans/{floorPlanId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"floorPlanId"+"}", url.PathEscape(parameterValueToString(r.floorPlanId, "floorPlanId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.updateNetworkFloorPlanRequest
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
