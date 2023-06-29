# \SplashAuthorizationStatusApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkClientSplashAuthorizationStatus**](SplashAuthorizationStatusApi.md#GetNetworkClientSplashAuthorizationStatus) | **Get** /networks/{networkId}/clients/{clientId}/splashAuthorizationStatus | Return the splash authorization for a client, for each SSID they&#39;ve associated with through splash
[**UpdateNetworkClientSplashAuthorizationStatus**](SplashAuthorizationStatusApi.md#UpdateNetworkClientSplashAuthorizationStatus) | **Put** /networks/{networkId}/clients/{clientId}/splashAuthorizationStatus | Update a client&#39;s splash authorization



## GetNetworkClientSplashAuthorizationStatus

> map[string]interface{} GetNetworkClientSplashAuthorizationStatus(ctx, networkId, clientId).Execute()

Return the splash authorization for a client, for each SSID they've associated with through splash



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
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SplashAuthorizationStatusApi.GetNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SplashAuthorizationStatusApi.GetNetworkClientSplashAuthorizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkClientSplashAuthorizationStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SplashAuthorizationStatusApi.GetNetworkClientSplashAuthorizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkClientSplashAuthorizationStatusRequest struct via the builder pattern


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


## UpdateNetworkClientSplashAuthorizationStatus

> map[string]interface{} UpdateNetworkClientSplashAuthorizationStatus(ctx, networkId, clientId).UpdateNetworkClientSplashAuthorizationStatusRequest(updateNetworkClientSplashAuthorizationStatusRequest).Execute()

Update a client's splash authorization



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
    networkId := "networkId_example" // string | Network ID
    clientId := "clientId_example" // string | Client ID
    updateNetworkClientSplashAuthorizationStatusRequest := *openapiclient.NewUpdateNetworkClientSplashAuthorizationStatusRequest(*openapiclient.NewUpdateNetworkClientSplashAuthorizationStatusRequestSsids()) // UpdateNetworkClientSplashAuthorizationStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SplashAuthorizationStatusApi.UpdateNetworkClientSplashAuthorizationStatus(context.Background(), networkId, clientId).UpdateNetworkClientSplashAuthorizationStatusRequest(updateNetworkClientSplashAuthorizationStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SplashAuthorizationStatusApi.UpdateNetworkClientSplashAuthorizationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkClientSplashAuthorizationStatus`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SplashAuthorizationStatusApi.UpdateNetworkClientSplashAuthorizationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkClientSplashAuthorizationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkClientSplashAuthorizationStatusRequest** | [**UpdateNetworkClientSplashAuthorizationStatusRequest**](UpdateNetworkClientSplashAuthorizationStatusRequest.md) |  | 

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

