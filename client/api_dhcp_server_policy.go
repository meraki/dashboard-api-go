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


// DhcpServerPolicyApiService DhcpServerPolicyApi service
type DhcpServerPolicyApiService service

type DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
}

func (r DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest = &createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
	return r
}

func (r DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	return r.ApiService.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Add a server to be trusted by Dynamic ARP Inspection on this network

Add a server to be trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *DhcpServerPolicyApiService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string) DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
func (a *DhcpServerPolicyApiService) CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r DhcpServerPolicyApiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest == nil {
		return localVarReturnValue, nil, reportError("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest is required and must be specified")
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
	localVarPostBody = r.createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
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

type DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	trustedServerId string
}

func (r DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Remove a server from being trusted by Dynamic ARP Inspection on this network

Remove a server from being trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param trustedServerId Trusted server ID
 @return DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *DhcpServerPolicyApiService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string, trustedServerId string) DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		trustedServerId: trustedServerId,
	}
}

// Execute executes the request
func (a *DhcpServerPolicyApiService) DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r DhcpServerPolicyApiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trustedServerId"+"}", url.PathEscape(parameterToString(r.trustedServerId, "")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
}

func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest) Execute() (*GetNetworkSwitchDhcpServerPolicy200Response, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicy Return the DHCP server settings

Return the DHCP server settings. Blocked/allowed servers are only applied when default policy is allow/block, respectively

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest
*/
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicy(ctx context.Context, networkId string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest {
	return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return GetNetworkSwitchDhcpServerPolicy200Response
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyExecute(r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyRequest) (*GetNetworkSwitchDhcpServerPolicy200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSwitchDhcpServerPolicy200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.GetNetworkSwitchDhcpServerPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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

type DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) PerPage(perPage int32) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) StartingAfter(startingAfter string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) EndingBefore(endingBefore string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) Execute() ([]GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers Return the list of servers trusted by Dynamic ARP Inspection on this network

Return the list of servers trusted by Dynamic ARP Inspection on this network. These are also known as allow listed snoop entries

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest
*/
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(ctx context.Context, networkId string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest {
	return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersExecute(r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest) ([]GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
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

type DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	perPage *int32
	startingAfter *string
	endingBefore *string
}

// The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) PerPage(perPage int32) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) StartingAfter(startingAfter string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) EndingBefore(endingBefore string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

func (r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) Execute() ([]GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner, *http.Response, error) {
	return r.ApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceExecute(r)
}

/*
GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice Return the devices that have a Dynamic ARP Inspection warning and their warnings

Return the devices that have a Dynamic ARP Inspection warning and their warnings

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest
*/
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(ctx context.Context, networkId string) DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest {
	return DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return []GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner
func (a *DhcpServerPolicyApiService) GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceExecute(r DhcpServerPolicyApiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest) ([]GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/warnings/byDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.startingAfter != nil {
		localVarQueryParams.Add("startingAfter", parameterToString(*r.startingAfter, ""))
	}
	if r.endingBefore != nil {
		localVarQueryParams.Add("endingBefore", parameterToString(*r.endingBefore, ""))
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

type DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	updateNetworkSwitchDhcpServerPolicyRequest *UpdateNetworkSwitchDhcpServerPolicyRequest
}

func (r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest) UpdateNetworkSwitchDhcpServerPolicyRequest(updateNetworkSwitchDhcpServerPolicyRequest UpdateNetworkSwitchDhcpServerPolicyRequest) DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest {
	r.updateNetworkSwitchDhcpServerPolicyRequest = &updateNetworkSwitchDhcpServerPolicyRequest
	return r
}

func (r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchDhcpServerPolicyExecute(r)
}

/*
UpdateNetworkSwitchDhcpServerPolicy Update the DHCP server settings

Update the DHCP server settings. Blocked/allowed servers are only applied when default policy is allow/block, respectively

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @return DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest
*/
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicy(ctx context.Context, networkId string) DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest {
	return DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicyExecute(r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.UpdateNetworkSwitchDhcpServerPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)

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
	localVarPostBody = r.updateNetworkSwitchDhcpServerPolicyRequest
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

type DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct {
	ctx context.Context
	ApiService *DhcpServerPolicyApiService
	networkId string
	trustedServerId string
	updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest *UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
}

func (r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	r.updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest = &updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
	return r
}

func (r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) Execute() (*GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	return r.ApiService.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r)
}

/*
UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer Update a server that is trusted by Dynamic ARP Inspection on this network

Update a server that is trusted by Dynamic ARP Inspection on this network

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param networkId Network ID
 @param trustedServerId Trusted server ID
 @return DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
*/
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx context.Context, networkId string, trustedServerId string) DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest {
	return DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest{
		ApiService: a,
		ctx: ctx,
		networkId: networkId,
		trustedServerId: trustedServerId,
	}
}

// Execute executes the request
//  @return GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
func (a *DhcpServerPolicyApiService) UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerExecute(r DhcpServerPolicyApiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest) (*GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DhcpServerPolicyApiService.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"networkId"+"}", url.PathEscape(parameterToString(r.networkId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trustedServerId"+"}", url.PathEscape(parameterToString(r.trustedServerId, "")), -1)

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
	localVarPostBody = r.updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest
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
