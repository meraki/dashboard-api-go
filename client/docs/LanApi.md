# \LanApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCellularGatewayLan**](LanApi.md#GetDeviceCellularGatewayLan) | **Get** /devices/{serial}/cellularGateway/lan | Show the LAN Settings of a MG
[**UpdateDeviceCellularGatewayLan**](LanApi.md#UpdateDeviceCellularGatewayLan) | **Put** /devices/{serial}/cellularGateway/lan | Update the LAN Settings for a single MG.



## GetDeviceCellularGatewayLan

> map[string]interface{} GetDeviceCellularGatewayLan(ctx, serial).Execute()

Show the LAN Settings of a MG



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
    serial := "serial_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanApi.GetDeviceCellularGatewayLan(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanApi.GetDeviceCellularGatewayLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCellularGatewayLan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LanApi.GetDeviceCellularGatewayLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularGatewayLanRequest struct via the builder pattern


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


## UpdateDeviceCellularGatewayLan

> map[string]interface{} UpdateDeviceCellularGatewayLan(ctx, serial).UpdateDeviceCellularGatewayLan(updateDeviceCellularGatewayLan).Execute()

Update the LAN Settings for a single MG.



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
    serial := "serial_example" // string | 
    updateDeviceCellularGatewayLan := *openapiclient.NewUpdateDeviceCellularGatewayLanRequest() // UpdateDeviceCellularGatewayLanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LanApi.UpdateDeviceCellularGatewayLan(context.Background(), serial).UpdateDeviceCellularGatewayLan(updateDeviceCellularGatewayLan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LanApi.UpdateDeviceCellularGatewayLan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCellularGatewayLan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LanApi.UpdateDeviceCellularGatewayLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularGatewayLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularGatewayLan** | [**UpdateDeviceCellularGatewayLanRequest**](UpdateDeviceCellularGatewayLanRequest.md) |  | 

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

