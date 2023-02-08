# \TrafficAnalysisApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkTrafficAnalysis**](TrafficAnalysisApi.md#GetNetworkTrafficAnalysis) | **Get** /networks/{networkId}/trafficAnalysis | Return the traffic analysis settings for a network
[**UpdateNetworkTrafficAnalysis**](TrafficAnalysisApi.md#UpdateNetworkTrafficAnalysis) | **Put** /networks/{networkId}/trafficAnalysis | Update the traffic analysis settings for a network



## GetNetworkTrafficAnalysis

> map[string]interface{} GetNetworkTrafficAnalysis(ctx, networkId).Execute()

Return the traffic analysis settings for a network



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
    networkId := "networkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrafficAnalysisApi.GetNetworkTrafficAnalysis(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrafficAnalysisApi.GetNetworkTrafficAnalysis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkTrafficAnalysis`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TrafficAnalysisApi.GetNetworkTrafficAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkTrafficAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkTrafficAnalysis

> map[string]interface{} UpdateNetworkTrafficAnalysis(ctx, networkId).UpdateNetworkTrafficAnalysis(updateNetworkTrafficAnalysis).Execute()

Update the traffic analysis settings for a network



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
    networkId := "networkId_example" // string | 
    updateNetworkTrafficAnalysis := *openapiclient.NewUpdateNetworkTrafficAnalysisRequest() // UpdateNetworkTrafficAnalysisRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TrafficAnalysisApi.UpdateNetworkTrafficAnalysis(context.Background(), networkId).UpdateNetworkTrafficAnalysis(updateNetworkTrafficAnalysis).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TrafficAnalysisApi.UpdateNetworkTrafficAnalysis``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkTrafficAnalysis`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TrafficAnalysisApi.UpdateNetworkTrafficAnalysis`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkTrafficAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkTrafficAnalysis** | [**UpdateNetworkTrafficAnalysisRequest**](UpdateNetworkTrafficAnalysisRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

