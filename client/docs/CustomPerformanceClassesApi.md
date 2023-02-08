# \CustomPerformanceClassesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkApplianceTrafficShapingCustomPerformanceClass**](CustomPerformanceClassesApi.md#CreateNetworkApplianceTrafficShapingCustomPerformanceClass) | **Post** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses | Add a custom performance class for an MX network
[**DeleteNetworkApplianceTrafficShapingCustomPerformanceClass**](CustomPerformanceClassesApi.md#DeleteNetworkApplianceTrafficShapingCustomPerformanceClass) | **Delete** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Delete a custom performance class from an MX network
[**GetNetworkApplianceTrafficShapingCustomPerformanceClass**](CustomPerformanceClassesApi.md#GetNetworkApplianceTrafficShapingCustomPerformanceClass) | **Get** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Return a custom performance class for an MX network
[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses**](CustomPerformanceClassesApi.md#GetNetworkApplianceTrafficShapingCustomPerformanceClasses) | **Get** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses | List all custom performance classes for an MX network
[**UpdateNetworkApplianceTrafficShapingCustomPerformanceClass**](CustomPerformanceClassesApi.md#UpdateNetworkApplianceTrafficShapingCustomPerformanceClass) | **Put** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Update a custom performance class for an MX network



## CreateNetworkApplianceTrafficShapingCustomPerformanceClass

> map[string]interface{} CreateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId).CreateNetworkApplianceTrafficShapingCustomPerformanceClass(createNetworkApplianceTrafficShapingCustomPerformanceClass).Execute()

Add a custom performance class for an MX network



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
    createNetworkApplianceTrafficShapingCustomPerformanceClass := *openapiclient.NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest("Name_example") // CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomPerformanceClassesApi.CreateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId).CreateNetworkApplianceTrafficShapingCustomPerformanceClass(createNetworkApplianceTrafficShapingCustomPerformanceClass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPerformanceClassesApi.CreateNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceTrafficShapingCustomPerformanceClass`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomPerformanceClassesApi.CreateNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceTrafficShapingCustomPerformanceClass** | [**CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest**](CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest.md) |  | 

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


## DeleteNetworkApplianceTrafficShapingCustomPerformanceClass

> DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).Execute()

Delete a custom performance class from an MX network



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
    customPerformanceClassId := "customPerformanceClassId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomPerformanceClassesApi.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPerformanceClassesApi.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**customPerformanceClassId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


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


## GetNetworkApplianceTrafficShapingCustomPerformanceClass

> map[string]interface{} GetNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).Execute()

Return a custom performance class for an MX network



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
    customPerformanceClassId := "customPerformanceClassId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomPerformanceClassesApi.GetNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPerformanceClassesApi.GetNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingCustomPerformanceClass`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomPerformanceClassesApi.GetNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**customPerformanceClassId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


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


## GetNetworkApplianceTrafficShapingCustomPerformanceClasses

> []map[string]interface{} GetNetworkApplianceTrafficShapingCustomPerformanceClasses(ctx, networkId).Execute()

List all custom performance classes for an MX network



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
    resp, r, err := apiClient.CustomPerformanceClassesApi.GetNetworkApplianceTrafficShapingCustomPerformanceClasses(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPerformanceClassesApi.GetNetworkApplianceTrafficShapingCustomPerformanceClasses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingCustomPerformanceClasses`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomPerformanceClassesApi.GetNetworkApplianceTrafficShapingCustomPerformanceClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest struct via the builder pattern


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


## UpdateNetworkApplianceTrafficShapingCustomPerformanceClass

> map[string]interface{} UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(updateNetworkApplianceTrafficShapingCustomPerformanceClass).Execute()

Update a custom performance class for an MX network



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
    customPerformanceClassId := "customPerformanceClassId_example" // string | 
    updateNetworkApplianceTrafficShapingCustomPerformanceClass := *openapiclient.NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest() // UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomPerformanceClassesApi.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(updateNetworkApplianceTrafficShapingCustomPerformanceClass).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomPerformanceClassesApi.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingCustomPerformanceClass`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CustomPerformanceClassesApi.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**customPerformanceClassId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceTrafficShapingCustomPerformanceClass** | [**UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest**](UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest.md) |  | 

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

