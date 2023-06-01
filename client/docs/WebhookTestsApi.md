# \WebhookTestsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWebhooksWebhookTest**](WebhookTestsApi.md#CreateNetworkWebhooksWebhookTest) | **Post** /networks/{networkId}/webhooks/webhookTests | Send a test webhook for a network
[**GetNetworkWebhooksWebhookTest**](WebhookTestsApi.md#GetNetworkWebhooksWebhookTest) | **Get** /networks/{networkId}/webhooks/webhookTests/{webhookTestId} | Return the status of a webhook test for a network



## CreateNetworkWebhooksWebhookTest

> InlineResponse2013 CreateNetworkWebhooksWebhookTest(ctx, networkId).CreateNetworkWebhooksWebhookTest(createNetworkWebhooksWebhookTest).Execute()

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
    networkId := "networkId_example" // string | Network ID
    createNetworkWebhooksWebhookTest := *openapiclient.NewInlineObject147("Url_example") // InlineObject147 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookTestsApi.CreateNetworkWebhooksWebhookTest(context.Background(), networkId).CreateNetworkWebhooksWebhookTest(createNetworkWebhooksWebhookTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookTestsApi.CreateNetworkWebhooksWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWebhooksWebhookTest`: InlineResponse2013
    fmt.Fprintf(os.Stdout, "Response from `WebhookTestsApi.CreateNetworkWebhooksWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWebhooksWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWebhooksWebhookTest** | [**InlineObject147**](InlineObject147.md) |  | 

### Return type

[**InlineResponse2013**](InlineResponse2013.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWebhooksWebhookTest

> InlineResponse2013 GetNetworkWebhooksWebhookTest(ctx, networkId, webhookTestId).Execute()

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
    networkId := "networkId_example" // string | Network ID
    webhookTestId := "webhookTestId_example" // string | Webhook test ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhookTestsApi.GetNetworkWebhooksWebhookTest(context.Background(), networkId, webhookTestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhookTestsApi.GetNetworkWebhooksWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWebhooksWebhookTest`: InlineResponse2013
    fmt.Fprintf(os.Stdout, "Response from `WebhookTestsApi.GetNetworkWebhooksWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**webhookTestId** | **string** | Webhook test ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWebhooksWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse2013**](InlineResponse2013.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

