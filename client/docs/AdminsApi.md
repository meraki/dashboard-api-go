# \AdminsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationAdmin**](AdminsApi.md#CreateOrganizationAdmin) | **Post** /organizations/{organizationId}/admins | Create a new dashboard administrator
[**DeleteOrganizationAdmin**](AdminsApi.md#DeleteOrganizationAdmin) | **Delete** /organizations/{organizationId}/admins/{adminId} | Revoke all access for a dashboard administrator within this organization
[**GetOrganizationAdmins**](AdminsApi.md#GetOrganizationAdmins) | **Get** /organizations/{organizationId}/admins | List the dashboard administrators in this organization
[**UpdateOrganizationAdmin**](AdminsApi.md#UpdateOrganizationAdmin) | **Put** /organizations/{organizationId}/admins/{adminId} | Update an administrator



## CreateOrganizationAdmin

> map[string]interface{} CreateOrganizationAdmin(ctx, organizationId).CreateOrganizationAdmin(createOrganizationAdmin).Execute()

Create a new dashboard administrator



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
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationAdmin := *openapiclient.NewInlineObject178("Email_example", "Name_example", "OrgAccess_example") // InlineObject178 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.CreateOrganizationAdmin(context.Background(), organizationId).CreateOrganizationAdmin(createOrganizationAdmin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.CreateOrganizationAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdmin`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.CreateOrganizationAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdmin** | [**InlineObject178**](InlineObject178.md) |  | 

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


## DeleteOrganizationAdmin

> DeleteOrganizationAdmin(ctx, organizationId, adminId).Execute()

Revoke all access for a dashboard administrator within this organization



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
    organizationId := "organizationId_example" // string | Organization ID
    adminId := "adminId_example" // string | Admin ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.DeleteOrganizationAdmin(context.Background(), organizationId, adminId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.DeleteOrganizationAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**adminId** | **string** | Admin ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdminRequest struct via the builder pattern


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


## GetOrganizationAdmins

> []map[string]interface{} GetOrganizationAdmins(ctx, organizationId).Execute()

List the dashboard administrators in this organization



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
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.GetOrganizationAdmins(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GetOrganizationAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdmins`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.GetOrganizationAdmins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdminsRequest struct via the builder pattern


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


## UpdateOrganizationAdmin

> map[string]interface{} UpdateOrganizationAdmin(ctx, organizationId, adminId).UpdateOrganizationAdmin(updateOrganizationAdmin).Execute()

Update an administrator



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
    organizationId := "organizationId_example" // string | Organization ID
    adminId := "adminId_example" // string | Admin ID
    updateOrganizationAdmin := *openapiclient.NewInlineObject179() // InlineObject179 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.UpdateOrganizationAdmin(context.Background(), organizationId, adminId).UpdateOrganizationAdmin(updateOrganizationAdmin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.UpdateOrganizationAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdmin`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.UpdateOrganizationAdmin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**adminId** | **string** | Admin ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdminRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdmin** | [**InlineObject179**](InlineObject179.md) |  | 

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

