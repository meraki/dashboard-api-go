# \RolesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationCameraRole**](RolesApi.md#CreateOrganizationCameraRole) | **Post** /organizations/{organizationId}/camera/roles | Creates new role for this organization.
[**DeleteOrganizationCameraRole**](RolesApi.md#DeleteOrganizationCameraRole) | **Delete** /organizations/{organizationId}/camera/roles/{roleId} | Delete an existing role for this organization.
[**GetOrganizationCameraRole**](RolesApi.md#GetOrganizationCameraRole) | **Get** /organizations/{organizationId}/camera/roles/{roleId} | Retrieve a single role.
[**GetOrganizationCameraRoles**](RolesApi.md#GetOrganizationCameraRoles) | **Get** /organizations/{organizationId}/camera/roles | List all the roles in this organization
[**UpdateOrganizationCameraRole**](RolesApi.md#UpdateOrganizationCameraRole) | **Put** /organizations/{organizationId}/camera/roles/{roleId} | Update an existing role in this organization.



## CreateOrganizationCameraRole

> map[string]interface{} CreateOrganizationCameraRole(ctx, organizationId).CreateOrganizationCameraRoleRequest(createOrganizationCameraRoleRequest).Execute()

Creates new role for this organization.



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
    createOrganizationCameraRoleRequest := *openapiclient.NewCreateOrganizationCameraRoleRequest("Name_example") // CreateOrganizationCameraRoleRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.CreateOrganizationCameraRole(context.Background(), organizationId).CreateOrganizationCameraRoleRequest(createOrganizationCameraRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.CreateOrganizationCameraRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCameraRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.CreateOrganizationCameraRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCameraRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCameraRoleRequest** | [**CreateOrganizationCameraRoleRequest**](CreateOrganizationCameraRoleRequest.md) |  | 

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


## DeleteOrganizationCameraRole

> DeleteOrganizationCameraRole(ctx, organizationId, roleId).Execute()

Delete an existing role for this organization.



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
    roleId := "roleId_example" // string | Role ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RolesApi.DeleteOrganizationCameraRole(context.Background(), organizationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.DeleteOrganizationCameraRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationCameraRoleRequest struct via the builder pattern


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


## GetOrganizationCameraRole

> map[string]interface{} GetOrganizationCameraRole(ctx, organizationId, roleId).Execute()

Retrieve a single role.



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
    roleId := "roleId_example" // string | Role ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetOrganizationCameraRole(context.Background(), organizationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetOrganizationCameraRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetOrganizationCameraRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationCameraRoles

> []map[string]interface{} GetOrganizationCameraRoles(ctx, organizationId).Execute()

List all the roles in this organization



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

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.GetOrganizationCameraRoles(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.GetOrganizationCameraRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraRoles`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.GetOrganizationCameraRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateOrganizationCameraRole

> map[string]interface{} UpdateOrganizationCameraRole(ctx, organizationId, roleId).UpdateOrganizationCameraRoleRequest(updateOrganizationCameraRoleRequest).Execute()

Update an existing role in this organization.



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
    roleId := "roleId_example" // string | Role ID
    updateOrganizationCameraRoleRequest := *openapiclient.NewUpdateOrganizationCameraRoleRequest() // UpdateOrganizationCameraRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RolesApi.UpdateOrganizationCameraRole(context.Background(), organizationId, roleId).UpdateOrganizationCameraRoleRequest(updateOrganizationCameraRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.UpdateOrganizationCameraRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCameraRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.UpdateOrganizationCameraRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCameraRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationCameraRoleRequest** | [**UpdateOrganizationCameraRoleRequest**](UpdateOrganizationCameraRoleRequest.md) |  | 

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

