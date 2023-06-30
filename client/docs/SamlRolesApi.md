# \SamlRolesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationSamlRole**](SamlRolesApi.md#CreateOrganizationSamlRole) | **Post** /organizations/{organizationId}/samlRoles | Create a SAML role
[**DeleteOrganizationSamlRole**](SamlRolesApi.md#DeleteOrganizationSamlRole) | **Delete** /organizations/{organizationId}/samlRoles/{samlRoleId} | Remove a SAML role
[**GetOrganizationSamlRole**](SamlRolesApi.md#GetOrganizationSamlRole) | **Get** /organizations/{organizationId}/samlRoles/{samlRoleId} | Return a SAML role
[**GetOrganizationSamlRoles**](SamlRolesApi.md#GetOrganizationSamlRoles) | **Get** /organizations/{organizationId}/samlRoles | List the SAML roles for this organization
[**UpdateOrganizationSamlRole**](SamlRolesApi.md#UpdateOrganizationSamlRole) | **Put** /organizations/{organizationId}/samlRoles/{samlRoleId} | Update a SAML role



## CreateOrganizationSamlRole

> map[string]interface{} CreateOrganizationSamlRole(ctx, organizationId).CreateOrganizationSamlRoleRequest(createOrganizationSamlRoleRequest).Execute()

Create a SAML role



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
    createOrganizationSamlRoleRequest := *openapiclient.NewCreateOrganizationSamlRoleRequest("Role_example", "OrgAccess_example") // CreateOrganizationSamlRoleRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SamlRolesApi.CreateOrganizationSamlRole(context.Background(), organizationId).CreateOrganizationSamlRoleRequest(createOrganizationSamlRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlRolesApi.CreateOrganizationSamlRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationSamlRole`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SamlRolesApi.CreateOrganizationSamlRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationSamlRoleRequest** | [**CreateOrganizationSamlRoleRequest**](CreateOrganizationSamlRoleRequest.md) |  | 

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


## DeleteOrganizationSamlRole

> DeleteOrganizationSamlRole(ctx, organizationId, samlRoleId).Execute()

Remove a SAML role



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
    samlRoleId := "samlRoleId_example" // string | Saml role ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SamlRolesApi.DeleteOrganizationSamlRole(context.Background(), organizationId, samlRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlRolesApi.DeleteOrganizationSamlRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**samlRoleId** | **string** | Saml role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationSamlRoleRequest struct via the builder pattern


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


## GetOrganizationSamlRole

> GetOrganizationSamlRoles200ResponseInner GetOrganizationSamlRole(ctx, organizationId, samlRoleId).Execute()

Return a SAML role



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
    samlRoleId := "samlRoleId_example" // string | Saml role ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SamlRolesApi.GetOrganizationSamlRole(context.Background(), organizationId, samlRoleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlRolesApi.GetOrganizationSamlRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSamlRole`: GetOrganizationSamlRoles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SamlRolesApi.GetOrganizationSamlRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**samlRoleId** | **string** | Saml role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetOrganizationSamlRoles200ResponseInner**](GetOrganizationSamlRoles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationSamlRoles

> []GetOrganizationSamlRoles200ResponseInner GetOrganizationSamlRoles(ctx, organizationId).Execute()

List the SAML roles for this organization



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
    resp, r, err := apiClient.SamlRolesApi.GetOrganizationSamlRoles(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlRolesApi.GetOrganizationSamlRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSamlRoles`: []GetOrganizationSamlRoles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SamlRolesApi.GetOrganizationSamlRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSamlRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationSamlRoles200ResponseInner**](GetOrganizationSamlRoles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationSamlRole

> GetOrganizationSamlRoles200ResponseInner UpdateOrganizationSamlRole(ctx, organizationId, samlRoleId).UpdateOrganizationSamlRoleRequest(updateOrganizationSamlRoleRequest).Execute()

Update a SAML role



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
    samlRoleId := "samlRoleId_example" // string | Saml role ID
    updateOrganizationSamlRoleRequest := *openapiclient.NewUpdateOrganizationSamlRoleRequest() // UpdateOrganizationSamlRoleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SamlRolesApi.UpdateOrganizationSamlRole(context.Background(), organizationId, samlRoleId).UpdateOrganizationSamlRoleRequest(updateOrganizationSamlRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlRolesApi.UpdateOrganizationSamlRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSamlRole`: GetOrganizationSamlRoles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SamlRolesApi.UpdateOrganizationSamlRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**samlRoleId** | **string** | Saml role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSamlRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationSamlRoleRequest** | [**UpdateOrganizationSamlRoleRequest**](UpdateOrganizationSamlRoleRequest.md) |  | 

### Return type

[**GetOrganizationSamlRoles200ResponseInner**](GetOrganizationSamlRoles200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

