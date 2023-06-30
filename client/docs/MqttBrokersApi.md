# \MqttBrokersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkMqttBroker**](MqttBrokersApi.md#CreateNetworkMqttBroker) | **Post** /networks/{networkId}/mqttBrokers | Add an MQTT broker
[**DeleteNetworkMqttBroker**](MqttBrokersApi.md#DeleteNetworkMqttBroker) | **Delete** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Delete an MQTT broker
[**GetNetworkMqttBroker**](MqttBrokersApi.md#GetNetworkMqttBroker) | **Get** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Return an MQTT broker
[**GetNetworkMqttBrokers**](MqttBrokersApi.md#GetNetworkMqttBrokers) | **Get** /networks/{networkId}/mqttBrokers | List the MQTT brokers for this network
[**GetNetworkSensorMqttBroker**](MqttBrokersApi.md#GetNetworkSensorMqttBroker) | **Get** /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId} | Return the sensor settings of an MQTT broker
[**GetNetworkSensorMqttBrokers**](MqttBrokersApi.md#GetNetworkSensorMqttBrokers) | **Get** /networks/{networkId}/sensor/mqttBrokers | List the sensor settings of all MQTT brokers for this network
[**UpdateNetworkMqttBroker**](MqttBrokersApi.md#UpdateNetworkMqttBroker) | **Put** /networks/{networkId}/mqttBrokers/{mqttBrokerId} | Update an MQTT broker
[**UpdateNetworkSensorMqttBroker**](MqttBrokersApi.md#UpdateNetworkSensorMqttBroker) | **Put** /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId} | Update the sensor settings of an MQTT broker



## CreateNetworkMqttBroker

> map[string]interface{} CreateNetworkMqttBroker(ctx, networkId).CreateNetworkMqttBrokerRequest(createNetworkMqttBrokerRequest).Execute()

Add an MQTT broker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    createNetworkMqttBrokerRequest := *openapiclient.NewCreateNetworkMqttBrokerRequest("Name_example", "Host_example", int32(123)) // CreateNetworkMqttBrokerRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttBrokersApi.CreateNetworkMqttBroker(context.Background(), networkId).CreateNetworkMqttBrokerRequest(createNetworkMqttBrokerRequest).Execute()
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
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkMqttBrokerRequest** | [**CreateNetworkMqttBrokerRequest**](CreateNetworkMqttBrokerRequest.md) |  | 

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
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MqttBrokersApi.DeleteNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
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
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

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
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID

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
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

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
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

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
**networkId** | **string** | Network ID | 

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


## GetNetworkSensorMqttBroker

> GetNetworkSensorMqttBrokers200ResponseInner GetNetworkSensorMqttBroker(ctx, networkId, mqttBrokerId).Execute()

Return the sensor settings of an MQTT broker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttBrokersApi.GetNetworkSensorMqttBroker(context.Background(), networkId, mqttBrokerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersApi.GetNetworkSensorMqttBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorMqttBroker`: GetNetworkSensorMqttBrokers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MqttBrokersApi.GetNetworkSensorMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSensorMqttBrokers200ResponseInner**](GetNetworkSensorMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorMqttBrokers

> []GetNetworkSensorMqttBrokers200ResponseInner GetNetworkSensorMqttBrokers(ctx, networkId).Execute()

List the sensor settings of all MQTT brokers for this network



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttBrokersApi.GetNetworkSensorMqttBrokers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersApi.GetNetworkSensorMqttBrokers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorMqttBrokers`: []GetNetworkSensorMqttBrokers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MqttBrokersApi.GetNetworkSensorMqttBrokers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorMqttBrokersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSensorMqttBrokers200ResponseInner**](GetNetworkSensorMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkMqttBroker

> map[string]interface{} UpdateNetworkMqttBroker(ctx, networkId, mqttBrokerId).UpdateNetworkMqttBrokerRequest(updateNetworkMqttBrokerRequest).Execute()

Update an MQTT broker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID
    updateNetworkMqttBrokerRequest := *openapiclient.NewUpdateNetworkMqttBrokerRequest() // UpdateNetworkMqttBrokerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttBrokersApi.UpdateNetworkMqttBroker(context.Background(), networkId, mqttBrokerId).UpdateNetworkMqttBrokerRequest(updateNetworkMqttBrokerRequest).Execute()
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
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkMqttBrokerRequest** | [**UpdateNetworkMqttBrokerRequest**](UpdateNetworkMqttBrokerRequest.md) |  | 

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


## UpdateNetworkSensorMqttBroker

> GetNetworkSensorMqttBrokers200ResponseInner UpdateNetworkSensorMqttBroker(ctx, networkId, mqttBrokerId).UpdateNetworkSensorMqttBrokerRequest(updateNetworkSensorMqttBrokerRequest).Execute()

Update the sensor settings of an MQTT broker



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    mqttBrokerId := "mqttBrokerId_example" // string | Mqtt broker ID
    updateNetworkSensorMqttBrokerRequest := *openapiclient.NewUpdateNetworkSensorMqttBrokerRequest(false) // UpdateNetworkSensorMqttBrokerRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MqttBrokersApi.UpdateNetworkSensorMqttBroker(context.Background(), networkId, mqttBrokerId).UpdateNetworkSensorMqttBrokerRequest(updateNetworkSensorMqttBrokerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MqttBrokersApi.UpdateNetworkSensorMqttBroker``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSensorMqttBroker`: GetNetworkSensorMqttBrokers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MqttBrokersApi.UpdateNetworkSensorMqttBroker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**mqttBrokerId** | **string** | Mqtt broker ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSensorMqttBrokerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSensorMqttBrokerRequest** | [**UpdateNetworkSensorMqttBrokerRequest**](UpdateNetworkSensorMqttBrokerRequest.md) |  | 

### Return type

[**GetNetworkSensorMqttBrokers200ResponseInner**](GetNetworkSensorMqttBrokers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

