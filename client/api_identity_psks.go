/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
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


// IdentityPsksApiService IdentityPsksApi service
type IdentityPsksApiService service

type IdentityPsksApiCreateNetworkWirelessSsidIdentityPskRequest struct {
	ctx context.Context
	ApiService *IdentityPsksApiService
	networkId string
	number string
	createNetworkWirelessSsidIdentityPskRequest *CreateNetworkWirelessSsidIdentityPskRequest
}

func (r IdentityPsksApiCreateNetworkWirelessSsidIdentityPskRequest) CreateNetworkWirelessSsidIdentityPskRequest(createNetworkWirelessSsidIdentityPskRequest CreateNetworkWirelessSsidIdentityPskRequest) IdentityPsksApiCreateNetworkWirelessSsidIdentityPskRequest {
	r.createNetworkWirelessSsidIdentityPskRequest = &createNetworkWirelessSsidIdentityPskRequest
	return r
}

func (r IdentityPsksApiCreateNetworkWirelessSsidIdentityPskRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CreateNetworkWirelessSsidIdentityPskExecute(r)
}

/*
CreateNetworkWirelessSsidIdentityPsk Create an Identity PSK

Create an Identity PSK

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param number Number
 @return IdentityPsksApiCreateNetworkWirelessSsidIdentityPskRequest
*/
func (a *IdentityPsksApiService) CreateNetworkWirelessSsidIdentityPsk(ctx context.Context, networkId string, number string) IdentityPsksApiCreateNetworkWirelessSsidIdentityPskRequest {
	return IdentityPsksApiCreateNetworkWirelessSsidIdentityPskRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *IdentityPsksApiService) CreateNetworkWirelessSsidIdentityPskExecute(r IdentityPsksApiCreateNetworkWirelessSsidIdentityPskRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPsksApiService.CreateNetworkWirelessSsidIdentityPsk")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/identityPsks"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", url.PathEscape(parameterValueToString(r.number, "number")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkWirelessSsidIdentityPskRequest == nil {
		return localVarReturnValue, nil, reportError("createNetworkWirelessSsidIdentityPskRequest is required and must be specified")
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
	localVarPostBody = r.createNetworkWirelessSsidIdentityPskRequest
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

type IdentityPsksApiDeleteNetworkWirelessSsidIdentityPskRequest struct {
	ctx context.Context
	ApiService *IdentityPsksApiService
	networkId string
	number string
	identityPskId string
}

func (r IdentityPsksApiDeleteNetworkWirelessSsidIdentityPskRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkWirelessSsidIdentityPskExecute(r)
}

/*
DeleteNetworkWirelessSsidIdentityPsk Delete an Identity PSK

Delete an Identity PSK

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param number Number
 @param identityPskId Identity psk ID
 @return IdentityPsksApiDeleteNetworkWirelessSsidIdentityPskRequest
*/
func (a *IdentityPsksApiService) DeleteNetworkWirelessSsidIdentityPsk(ctx context.Context, networkId string, number string, identityPskId string) IdentityPsksApiDeleteNetworkWirelessSsidIdentityPskRequest {
	return IdentityPsksApiDeleteNetworkWirelessSsidIdentityPskRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
		identityPskId: identityPskId,
	}
}

// Execute executes the request
func (a *IdentityPsksApiService) DeleteNetworkWirelessSsidIdentityPskExecute(r IdentityPsksApiDeleteNetworkWirelessSsidIdentityPskRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPsksApiService.DeleteNetworkWirelessSsidIdentityPsk")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", url.PathEscape(parameterValueToString(r.number, "number")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identityPskId"+"}", url.PathEscape(parameterValueToString(r.identityPskId, "identityPskId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type IdentityPsksApiGetNetworkWirelessSsidIdentityPskRequest struct {
	ctx context.Context
	ApiService *IdentityPsksApiService
	networkId string
	number string
	identityPskId string
}

func (r IdentityPsksApiGetNetworkWirelessSsidIdentityPskRequest) Execute() (*GetNetworkWirelessSsidIdentityPsks200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessSsidIdentityPskExecute(r)
}

/*
GetNetworkWirelessSsidIdentityPsk Return an Identity PSK

Return an Identity PSK

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param number Number
 @param identityPskId Identity psk ID
 @return IdentityPsksApiGetNetworkWirelessSsidIdentityPskRequest
*/
func (a *IdentityPsksApiService) GetNetworkWirelessSsidIdentityPsk(ctx context.Context, networkId string, number string, identityPskId string) IdentityPsksApiGetNetworkWirelessSsidIdentityPskRequest {
	return IdentityPsksApiGetNetworkWirelessSsidIdentityPskRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
		identityPskId: identityPskId,
	}
}

