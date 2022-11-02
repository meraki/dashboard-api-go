# \WebhookTestsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkHttpServersWebhookTest**](WebhookTestsApi.md#CreateNetworkHttpServersWebhookTest) | **Post** /networks/{networkId}/httpServers/webhookTests | Send a test webhook for a network
[**CreateNetworkWebhooksWebhookTest**](WebhookTestsApi.md#CreateNetworkWebhooksWebhookTest) | **Post** /networks/{networkId}/webhooks/webhookTests | Send a test webhook for a network
[**GetNetworkHttpServersWebhookTest**](WebhookTestsApi.md#GetNetworkHttpServersWebhookTest) | **Get** /networks/{networkId}/httpServers/webhookTests/{id} | Return the status of a webhook test for a network
[**GetNetworkWebhooksWebhookTest**](WebhookTestsApi.md#GetNetworkWebhooksWebhookTest) | **Get** /networks/{networkId}/webhooks/webhookTests/{webhookTestId} | Return the status of a webhook test for a network



## CreateNetworkHttpServersWebhookTest

> map[string]interface{} CreateNetworkHttpServersWebhookTest(ctx, networkId).CreateNetworkHttpServersWebhookTest(createNetworkHttpServersWebhookTest).Execute()

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
    createNetworkHttpServersWebhookTest := *openapiclient.NewCreateNetworkHttpServersWebhookTestRequest("Url_example") // CreateNetworkHttpServersWebhookTestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookTestsApi.CreateNetworkHttpServersWebhookTest(context.Background(), networkId).CreateNetworkHttpServersWebhookTest(createNetworkHttpServersWebhookTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookTestsApi.CreateNetworkHttpServersWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkHttpServersWebhookTest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WebhookTestsApi.CreateNetworkHttpServersWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkHttpServersWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkHttpServersWebhookTest** | [**CreateNetworkHttpServersWebhookTestRequest**](CreateNetworkHttpServersWebhookTestRequest.md) |  | 

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
    resp, r, err := apiClient.WebhookTestsApi.CreateNetworkWebhooksWebhookTest(context.Background(), networkId).CreateNetworkWebhooksWebhookTest(createNetworkWebhooksWebhookTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookTestsApi.CreateNetworkWebhooksWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWebhooksWebhookTest`: CreateNetworkWebhooksWebhookTest201Response
    fmt.Fprintf(os.Stdout, "Response from `WebhookTestsApi.CreateNetworkWebhooksWebhookTest`: %v\n", resp)
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


## GetNetworkHttpServersWebhookTest

> map[string]interface{} GetNetworkHttpServersWebhookTest(ctx, networkId, id).Execute()

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookTestsApi.GetNetworkHttpServersWebhookTest(context.Background(), networkId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookTestsApi.GetNetworkHttpServersWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkHttpServersWebhookTest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WebhookTestsApi.GetNetworkHttpServersWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkHttpServersWebhookTestRequest struct via the builder pattern


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
    resp, r, err := apiClient.WebhookTestsApi.GetNetworkWebhooksWebhookTest(context.Background(), networkId, webhookTestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookTestsApi.GetNetworkWebhooksWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksWebhookTest`: CreateNetworkWebhooksWebhookTest201Response
    fmt.Fprintf(os.Stdout, "Response from `WebhookTestsApi.GetNetworkWebhooksWebhookTest`: %v\n", resp)
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

