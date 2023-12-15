# \PowerApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSummarySwitchPowerHistory**](PowerApi.md#GetOrganizationSummarySwitchPowerHistory) | **Get** /organizations/{organizationId}/summary/switch/power/history | Returns the total PoE power draw for all switch ports in the organization over the requested timespan (by default the last 24 hours)



## GetOrganizationSummarySwitchPowerHistory

> []GetOrganizationSummarySwitchPowerHistory200ResponseInner GetOrganizationSummarySwitchPowerHistory(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Returns the total PoE power draw for all switch ports in the organization over the requested timespan (by default the last 24 hours)



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PowerApi.GetOrganizationSummarySwitchPowerHistory(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PowerApi.GetOrganizationSummarySwitchPowerHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummarySwitchPowerHistory`: []GetOrganizationSummarySwitchPowerHistory200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `PowerApi.GetOrganizationSummarySwitchPowerHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummarySwitchPowerHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationSummarySwitchPowerHistory200ResponseInner**](GetOrganizationSummarySwitchPowerHistory200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

