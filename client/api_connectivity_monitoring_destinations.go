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


// ConnectivityMonitoringDestinationsApiService ConnectivityMonitoringDestinationsApi service
type ConnectivityMonitoringDestinationsApiService service

type ConnectivityMonitoringDestinationsApiGetNetworkApplianceConnectivityMonitoringDestinationsRequest struct {
	ctx context.Context
	ApiService *ConnectivityMonitoringDestinationsApiService
	networkId string
}

func (r ConnectivityMonitoringDestinationsApiGetNetworkApplianceConnectivityMonitoringDestinationsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkApplianceConnectivityMonitoringDestinationsExecute(r)
}

/*
GetNetworkApplianceConnectivityMonitoringDestinations Return the connectivity testing destinations for an MX network

Return the connectivity testing destinations for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ConnectivityMonitoringDestinationsApiGetNetworkApplianceConnectivityMonitoringDestinationsRequest
*/
func (a *ConnectivityMonitoringDestinationsApiService) GetNetworkApplianceConnectivityMonitoringDestinations(ctx context.Context, networkId string) ConnectivityMonitoringDestinationsApiGetNetworkApplianceConnectivityMonitoringDestinationsRequest {
	return ConnectivityMonitoringDestinationsApiGetNetworkApplianceConnectivityMonitoringDestinationsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConnectivityMonitoringDestinationsApiService) GetNetworkApplianceConnectivityMonitoringDestinationsExecute(r ConnectivityMonitoringDestinationsApiGetNetworkApplianceConnectivityMonitoringDestinationsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectivityMonitoringDestinationsApiService.GetNetworkApplianceConnectivityMonitoringDestinations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/connectivityMonitoringDestinations"
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

type ConnectivityMonitoringDestinationsApiGetNetworkCellularGatewayConnectivityMonitoringDestinationsRequest struct {
	ctx context.Context
	ApiService *ConnectivityMonitoringDestinationsApiService
	networkId string
}

func (r ConnectivityMonitoringDestinationsApiGetNetworkCellularGatewayConnectivityMonitoringDestinationsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.GetNetworkCellularGatewayConnectivityMonitoringDestinationsExecute(r)
}

/*
GetNetworkCellularGatewayConnectivityMonitoringDestinations Return the connectivity testing destinations for an MG network

Return the connectivity testing destinations for an MG network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ConnectivityMonitoringDestinationsApiGetNetworkCellularGatewayConnectivityMonitoringDestinationsRequest
*/
func (a *ConnectivityMonitoringDestinationsApiService) GetNetworkCellularGatewayConnectivityMonitoringDestinations(ctx context.Context, networkId string) ConnectivityMonitoringDestinationsApiGetNetworkCellularGatewayConnectivityMonitoringDestinationsRequest {
	return ConnectivityMonitoringDestinationsApiGetNetworkCellularGatewayConnectivityMonitoringDestinationsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConnectivityMonitoringDestinationsApiService) GetNetworkCellularGatewayConnectivityMonitoringDestinationsExecute(r ConnectivityMonitoringDestinationsApiGetNetworkCellularGatewayConnectivityMonitoringDestinationsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectivityMonitoringDestinationsApiService.GetNetworkCellularGatewayConnectivityMonitoringDestinations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/cellularGateway/connectivityMonitoringDestinations"
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

type ConnectivityMonitoringDestinationsApiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest struct {
	ctx context.Context
	ApiService *ConnectivityMonitoringDestinationsApiService
	networkId string
	updateNetworkApplianceConnectivityMonitoringDestinationsRequest *UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest
}

func (r ConnectivityMonitoringDestinationsApiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest(updateNetworkApplianceConnectivityMonitoringDestinationsRequest UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) ConnectivityMonitoringDestinationsApiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest {
	r.updateNetworkApplianceConnectivityMonitoringDestinationsRequest = &updateNetworkApplianceConnectivityMonitoringDestinationsRequest
	return r
}

func (r ConnectivityMonitoringDestinationsApiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkApplianceConnectivityMonitoringDestinationsExecute(r)
}

/*
UpdateNetworkApplianceConnectivityMonitoringDestinations Update the connectivity testing destinations for an MX network

Update the connectivity testing destinations for an MX network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ConnectivityMonitoringDestinationsApiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest
*/
func (a *ConnectivityMonitoringDestinationsApiService) UpdateNetworkApplianceConnectivityMonitoringDestinations(ctx context.Context, networkId string) ConnectivityMonitoringDestinationsApiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest {
	return ConnectivityMonitoringDestinationsApiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConnectivityMonitoringDestinationsApiService) UpdateNetworkApplianceConnectivityMonitoringDestinationsExecute(r ConnectivityMonitoringDestinationsApiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectivityMonitoringDestinationsApiService.UpdateNetworkApplianceConnectivityMonitoringDestinations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/appliance/connectivityMonitoringDestinations"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

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
	localVarPostBody = r.updateNetworkApplianceConnectivityMonitoringDestinationsRequest
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

type ConnectivityMonitoringDestinationsApiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest struct {
	ctx context.Context
	ApiService *ConnectivityMonitoringDestinationsApiService
	networkId string
	updateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest
}

func (r ConnectivityMonitoringDestinationsApiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest) UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest(updateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest) ConnectivityMonitoringDestinationsApiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest {
	r.updateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest = &updateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest
	return r
}

func (r ConnectivityMonitoringDestinationsApiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsExecute(r)
}

/*
UpdateNetworkCellularGatewayConnectivityMonitoringDestinations Update the connectivity testing destinations for an MG network

Update the connectivity testing destinations for an MG network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return ConnectivityMonitoringDestinationsApiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest
*/
func (a *ConnectivityMonitoringDestinationsApiService) UpdateNetworkCellularGatewayConnectivityMonitoringDestinations(ctx context.Context, networkId string) ConnectivityMonitoringDestinationsApiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest {
	return ConnectivityMonitoringDestinationsApiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *ConnectivityMonitoringDestinationsApiService) UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsExecute(r ConnectivityMonitoringDestinationsApiUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConnectivityMonitoringDestinationsApiService.UpdateNetworkCellularGatewayConnectivityMonitoringDestinations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/cellularGateway/connectivityMonitoringDestinations"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

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
	localVarPostBody = r.updateNetworkCellularGatewayConnectivityMonitoringDestinationsRequest
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
