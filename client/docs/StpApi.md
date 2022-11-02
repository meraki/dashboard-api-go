# \StpApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchStp**](StpApi.md#GetNetworkSwitchStp) | **Get** /networks/{networkId}/switch/stp | Returns STP settings
[**UpdateNetworkSwitchStp**](StpApi.md#UpdateNetworkSwitchStp) | **Put** /networks/{networkId}/switch/stp | Updates STP settings



## GetNetworkSwitchStp

> map[string]interface{} GetNetworkSwitchStp(ctx, networkId).Execute()

Returns STP settings



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
    resp, r, err := apiClient.StpApi.GetNetworkSwitchStp(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StpApi.GetNetworkSwitchStp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchStp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StpApi.GetNetworkSwitchStp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchStpRequest struct via the builder pattern


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


## UpdateNetworkSwitchStp

> map[string]interface{} UpdateNetworkSwitchStp(ctx, networkId).UpdateNetworkSwitchStp(updateNetworkSwitchStp).Execute()

Updates STP settings



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
    updateNetworkSwitchStp := *openapiclient.NewUpdateNetworkSwitchStpRequest() // UpdateNetworkSwitchStpRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StpApi.UpdateNetworkSwitchStp(context.Background(), networkId).UpdateNetworkSwitchStp(updateNetworkSwitchStp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StpApi.UpdateNetworkSwitchStp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchStp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StpApi.UpdateNetworkSwitchStp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchStpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchStp** | [**UpdateNetworkSwitchStpRequest**](UpdateNetworkSwitchStpRequest.md) |  | 

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

