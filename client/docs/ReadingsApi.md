# \ReadingsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSensorReadingsHistory**](ReadingsApi.md#GetOrganizationSensorReadingsHistory) | **Get** /organizations/{organizationId}/sensor/readings/history | Return all reported readings from sensors in a given timespan, sorted by timestamp
[**GetOrganizationSensorReadingsLatest**](ReadingsApi.md#GetOrganizationSensorReadingsLatest) | **Get** /organizations/{organizationId}/sensor/readings/latest | Return the latest available reading for each metric from each sensor, sorted by sensor serial



## GetOrganizationSensorReadingsHistory

> []InlineResponse200129 GetOrganizationSensorReadingsHistory(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).Serials(serials).Metrics(metrics).Execute()

Return all reported readings from sensors in a given timespan, sorted by timestamp



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days and 6 hours from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter readings by network. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter readings by sensor. (optional)
    metrics := []string{"Inner_example"} // []string | Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are battery, button, door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReadingsApi.GetOrganizationSensorReadingsHistory(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).T0(t0).T1(t1).Timespan(timespan).NetworkIds(networkIds).Serials(serials).Metrics(metrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReadingsApi.GetOrganizationSensorReadingsHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSensorReadingsHistory`: []InlineResponse200129
    fmt.Fprintf(os.Stdout, "Response from `ReadingsApi.GetOrganizationSensorReadingsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSensorReadingsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days and 6 hours from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours. | 
 **networkIds** | **[]string** | Optional parameter to filter readings by network. | 
 **serials** | **[]string** | Optional parameter to filter readings by sensor. | 
 **metrics** | **[]string** | Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are battery, button, door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water. | 

### Return type

[**[]InlineResponse200129**](InlineResponse200129.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSensorReadingsLatest

> []InlineResponse200130 GetOrganizationSensorReadingsLatest(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Metrics(metrics).Execute()

Return the latest available reading for each metric from each sensor, sorted by sensor serial



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter readings by network. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter readings by sensor. (optional)
    metrics := []string{"Inner_example"} // []string | Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are battery, button, door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReadingsApi.GetOrganizationSensorReadingsLatest(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Metrics(metrics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReadingsApi.GetOrganizationSensorReadingsLatest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSensorReadingsLatest`: []InlineResponse200130
    fmt.Fprintf(os.Stdout, "Response from `ReadingsApi.GetOrganizationSensorReadingsLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSensorReadingsLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter readings by network. | 
 **serials** | **[]string** | Optional parameter to filter readings by sensor. | 
 **metrics** | **[]string** | Types of sensor readings to retrieve. If no metrics are supplied, all available types of readings will be retrieved. Allowed values are battery, button, door, humidity, indoorAirQuality, noise, pm25, temperature, tvoc, and water. | 

### Return type

[**[]InlineResponse200130**](InlineResponse200130.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

