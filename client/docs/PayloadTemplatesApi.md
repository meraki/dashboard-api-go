# \PayloadTemplatesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWebhooksPayloadTemplate**](PayloadTemplatesApi.md#CreateNetworkWebhooksPayloadTemplate) | **Post** /networks/{networkId}/webhooks/payloadTemplates | Create a webhook payload template for a network
[**DeleteNetworkWebhooksPayloadTemplate**](PayloadTemplatesApi.md#DeleteNetworkWebhooksPayloadTemplate) | **Delete** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Destroy a webhook payload template for a network
[**GetNetworkWebhooksPayloadTemplate**](PayloadTemplatesApi.md#GetNetworkWebhooksPayloadTemplate) | **Get** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Get the webhook payload template for a network
[**GetNetworkWebhooksPayloadTemplates**](PayloadTemplatesApi.md#GetNetworkWebhooksPayloadTemplates) | **Get** /networks/{networkId}/webhooks/payloadTemplates | List the webhook payload templates for a network
[**UpdateNetworkWebhooksPayloadTemplate**](PayloadTemplatesApi.md#UpdateNetworkWebhooksPayloadTemplate) | **Put** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Update a webhook payload template for a network



## CreateNetworkWebhooksPayloadTemplate

> InlineResponse20075 CreateNetworkWebhooksPayloadTemplate(ctx, networkId).CreateNetworkWebhooksPayloadTemplate(createNetworkWebhooksPayloadTemplate).Execute()

Create a webhook payload template for a network



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
    createNetworkWebhooksPayloadTemplate := *openapiclient.NewInlineObject145("Name_example") // InlineObject145 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayloadTemplatesApi.CreateNetworkWebhooksPayloadTemplate(context.Background(), networkId).CreateNetworkWebhooksPayloadTemplate(createNetworkWebhooksPayloadTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayloadTemplatesApi.CreateNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWebhooksPayloadTemplate`: InlineResponse20075
    fmt.Fprintf(os.Stdout, "Response from `PayloadTemplatesApi.CreateNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksPayloadTemplate** | [**InlineObject145**](InlineObject145.md) |  | 

### Return type

[**InlineResponse20075**](InlineResponse20075.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkWebhooksPayloadTemplate

> DeleteNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).Execute()

Destroy a webhook payload template for a network



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
    payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayloadTemplatesApi.DeleteNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayloadTemplatesApi.DeleteNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


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


## GetNetworkWebhooksPayloadTemplate

> InlineResponse20075 GetNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).Execute()

Get the webhook payload template for a network



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
    payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayloadTemplatesApi.GetNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayloadTemplatesApi.GetNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksPayloadTemplate`: InlineResponse20075
    fmt.Fprintf(os.Stdout, "Response from `PayloadTemplatesApi.GetNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20075**](InlineResponse20075.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksPayloadTemplates

> []InlineResponse20075 GetNetworkWebhooksPayloadTemplates(ctx, networkId).Execute()

List the webhook payload templates for a network



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
    resp, r, err := apiClient.PayloadTemplatesApi.GetNetworkWebhooksPayloadTemplates(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayloadTemplatesApi.GetNetworkWebhooksPayloadTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksPayloadTemplates`: []InlineResponse20075
    fmt.Fprintf(os.Stdout, "Response from `PayloadTemplatesApi.GetNetworkWebhooksPayloadTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksPayloadTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20075**](InlineResponse20075.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWebhooksPayloadTemplate

> InlineResponse20075 UpdateNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).UpdateNetworkWebhooksPayloadTemplate(updateNetworkWebhooksPayloadTemplate).Execute()

Update a webhook payload template for a network



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
    payloadTemplateId := "payloadTemplateId_example" // string | Payload template ID
    updateNetworkWebhooksPayloadTemplate := *openapiclient.NewInlineObject146() // InlineObject146 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayloadTemplatesApi.UpdateNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).UpdateNetworkWebhooksPayloadTemplate(updateNetworkWebhooksPayloadTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayloadTemplatesApi.UpdateNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWebhooksPayloadTemplate`: InlineResponse20075
    fmt.Fprintf(os.Stdout, "Response from `PayloadTemplatesApi.UpdateNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**payloadTemplateId** | **string** | Payload template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWebhooksPayloadTemplate** | [**InlineObject146**](InlineObject146.md) |  | 

### Return type

[**InlineResponse20075**](InlineResponse20075.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

