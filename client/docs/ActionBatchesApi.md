# \ActionBatchesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationActionBatch**](ActionBatchesApi.md#CreateOrganizationActionBatch) | **Post** /organizations/{organizationId}/actionBatches | Create an action batch
[**DeleteOrganizationActionBatch**](ActionBatchesApi.md#DeleteOrganizationActionBatch) | **Delete** /organizations/{organizationId}/actionBatches/{actionBatchId} | Delete an action batch
[**GetOrganizationActionBatch**](ActionBatchesApi.md#GetOrganizationActionBatch) | **Get** /organizations/{organizationId}/actionBatches/{actionBatchId} | Return an action batch
[**GetOrganizationActionBatches**](ActionBatchesApi.md#GetOrganizationActionBatches) | **Get** /organizations/{organizationId}/actionBatches | Return the list of action batches in the organization
[**UpdateOrganizationActionBatch**](ActionBatchesApi.md#UpdateOrganizationActionBatch) | **Put** /organizations/{organizationId}/actionBatches/{actionBatchId} | Update an action batch



## CreateOrganizationActionBatch

> CreateOrganizationActionBatch201Response CreateOrganizationActionBatch(ctx, organizationId).CreateOrganizationActionBatchRequest(createOrganizationActionBatchRequest).Execute()

Create an action batch



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
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationActionBatchRequest := *openapiclient.NewCreateOrganizationActionBatchRequest([]openapiclient.CreateOrganizationActionBatchRequestActionsInner{*openapiclient.NewCreateOrganizationActionBatchRequestActionsInner("Resource_example", "Operation_example")}) // CreateOrganizationActionBatchRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionBatchesApi.CreateOrganizationActionBatch(context.Background(), organizationId).CreateOrganizationActionBatchRequest(createOrganizationActionBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionBatchesApi.CreateOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationActionBatch`: CreateOrganizationActionBatch201Response
    fmt.Fprintf(os.Stdout, "Response from `ActionBatchesApi.CreateOrganizationActionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationActionBatchRequest** | [**CreateOrganizationActionBatchRequest**](CreateOrganizationActionBatchRequest.md) |  | 

### Return type

[**CreateOrganizationActionBatch201Response**](CreateOrganizationActionBatch201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationActionBatch

> DeleteOrganizationActionBatch(ctx, organizationId, actionBatchId).Execute()

Delete an action batch



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
    organizationId := "organizationId_example" // string | Organization ID
    actionBatchId := "actionBatchId_example" // string | Action batch ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActionBatchesApi.DeleteOrganizationActionBatch(context.Background(), organizationId, actionBatchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionBatchesApi.DeleteOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**actionBatchId** | **string** | Action batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationActionBatch

> CreateOrganizationActionBatch201Response GetOrganizationActionBatch(ctx, organizationId, actionBatchId).Execute()

Return an action batch



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
    organizationId := "organizationId_example" // string | Organization ID
    actionBatchId := "actionBatchId_example" // string | Action batch ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionBatchesApi.GetOrganizationActionBatch(context.Background(), organizationId, actionBatchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionBatchesApi.GetOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationActionBatch`: CreateOrganizationActionBatch201Response
    fmt.Fprintf(os.Stdout, "Response from `ActionBatchesApi.GetOrganizationActionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**actionBatchId** | **string** | Action batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateOrganizationActionBatch201Response**](CreateOrganizationActionBatch201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationActionBatches

> []GetOrganizationActionBatches200ResponseInner GetOrganizationActionBatches(ctx, organizationId).Status(status).Execute()

Return the list of action batches in the organization



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
    organizationId := "organizationId_example" // string | Organization ID
    status := "status_example" // string | Filter batches by status. Valid types are pending, completed, and failed. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionBatchesApi.GetOrganizationActionBatches(context.Background(), organizationId).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionBatchesApi.GetOrganizationActionBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationActionBatches`: []GetOrganizationActionBatches200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ActionBatchesApi.GetOrganizationActionBatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationActionBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Filter batches by status. Valid types are pending, completed, and failed. | 

### Return type

[**[]GetOrganizationActionBatches200ResponseInner**](GetOrganizationActionBatches200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationActionBatch

> GetOrganizationActionBatches200ResponseInner UpdateOrganizationActionBatch(ctx, organizationId, actionBatchId).UpdateOrganizationActionBatchRequest(updateOrganizationActionBatchRequest).Execute()

Update an action batch



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
    organizationId := "organizationId_example" // string | Organization ID
    actionBatchId := "actionBatchId_example" // string | Action batch ID
    updateOrganizationActionBatchRequest := *openapiclient.NewUpdateOrganizationActionBatchRequest() // UpdateOrganizationActionBatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActionBatchesApi.UpdateOrganizationActionBatch(context.Background(), organizationId, actionBatchId).UpdateOrganizationActionBatchRequest(updateOrganizationActionBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActionBatchesApi.UpdateOrganizationActionBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationActionBatch`: GetOrganizationActionBatches200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ActionBatchesApi.UpdateOrganizationActionBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**actionBatchId** | **string** | Action batch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationActionBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationActionBatchRequest** | [**UpdateOrganizationActionBatchRequest**](UpdateOrganizationActionBatchRequest.md) |  | 

### Return type

[**GetOrganizationActionBatches200ResponseInner**](GetOrganizationActionBatches200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

