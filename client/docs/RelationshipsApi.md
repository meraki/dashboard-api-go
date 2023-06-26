# \RelationshipsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceSensorRelationships**](RelationshipsApi.md#GetDeviceSensorRelationships) | **Get** /devices/{serial}/sensor/relationships | List the sensor roles for a given sensor or camera device.
[**GetNetworkSensorRelationships**](RelationshipsApi.md#GetNetworkSensorRelationships) | **Get** /networks/{networkId}/sensor/relationships | List the sensor roles for devices in a given network
[**UpdateDeviceSensorRelationships**](RelationshipsApi.md#UpdateDeviceSensorRelationships) | **Put** /devices/{serial}/sensor/relationships | Assign one or more sensor roles to a given sensor or camera device.



## GetDeviceSensorRelationships

> []InlineResponse2004 GetDeviceSensorRelationships(ctx, serial).Execute()

List the sensor roles for a given sensor or camera device.



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipsApi.GetDeviceSensorRelationships(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsApi.GetDeviceSensorRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceSensorRelationships`: []InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `RelationshipsApi.GetDeviceSensorRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceSensorRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse2004**](InlineResponse2004.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorRelationships

> []InlineResponse20043 GetNetworkSensorRelationships(ctx, networkId).Execute()

List the sensor roles for devices in a given network



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
    resp, r, err := apiClient.RelationshipsApi.GetNetworkSensorRelationships(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsApi.GetNetworkSensorRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorRelationships`: []InlineResponse20043
    fmt.Fprintf(os.Stdout, "Response from `RelationshipsApi.GetNetworkSensorRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20043**](InlineResponse20043.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceSensorRelationships

> InlineResponse2004 UpdateDeviceSensorRelationships(ctx, serial).UpdateDeviceSensorRelationships(updateDeviceSensorRelationships).Execute()

Assign one or more sensor roles to a given sensor or camera device.



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
    serial := "serial_example" // string | Serial
    updateDeviceSensorRelationships := *openapiclient.NewInlineObject16() // InlineObject16 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RelationshipsApi.UpdateDeviceSensorRelationships(context.Background(), serial).UpdateDeviceSensorRelationships(updateDeviceSensorRelationships).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RelationshipsApi.UpdateDeviceSensorRelationships``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceSensorRelationships`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `RelationshipsApi.UpdateDeviceSensorRelationships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceSensorRelationshipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceSensorRelationships** | [**InlineObject16**](InlineObject16.md) |  | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

