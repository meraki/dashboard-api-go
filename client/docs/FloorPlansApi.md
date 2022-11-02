# \FloorPlansApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkFloorPlan**](FloorPlansApi.md#CreateNetworkFloorPlan) | **Post** /networks/{networkId}/floorPlans | Upload a floor plan
[**DeleteNetworkFloorPlan**](FloorPlansApi.md#DeleteNetworkFloorPlan) | **Delete** /networks/{networkId}/floorPlans/{floorPlanId} | Destroy a floor plan
[**GetNetworkFloorPlan**](FloorPlansApi.md#GetNetworkFloorPlan) | **Get** /networks/{networkId}/floorPlans/{floorPlanId} | Find a floor plan by ID
[**GetNetworkFloorPlans**](FloorPlansApi.md#GetNetworkFloorPlans) | **Get** /networks/{networkId}/floorPlans | List the floor plans that belong to your network
[**UpdateNetworkFloorPlan**](FloorPlansApi.md#UpdateNetworkFloorPlan) | **Put** /networks/{networkId}/floorPlans/{floorPlanId} | Update a floor plan&#39;s geolocation and other meta data



## CreateNetworkFloorPlan

> map[string]interface{} CreateNetworkFloorPlan(ctx, networkId).CreateNetworkFloorPlan(createNetworkFloorPlan).Execute()

Upload a floor plan



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
    createNetworkFloorPlan := *openapiclient.NewCreateNetworkFloorPlanRequest("Name_example", string(123)) // CreateNetworkFloorPlanRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.CreateNetworkFloorPlan(context.Background(), networkId).CreateNetworkFloorPlan(createNetworkFloorPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.CreateNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFloorPlan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.CreateNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFloorPlan** | [**CreateNetworkFloorPlanRequest**](CreateNetworkFloorPlanRequest.md) |  | 

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


## DeleteNetworkFloorPlan

> DeleteNetworkFloorPlan(ctx, networkId, floorPlanId).Execute()

Destroy a floor plan



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
    floorPlanId := "floorPlanId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.DeleteNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.DeleteNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**floorPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFloorPlanRequest struct via the builder pattern


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


## GetNetworkFloorPlan

> map[string]interface{} GetNetworkFloorPlan(ctx, networkId, floorPlanId).Execute()

Find a floor plan by ID



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
    floorPlanId := "floorPlanId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.GetNetworkFloorPlan(context.Background(), networkId, floorPlanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.GetNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFloorPlan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.GetNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**floorPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFloorPlanRequest struct via the builder pattern


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


## GetNetworkFloorPlans

> []map[string]interface{} GetNetworkFloorPlans(ctx, networkId).Execute()

List the floor plans that belong to your network



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
    resp, r, err := apiClient.FloorPlansApi.GetNetworkFloorPlans(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.GetNetworkFloorPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFloorPlans`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.GetNetworkFloorPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFloorPlansRequest struct via the builder pattern


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


## UpdateNetworkFloorPlan

> map[string]interface{} UpdateNetworkFloorPlan(ctx, networkId, floorPlanId).UpdateNetworkFloorPlan(updateNetworkFloorPlan).Execute()

Update a floor plan's geolocation and other meta data



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
    floorPlanId := "floorPlanId_example" // string | 
    updateNetworkFloorPlan := *openapiclient.NewUpdateNetworkFloorPlanRequest() // UpdateNetworkFloorPlanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloorPlansApi.UpdateNetworkFloorPlan(context.Background(), networkId, floorPlanId).UpdateNetworkFloorPlan(updateNetworkFloorPlan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloorPlansApi.UpdateNetworkFloorPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFloorPlan`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FloorPlansApi.UpdateNetworkFloorPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**floorPlanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFloorPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkFloorPlan** | [**UpdateNetworkFloorPlanRequest**](UpdateNetworkFloorPlanRequest.md) |  | 

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

