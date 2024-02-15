# \WakeOnLanApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceLiveToolsWakeOnLan**](WakeOnLanApi.md#CreateDeviceLiveToolsWakeOnLan) | **Post** /devices/{serial}/liveTools/wakeOnLan | Enqueue a job to send a Wake-on-LAN packet from the device
[**GetDeviceLiveToolsWakeOnLan**](WakeOnLanApi.md#GetDeviceLiveToolsWakeOnLan) | **Get** /devices/{serial}/liveTools/wakeOnLan/{wakeOnLanId} | Return a Wake-on-LAN job



## CreateDeviceLiveToolsWakeOnLan

> CreateDeviceLiveToolsWakeOnLan201Response CreateDeviceLiveToolsWakeOnLan(ctx, serial).CreateDeviceLiveToolsWakeOnLanRequest(createDeviceLiveToolsWakeOnLanRequest).Execute()

Enqueue a job to send a Wake-on-LAN packet from the device



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
    serial := "serial_example" // string | Serial
    createDeviceLiveToolsWakeOnLanRequest := *openapiclient.NewCreateDeviceLiveToolsWakeOnLanRequest(int32(123), "Mac_example") // CreateDeviceLiveToolsWakeOnLanRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WakeOnLanApi.CreateDeviceLiveToolsWakeOnLan(context.Background(), serial).CreateDeviceLiveToolsWakeOnLanRequest(createDeviceLiveToolsWakeOnLanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WakeOnLanApi.CreateDeviceLiveToolsWakeOnLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceLiveToolsWakeOnLan`: CreateDeviceLiveToolsWakeOnLan201Response
    fmt.Fprintf(os.Stdout, "Response from `WakeOnLanApi.CreateDeviceLiveToolsWakeOnLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceLiveToolsWakeOnLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeviceLiveToolsWakeOnLanRequest** | [**CreateDeviceLiveToolsWakeOnLanRequest**](CreateDeviceLiveToolsWakeOnLanRequest.md) |  | 

### Return type

[**CreateDeviceLiveToolsWakeOnLan201Response**](CreateDeviceLiveToolsWakeOnLan201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceLiveToolsWakeOnLan

> DevicesSerialLiveToolsWakeOnLanPostRequestMessage GetDeviceLiveToolsWakeOnLan(ctx, serial, wakeOnLanId).Execute()

Return a Wake-on-LAN job



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
    serial := "serial_example" // string | Serial
    wakeOnLanId := "wakeOnLanId_example" // string | Wake on lan ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WakeOnLanApi.GetDeviceLiveToolsWakeOnLan(context.Background(), serial, wakeOnLanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WakeOnLanApi.GetDeviceLiveToolsWakeOnLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceLiveToolsWakeOnLan`: DevicesSerialLiveToolsWakeOnLanPostRequestMessage
    fmt.Fprintf(os.Stdout, "Response from `WakeOnLanApi.GetDeviceLiveToolsWakeOnLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 
**wakeOnLanId** | **string** | Wake on lan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLiveToolsWakeOnLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevicesSerialLiveToolsWakeOnLanPostRequestMessage**](DevicesSerialLiveToolsWakeOnLanPostRequestMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

