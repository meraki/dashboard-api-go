# \MerakiAuthUsersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkMerakiAuthUser**](MerakiAuthUsersApi.md#CreateNetworkMerakiAuthUser) | **Post** /networks/{networkId}/merakiAuthUsers | Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)
[**DeleteNetworkMerakiAuthUser**](MerakiAuthUsersApi.md#DeleteNetworkMerakiAuthUser) | **Delete** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Deauthorize a user
[**GetNetworkMerakiAuthUser**](MerakiAuthUsersApi.md#GetNetworkMerakiAuthUser) | **Get** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Return the Meraki Auth splash guest, RADIUS, or client VPN user
[**GetNetworkMerakiAuthUsers**](MerakiAuthUsersApi.md#GetNetworkMerakiAuthUsers) | **Get** /networks/{networkId}/merakiAuthUsers | List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network)
[**UpdateNetworkMerakiAuthUser**](MerakiAuthUsersApi.md#UpdateNetworkMerakiAuthUser) | **Put** /networks/{networkId}/merakiAuthUsers/{merakiAuthUserId} | Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)



## CreateNetworkMerakiAuthUser

> GetNetworkMerakiAuthUsers200ResponseInner CreateNetworkMerakiAuthUser(ctx, networkId).CreateNetworkMerakiAuthUserRequest(createNetworkMerakiAuthUserRequest).Execute()

Authorize a user configured with Meraki Authentication for a network (currently supports 802.1X, splash guest, and client VPN users, and currently, organizations have a 50,000 user cap)



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
    networkId := "networkId_example" // string | Network ID
    createNetworkMerakiAuthUserRequest := *openapiclient.NewCreateNetworkMerakiAuthUserRequest("Email_example", []openapiclient.CreateNetworkMerakiAuthUserRequestAuthorizationsInner{*openapiclient.NewCreateNetworkMerakiAuthUserRequestAuthorizationsInner()}) // CreateNetworkMerakiAuthUserRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerakiAuthUsersApi.CreateNetworkMerakiAuthUser(context.Background(), networkId).CreateNetworkMerakiAuthUserRequest(createNetworkMerakiAuthUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerakiAuthUsersApi.CreateNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkMerakiAuthUser`: GetNetworkMerakiAuthUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MerakiAuthUsersApi.CreateNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkMerakiAuthUserRequest** | [**CreateNetworkMerakiAuthUserRequest**](CreateNetworkMerakiAuthUserRequest.md) |  | 

### Return type

[**GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkMerakiAuthUser

> DeleteNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).Execute()

Deauthorize a user



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
    networkId := "networkId_example" // string | Network ID
    merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MerakiAuthUsersApi.DeleteNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerakiAuthUsersApi.DeleteNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkMerakiAuthUserRequest struct via the builder pattern


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


## GetNetworkMerakiAuthUser

> GetNetworkMerakiAuthUsers200ResponseInner GetNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).Execute()

Return the Meraki Auth splash guest, RADIUS, or client VPN user



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
    networkId := "networkId_example" // string | Network ID
    merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerakiAuthUsersApi.GetNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerakiAuthUsersApi.GetNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMerakiAuthUser`: GetNetworkMerakiAuthUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MerakiAuthUsersApi.GetNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkMerakiAuthUsers

> []GetNetworkMerakiAuthUsers200ResponseInner GetNetworkMerakiAuthUsers(ctx, networkId).Execute()

List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a wireless network, or client VPN users for a wired network)



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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerakiAuthUsersApi.GetNetworkMerakiAuthUsers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerakiAuthUsersApi.GetNetworkMerakiAuthUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkMerakiAuthUsers`: []GetNetworkMerakiAuthUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MerakiAuthUsersApi.GetNetworkMerakiAuthUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkMerakiAuthUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkMerakiAuthUser

> GetNetworkMerakiAuthUsers200ResponseInner UpdateNetworkMerakiAuthUser(ctx, networkId, merakiAuthUserId).UpdateNetworkMerakiAuthUserRequest(updateNetworkMerakiAuthUserRequest).Execute()

Update a user configured with Meraki Authentication (currently, 802.1X RADIUS, splash guest, and client VPN users can be updated)



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
    networkId := "networkId_example" // string | Network ID
    merakiAuthUserId := "merakiAuthUserId_example" // string | Meraki auth user ID
    updateNetworkMerakiAuthUserRequest := *openapiclient.NewUpdateNetworkMerakiAuthUserRequest() // UpdateNetworkMerakiAuthUserRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MerakiAuthUsersApi.UpdateNetworkMerakiAuthUser(context.Background(), networkId, merakiAuthUserId).UpdateNetworkMerakiAuthUserRequest(updateNetworkMerakiAuthUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MerakiAuthUsersApi.UpdateNetworkMerakiAuthUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkMerakiAuthUser`: GetNetworkMerakiAuthUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MerakiAuthUsersApi.UpdateNetworkMerakiAuthUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**merakiAuthUserId** | **string** | Meraki auth user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkMerakiAuthUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkMerakiAuthUserRequest** | [**UpdateNetworkMerakiAuthUserRequest**](UpdateNetworkMerakiAuthUserRequest.md) |  | 

### Return type

[**GetNetworkMerakiAuthUsers200ResponseInner**](GetNetworkMerakiAuthUsers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

