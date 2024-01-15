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
)


// HttpServersApiService HttpServersApi service
type HttpServersApiService service

type HttpServersApiCreateNetworkWebhooksHttpServerRequest struct {
	ctx context.Context
	ApiService *HttpServersApiService
	networkId string
	createNetworkWebhooksHttpServerRequest *CreateNetworkWebhooksHttpServerRequest
}

func (r HttpServersApiCreateNetworkWebhooksHttpServerRequest) CreateNetworkWebhooksHttpServerRequest(createNetworkWebhooksHttpServerRequest CreateNetworkWebhooksHttpServerRequest) HttpServersApiCreateNetworkWebhooksHttpServerRequest {
	r.createNetworkWebhooksHttpServerRequest = &createNetworkWebhooksHttpServerRequest
	return r
}

func (r HttpServersApiCreateNetworkWebhooksHttpServerRequest) Execute() (*GetNetworkWebhooksHttpServers200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateNetworkWebhooksHttpServerExecute(r)
}

/*
CreateNetworkWebhooksHttpServer Add an HTTP server to a network

Add an HTTP server to a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return HttpServersApiCreateNetworkWebhooksHttpServerRequest
*/
func (a *HttpServersApiService) CreateNetworkWebhooksHttpServer(ctx context.Context, networkId string) HttpServersApiCreateNetworkWebhooksHttpServerRequest {
	return HttpServersApiCreateNetworkWebhooksHttpServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkWebhooksHttpServers200ResponseInner
func (a *HttpServersApiService) CreateNetworkWebhooksHttpServerExecute(r HttpServersApiCreateNetworkWebhooksHttpServerRequest) (*GetNetworkWebhooksHttpServers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkWebhooksHttpServers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServersApiService.CreateNetworkWebhooksHttpServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/webhooks/httpServers"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkWebhooksHttpServerRequest == nil {
		return localVarReturnValue, nil, reportError("createNetworkWebhooksHttpServerRequest is required and must be specified")
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
	localVarPostBody = r.createNetworkWebhooksHttpServerRequest
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

type HttpServersApiDeleteNetworkWebhooksHttpServerRequest struct {
	ctx context.Context
	ApiService *HttpServersApiService
	networkId string
	httpServerId string
}

func (r HttpServersApiDeleteNetworkWebhooksHttpServerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkWebhooksHttpServerExecute(r)
}

/*
DeleteNetworkWebhooksHttpServer Delete an HTTP server from a network

Delete an HTTP server from a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param httpServerId Http server ID
 @return HttpServersApiDeleteNetworkWebhooksHttpServerRequest
*/
func (a *HttpServersApiService) DeleteNetworkWebhooksHttpServer(ctx context.Context, networkId string, httpServerId string) HttpServersApiDeleteNetworkWebhooksHttpServerRequest {
	return HttpServersApiDeleteNetworkWebhooksHttpServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		httpServerId: httpServerId,
	}
}

// Execute executes the request
func (a *HttpServersApiService) DeleteNetworkWebhooksHttpServerExecute(r HttpServersApiDeleteNetworkWebhooksHttpServerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServersApiService.DeleteNetworkWebhooksHttpServer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/webhooks/httpServers/{httpServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"httpServerId"+"}", url.PathEscape(parameterValueToString(r.httpServerId, "httpServerId")), -1)

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

type HttpServersApiGetNetworkWebhooksHttpServerRequest struct {
	ctx context.Context
	ApiService *HttpServersApiService
	networkId string
	httpServerId string
}

func (r HttpServersApiGetNetworkWebhooksHttpServerRequest) Execute() (*GetNetworkWebhooksHttpServers200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkWebhooksHttpServerExecute(r)
}

/*
GetNetworkWebhooksHttpServer Return an HTTP server for a network

Return an HTTP server for a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param httpServerId Http server ID
 @return HttpServersApiGetNetworkWebhooksHttpServerRequest
*/
func (a *HttpServersApiService) GetNetworkWebhooksHttpServer(ctx context.Context, networkId string, httpServerId string) HttpServersApiGetNetworkWebhooksHttpServerRequest {
	return HttpServersApiGetNetworkWebhooksHttpServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		httpServerId: httpServerId,
	}
}

// Execute executes the request
//  @return GetNetworkWebhooksHttpServers200ResponseInner
func (a *HttpServersApiService) GetNetworkWebhooksHttpServerExecute(r HttpServersApiGetNetworkWebhooksHttpServerRequest) (*GetNetworkWebhooksHttpServers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkWebhooksHttpServers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServersApiService.GetNetworkWebhooksHttpServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/webhooks/httpServers/{httpServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"httpServerId"+"}", url.PathEscape(parameterValueToString(r.httpServerId, "httpServerId")), -1)

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

type HttpServersApiGetNetworkWebhooksHttpServersRequest struct {
	ctx context.Context
	ApiService *HttpServersApiService
	networkId string
}

func (r HttpServersApiGetNetworkWebhooksHttpServersRequest) Execute() ([]GetNetworkWebhooksHttpServers200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkWebhooksHttpServersExecute(r)
}

/*
GetNetworkWebhooksHttpServers List the HTTP servers for a network

List the HTTP servers for a network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return HttpServersApiGetNetworkWebhooksHttpServersRequest
*/
func (a *HttpServersApiService) GetNetworkWebhooksHttpServers(ctx context.Context, networkId string) HttpServersApiGetNetworkWebhooksHttpServersRequest {
	return HttpServersApiGetNetworkWebhooksHttpServersRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkWebhooksHttpServers200ResponseInner
func (a *HttpServersApiService) GetNetworkWebhooksHttpServersExecute(r HttpServersApiGetNetworkWebhooksHttpServersRequest) ([]GetNetworkWebhooksHttpServers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkWebhooksHttpServers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServersApiService.GetNetworkWebhooksHttpServers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/webhooks/httpServers"
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

type HttpServersApiUpdateNetworkWebhooksHttpServerRequest struct {
	ctx context.Context
	ApiService *HttpServersApiService
	networkId string
	httpServerId string
	updateNetworkWebhooksHttpServerRequest *UpdateNetworkWebhooksHttpServerRequest
}

func (r HttpServersApiUpdateNetworkWebhooksHttpServerRequest) UpdateNetworkWebhooksHttpServerRequest(updateNetworkWebhooksHttpServerRequest UpdateNetworkWebhooksHttpServerRequest) HttpServersApiUpdateNetworkWebhooksHttpServerRequest {
	r.updateNetworkWebhooksHttpServerRequest = &updateNetworkWebhooksHttpServerRequest
	return r
}

func (r HttpServersApiUpdateNetworkWebhooksHttpServerRequest) Execute() (*GetNetworkWebhooksHttpServers200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateNetworkWebhooksHttpServerExecute(r)
}

/*
UpdateNetworkWebhooksHttpServer Update an HTTP server

Update an HTTP server. To change a URL, create a new HTTP server.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param httpServerId Http server ID
 @return HttpServersApiUpdateNetworkWebhooksHttpServerRequest
*/
func (a *HttpServersApiService) UpdateNetworkWebhooksHttpServer(ctx context.Context, networkId string, httpServerId string) HttpServersApiUpdateNetworkWebhooksHttpServerRequest {
	return HttpServersApiUpdateNetworkWebhooksHttpServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		httpServerId: httpServerId,
	}
}

// Execute executes the request
//  @return GetNetworkWebhooksHttpServers200ResponseInner
func (a *HttpServersApiService) UpdateNetworkWebhooksHttpServerExecute(r HttpServersApiUpdateNetworkWebhooksHttpServerRequest) (*GetNetworkWebhooksHttpServers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkWebhooksHttpServers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HttpServersApiService.UpdateNetworkWebhooksHttpServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/webhooks/httpServers/{httpServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"httpServerId"+"}", url.PathEscape(parameterValueToString(r.httpServerId, "httpServerId")), -1)

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
	localVarPostBody = r.updateNetworkWebhooksHttpServerRequest
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
