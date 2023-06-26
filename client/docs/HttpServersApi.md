# \HttpServersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWebhooksHttpServer**](HttpServersApi.md#CreateNetworkWebhooksHttpServer) | **Post** /networks/{networkId}/webhooks/httpServers | Add an HTTP server to a network
[**DeleteNetworkWebhooksHttpServer**](HttpServersApi.md#DeleteNetworkWebhooksHttpServer) | **Delete** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Delete an HTTP server from a network
[**GetNetworkWebhooksHttpServer**](HttpServersApi.md#GetNetworkWebhooksHttpServer) | **Get** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Return an HTTP server for a network
[**GetNetworkWebhooksHttpServers**](HttpServersApi.md#GetNetworkWebhooksHttpServers) | **Get** /networks/{networkId}/webhooks/httpServers | List the HTTP servers for a network
[**UpdateNetworkWebhooksHttpServer**](HttpServersApi.md#UpdateNetworkWebhooksHttpServer) | **Put** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Update an HTTP server



## CreateNetworkWebhooksHttpServer

> InlineResponse20083 CreateNetworkWebhooksHttpServer(ctx, networkId).CreateNetworkWebhooksHttpServer(createNetworkWebhooksHttpServer).Execute()

Add an HTTP server to a network



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
    createNetworkWebhooksHttpServer := *openapiclient.NewInlineObject146("Name_example", "Url_example") // InlineObject146 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpServersApi.CreateNetworkWebhooksHttpServer(context.Background(), networkId).CreateNetworkWebhooksHttpServer(createNetworkWebhooksHttpServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServersApi.CreateNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWebhooksHttpServer`: InlineResponse20083
    fmt.Fprintf(os.Stdout, "Response from `HttpServersApi.CreateNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksHttpServer** | [**InlineObject146**](InlineObject146.md) |  | 

### Return type

[**InlineResponse20083**](InlineResponse20083.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkWebhooksHttpServer

> DeleteNetworkWebhooksHttpServer(ctx, networkId, httpServerId).Execute()

Delete an HTTP server from a network



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
    httpServerId := "httpServerId_example" // string | Http server ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpServersApi.DeleteNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServersApi.DeleteNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**httpServerId** | **string** | Http server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWebhooksHttpServerRequest struct via the builder pattern


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


## GetNetworkWebhooksHttpServer

> InlineResponse20083 GetNetworkWebhooksHttpServer(ctx, networkId, httpServerId).Execute()

Return an HTTP server for a network



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
    httpServerId := "httpServerId_example" // string | Http server ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpServersApi.GetNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServersApi.GetNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksHttpServer`: InlineResponse20083
    fmt.Fprintf(os.Stdout, "Response from `HttpServersApi.GetNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**httpServerId** | **string** | Http server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20083**](InlineResponse20083.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksHttpServers

> []InlineResponse20083 GetNetworkWebhooksHttpServers(ctx, networkId).Execute()

List the HTTP servers for a network



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
    resp, r, err := apiClient.HttpServersApi.GetNetworkWebhooksHttpServers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServersApi.GetNetworkWebhooksHttpServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksHttpServers`: []InlineResponse20083
    fmt.Fprintf(os.Stdout, "Response from `HttpServersApi.GetNetworkWebhooksHttpServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksHttpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20083**](InlineResponse20083.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWebhooksHttpServer

> InlineResponse20083 UpdateNetworkWebhooksHttpServer(ctx, networkId, httpServerId).UpdateNetworkWebhooksHttpServer(updateNetworkWebhooksHttpServer).Execute()

Update an HTTP server



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
    httpServerId := "httpServerId_example" // string | Http server ID
    updateNetworkWebhooksHttpServer := *openapiclient.NewInlineObject147() // InlineObject147 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HttpServersApi.UpdateNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).UpdateNetworkWebhooksHttpServer(updateNetworkWebhooksHttpServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HttpServersApi.UpdateNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWebhooksHttpServer`: InlineResponse20083
    fmt.Fprintf(os.Stdout, "Response from `HttpServersApi.UpdateNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**httpServerId** | **string** | Http server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWebhooksHttpServer** | [**InlineObject147**](InlineObject147.md) |  | 

### Return type

[**InlineResponse20083**](InlineResponse20083.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

