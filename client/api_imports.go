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
	"reflect"
)


// ImportsApiService ImportsApi service
type ImportsApiService service

type ImportsApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct {
	ctx context.Context
	ApiService *ImportsApiService
	organizationId string
	createOrganizationInventoryOnboardingCloudMonitoringImportRequest *CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest
}

func (r ImportsApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest(createOrganizationInventoryOnboardingCloudMonitoringImportRequest CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) ImportsApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	r.createOrganizationInventoryOnboardingCloudMonitoringImportRequest = &createOrganizationInventoryOnboardingCloudMonitoringImportRequest
	return r
}

func (r ImportsApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) Execute() ([]CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner, *http.Response, error) {
	return r.ApiService.CreateOrganizationInventoryOnboardingCloudMonitoringImportExecute(r)
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringImport Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ImportsApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest
*/
func (a *ImportsApiService) CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx context.Context, organizationId string) ImportsApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest {
	return ImportsApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner
func (a *ImportsApiService) CreateOrganizationInventoryOnboardingCloudMonitoringImportExecute(r ImportsApiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest) ([]CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.CreateOrganizationInventoryOnboardingCloudMonitoringImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOrganizationInventoryOnboardingCloudMonitoringImportRequest == nil {
		return localVarReturnValue, nil, reportError("createOrganizationInventoryOnboardingCloudMonitoringImportRequest is required and must be specified")
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
	localVarPostBody = r.createOrganizationInventoryOnboardingCloudMonitoringImportRequest
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

type ImportsApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct {
	ctx context.Context
	ApiService *ImportsApiService
	organizationId string
	importIds *[]string
}

// import ids from an imports
func (r ImportsApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) ImportIds(importIds []string) ImportsApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest {
	r.importIds = &importIds
	return r
}

func (r ImportsApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) Execute() ([]GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationInventoryOnboardingCloudMonitoringImportsExecute(r)
}

/*
GetOrganizationInventoryOnboardingCloudMonitoringImports Check the status of a committed Import operation

Check the status of a committed Import operation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return ImportsApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest
*/
func (a *ImportsApiService) GetOrganizationInventoryOnboardingCloudMonitoringImports(ctx context.Context, organizationId string) ImportsApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest {
	return ImportsApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner
func (a *ImportsApiService) GetOrganizationInventoryOnboardingCloudMonitoringImportsExecute(r ImportsApiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest) ([]GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ImportsApiService.GetOrganizationInventoryOnboardingCloudMonitoringImports")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importIds == nil {
		return localVarReturnValue, nil, reportError("importIds is required and must be specified")
	}

	{
		t := *r.importIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "importIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "importIds", t, "multi")
		}
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
