# \OnboardingApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationInventoryOnboardingCloudMonitoringImport**](OnboardingApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringImport) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.
[**CreateOrganizationInventoryOnboardingCloudMonitoringPrepare**](OnboardingApi.md#CreateOrganizationInventoryOnboardingCloudMonitoringPrepare) | **Post** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/prepare | Initiates or updates an import session
[**GetOrganizationCameraOnboardingStatuses**](OnboardingApi.md#GetOrganizationCameraOnboardingStatuses) | **Get** /organizations/{organizationId}/camera/onboarding/statuses | Fetch onboarding status of cameras
[**GetOrganizationInventoryOnboardingCloudMonitoringImports**](OnboardingApi.md#GetOrganizationInventoryOnboardingCloudMonitoringImports) | **Get** /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports | Check the status of a committed Import operation
[**UpdateOrganizationCameraOnboardingStatuses**](OnboardingApi.md#UpdateOrganizationCameraOnboardingStatuses) | **Put** /organizations/{organizationId}/camera/onboarding/statuses | Notify that credential handoff to camera has completed



## CreateOrganizationInventoryOnboardingCloudMonitoringImport

> []CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner CreateOrganizationInventoryOnboardingCloudMonitoringImport(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport).Execute()

Commits the import operation to complete the onboarding of a device into Dashboard for monitoring.



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
    organizationId := "organizationId_example" // string | 
    createOrganizationInventoryOnboardingCloudMonitoringImport := *openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest([]openapiclient.CreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner{*openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringImportRequestDevicesInner("DeviceId_example", "Udi_example", "NetworkId_example")}) // CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringImport(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringImport(createOrganizationInventoryOnboardingCloudMonitoringImport).Execute()
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
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringImport** | [**CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest**](CreateOrganizationInventoryOnboardingCloudMonitoringImportRequest.md) |  | 

### Return type

[**[]CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner**](CreateOrganizationInventoryOnboardingCloudMonitoringImport201ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationInventoryOnboardingCloudMonitoringPrepare

> []CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(ctx, organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(createOrganizationInventoryOnboardingCloudMonitoringPrepare).Execute()

Initiates or updates an import session



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
    organizationId := "organizationId_example" // string | 
    createOrganizationInventoryOnboardingCloudMonitoringPrepare := *openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest([]openapiclient.CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner{*openapiclient.NewCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequestDevicesInner("Sudi_example")}) // CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnboardingApi.CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(context.Background(), organizationId).CreateOrganizationInventoryOnboardingCloudMonitoringPrepare(createOrganizationInventoryOnboardingCloudMonitoringPrepare).Execute()
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
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInventoryOnboardingCloudMonitoringPrepare** | [**CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest**](CreateOrganizationInventoryOnboardingCloudMonitoringPrepareRequest.md) |  | 

### Return type

[**[]CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner**](CreateOrganizationInventoryOnboardingCloudMonitoringPrepare201ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

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
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | 
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
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraOnboardingStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned cameras will be filtered to only include these serials. | 
 **networkIds** | **[]string** | A list of network IDs. The returned cameras will be filtered to only include these networks. | 

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
    openapiclient "./openapi"
)

func main() {
    organizationId := "organizationId_example" // string | 
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
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInventoryOnboardingCloudMonitoringImportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importIds** | **[]string** | import ids from an imports | 

### Return type

[**[]GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner**](GetOrganizationInventoryOnboardingCloudMonitoringImports200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationCameraOnboardingStatuses

> map[string]interface{} UpdateOrganizationCameraOnboardingStatuses(ctx, organizationId).UpdateOrganizationCameraOnboardingStatuses(updateOrganizationCameraOnboardingStatuses).Execute()

Notify that credential handoff to camera has completed



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
    organizationId := "organizationId_example" // string | 
    updateOrganizationCameraOnboardingStatuses := *openapiclient.NewUpdateOrganizationCameraOnboardingStatusesRequest() // UpdateOrganizationCameraOnboardingStatusesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OnboardingApi.UpdateOrganizationCameraOnboardingStatuses(context.Background(), organizationId).UpdateOrganizationCameraOnboardingStatuses(updateOrganizationCameraOnboardingStatuses).Execute()
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
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCameraOnboardingStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationCameraOnboardingStatuses** | [**UpdateOrganizationCameraOnboardingStatusesRequest**](UpdateOrganizationCameraOnboardingStatusesRequest.md) |  | 

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

