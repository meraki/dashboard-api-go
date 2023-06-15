# \UsersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrganizationUser**](UsersApi.md#DeleteOrganizationUser) | **Delete** /organizations/{organizationId}/users/{userId} | Delete a user and all of its authentication methods.
[**GetNetworkSmUserDeviceProfiles**](UsersApi.md#GetNetworkSmUserDeviceProfiles) | **Get** /networks/{networkId}/sm/users/{userId}/deviceProfiles | Get the profiles associated with a user
[**GetNetworkSmUserSoftwares**](UsersApi.md#GetNetworkSmUserSoftwares) | **Get** /networks/{networkId}/sm/users/{userId}/softwares | Get a list of softwares associated with a user
[**GetNetworkSmUsers**](UsersApi.md#GetNetworkSmUsers) | **Get** /networks/{networkId}/sm/users | List the owners in an SM network with various specified fields and filters



## DeleteOrganizationUser

> DeleteOrganizationUser(ctx, organizationId, userId).Execute()

Delete a user and all of its authentication methods.



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
    userId := "userId_example" // string | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.DeleteOrganizationUser(context.Background(), organizationId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.DeleteOrganizationUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationUserRequest struct via the builder pattern


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


## GetNetworkSmUserDeviceProfiles

> []InlineResponse20056 GetNetworkSmUserDeviceProfiles(ctx, networkId, userId).Execute()

Get the profiles associated with a user



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
    networkId := "networkId_example" // string | Network ID
    userId := "userId_example" // string | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetNetworkSmUserDeviceProfiles(context.Background(), networkId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetNetworkSmUserDeviceProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmUserDeviceProfiles`: []InlineResponse20056
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetNetworkSmUserDeviceProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUserDeviceProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse20056**](InlineResponse20056.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUserSoftwares

> []InlineResponse20060 GetNetworkSmUserSoftwares(ctx, networkId, userId).Execute()

Get a list of softwares associated with a user



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
    networkId := "networkId_example" // string | Network ID
    userId := "userId_example" // string | User ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetNetworkSmUserSoftwares(context.Background(), networkId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetNetworkSmUserSoftwares``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmUserSoftwares`: []InlineResponse20060
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetNetworkSmUserSoftwares`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUserSoftwaresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse20060**](InlineResponse20060.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmUsers

> []InlineResponse20065 GetNetworkSmUsers(ctx, networkId).Ids(ids).Usernames(usernames).Emails(emails).Scope(scope).Execute()

List the owners in an SM network with various specified fields and filters



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
    networkId := "networkId_example" // string | Network ID
    ids := []string{"Inner_example"} // []string | Filter users by id(s). (optional)
    usernames := []string{"Inner_example"} // []string | Filter users by username(s). (optional)
    emails := []string{"Inner_example"} // []string | Filter users by email(s). (optional)
    scope := []string{"Inner_example"} // []string | Specifiy a scope (one of all, none, withAny, withAll, withoutAny, withoutAll) and a set of tags. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetNetworkSmUsers(context.Background(), networkId).Ids(ids).Usernames(usernames).Emails(emails).Scope(scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetNetworkSmUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmUsers`: []InlineResponse20065
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetNetworkSmUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **[]string** | Filter users by id(s). | 
 **usernames** | **[]string** | Filter users by username(s). | 
 **emails** | **[]string** | Filter users by email(s). | 
 **scope** | **[]string** | Specifiy a scope (one of all, none, withAny, withAll, withoutAny, withoutAll) and a set of tags. | 

### Return type

[**[]InlineResponse20065**](InlineResponse20065.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

