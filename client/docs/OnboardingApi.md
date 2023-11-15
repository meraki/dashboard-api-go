# \OnboardingApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent**](OnboardingApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/exportEvents | Imports event logs related to the onboarding app into elastisearch
[**CreateOrganizationInventoryOnboardingCloudMonitoringImport**](OnboardingApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringImport) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.
[**CreateOrganizationInventoryOnboardingCloudMonitoringPrepare**](OnboardingApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringPrepare) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/prepare | Initiates or updates an import session
[**GetOrganizationCameraOnboardingStatuses**](OnboardingApi.md#GetOrganizationCameraOnboardingStatuses) | **Get** /organizations/{organizationId}/camera/onboarding/statuses | Fetch onboarding status of cameras
[**GetOrganizationInventoryOnboardingCloudMonitoringImports**](OnboardingApi.md#GetOrganizationInventoryOnboardingCloudMonitoringImports) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Check the status of a committed Import operation
[**GetOrganizationInventoryOnboardingCloudMonitoringNetworks**](OnboardingApi.md#GetOrganizationInventoryOnboardingCloudMonitoringNetworks) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/networks | Returns list of networks eligible for adding cloud monitored device
[**UpdateOrganizationCameraOnboardingStatuses**](OnboardingApi.md#UpdateOrganizationCameraOnboardingStatuses) | **Put** /organizations/{organizationId}/camera/onboarding/statuses | Notify that credential handoff to camera has completed



## CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent

> map[string]interface{} CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest(createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest).Execute()

Imports event logs related to the onboarding app into elastisearch



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
    createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest := *openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest("LogEvent_example", int32(123)) // CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest(createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringExportEventRequest** | [**CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest**](CreateOrganizationInventoryOnboardingCloudMonitoringExportEventRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryOnboardingCloudMonitoringImport

> []CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest(createOrganizationInventoryOnboardingCloudMonitoringImportRequest).Execute()

Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.



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
    createOrganizationInventoryOnboardingCloudMonitoringImportRequest := *openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest([]openapiclient.CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner{*openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner("DeviceId_example", "Udi_example", "NetworkId_example")}) // CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest(createOrganizationInventoryOnboardingCloudMonitoringImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryOnboardingCloudMonitoringImport`: []CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringImportRequest** | [**CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest**](CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest.md) |  | 

### Return type

[**[]CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner**](CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryOnboardingCloudMonitoringPrepare

> []CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest).Execute()

Initiates or updates an import session



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
    createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest := *openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest([]openapiclient.CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner{*openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner("Sudi_example")}) // CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest(createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInventoryOnboardingCloudMonitoringPrepare`: []CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringPrepareRequest** | [**CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest**](CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest.md) |  | 

### Return type

[**[]CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner**](CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraOnboardingStatuses

> []map[string]interface{} GetOrganizationCameraOnboardingStatuses(ctx, organizationId).Serials(serials).NetworkIds(networkIds).Execute()

Fetch onboarding status of cameras



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
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned cameras will be filtered to only include these serials. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned cameras will be filtered to only include these networks. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnboardingApi.GetOrganizationCameraOnboardingStatuses(context.Background(), organizationId).Serials(serials).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnboardingApi.GetOrganizationCameraOnboardingStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraOnboardingStatuses`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OnboardingApi.GetOrganizationCameraOnboardingStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraOnboardingStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned cameras will be filtered to only include these serials. | 
 **networkIds** | **[]string** | A list of network IDs. The returned cameras will be filtered to only include these networks. | 

### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryOnboardingCloudMonitoringImports

> []GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner GetOrganizationInventoryOnboardingCloudMonitoringImports(ctx, organizationId).ImportIds(importIds).Execute()

Check the status of a committed Import operation



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
    importIds := []string{"Inner_example"} // []string | import ids from an imports

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnboardingApi.GetOrganizationInventoryOnboardingCloudMonitoringImports(context.Background(), organizationId).ImportIds(importIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnboardingApi.GetOrganizationInventoryOnboardingCloudMonitoringImports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryOnboardingCloudMonitoringImports`: []GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `OnboardingApi.GetOrganizationInventoryOnboardingCloudMonitoringImports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importIds** | **[]string** | import ids from an imports | 

### Return type

[**[]GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner**](GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInventoryOnboardingCloudMonitoringNetworks

> []GetNetwork200Response GetOrganizationInventoryOnboardingCloudMonitoringNetworks(ctx, organizationId).DeviceType(deviceType).Search(search).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Returns list of networks eligible for adding cloud monitored device



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
    deviceType := "deviceType_example" // string | Device Type switch or wireless controller
    search := "search_example" // string | Optional parameter to search on network name (optional)
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnboardingApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks(context.Background(), organizationId).DeviceType(deviceType).Search(search).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnboardingApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: []GetNetwork200Response
    fmt.Fprintf(os.Stdout, "Response from `OnboardingApi.GetOrganizationInventoryOnboardingCloudMonitoringNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceType** | **string** | Device Type switch or wireless controller | 
 **search** | **string** | Optional parameter to search on network name | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 100000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetwork200Response**](GetNetwork200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationCameraOnboardingStatuses

> map[string]interface{} UpdateOrganizationCameraOnboardingStatuses(ctx, organizationId).UpdateOrganizationCameraOnboardingStatusesRequest(updateOrganizationCameraOnboardingStatusesRequest).Execute()

Notify that credential handoff to camera has completed



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
    updateOrganizationCameraOnboardingStatusesRequest := *openapiclient.NewUpdateOrganizationCameraOnboardingStatusesRequest() // UpdateOrganizationCameraOnboardingStatusesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnboardingApi.UpdateOrganizationCameraOnboardingStatuses(context.Background(), organizationId).UpdateOrganizationCameraOnboardingStatusesRequest(updateOrganizationCameraOnboardingStatusesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OnboardingApi.UpdateOrganizationCameraOnboardingStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCameraOnboardingStatuses`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OnboardingApi.UpdateOrganizationCameraOnboardingStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCameraOnboardingStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationCameraOnboardingStatusesRequest** | [**UpdateOrganizationCameraOnboardingStatusesRequest**](UpdateOrganizationCameraOnboardingStatusesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

