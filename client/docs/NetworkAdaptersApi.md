# \NetworkAdaptersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSmDeviceNetworkAdapters**](NetworkAdaptersApi.md#GetNetworkSmDeviceNetworkAdapters) | **Get** /networks/{networkId}/sm/devices/{deviceId}/networkAdapters | List the network adapters of a device



## GetNetworkSmDeviceNetworkAdapters

> []InlineResponse20054 GetNetworkSmDeviceNetworkAdapters(ctx, networkId, deviceId).Execute()

List the network adapters of a device



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
    deviceId := "deviceId_example" // string | Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkAdaptersApi.GetNetworkSmDeviceNetworkAdapters(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkAdaptersApi.GetNetworkSmDeviceNetworkAdapters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceNetworkAdapters`: []InlineResponse20054
    fmt.Fprintf(os.Stdout, "Response from `NetworkAdaptersApi.GetNetworkSmDeviceNetworkAdapters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**deviceId** | **string** | Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceNetworkAdaptersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse20054**](InlineResponse20054.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

