# \WebhooksApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWebhooksHttpServer**](WebhooksApi.md#CreateNetworkWebhooksHttpServer) | **Post** /networks/{networkId}/webhooks/httpServers | Add an HTTP server to a network
[**CreateNetworkWebhooksPayloadTemplate**](WebhooksApi.md#CreateNetworkWebhooksPayloadTemplate) | **Post** /networks/{networkId}/webhooks/payloadTemplates | Create a webhook payload template for a network
[**CreateNetworkWebhooksWebhookTest**](WebhooksApi.md#CreateNetworkWebhooksWebhookTest) | **Post** /networks/{networkId}/webhooks/webhookTests | Send a test webhook for a network
[**DeleteNetworkWebhooksHttpServer**](WebhooksApi.md#DeleteNetworkWebhooksHttpServer) | **Delete** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Delete an HTTP server from a network
[**DeleteNetworkWebhooksPayloadTemplate**](WebhooksApi.md#DeleteNetworkWebhooksPayloadTemplate) | **Delete** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Destroy a webhook payload template for a network
[**GetNetworkWebhooksHttpServer**](WebhooksApi.md#GetNetworkWebhooksHttpServer) | **Get** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Return an HTTP server for a network
[**GetNetworkWebhooksHttpServers**](WebhooksApi.md#GetNetworkWebhooksHttpServers) | **Get** /networks/{networkId}/webhooks/httpServers | List the HTTP servers for a network
[**GetNetworkWebhooksPayloadTemplate**](WebhooksApi.md#GetNetworkWebhooksPayloadTemplate) | **Get** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Get the webhook payload template for a network
[**GetNetworkWebhooksPayloadTemplates**](WebhooksApi.md#GetNetworkWebhooksPayloadTemplates) | **Get** /networks/{networkId}/webhooks/payloadTemplates | List the webhook payload templates for a network
[**GetNetworkWebhooksWebhookTest**](WebhooksApi.md#GetNetworkWebhooksWebhookTest) | **Get** /networks/{networkId}/webhooks/webhookTests/{webhookTestId} | Return the status of a webhook test for a network
[**GetOrganizationWebhooksAlertTypes**](WebhooksApi.md#GetOrganizationWebhooksAlertTypes) | **Get** /organizations/{organizationId}/webhooks/alertTypes | Return a list of alert types to be used with managing webhook alerts
[**GetOrganizationWebhooksLogs**](WebhooksApi.md#GetOrganizationWebhooksLogs) | **Get** /organizations/{organizationId}/webhooks/logs | Return the log of webhook POSTs sent
[**UpdateNetworkWebhooksHttpServer**](WebhooksApi.md#UpdateNetworkWebhooksHttpServer) | **Put** /networks/{networkId}/webhooks/httpServers/{httpServerId} | Update an HTTP server
[**UpdateNetworkWebhooksPayloadTemplate**](WebhooksApi.md#UpdateNetworkWebhooksPayloadTemplate) | **Put** /networks/{networkId}/webhooks/payloadTemplates/{payloadTemplateId} | Update a webhook payload template for a network



## CreateNetworkWebhooksHttpServer

> GetNetworkWebhooksHttpServers200ResponseInner CreateNetworkWebhooksHttpServer(ctx, networkId).CreateNetworkWebhooksHttpServer(createNetworkWebhooksHttpServer).Execute()

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
    networkId := "networkId_example" // string | 
    createNetworkWebhooksHttpServer := *openapiclient.NewCreateNetworkWebhooksHttpServerRequest("Name_example", "Url_example") // CreateNetworkWebhooksHttpServerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.CreateNetworkWebhooksHttpServer(context.Background(), networkId).CreateNetworkWebhooksHttpServer(createNetworkWebhooksHttpServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.CreateNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWebhooksHttpServer`: GetNetworkWebhooksHttpServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.CreateNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksHttpServer** | [**CreateNetworkWebhooksHttpServerRequest**](CreateNetworkWebhooksHttpServerRequest.md) |  | 

### Return type

[**GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner CreateNetworkWebhooksPayloadTemplate(ctx, networkId).CreateNetworkWebhooksPayloadTemplate(createNetworkWebhooksPayloadTemplate).Execute()

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
    networkId := "networkId_example" // string | 
    createNetworkWebhooksPayloadTemplate := *openapiclient.NewCreateNetworkWebhooksPayloadTemplateRequest("Name_example") // CreateNetworkWebhooksPayloadTemplateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.CreateNetworkWebhooksPayloadTemplate(context.Background(), networkId).CreateNetworkWebhooksPayloadTemplate(createNetworkWebhooksPayloadTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.CreateNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.CreateNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksPayloadTemplate** | [**CreateNetworkWebhooksPayloadTemplateRequest**](CreateNetworkWebhooksPayloadTemplateRequest.md) |  | 

### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWebhooksWebhookTest

> CreateNetworkWebhooksWebhookTest201Response CreateNetworkWebhooksWebhookTest(ctx, networkId).CreateNetworkWebhooksWebhookTest(createNetworkWebhooksWebhookTest).Execute()

Send a test webhook for a network



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
    createNetworkWebhooksWebhookTest := *openapiclient.NewCreateNetworkWebhooksWebhookTestRequest("Url_example") // CreateNetworkWebhooksWebhookTestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.CreateNetworkWebhooksWebhookTest(context.Background(), networkId).CreateNetworkWebhooksWebhookTest(createNetworkWebhooksWebhookTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.CreateNetworkWebhooksWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWebhooksWebhookTest`: CreateNetworkWebhooksWebhookTest201Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.CreateNetworkWebhooksWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksWebhookTest** | [**CreateNetworkWebhooksWebhookTestRequest**](CreateNetworkWebhooksWebhookTestRequest.md) |  | 

### Return type

[**CreateNetworkWebhooksWebhookTest201Response**](CreateNetworkWebhooksWebhookTest201Response.md)

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
    networkId := "networkId_example" // string | 
    httpServerId := "httpServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.DeleteNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**httpServerId** | **string** |  | 

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
    networkId := "networkId_example" // string | 
    payloadTemplateId := "payloadTemplateId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.DeleteNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**payloadTemplateId** | **string** |  | 

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


## GetNetworkWebhooksHttpServer

> GetNetworkWebhooksHttpServers200ResponseInner GetNetworkWebhooksHttpServer(ctx, networkId, httpServerId).Execute()

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
    networkId := "networkId_example" // string | 
    httpServerId := "httpServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksHttpServer`: GetNetworkWebhooksHttpServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**httpServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksHttpServers

> []GetNetworkWebhooksHttpServers200ResponseInner GetNetworkWebhooksHttpServers(ctx, networkId).Execute()

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
    networkId := "networkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetNetworkWebhooksHttpServers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetNetworkWebhooksHttpServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksHttpServers`: []GetNetworkWebhooksHttpServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetNetworkWebhooksHttpServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksHttpServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner GetNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).Execute()

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
    networkId := "networkId_example" // string | 
    payloadTemplateId := "payloadTemplateId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**payloadTemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksPayloadTemplates

> []GetNetworkWebhooksPayloadTemplates200ResponseInner GetNetworkWebhooksPayloadTemplates(ctx, networkId).Execute()

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
    networkId := "networkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetNetworkWebhooksPayloadTemplates(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetNetworkWebhooksPayloadTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksPayloadTemplates`: []GetNetworkWebhooksPayloadTemplates200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetNetworkWebhooksPayloadTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksPayloadTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksWebhookTest

> CreateNetworkWebhooksWebhookTest201Response GetNetworkWebhooksWebhookTest(ctx, networkId, webhookTestId).Execute()

Return the status of a webhook test for a network



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
    webhookTestId := "webhookTestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetNetworkWebhooksWebhookTest(context.Background(), networkId, webhookTestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetNetworkWebhooksWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksWebhookTest`: CreateNetworkWebhooksWebhookTest201Response
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetNetworkWebhooksWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**webhookTestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateNetworkWebhooksWebhookTest201Response**](CreateNetworkWebhooksWebhookTest201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWebhooksAlertTypes

> []map[string]interface{} GetOrganizationWebhooksAlertTypes(ctx, organizationId).ProductType(productType).Execute()

Return a list of alert types to be used with managing webhook alerts



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
    organizationId := "organizationId_example" // string | 
    productType := "productType_example" // string | Filter sample alerts to a specific product type (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetOrganizationWebhooksAlertTypes(context.Background(), organizationId).ProductType(productType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetOrganizationWebhooksAlertTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWebhooksAlertTypes`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetOrganizationWebhooksAlertTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWebhooksAlertTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **productType** | **string** | Filter sample alerts to a specific product type | 

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


## GetOrganizationWebhooksLogs

> []GetOrganizationWebhooksLogs200ResponseInner GetOrganizationWebhooksLogs(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Url(url).Execute()

Return the log of webhook POSTs sent



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
    organizationId := "organizationId_example" // string | 
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    url := "url_example" // string | The URL the webhook was sent to (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.GetOrganizationWebhooksLogs(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Url(url).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetOrganizationWebhooksLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWebhooksLogs`: []GetOrganizationWebhooksLogs200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetOrganizationWebhooksLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWebhooksLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 90 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **url** | **string** | The URL the webhook was sent to | 

### Return type

[**[]GetOrganizationWebhooksLogs200ResponseInner**](GetOrganizationWebhooksLogs200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWebhooksHttpServer

> GetNetworkWebhooksHttpServers200ResponseInner UpdateNetworkWebhooksHttpServer(ctx, networkId, httpServerId).UpdateNetworkWebhooksHttpServer(updateNetworkWebhooksHttpServer).Execute()

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
    networkId := "networkId_example" // string | 
    httpServerId := "httpServerId_example" // string | 
    updateNetworkWebhooksHttpServer := *openapiclient.NewUpdateNetworkWebhooksHttpServerRequest() // UpdateNetworkWebhooksHttpServerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.UpdateNetworkWebhooksHttpServer(context.Background(), networkId, httpServerId).UpdateNetworkWebhooksHttpServer(updateNetworkWebhooksHttpServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.UpdateNetworkWebhooksHttpServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWebhooksHttpServer`: GetNetworkWebhooksHttpServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.UpdateNetworkWebhooksHttpServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**httpServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWebhooksHttpServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWebhooksHttpServer** | [**UpdateNetworkWebhooksHttpServerRequest**](UpdateNetworkWebhooksHttpServerRequest.md) |  | 

### Return type

[**GetNetworkWebhooksHttpServers200ResponseInner**](GetNetworkWebhooksHttpServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWebhooksPayloadTemplate

> GetNetworkWebhooksPayloadTemplates200ResponseInner UpdateNetworkWebhooksPayloadTemplate(ctx, networkId, payloadTemplateId).UpdateNetworkWebhooksPayloadTemplate(updateNetworkWebhooksPayloadTemplate).Execute()

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
    networkId := "networkId_example" // string | 
    payloadTemplateId := "payloadTemplateId_example" // string | 
    updateNetworkWebhooksPayloadTemplate := *openapiclient.NewUpdateNetworkWebhooksPayloadTemplateRequest() // UpdateNetworkWebhooksPayloadTemplateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.UpdateNetworkWebhooksPayloadTemplate(context.Background(), networkId, payloadTemplateId).UpdateNetworkWebhooksPayloadTemplate(updateNetworkWebhooksPayloadTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.UpdateNetworkWebhooksPayloadTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWebhooksPayloadTemplate`: GetNetworkWebhooksPayloadTemplates200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.UpdateNetworkWebhooksPayloadTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**payloadTemplateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWebhooksPayloadTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWebhooksPayloadTemplate** | [**UpdateNetworkWebhooksPayloadTemplateRequest**](UpdateNetworkWebhooksPayloadTemplateRequest.md) |  | 

### Return type

[**GetNetworkWebhooksPayloadTemplates200ResponseInner**](GetNetworkWebhooksPayloadTemplates200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

