# \AlternateManagementInterfaceApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchAlternateManagementInterface**](AlternateManagementInterfaceApi.md#GetNetworkSwitchAlternateManagementInterface) | **Get** /networks/{networkId}/switch/alternateManagementInterface | Return the switch alternate management interface for the network
[**GetNetworkWirelessAlternateManagementInterface**](AlternateManagementInterfaceApi.md#GetNetworkWirelessAlternateManagementInterface) | **Get** /networks/{networkId}/wireless/alternateManagementInterface | Return alternate management interface and devices with IP assigned
[**UpdateNetworkSwitchAlternateManagementInterface**](AlternateManagementInterfaceApi.md#UpdateNetworkSwitchAlternateManagementInterface) | **Put** /networks/{networkId}/switch/alternateManagementInterface | Update the switch alternate management interface for the network
[**UpdateNetworkWirelessAlternateManagementInterface**](AlternateManagementInterfaceApi.md#UpdateNetworkWirelessAlternateManagementInterface) | **Put** /networks/{networkId}/wireless/alternateManagementInterface | Update alternate management interface and device static IP



## GetNetworkSwitchAlternateManagementInterface

> map[string]interface{} GetNetworkSwitchAlternateManagementInterface(ctx, networkId).Execute()

Return the switch alternate management interface for the network



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
    resp, r, err := apiClient.AlternateManagementInterfaceApi.GetNetworkSwitchAlternateManagementInterface(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlternateManagementInterfaceApi.GetNetworkSwitchAlternateManagementInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchAlternateManagementInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AlternateManagementInterfaceApi.GetNetworkSwitchAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAlternateManagementInterfaceRequest struct via the builder pattern


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


## GetNetworkWirelessAlternateManagementInterface

> map[string]interface{} GetNetworkWirelessAlternateManagementInterface(ctx, networkId).Execute()

Return alternate management interface and devices with IP assigned



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
    resp, r, err := apiClient.AlternateManagementInterfaceApi.GetNetworkWirelessAlternateManagementInterface(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlternateManagementInterfaceApi.GetNetworkWirelessAlternateManagementInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessAlternateManagementInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AlternateManagementInterfaceApi.GetNetworkWirelessAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessAlternateManagementInterfaceRequest struct via the builder pattern


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


## UpdateNetworkSwitchAlternateManagementInterface

> map[string]interface{} UpdateNetworkSwitchAlternateManagementInterface(ctx, networkId).UpdateNetworkSwitchAlternateManagementInterface(updateNetworkSwitchAlternateManagementInterface).Execute()

Update the switch alternate management interface for the network



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
    updateNetworkSwitchAlternateManagementInterface := *openapiclient.NewUpdateNetworkSwitchAlternateManagementInterfaceRequest() // UpdateNetworkSwitchAlternateManagementInterfaceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlternateManagementInterfaceApi.UpdateNetworkSwitchAlternateManagementInterface(context.Background(), networkId).UpdateNetworkSwitchAlternateManagementInterface(updateNetworkSwitchAlternateManagementInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlternateManagementInterfaceApi.UpdateNetworkSwitchAlternateManagementInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchAlternateManagementInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AlternateManagementInterfaceApi.UpdateNetworkSwitchAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchAlternateManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchAlternateManagementInterface** | [**UpdateNetworkSwitchAlternateManagementInterfaceRequest**](UpdateNetworkSwitchAlternateManagementInterfaceRequest.md) |  | 

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


## UpdateNetworkWirelessAlternateManagementInterface

> map[string]interface{} UpdateNetworkWirelessAlternateManagementInterface(ctx, networkId).UpdateNetworkWirelessAlternateManagementInterface(updateNetworkWirelessAlternateManagementInterface).Execute()

Update alternate management interface and device static IP



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
    updateNetworkWirelessAlternateManagementInterface := *openapiclient.NewUpdateNetworkWirelessAlternateManagementInterfaceRequest() // UpdateNetworkWirelessAlternateManagementInterfaceRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlternateManagementInterfaceApi.UpdateNetworkWirelessAlternateManagementInterface(context.Background(), networkId).UpdateNetworkWirelessAlternateManagementInterface(updateNetworkWirelessAlternateManagementInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlternateManagementInterfaceApi.UpdateNetworkWirelessAlternateManagementInterface``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessAlternateManagementInterface`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AlternateManagementInterfaceApi.UpdateNetworkWirelessAlternateManagementInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessAlternateManagementInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessAlternateManagementInterface** | [**UpdateNetworkWirelessAlternateManagementInterfaceRequest**](UpdateNetworkWirelessAlternateManagementInterfaceRequest.md) |  | 

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

