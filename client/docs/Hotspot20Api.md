# \Hotspot20Api

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessSsidHotspot20**](Hotspot20Api.md#GetNetworkWirelessSsidHotspot20) | **Get** /networks/{networkId}/wireless/ssids/{number}/hotspot20 | Return the Hotspot 2.0 settings for an SSID
[**UpdateNetworkWirelessSsidHotspot20**](Hotspot20Api.md#UpdateNetworkWirelessSsidHotspot20) | **Put** /networks/{networkId}/wireless/ssids/{number}/hotspot20 | Update the Hotspot 2.0 settings of an SSID



## GetNetworkWirelessSsidHotspot20

> map[string]interface{} GetNetworkWirelessSsidHotspot20(ctx, networkId, number).Execute()

Return the Hotspot 2.0 settings for an SSID



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
    number := "number_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Hotspot20Api.GetNetworkWirelessSsidHotspot20(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Hotspot20Api.GetNetworkWirelessSsidHotspot20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidHotspot20`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Hotspot20Api.GetNetworkWirelessSsidHotspot20`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidHotspot20Request struct via the builder pattern


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


## UpdateNetworkWirelessSsidHotspot20

> map[string]interface{} UpdateNetworkWirelessSsidHotspot20(ctx, networkId, number).UpdateNetworkWirelessSsidHotspot20(updateNetworkWirelessSsidHotspot20).Execute()

Update the Hotspot 2.0 settings of an SSID



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
    number := "number_example" // string | 
    updateNetworkWirelessSsidHotspot20 := *openapiclient.NewUpdateNetworkWirelessSsidHotspot20Request() // UpdateNetworkWirelessSsidHotspot20Request |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Hotspot20Api.UpdateNetworkWirelessSsidHotspot20(context.Background(), networkId, number).UpdateNetworkWirelessSsidHotspot20(updateNetworkWirelessSsidHotspot20).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Hotspot20Api.UpdateNetworkWirelessSsidHotspot20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidHotspot20`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Hotspot20Api.UpdateNetworkWirelessSsidHotspot20`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidHotspot20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidHotspot20** | [**UpdateNetworkWirelessSsidHotspot20Request**](UpdateNetworkWirelessSsidHotspot20Request.md) |  | 

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

