# \BypassActivationLockAttemptsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSmBypassActivationLockAttempt**](BypassActivationLockAttemptsApi.md#CreateNetworkSmBypassActivationLockAttempt) | **Post** /networks/{networkId}/sm/bypassActivationLockAttempts | Bypass activation lock attempt
[**GetNetworkSmBypassActivationLockAttempt**](BypassActivationLockAttemptsApi.md#GetNetworkSmBypassActivationLockAttempt) | **Get** /networks/{networkId}/sm/bypassActivationLockAttempts/{attemptId} | Bypass activation lock attempt status



## CreateNetworkSmBypassActivationLockAttempt

> map[string]interface{} CreateNetworkSmBypassActivationLockAttempt(ctx, networkId).CreateNetworkSmBypassActivationLockAttempt(createNetworkSmBypassActivationLockAttempt).Execute()

Bypass activation lock attempt



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
    createNetworkSmBypassActivationLockAttempt := *openapiclient.NewCreateNetworkSmBypassActivationLockAttemptRequest([]string{"Ids_example"}) // CreateNetworkSmBypassActivationLockAttemptRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BypassActivationLockAttemptsApi.CreateNetworkSmBypassActivationLockAttempt(context.Background(), networkId).CreateNetworkSmBypassActivationLockAttempt(createNetworkSmBypassActivationLockAttempt).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BypassActivationLockAttemptsApi.CreateNetworkSmBypassActivationLockAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSmBypassActivationLockAttempt`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BypassActivationLockAttemptsApi.CreateNetworkSmBypassActivationLockAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSmBypassActivationLockAttemptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSmBypassActivationLockAttempt** | [**CreateNetworkSmBypassActivationLockAttemptRequest**](CreateNetworkSmBypassActivationLockAttemptRequest.md) |  | 

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


## GetNetworkSmBypassActivationLockAttempt

> map[string]interface{} GetNetworkSmBypassActivationLockAttempt(ctx, networkId, attemptId).Execute()

Bypass activation lock attempt status



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
    attemptId := "attemptId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BypassActivationLockAttemptsApi.GetNetworkSmBypassActivationLockAttempt(context.Background(), networkId, attemptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BypassActivationLockAttemptsApi.GetNetworkSmBypassActivationLockAttempt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmBypassActivationLockAttempt`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BypassActivationLockAttemptsApi.GetNetworkSmBypassActivationLockAttempt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**attemptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmBypassActivationLockAttemptRequest struct via the builder pattern


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

