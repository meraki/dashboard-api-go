# \IdentityPsksApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWirelessSsidIdentityPsk**](IdentityPsksApi.md#CreateNetworkWirelessSsidIdentityPsk) | **Post** /networks/{networkId}/wireless/ssids/{number}/identityPsks | Create an Identity PSK
[**DeleteNetworkWirelessSsidIdentityPsk**](IdentityPsksApi.md#DeleteNetworkWirelessSsidIdentityPsk) | **Delete** /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} | Delete an Identity PSK
[**GetNetworkWirelessSsidIdentityPsk**](IdentityPsksApi.md#GetNetworkWirelessSsidIdentityPsk) | **Get** /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} | Return an Identity PSK
[**GetNetworkWirelessSsidIdentityPsks**](IdentityPsksApi.md#GetNetworkWirelessSsidIdentityPsks) | **Get** /networks/{networkId}/wireless/ssids/{number}/identityPsks | List all Identity PSKs in a wireless network
[**UpdateNetworkWirelessSsidIdentityPsk**](IdentityPsksApi.md#UpdateNetworkWirelessSsidIdentityPsk) | **Put** /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} | Update an Identity PSK



## CreateNetworkWirelessSsidIdentityPsk

> map[string]interface{} CreateNetworkWirelessSsidIdentityPsk(ctx, networkId, number).CreateNetworkWirelessSsidIdentityPsk(createNetworkWirelessSsidIdentityPsk).Execute()

Create an Identity PSK



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
    createNetworkWirelessSsidIdentityPsk := *openapiclient.NewCreateNetworkWirelessSsidIdentityPskRequest("Name_example", "GroupPolicyId_example") // CreateNetworkWirelessSsidIdentityPskRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPsksApi.CreateNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number).CreateNetworkWirelessSsidIdentityPsk(createNetworkWirelessSsidIdentityPsk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPsksApi.CreateNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWirelessSsidIdentityPsk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdentityPsksApi.CreateNetworkWirelessSsidIdentityPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkWirelessSsidIdentityPsk** | [**CreateNetworkWirelessSsidIdentityPskRequest**](CreateNetworkWirelessSsidIdentityPskRequest.md) |  | 

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


## DeleteNetworkWirelessSsidIdentityPsk

> DeleteNetworkWirelessSsidIdentityPsk(ctx, networkId, number, identityPskId).Execute()

Delete an Identity PSK



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
    identityPskId := "identityPskId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPsksApi.DeleteNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPsksApi.DeleteNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**number** | **string** |  | 
**identityPskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidIdentityPsk

> GetNetworkWirelessSsidIdentityPsks200ResponseInner GetNetworkWirelessSsidIdentityPsk(ctx, networkId, number, identityPskId).Execute()

Return an Identity PSK



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
    identityPskId := "identityPskId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPsksApi.GetNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPsksApi.GetNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidIdentityPsk`: GetNetworkWirelessSsidIdentityPsks200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `IdentityPsksApi.GetNetworkWirelessSsidIdentityPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**number** | **string** |  | 
**identityPskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetNetworkWirelessSsidIdentityPsks200ResponseInner**](GetNetworkWirelessSsidIdentityPsks200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidIdentityPsks

> []GetNetworkWirelessSsidIdentityPsks200ResponseInner GetNetworkWirelessSsidIdentityPsks(ctx, networkId, number).Execute()

List all Identity PSKs in a wireless network



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
    resp, r, err := apiClient.IdentityPsksApi.GetNetworkWirelessSsidIdentityPsks(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPsksApi.GetNetworkWirelessSsidIdentityPsks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidIdentityPsks`: []GetNetworkWirelessSsidIdentityPsks200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `IdentityPsksApi.GetNetworkWirelessSsidIdentityPsks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidIdentityPsksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkWirelessSsidIdentityPsks200ResponseInner**](GetNetworkWirelessSsidIdentityPsks200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidIdentityPsk

> map[string]interface{} UpdateNetworkWirelessSsidIdentityPsk(ctx, networkId, number, identityPskId).UpdateNetworkWirelessSsidIdentityPsk(updateNetworkWirelessSsidIdentityPsk).Execute()

Update an Identity PSK



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
    identityPskId := "identityPskId_example" // string | 
    updateNetworkWirelessSsidIdentityPsk := *openapiclient.NewUpdateNetworkWirelessSsidIdentityPskRequest() // UpdateNetworkWirelessSsidIdentityPskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IdentityPsksApi.UpdateNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).UpdateNetworkWirelessSsidIdentityPsk(updateNetworkWirelessSsidIdentityPsk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IdentityPsksApi.UpdateNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidIdentityPsk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `IdentityPsksApi.UpdateNetworkWirelessSsidIdentityPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**number** | **string** |  | 
**identityPskId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkWirelessSsidIdentityPsk** | [**UpdateNetworkWirelessSsidIdentityPskRequest**](UpdateNetworkWirelessSsidIdentityPskRequest.md) |  | 

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

