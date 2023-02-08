# \MqttBrokersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkMqttBroker**](MqttBrokersApi.md#CreateNetworkMqttBroker) | **Post** /networks/{networkId}/mqttBrokers | Add an MQTT broker
[**DeleteNetworkMqttBroker**](MqttBrokersApi.md#DeleteNetworkMqttBroker) | **Delete** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Delete an MQTT broker
[**GetNetworkMqttBroker**](MqttBrokersApi.md#GetNetworkMqttBroker) | **Get** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Return an MQTT broker
[**GetNetworkMqttBrokers**](MqttBrokersApi.md#GetNetworkMqttBrokers) | **Get** /networks/{networkId}/mqttBrokers | List the MQTT brokers for this network
[**UpdateNetworkMqttBroker**](MqttBrokersApi.md#UpdateNetworkMqttBroker) | **Put** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Update an MQTT broker



## CreateNetworkMqttBroker

> map[string]interface{} CreateNetworkMqttBroker(ctx, networkId).CreateNetworkMqttBroker(createNetworkMqttBroker).Execute()

Add an MQTT broker



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
    createNetworkMqttBroker := *openapiclient.NewCreateNetworkMqttBrokerRequest("Name_example", "Host_example", int32(123)) // CreateNetworkMqttBrokerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttBrokersApi.CreateNetworkMqttBroker(context.Background(), networkId).CreateNetworkMqttBroker(createNetworkMqttBroker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersApi.CreateNetworkMqttBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkMqttBroker`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MqttBrokersApi.CreateNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkMqttBroker** | [**CreateNetworkMqttBrokerRequest**](CreateNetworkMqttBrokerRequest.md) |  | 

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


## DeleteNetworkMqttBroker

> DeleteNetworkMqttBroker(ctx, networkId, mqttBrokerId).Execute()

Delete an MQTT broker



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
    mqttBrokerId := "mqttBrokerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttBrokersApi.DeleteNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersApi.DeleteNetworkMqttBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**mqttBrokerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkMqttBrokerRequest struct via the builder pattern


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


## GetNetworkMqttBroker

> map[string]interface{} GetNetworkMqttBroker(ctx, networkId, mqttBrokerId).Execute()

Return an MQTT broker



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
    mqttBrokerId := "mqttBrokerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttBrokersApi.GetNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersApi.GetNetworkMqttBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMqttBroker`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MqttBrokersApi.GetNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**mqttBrokerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMqttBrokerRequest struct via the builder pattern


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


## GetNetworkMqttBrokers

> []map[string]interface{} GetNetworkMqttBrokers(ctx, networkId).Execute()

List the MQTT brokers for this network



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
    resp, r, err := apiClient.MqttBrokersApi.GetNetworkMqttBrokers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersApi.GetNetworkMqttBrokers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMqttBrokers`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MqttBrokersApi.GetNetworkMqttBrokers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMqttBrokersRequest struct via the builder pattern


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


## UpdateNetworkMqttBroker

> map[string]interface{} UpdateNetworkMqttBroker(ctx, networkId, mqttBrokerId).UpdateNetworkMqttBroker(updateNetworkMqttBroker).Execute()

Update an MQTT broker



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
    mqttBrokerId := "mqttBrokerId_example" // string | 
    updateNetworkMqttBroker := *openapiclient.NewUpdateNetworkMqttBrokerRequest() // UpdateNetworkMqttBrokerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttBrokersApi.UpdateNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).UpdateNetworkMqttBroker(updateNetworkMqttBroker).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersApi.UpdateNetworkMqttBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkMqttBroker`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MqttBrokersApi.UpdateNetworkMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**mqttBrokerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkMqttBroker** | [**UpdateNetworkMqttBrokerRequest**](UpdateNetworkMqttBrokerRequest.md) |  | 

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

