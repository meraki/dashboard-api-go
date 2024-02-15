# \ByBoundaryApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCameraDetectionsHistoryByBoundaryByInterval**](ByBoundaryApi.md#GetOrganizationCameraDetectionsHistoryByBoundaryByInterval) | **Get** /organizations/{organizationId}/camera/detections/history/byBoundary/byInterval | Returns analytics data for timespans



## GetOrganizationCameraDetectionsHistoryByBoundaryByInterval

> []GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner GetOrganizationCameraDetectionsHistoryByBoundaryByInterval(ctx, organizationId).BoundaryIds(boundaryIds).Ranges(ranges).Duration(duration).PerPage(perPage).BoundaryTypes(boundaryTypes).Execute()

Returns analytics data for timespans



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    boundaryIds := []string{"Inner_example"} // []string | A list of boundary ids. The returned cameras will be filtered to only include these ids.
    ranges := []openapiclient.GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner{*openapiclient.NewGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner("StartTime_example", "EndTime_example", int32(123))} // []GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner | A list of time ranges with intervals
    duration := int32(56) // int32 | The minimum time, in seconds, that the person or car remains in the area to be counted. Defaults to boundary configuration or 60. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 1 - 1000. Defaults to 1000. (optional)
    boundaryTypes := []string{"BoundaryTypes_example"} // []string | The detection types. Defaults to 'person'. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ByBoundaryApi.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval(context.Background(), organizationId).BoundaryIds(boundaryIds).Ranges(ranges).Duration(duration).PerPage(perPage).BoundaryTypes(boundaryTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ByBoundaryApi.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraDetectionsHistoryByBoundaryByInterval`: []GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ByBoundaryApi.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **boundaryIds** | **[]string** | A list of boundary ids. The returned cameras will be filtered to only include these ids. | 
 **ranges** | [**[]GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner**](GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalRangesParameterInner.md) | A list of time ranges with intervals | 
 **duration** | **int32** | The minimum time, in seconds, that the person or car remains in the area to be counted. Defaults to boundary configuration or 60. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 1 - 1000. Defaults to 1000. | 
 **boundaryTypes** | **[]string** | The detection types. Defaults to &#39;person&#39;. | 

### Return type

[**[]GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner**](GetOrganizationCameraDetectionsHistoryByBoundaryByInterval200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

