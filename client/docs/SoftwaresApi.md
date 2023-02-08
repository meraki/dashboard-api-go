# \SoftwaresApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSmDeviceSoftwares**](SoftwaresApi.md#GetNetworkSmDeviceSoftwares) | **Get** /networks/{networkId}/sm/devices/{deviceId}/softwares | Get a list of softwares associated with a device
[**GetNetworkSmUserSoftwares**](SoftwaresApi.md#GetNetworkSmUserSoftwares) | **Get** /networks/{networkId}/sm/users/{userId}/softwares | Get a list of softwares associated with a user



## GetNetworkSmDeviceSoftwares

> []GetNetworkSmDeviceSoftwares200ResponseInner GetNetworkSmDeviceSoftwares(ctx, networkId, deviceId).Execute()

Get a list of softwares associated with a device



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
    deviceId := "deviceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwaresApi.GetNetworkSmDeviceSoftwares(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwaresApi.GetNetworkSmDeviceSoftwares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceSoftwares`: []GetNetworkSmDeviceSoftwares200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SoftwaresApi.GetNetworkSmDeviceSoftwares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceSoftwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceSoftwares200ResponseInner**](GetNetworkSmDeviceSoftwares200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUserSoftwares

> []GetNetworkSmDeviceSoftwares200ResponseInner GetNetworkSmUserSoftwares(ctx, networkId, userId).Execute()

Get a list of softwares associated with a user



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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoftwaresApi.GetNetworkSmUserSoftwares(context.Background(), networkId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoftwaresApi.GetNetworkSmUserSoftwares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmUserSoftwares`: []GetNetworkSmDeviceSoftwares200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SoftwaresApi.GetNetworkSmUserSoftwares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUserSoftwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceSoftwares200ResponseInner**](GetNetworkSmDeviceSoftwares200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

