# \LinkAggregationsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchLinkAggregation**](LinkAggregationsApi.md#CreateNetworkSwitchLinkAggregation) | **Post** /networks/{networkId}/switch/linkAggregations | Create a link aggregation group
[**DeleteNetworkSwitchLinkAggregation**](LinkAggregationsApi.md#DeleteNetworkSwitchLinkAggregation) | **Delete** /networks/{networkId}/switch/linkAggregations/{linkAggregationId} | Split a link aggregation group into separate ports
[**GetNetworkSwitchLinkAggregations**](LinkAggregationsApi.md#GetNetworkSwitchLinkAggregations) | **Get** /networks/{networkId}/switch/linkAggregations | List link aggregation groups
[**UpdateNetworkSwitchLinkAggregation**](LinkAggregationsApi.md#UpdateNetworkSwitchLinkAggregation) | **Put** /networks/{networkId}/switch/linkAggregations/{linkAggregationId} | Update a link aggregation group



## CreateNetworkSwitchLinkAggregation

> map[string]interface{} CreateNetworkSwitchLinkAggregation(ctx, networkId).CreateNetworkSwitchLinkAggregation(createNetworkSwitchLinkAggregation).Execute()

Create a link aggregation group



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
    createNetworkSwitchLinkAggregation := *openapiclient.NewCreateNetworkSwitchLinkAggregationRequest() // CreateNetworkSwitchLinkAggregationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinkAggregationsApi.CreateNetworkSwitchLinkAggregation(context.Background(), networkId).CreateNetworkSwitchLinkAggregation(createNetworkSwitchLinkAggregation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsApi.CreateNetworkSwitchLinkAggregation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchLinkAggregation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinkAggregationsApi.CreateNetworkSwitchLinkAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchLinkAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchLinkAggregation** | [**CreateNetworkSwitchLinkAggregationRequest**](CreateNetworkSwitchLinkAggregationRequest.md) |  | 

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


## DeleteNetworkSwitchLinkAggregation

> DeleteNetworkSwitchLinkAggregation(ctx, networkId, linkAggregationId).Execute()

Split a link aggregation group into separate ports



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
    linkAggregationId := "linkAggregationId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinkAggregationsApi.DeleteNetworkSwitchLinkAggregation(context.Background(), networkId, linkAggregationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsApi.DeleteNetworkSwitchLinkAggregation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**linkAggregationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchLinkAggregationRequest struct via the builder pattern


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


## GetNetworkSwitchLinkAggregations

> []map[string]interface{} GetNetworkSwitchLinkAggregations(ctx, networkId).Execute()

List link aggregation groups



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
    resp, r, err := apiClient.LinkAggregationsApi.GetNetworkSwitchLinkAggregations(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsApi.GetNetworkSwitchLinkAggregations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchLinkAggregations`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinkAggregationsApi.GetNetworkSwitchLinkAggregations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchLinkAggregationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchLinkAggregation

> map[string]interface{} UpdateNetworkSwitchLinkAggregation(ctx, networkId, linkAggregationId).UpdateNetworkSwitchLinkAggregation(updateNetworkSwitchLinkAggregation).Execute()

Update a link aggregation group



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
    linkAggregationId := "linkAggregationId_example" // string | 
    updateNetworkSwitchLinkAggregation := *openapiclient.NewUpdateNetworkSwitchLinkAggregationRequest() // UpdateNetworkSwitchLinkAggregationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LinkAggregationsApi.UpdateNetworkSwitchLinkAggregation(context.Background(), networkId, linkAggregationId).UpdateNetworkSwitchLinkAggregation(updateNetworkSwitchLinkAggregation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LinkAggregationsApi.UpdateNetworkSwitchLinkAggregation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchLinkAggregation`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `LinkAggregationsApi.UpdateNetworkSwitchLinkAggregation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**linkAggregationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchLinkAggregationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchLinkAggregation** | [**UpdateNetworkSwitchLinkAggregationRequest**](UpdateNetworkSwitchLinkAggregationRequest.md) |  | 

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

