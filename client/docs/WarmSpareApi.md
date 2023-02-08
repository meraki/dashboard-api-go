# \WarmSpareApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceSwitchWarmSpare**](WarmSpareApi.md#GetDeviceSwitchWarmSpare) | **Get** /devices/{serial}/switch/warmSpare | Return warm spare configuration for a switch
[**GetNetworkApplianceWarmSpare**](WarmSpareApi.md#GetNetworkApplianceWarmSpare) | **Get** /networks/{networkId}/appliance/warmSpare | Return MX warm spare settings
[**SwapNetworkApplianceWarmSpare**](WarmSpareApi.md#SwapNetworkApplianceWarmSpare) | **Post** /networks/{networkId}/appliance/warmSpare/swap | Swap MX primary and warm spare appliances
[**UpdateDeviceSwitchWarmSpare**](WarmSpareApi.md#UpdateDeviceSwitchWarmSpare) | **Put** /devices/{serial}/switch/warmSpare | Update warm spare configuration for a switch
[**UpdateNetworkApplianceWarmSpare**](WarmSpareApi.md#UpdateNetworkApplianceWarmSpare) | **Put** /networks/{networkId}/appliance/warmSpare | Update MX warm spare settings



## GetDeviceSwitchWarmSpare

> map[string]interface{} GetDeviceSwitchWarmSpare(ctx, serial).Execute()

Return warm spare configuration for a switch



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
    resp, r, err := apiClient.WarmSpareApi.GetDeviceSwitchWarmSpare(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarmSpareApi.GetDeviceSwitchWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceSwitchWarmSpare`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WarmSpareApi.GetDeviceSwitchWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSwitchWarmSpareRequest struct via the builder pattern


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


## GetNetworkApplianceWarmSpare

> map[string]interface{} GetNetworkApplianceWarmSpare(ctx, networkId).Execute()

Return MX warm spare settings



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
    resp, r, err := apiClient.WarmSpareApi.GetNetworkApplianceWarmSpare(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarmSpareApi.GetNetworkApplianceWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceWarmSpare`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WarmSpareApi.GetNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceWarmSpareRequest struct via the builder pattern


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


## SwapNetworkApplianceWarmSpare

> map[string]interface{} SwapNetworkApplianceWarmSpare(ctx, networkId).Execute()

Swap MX primary and warm spare appliances



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
    resp, r, err := apiClient.WarmSpareApi.SwapNetworkApplianceWarmSpare(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarmSpareApi.SwapNetworkApplianceWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SwapNetworkApplianceWarmSpare`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WarmSpareApi.SwapNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwapNetworkApplianceWarmSpareRequest struct via the builder pattern


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


## UpdateDeviceSwitchWarmSpare

> map[string]interface{} UpdateDeviceSwitchWarmSpare(ctx, serial).UpdateDeviceSwitchWarmSpare(updateDeviceSwitchWarmSpare).Execute()

Update warm spare configuration for a switch



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
    updateDeviceSwitchWarmSpare := *openapiclient.NewUpdateDeviceSwitchWarmSpareRequest(false) // UpdateDeviceSwitchWarmSpareRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WarmSpareApi.UpdateDeviceSwitchWarmSpare(context.Background(), serial).UpdateDeviceSwitchWarmSpare(updateDeviceSwitchWarmSpare).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarmSpareApi.UpdateDeviceSwitchWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceSwitchWarmSpare`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WarmSpareApi.UpdateDeviceSwitchWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSwitchWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceSwitchWarmSpare** | [**UpdateDeviceSwitchWarmSpareRequest**](UpdateDeviceSwitchWarmSpareRequest.md) |  | 

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


## UpdateNetworkApplianceWarmSpare

> map[string]interface{} UpdateNetworkApplianceWarmSpare(ctx, networkId).UpdateNetworkApplianceWarmSpare(updateNetworkApplianceWarmSpare).Execute()

Update MX warm spare settings



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
    updateNetworkApplianceWarmSpare := *openapiclient.NewUpdateNetworkApplianceWarmSpareRequest(false) // UpdateNetworkApplianceWarmSpareRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WarmSpareApi.UpdateNetworkApplianceWarmSpare(context.Background(), networkId).UpdateNetworkApplianceWarmSpare(updateNetworkApplianceWarmSpare).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WarmSpareApi.UpdateNetworkApplianceWarmSpare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceWarmSpare`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WarmSpareApi.UpdateNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceWarmSpare** | [**UpdateNetworkApplianceWarmSpareRequest**](UpdateNetworkApplianceWarmSpareRequest.md) |  | 

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

