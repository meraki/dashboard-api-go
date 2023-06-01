# \MtuApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchMtu**](MtuApi.md#GetNetworkSwitchMtu) | **Get** /networks/{networkId}/switch/mtu | Return the MTU configuration
[**UpdateNetworkSwitchMtu**](MtuApi.md#UpdateNetworkSwitchMtu) | **Put** /networks/{networkId}/switch/mtu | Update the MTU configuration



## GetNetworkSwitchMtu

> InlineResponse20069 GetNetworkSwitchMtu(ctx, networkId).Execute()

Return the MTU configuration



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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MtuApi.GetNetworkSwitchMtu(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MtuApi.GetNetworkSwitchMtu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchMtu`: InlineResponse20069
    fmt.Fprintf(os.Stdout, "Response from `MtuApi.GetNetworkSwitchMtu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchMtuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20069**](InlineResponse20069.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchMtu

> map[string]interface{} UpdateNetworkSwitchMtu(ctx, networkId).UpdateNetworkSwitchMtu(updateNetworkSwitchMtu).Execute()

Update the MTU configuration



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
    networkId := "networkId_example" // string | Network ID
    updateNetworkSwitchMtu := *openapiclient.NewInlineObject119() // InlineObject119 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MtuApi.UpdateNetworkSwitchMtu(context.Background(), networkId).UpdateNetworkSwitchMtu(updateNetworkSwitchMtu).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MtuApi.UpdateNetworkSwitchMtu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchMtu`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MtuApi.UpdateNetworkSwitchMtu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchMtuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchMtu** | [**InlineObject119**](InlineObject119.md) |  | 

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

