# \StaticsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkAppliancePrefixesDelegatedStatic**](StaticsApi.md#CreateNetworkAppliancePrefixesDelegatedStatic) | **Post** /networks/{networkId}/appliance/prefixes/delegated/statics | Add a static delegated prefix from a network
[**DeleteNetworkAppliancePrefixesDelegatedStatic**](StaticsApi.md#DeleteNetworkAppliancePrefixesDelegatedStatic) | **Delete** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Delete a static delegated prefix from a network
[**GetNetworkAppliancePrefixesDelegatedStatic**](StaticsApi.md#GetNetworkAppliancePrefixesDelegatedStatic) | **Get** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Return a static delegated prefix from a network
[**GetNetworkAppliancePrefixesDelegatedStatics**](StaticsApi.md#GetNetworkAppliancePrefixesDelegatedStatics) | **Get** /networks/{networkId}/appliance/prefixes/delegated/statics | List static delegated prefixes for a network
[**UpdateNetworkAppliancePrefixesDelegatedStatic**](StaticsApi.md#UpdateNetworkAppliancePrefixesDelegatedStatic) | **Put** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Update a static delegated prefix from a network



## CreateNetworkAppliancePrefixesDelegatedStatic

> map[string]interface{} CreateNetworkAppliancePrefixesDelegatedStatic(ctx, networkId).CreateNetworkAppliancePrefixesDelegatedStatic(createNetworkAppliancePrefixesDelegatedStatic).Execute()

Add a static delegated prefix from a network



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
    createNetworkAppliancePrefixesDelegatedStatic := *openapiclient.NewInlineObject41("Prefix_example", *openapiclient.NewNetworksNetworkIdAppliancePrefixesDelegatedStaticsOrigin1()) // InlineObject41 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticsApi.CreateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId).CreateNetworkAppliancePrefixesDelegatedStatic(createNetworkAppliancePrefixesDelegatedStatic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticsApi.CreateNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkAppliancePrefixesDelegatedStatic`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StaticsApi.CreateNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkAppliancePrefixesDelegatedStatic** | [**InlineObject41**](InlineObject41.md) |  | 

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


## DeleteNetworkAppliancePrefixesDelegatedStatic

> DeleteNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).Execute()

Delete a static delegated prefix from a network



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
    staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticsApi.DeleteNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticsApi.DeleteNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


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


## GetNetworkAppliancePrefixesDelegatedStatic

> InlineResponse20014 GetNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).Execute()

Return a static delegated prefix from a network



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
    staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticsApi.GetNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticsApi.GetNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAppliancePrefixesDelegatedStatic`: InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `StaticsApi.GetNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePrefixesDelegatedStatics

> []InlineResponse20014 GetNetworkAppliancePrefixesDelegatedStatics(ctx, networkId).Execute()

List static delegated prefixes for a network



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticsApi.GetNetworkAppliancePrefixesDelegatedStatics(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticsApi.GetNetworkAppliancePrefixesDelegatedStatics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkAppliancePrefixesDelegatedStatics`: []InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `StaticsApi.GetNetworkAppliancePrefixesDelegatedStatics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePrefixesDelegatedStaticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20014**](InlineResponse20014.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAppliancePrefixesDelegatedStatic

> map[string]interface{} UpdateNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).UpdateNetworkAppliancePrefixesDelegatedStatic(updateNetworkAppliancePrefixesDelegatedStatic).Execute()

Update a static delegated prefix from a network



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
    staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID
    updateNetworkAppliancePrefixesDelegatedStatic := *openapiclient.NewInlineObject42() // InlineObject42 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StaticsApi.UpdateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).UpdateNetworkAppliancePrefixesDelegatedStatic(updateNetworkAppliancePrefixesDelegatedStatic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StaticsApi.UpdateNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkAppliancePrefixesDelegatedStatic`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StaticsApi.UpdateNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkAppliancePrefixesDelegatedStatic** | [**InlineObject42**](InlineObject42.md) |  | 

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

