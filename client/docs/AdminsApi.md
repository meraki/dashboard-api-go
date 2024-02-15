# \AdminsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationAdmin**](AdminsApi.md#CreateOrganizationAdmin) | **Post** /organizations/{organizationId}/admins | Create a new dashboard administrator
[**CreateOrganizationSmAdminsRole**](AdminsApi.md#CreateOrganizationSmAdminsRole) | **Post** /organizations/{organizationId}/sm/admins/roles | Create a Limited Access Role
[**DeleteOrganizationAdmin**](AdminsApi.md#DeleteOrganizationAdmin) | **Delete** /organizations/{organizationId}/admins/{adminId} | Revoke all access for a dashboard administrator within this organization
[**DeleteOrganizationSmAdminsRole**](AdminsApi.md#DeleteOrganizationSmAdminsRole) | **Delete** /organizations/{organizationId}/sm/admins/roles/{roleId} | Delete a Limited Access Role
[**GetOrganizationAdmins**](AdminsApi.md#GetOrganizationAdmins) | **Get** /organizations/{organizationId}/admins | List the dashboard administrators in this organization
[**GetOrganizationSmAdminsRole**](AdminsApi.md#GetOrganizationSmAdminsRole) | **Get** /organizations/{organizationId}/sm/admins/roles/{roleId} | Return a Limited Access Role
[**GetOrganizationSmAdminsRoles**](AdminsApi.md#GetOrganizationSmAdminsRoles) | **Get** /organizations/{organizationId}/sm/admins/roles | List the Limited Access Roles for an organization
[**UpdateOrganizationAdmin**](AdminsApi.md#UpdateOrganizationAdmin) | **Put** /organizations/{organizationId}/admins/{adminId} | Update an administrator
[**UpdateOrganizationSmAdminsRole**](AdminsApi.md#UpdateOrganizationSmAdminsRole) | **Put** /organizations/{organizationId}/sm/admins/roles/{roleId} | Update a Limited Access Role



## CreateOrganizationAdmin

> GetOrganizationAdmins200ResponseInner CreateOrganizationAdmin(ctx, organizationId).CreateOrganizationAdminRequest(createOrganizationAdminRequest).Execute()

Create a new dashboard administrator



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
    createOrganizationAdminRequest := *openapiclient.NewCreateOrganizationAdminRequest("Email_example", "Name_example", "OrgAccess_example") // CreateOrganizationAdminRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.CreateOrganizationAdmin(context.Background(), organizationId).CreateOrganizationAdminRequest(createOrganizationAdminRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.CreateOrganizationAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdmin`: GetOrganizationAdmins200ResponseInner
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

 **createOrganizationAdminRequest** | [**CreateOrganizationAdminRequest**](CreateOrganizationAdminRequest.md) |  | 

### Return type

[**GetOrganizationAdmins200ResponseInner**](GetOrganizationAdmins200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganizationSmAdminsRole

> GetOrganizationSmAdminsRoles200ResponseItemsInner CreateOrganizationSmAdminsRole(ctx, organizationId).CreateOrganizationSmAdminsRoleRequest(createOrganizationSmAdminsRoleRequest).Execute()

Create a Limited Access Role



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
    createOrganizationSmAdminsRoleRequest := *openapiclient.NewCreateOrganizationSmAdminsRoleRequest("Name_example") // CreateOrganizationSmAdminsRoleRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.CreateOrganizationSmAdminsRole(context.Background(), organizationId).CreateOrganizationSmAdminsRoleRequest(createOrganizationSmAdminsRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.CreateOrganizationSmAdminsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationSmAdminsRole`: GetOrganizationSmAdminsRoles200ResponseItemsInner
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.CreateOrganizationSmAdminsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSmAdminsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSmAdminsRoleRequest** | [**CreateOrganizationSmAdminsRoleRequest**](CreateOrganizationSmAdminsRoleRequest.md) |  | 

### Return type

[**GetOrganizationSmAdminsRoles200ResponseItemsInner**](GetOrganizationSmAdminsRoles200ResponseItemsInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

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
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    adminId := "adminId_example" // string | Admin ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminsApi.DeleteOrganizationAdmin(context.Background(), organizationId, adminId).Execute()
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

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationSmAdminsRole

> DeleteOrganizationSmAdminsRole(ctx, organizationId, roleId).Execute()

Delete a Limited Access Role



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
    r, err := apiClient.AdminsApi.DeleteOrganizationSmAdminsRole(context.Background(), organizationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.DeleteOrganizationSmAdminsRole``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOrganizationSmAdminsRoleRequest struct via the builder pattern


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


## GetOrganizationAdmins

> []GetOrganizationAdmins200ResponseInner GetOrganizationAdmins(ctx, organizationId).Execute()

List the dashboard administrators in this organization



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
    resp, r, err := apiClient.AdminsApi.GetOrganizationAdmins(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GetOrganizationAdmins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdmins`: []GetOrganizationAdmins200ResponseInner
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

[**[]GetOrganizationAdmins200ResponseInner**](GetOrganizationAdmins200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmAdminsRole

> GetOrganizationSmAdminsRoles200ResponseItemsInner GetOrganizationSmAdminsRole(ctx, organizationId, roleId).Execute()

Return a Limited Access Role



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
    resp, r, err := apiClient.AdminsApi.GetOrganizationSmAdminsRole(context.Background(), organizationId, roleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GetOrganizationSmAdminsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSmAdminsRole`: GetOrganizationSmAdminsRoles200ResponseItemsInner
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.GetOrganizationSmAdminsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmAdminsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationSmAdminsRoles200ResponseItemsInner**](GetOrganizationSmAdminsRoles200ResponseItemsInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSmAdminsRoles

> GetOrganizationSmAdminsRoles200Response GetOrganizationSmAdminsRoles(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List the Limited Access Roles for an organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.GetOrganizationSmAdminsRoles(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.GetOrganizationSmAdminsRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSmAdminsRoles`: GetOrganizationSmAdminsRoles200Response
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.GetOrganizationSmAdminsRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmAdminsRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**GetOrganizationSmAdminsRoles200Response**](GetOrganizationSmAdminsRoles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationAdmin

> GetOrganizationAdmins200ResponseInner UpdateOrganizationAdmin(ctx, organizationId, adminId).UpdateOrganizationAdminRequest(updateOrganizationAdminRequest).Execute()

Update an administrator



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
    adminId := "adminId_example" // string | Admin ID
    updateOrganizationAdminRequest := *openapiclient.NewUpdateOrganizationAdminRequest() // UpdateOrganizationAdminRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.UpdateOrganizationAdmin(context.Background(), organizationId, adminId).UpdateOrganizationAdminRequest(updateOrganizationAdminRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.UpdateOrganizationAdmin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdmin`: GetOrganizationAdmins200ResponseInner
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


 **updateOrganizationAdminRequest** | [**UpdateOrganizationAdminRequest**](UpdateOrganizationAdminRequest.md) |  | 

### Return type

[**GetOrganizationAdmins200ResponseInner**](GetOrganizationAdmins200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSmAdminsRole

> GetOrganizationSmAdminsRoles200ResponseItemsInner UpdateOrganizationSmAdminsRole(ctx, organizationId, roleId).UpdateOrganizationSmAdminsRoleRequest(updateOrganizationSmAdminsRoleRequest).Execute()

Update a Limited Access Role



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
    updateOrganizationSmAdminsRoleRequest := *openapiclient.NewUpdateOrganizationSmAdminsRoleRequest() // UpdateOrganizationSmAdminsRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminsApi.UpdateOrganizationSmAdminsRole(context.Background(), organizationId, roleId).UpdateOrganizationSmAdminsRoleRequest(updateOrganizationSmAdminsRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminsApi.UpdateOrganizationSmAdminsRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSmAdminsRole`: GetOrganizationSmAdminsRoles200ResponseItemsInner
    fmt.Fprintf(os.Stdout, "Response from `AdminsApi.UpdateOrganizationSmAdminsRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**roleId** | **string** | Role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSmAdminsRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationSmAdminsRoleRequest** | [**UpdateOrganizationSmAdminsRoleRequest**](UpdateOrganizationSmAdminsRoleRequest.md) |  | 

### Return type

[**GetOrganizationSmAdminsRoles200ResponseItemsInner**](GetOrganizationSmAdminsRoles200ResponseItemsInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

