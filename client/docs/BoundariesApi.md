# \BoundariesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationCameraBoundariesAreasByDevice**](BoundariesApi.md#GetOrganizationCameraBoundariesAreasByDevice) | **Get** /organizations/{organizationId}/camera/boundaries/areas/byDevice | Returns all configured area boundaries of cameras
[**GetOrganizationCameraBoundariesLinesByDevice**](BoundariesApi.md#GetOrganizationCameraBoundariesLinesByDevice) | **Get** /organizations/{organizationId}/camera/boundaries/lines/byDevice | Returns all configured crossingline boundaries of cameras



## GetOrganizationCameraBoundariesAreasByDevice

> []GetOrganizationCameraBoundariesAreasByDevice200ResponseInner GetOrganizationCameraBoundariesAreasByDevice(ctx, organizationId).Serials(serials).Execute()

Returns all configured area boundaries of cameras



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

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoundariesApi.GetOrganizationCameraBoundariesAreasByDevice(context.Background(), organizationId).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoundariesApi.GetOrganizationCameraBoundariesAreasByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraBoundariesAreasByDevice`: []GetOrganizationCameraBoundariesAreasByDevice200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `BoundariesApi.GetOrganizationCameraBoundariesAreasByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraBoundariesAreasByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned cameras will be filtered to only include these serials. | 

### Return type

[**[]GetOrganizationCameraBoundariesAreasByDevice200ResponseInner**](GetOrganizationCameraBoundariesAreasByDevice200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraBoundariesLinesByDevice

> []GetOrganizationCameraBoundariesLinesByDevice200ResponseInner GetOrganizationCameraBoundariesLinesByDevice(ctx, organizationId).Serials(serials).Execute()

Returns all configured crossingline boundaries of cameras



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

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BoundariesApi.GetOrganizationCameraBoundariesLinesByDevice(context.Background(), organizationId).Serials(serials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BoundariesApi.GetOrganizationCameraBoundariesLinesByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraBoundariesLinesByDevice`: []GetOrganizationCameraBoundariesLinesByDevice200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `BoundariesApi.GetOrganizationCameraBoundariesLinesByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraBoundariesLinesByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned cameras will be filtered to only include these serials. | 

### Return type

[**[]GetOrganizationCameraBoundariesLinesByDevice200ResponseInner**](GetOrganizationCameraBoundariesLinesByDevice200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

