/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
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


// MerakiAuthUsersApiService MerakiAuthUsersApi service
type MerakiAuthUsersApiService service

type MerakiAuthUsersApiCreateNetworkMerakiAuthUserRequest struct {
	ctx context.Context
	ApiService *MerakiAuthUsersApiService
	networkId string
	createNetworkMerakiAuthUserRequest *CreateNetworkMerakiAuthUserRequest
}

func (r MerakiAuthUsersApiCreateNetworkMerakiAuthUserRequest) CreateNetworkMerakiAuthUserRequest(createNetworkMerakiAuthUserRequest CreateNetworkMerakiAuthUserRequest) MerakiAuthUsersApiCreateNetworkMerakiAuthUserRequest {
	r.createNetworkMerakiAuthUserRequest = &createNetworkMerakiAuthUserRequest
	return r
}

func (r MerakiAuthUsersApiCreateNetworkMerakiAuthUserRequest) Execute() (*GetNetworkMerakiAuthUsers200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateNetworkMerakiAuthUserExecute(r)
}

/*
CreateNetworkMerakiAuthUser Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)

Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return MerakiAuthUsersApiCreateNetworkMerakiAuthUserRequest
*/
func (a *MerakiAuthUsersApiService) CreateNetworkMerakiAuthUser(ctx context.Context, networkId string) MerakiAuthUsersApiCreateNetworkMerakiAuthUserRequest {
	return MerakiAuthUsersApiCreateNetworkMerakiAuthUserRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkMerakiAuthUsers200ResponseInner
func (a *MerakiAuthUsersApiService) CreateNetworkMerakiAuthUserExecute(r MerakiAuthUsersApiCreateNetworkMerakiAuthUserRequest) (*GetNetworkMerakiAuthUsers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkMerakiAuthUsers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerakiAuthUsersApiService.CreateNetworkMerakiAuthUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/merakiAuthUsers"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkMerakiAuthUserRequest == nil {
		return localVarReturnValue, nil, reportError("createNetworkMerakiAuthUserRequest is required and must be specified")
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
	localVarPostBody = r.createNetworkMerakiAuthUserRequest
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

type MerakiAuthUsersApiDeleteNetworkMerakiAuthUserRequest struct {
	ctx context.Context
	ApiService *MerakiAuthUsersApiService
	networkId string
	merakiAuthUserId string
	delete *bool
}

// If the ID supplied is for a splash guest or client VPN user, and that user is not authorized for any other networks in the organization, then also delete the user. 802.1X RADIUS users are always deleted regardless of this optional attribute.
func (r MerakiAuthUsersApiDeleteNetworkMerakiAuthUserRequest) Delete(delete bool) MerakiAuthUsersApiDeleteNetworkMerakiAuthUserRequest {
	r.delete = &delete
	return r
}

func (r MerakiAuthUsersApiDeleteNetworkMerakiAuthUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkMerakiAuthUserExecute(r)
}

/*
DeleteNetworkMerakiAuthUser Delete an 802.1X RADIUS user, or deauthorize and optionally delete a splash guest or client VPN user.

Delete an 802.1X RADIUS user, or deauthorize and optionally delete a splash guest or client VPN user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param merakiAuthUserId Meraki auth user ID
 @return MerakiAuthUsersApiDeleteNetworkMerakiAuthUserRequest
*/
func (a *MerakiAuthUsersApiService) DeleteNetworkMerakiAuthUser(ctx context.Context, networkId string, merakiAuthUserId string) MerakiAuthUsersApiDeleteNetworkMerakiAuthUserRequest {
	return MerakiAuthUsersApiDeleteNetworkMerakiAuthUserRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		merakiAuthUserId: merakiAuthUserId,
	}
}

// Execute executes the request
func (a *MerakiAuthUsersApiService) DeleteNetworkMerakiAuthUserExecute(r MerakiAuthUsersApiDeleteNetworkMerakiAuthUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerakiAuthUsersApiService.DeleteNetworkMerakiAuthUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/merakiAuthUsers/{merakiAuthUserId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"merakiAuthUserId"+"}", url.PathEscape(parameterValueToString(r.merakiAuthUserId, "merakiAuthUserId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.delete != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "delete", r.delete, "")
	}
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

type MerakiAuthUsersApiGetNetworkMerakiAuthUserRequest struct {
	ctx context.Context
	ApiService *MerakiAuthUsersApiService
	networkId string
	merakiAuthUserId string
}

func (r MerakiAuthUsersApiGetNetworkMerakiAuthUserRequest) Execute() (*GetNetworkMerakiAuthUsers200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkMerakiAuthUserExecute(r)
}

/*
GetNetworkMerakiAuthUser Return the Meraki Auth splash guest, RADIUS, or client VPN user

Return the Meraki Auth splash guest, RADIUS, or client VPN user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param merakiAuthUserId Meraki auth user ID
 @return MerakiAuthUsersApiGetNetworkMerakiAuthUserRequest
*/
func (a *MerakiAuthUsersApiService) GetNetworkMerakiAuthUser(ctx context.Context, networkId string, merakiAuthUserId string) MerakiAuthUsersApiGetNetworkMerakiAuthUserRequest {
	return MerakiAuthUsersApiGetNetworkMerakiAuthUserRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		merakiAuthUserId: merakiAuthUserId,
	}
}

// Execute executes the request
//  @return GetNetworkMerakiAuthUsers200ResponseInner
func (a *MerakiAuthUsersApiService) GetNetworkMerakiAuthUserExecute(r MerakiAuthUsersApiGetNetworkMerakiAuthUserRequest) (*GetNetworkMerakiAuthUsers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkMerakiAuthUsers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerakiAuthUsersApiService.GetNetworkMerakiAuthUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/merakiAuthUsers/{merakiAuthUserId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"merakiAuthUserId"+"}", url.PathEscape(parameterValueToString(r.merakiAuthUserId, "merakiAuthUserId")), -1)

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

type MerakiAuthUsersApiGetNetworkMerakiAuthUsersRequest struct {
	ctx context.Context
	ApiService *MerakiAuthUsersApiService
	networkId string
}

func (r MerakiAuthUsersApiGetNetworkMerakiAuthUsersRequest) Execute() ([]GetNetworkMerakiAuthUsers200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkMerakiAuthUsersExecute(r)
}

/*
GetNetworkMerakiAuthUsers List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a MX network)

List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a MX network)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return MerakiAuthUsersApiGetNetworkMerakiAuthUsersRequest
*/
func (a *MerakiAuthUsersApiService) GetNetworkMerakiAuthUsers(ctx context.Context, networkId string) MerakiAuthUsersApiGetNetworkMerakiAuthUsersRequest {
	return MerakiAuthUsersApiGetNetworkMerakiAuthUsersRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkMerakiAuthUsers200ResponseInner
func (a *MerakiAuthUsersApiService) GetNetworkMerakiAuthUsersExecute(r MerakiAuthUsersApiGetNetworkMerakiAuthUsersRequest) ([]GetNetworkMerakiAuthUsers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkMerakiAuthUsers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerakiAuthUsersApiService.GetNetworkMerakiAuthUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/merakiAuthUsers"
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

type MerakiAuthUsersApiUpdateNetworkMerakiAuthUserRequest struct {
	ctx context.Context
	ApiService *MerakiAuthUsersApiService
	networkId string
	merakiAuthUserId string
	updateNetworkMerakiAuthUserRequest *UpdateNetworkMerakiAuthUserRequest
}

func (r MerakiAuthUsersApiUpdateNetworkMerakiAuthUserRequest) UpdateNetworkMerakiAuthUserRequest(updateNetworkMerakiAuthUserRequest UpdateNetworkMerakiAuthUserRequest) MerakiAuthUsersApiUpdateNetworkMerakiAuthUserRequest {
	r.updateNetworkMerakiAuthUserRequest = &updateNetworkMerakiAuthUserRequest
	return r
}

func (r MerakiAuthUsersApiUpdateNetworkMerakiAuthUserRequest) Execute() (*GetNetworkMerakiAuthUsers200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateNetworkMerakiAuthUserExecute(r)
}

/*
UpdateNetworkMerakiAuthUser Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)

Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param merakiAuthUserId Meraki auth user ID
 @return MerakiAuthUsersApiUpdateNetworkMerakiAuthUserRequest
*/
func (a *MerakiAuthUsersApiService) UpdateNetworkMerakiAuthUser(ctx context.Context, networkId string, merakiAuthUserId string) MerakiAuthUsersApiUpdateNetworkMerakiAuthUserRequest {
	return MerakiAuthUsersApiUpdateNetworkMerakiAuthUserRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		merakiAuthUserId: merakiAuthUserId,
	}
}

// Execute executes the request
//  @return GetNetworkMerakiAuthUsers200ResponseInner
func (a *MerakiAuthUsersApiService) UpdateNetworkMerakiAuthUserExecute(r MerakiAuthUsersApiUpdateNetworkMerakiAuthUserRequest) (*GetNetworkMerakiAuthUsers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkMerakiAuthUsers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MerakiAuthUsersApiService.UpdateNetworkMerakiAuthUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/merakiAuthUsers/{merakiAuthUserId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterValueToString(r.networkId, "networkId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"merakiAuthUserId"+"}", url.PathEscape(parameterValueToString(r.merakiAuthUserId, "merakiAuthUserId")), -1)

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
	localVarPostBody = r.updateNetworkMerakiAuthUserRequest
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