// Execute executes the request
//  @return GetNetworkWirelessSsidIdentityPsks200ResponseInner
func (a *IdentityPsksApiService) GetNetworkWirelessSsidIdentityPskExecute(r IdentityPsksApiGetNetworkWirelessSsidIdentityPskRequest) (*GetNetworkWirelessSsidIdentityPsks200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkWirelessSsidIdentityPsks200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPsksApiService.GetNetworkWirelessSsidIdentityPsk")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", url.PathEscape(parameterValueToString(r.number, "number")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identityPskId"+"}", url.PathEscape(parameterValueToString(r.identityPskId, "identityPskId")), -1)

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

type IdentityPsksApiGetNetworkWirelessSsidIdentityPsksRequest struct {
	ctx context.Context
	ApiService *IdentityPsksApiService
	networkId string
	number string
}

func (r IdentityPsksApiGetNetworkWirelessSsidIdentityPsksRequest) Execute() ([]GetNetworkWirelessSsidIdentityPsks200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkWirelessSsidIdentityPsksExecute(r)
}

/*
GetNetworkWirelessSsidIdentityPsks List all Identity PSKs in a wireless network

List all Identity PSKs in a wireless network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param number Number
 @return IdentityPsksApiGetNetworkWirelessSsidIdentityPsksRequest
*/
func (a *IdentityPsksApiService) GetNetworkWirelessSsidIdentityPsks(ctx context.Context, networkId string, number string) IdentityPsksApiGetNetworkWirelessSsidIdentityPsksRequest {
	return IdentityPsksApiGetNetworkWirelessSsidIdentityPsksRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
	}
}

// Execute executes the request
//  @return []GetNetworkWirelessSsidIdentityPsks200ResponseInner
func (a *IdentityPsksApiService) GetNetworkWirelessSsidIdentityPsksExecute(r IdentityPsksApiGetNetworkWirelessSsidIdentityPsksRequest) ([]GetNetworkWirelessSsidIdentityPsks200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkWirelessSsidIdentityPsks200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPsksApiService.GetNetworkWirelessSsidIdentityPsks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/identityPsks"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", url.PathEscape(parameterValueToString(r.number, "number")), -1)

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

type IdentityPsksApiUpdateNetworkWirelessSsidIdentityPskRequest struct {
	ctx context.Context
	ApiService *IdentityPsksApiService
	networkId string
	number string
	identityPskId string
	updateNetworkWirelessSsidIdentityPskRequest *UpdateNetworkWirelessSsidIdentityPskRequest
}

func (r IdentityPsksApiUpdateNetworkWirelessSsidIdentityPskRequest) UpdateNetworkWirelessSsidIdentityPskRequest(updateNetworkWirelessSsidIdentityPskRequest UpdateNetworkWirelessSsidIdentityPskRequest) IdentityPsksApiUpdateNetworkWirelessSsidIdentityPskRequest {
	r.updateNetworkWirelessSsidIdentityPskRequest = &updateNetworkWirelessSsidIdentityPskRequest
	return r
}

func (r IdentityPsksApiUpdateNetworkWirelessSsidIdentityPskRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkWirelessSsidIdentityPskExecute(r)
}

/*
UpdateNetworkWirelessSsidIdentityPsk Update an Identity PSK

Update an Identity PSK

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param number Number
 @param identityPskId Identity psk ID
 @return IdentityPsksApiUpdateNetworkWirelessSsidIdentityPskRequest
*/
func (a *IdentityPsksApiService) UpdateNetworkWirelessSsidIdentityPsk(ctx context.Context, networkId string, number string, identityPskId string) IdentityPsksApiUpdateNetworkWirelessSsidIdentityPskRequest {
	return IdentityPsksApiUpdateNetworkWirelessSsidIdentityPskRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		number: number,
		identityPskId: identityPskId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *IdentityPsksApiService) UpdateNetworkWirelessSsidIdentityPskExecute(r IdentityPsksApiUpdateNetworkWirelessSsidIdentityPskRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityPsksApiService.UpdateNetworkWirelessSsidIdentityPsk")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"number"+"}", url.PathEscape(parameterValueToString(r.number, "number")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"identityPskId"+"}", url.PathEscape(parameterValueToString(r.identityPskId, "identityPskId")), -1)

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
	localVarPostBody = r.updateNetworkWirelessSsidIdentityPskRequest
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
