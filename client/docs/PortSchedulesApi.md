# \PortSchedulesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchPortSchedule**](PortSchedulesApi.md#CreateNetworkSwitchPortSchedule) | **Post** /networks/{networkId}/switch/portSchedules | Add a switch port schedule
[**DeleteNetworkSwitchPortSchedule**](PortSchedulesApi.md#DeleteNetworkSwitchPortSchedule) | **Delete** /networks/{networkId}/switch/portSchedules/{portScheduleId} | Delete a switch port schedule
[**GetNetworkSwitchPortSchedules**](PortSchedulesApi.md#GetNetworkSwitchPortSchedules) | **Get** /networks/{networkId}/switch/portSchedules | List switch port schedules
[**UpdateNetworkSwitchPortSchedule**](PortSchedulesApi.md#UpdateNetworkSwitchPortSchedule) | **Put** /networks/{networkId}/switch/portSchedules/{portScheduleId} | Update a switch port schedule



## CreateNetworkSwitchPortSchedule

> map[string]interface{} CreateNetworkSwitchPortSchedule(ctx, networkId).CreateNetworkSwitchPortSchedule(createNetworkSwitchPortSchedule).Execute()

Add a switch port schedule



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
    createNetworkSwitchPortSchedule := *openapiclient.NewCreateNetworkSwitchPortScheduleRequest("Name_example") // CreateNetworkSwitchPortScheduleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortSchedulesApi.CreateNetworkSwitchPortSchedule(context.Background(), networkId).CreateNetworkSwitchPortSchedule(createNetworkSwitchPortSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortSchedulesApi.CreateNetworkSwitchPortSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchPortSchedule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PortSchedulesApi.CreateNetworkSwitchPortSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchPortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchPortSchedule** | [**CreateNetworkSwitchPortScheduleRequest**](CreateNetworkSwitchPortScheduleRequest.md) |  | 

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


## DeleteNetworkSwitchPortSchedule

> DeleteNetworkSwitchPortSchedule(ctx, networkId, portScheduleId).Execute()

Delete a switch port schedule



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
    portScheduleId := "portScheduleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortSchedulesApi.DeleteNetworkSwitchPortSchedule(context.Background(), networkId, portScheduleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortSchedulesApi.DeleteNetworkSwitchPortSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**portScheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchPortScheduleRequest struct via the builder pattern


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


## GetNetworkSwitchPortSchedules

> []map[string]interface{} GetNetworkSwitchPortSchedules(ctx, networkId).Execute()

List switch port schedules



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
    resp, r, err := apiClient.PortSchedulesApi.GetNetworkSwitchPortSchedules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortSchedulesApi.GetNetworkSwitchPortSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchPortSchedules`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PortSchedulesApi.GetNetworkSwitchPortSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchPortSchedulesRequest struct via the builder pattern


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


## UpdateNetworkSwitchPortSchedule

> map[string]interface{} UpdateNetworkSwitchPortSchedule(ctx, networkId, portScheduleId).UpdateNetworkSwitchPortSchedule(updateNetworkSwitchPortSchedule).Execute()

Update a switch port schedule



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
    portScheduleId := "portScheduleId_example" // string | 
    updateNetworkSwitchPortSchedule := *openapiclient.NewUpdateNetworkSwitchPortScheduleRequest() // UpdateNetworkSwitchPortScheduleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortSchedulesApi.UpdateNetworkSwitchPortSchedule(context.Background(), networkId, portScheduleId).UpdateNetworkSwitchPortSchedule(updateNetworkSwitchPortSchedule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortSchedulesApi.UpdateNetworkSwitchPortSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchPortSchedule`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PortSchedulesApi.UpdateNetworkSwitchPortSchedule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**portScheduleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchPortScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchPortSchedule** | [**UpdateNetworkSwitchPortScheduleRequest**](UpdateNetworkSwitchPortScheduleRequest.md) |  | 

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

