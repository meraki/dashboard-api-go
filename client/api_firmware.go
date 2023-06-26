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
	"reflect"
)


// FirmwareApiService FirmwareApi service
type FirmwareApiService service

type FirmwareApiGetOrganizationFirmwareUpgradesRequest struct {
	ctx context.Context
	ApiService *FirmwareApiService
	organizationId string
	status *[]string
	productType *[]string
}

// The status of an upgrade 
func (r FirmwareApiGetOrganizationFirmwareUpgradesRequest) Status(status []string) FirmwareApiGetOrganizationFirmwareUpgradesRequest {
	r.status = &status
	return r
}

// The product type in a given upgrade ID
func (r FirmwareApiGetOrganizationFirmwareUpgradesRequest) ProductType(productType []string) FirmwareApiGetOrganizationFirmwareUpgradesRequest {
	r.productType = &productType
	return r
}

func (r FirmwareApiGetOrganizationFirmwareUpgradesRequest) Execute() ([]GetOrganizationFirmwareUpgrades200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationFirmwareUpgradesExecute(r)
}

/*
GetOrganizationFirmwareUpgrades Get firmware upgrade information for an organization

Get firmware upgrade information for an organization

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return FirmwareApiGetOrganizationFirmwareUpgradesRequest
*/
func (a *FirmwareApiService) GetOrganizationFirmwareUpgrades(ctx context.Context, organizationId string) FirmwareApiGetOrganizationFirmwareUpgradesRequest {
	return FirmwareApiGetOrganizationFirmwareUpgradesRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationFirmwareUpgrades200ResponseInner
func (a *FirmwareApiService) GetOrganizationFirmwareUpgradesExecute(r FirmwareApiGetOrganizationFirmwareUpgradesRequest) ([]GetOrganizationFirmwareUpgrades200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationFirmwareUpgrades200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareApiService.GetOrganizationFirmwareUpgrades")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/firmware/upgrades"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		t := *r.status
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("status", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("status", parameterToString(t, "multi"))
		}
	}
	if r.productType != nil {
		t := *r.productType
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("productType", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("productType", parameterToString(t, "multi"))
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

type FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest struct {
	ctx context.Context
	ApiService *FirmwareApiService
	organizationId string
	perPage *int32
	startingAfter *string
	endingBefore *string
	networkIds *[]string
	serials *[]string
	macs *[]string
	firmwareUpgradeIds *[]string
	firmwareUpgradeBatchIds *[]string
}

// The number of entries per page returned. Acceptable range is 3 - 50. Default is 50.
func (r FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest) PerPage(perPage int32) FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.perPage = &perPage
	return r
}

// A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest) StartingAfter(startingAfter string) FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.startingAfter = &startingAfter
	return r
}

// A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.
func (r FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest) EndingBefore(endingBefore string) FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.endingBefore = &endingBefore
	return r
}

// Optional parameter to filter by network
func (r FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest) NetworkIds(networkIds []string) FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.networkIds = &networkIds
	return r
}

// Optional parameter to filter by serial number.  All returned devices will have a serial number that is an exact match.
func (r FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest) Serials(serials []string) FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.serials = &serials
	return r
}

// Optional parameter to filter by one or more MAC addresses belonging to devices. All devices returned belong to MAC addresses that are an exact match.
func (r FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest) Macs(macs []string) FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.macs = &macs
	return r
}

// Optional parameter to filter by firmware upgrade ids.
func (r FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest) FirmwareUpgradeIds(firmwareUpgradeIds []string) FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.firmwareUpgradeIds = &firmwareUpgradeIds
	return r
}

// Optional parameter to filter by firmware upgrade batch ids.
func (r FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest) FirmwareUpgradeBatchIds(firmwareUpgradeBatchIds []string) FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	r.firmwareUpgradeBatchIds = &firmwareUpgradeBatchIds
	return r
}

func (r FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest) Execute() ([]GetOrganizationFirmwareUpgradesByDevice200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationFirmwareUpgradesByDeviceExecute(r)
}

/*
GetOrganizationFirmwareUpgradesByDevice Get firmware upgrade status for the filtered devices

Get firmware upgrade status for the filtered devices

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Organization ID
 @return FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest
*/
func (a *FirmwareApiService) GetOrganizationFirmwareUpgradesByDevice(ctx context.Context, organizationId string) FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest {
	return FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationFirmwareUpgradesByDevice200ResponseInner
func (a *FirmwareApiService) GetOrganizationFirmwareUpgradesByDeviceExecute(r FirmwareApiGetOrganizationFirmwareUpgradesByDeviceRequest) ([]GetOrganizationFirmwareUpgradesByDevice200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationFirmwareUpgradesByDevice200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FirmwareApiService.GetOrganizationFirmwareUpgradesByDevice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/firmware/upgrades/byDevice"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

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
	if r.networkIds != nil {
		t := *r.networkIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("networkIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("networkIds", parameterToString(t, "multi"))
		}
	}
	if r.serials != nil {
		t := *r.serials
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("serials", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("serials", parameterToString(t, "multi"))
		}
	}
	if r.macs != nil {
		t := *r.macs
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("macs", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("macs", parameterToString(t, "multi"))
		}
	}
	if r.firmwareUpgradeIds != nil {
		t := *r.firmwareUpgradeIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("firmwareUpgradeIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("firmwareUpgradeIds", parameterToString(t, "multi"))
		}
	}
	if r.firmwareUpgradeBatchIds != nil {
		t := *r.firmwareUpgradeBatchIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("firmwareUpgradeBatchIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("firmwareUpgradeBatchIds", parameterToString(t, "multi"))
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
