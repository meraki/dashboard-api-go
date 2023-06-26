# \AccessPoliciesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchAccessPolicy**](AccessPoliciesApi.md#CreateNetworkSwitchAccessPolicy) | **Post** /networks/{networkId}/switch/accessPolicies | Create an access policy for a switch network
[**DeleteNetworkSwitchAccessPolicy**](AccessPoliciesApi.md#DeleteNetworkSwitchAccessPolicy) | **Delete** /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber} | Delete an access policy for a switch network
[**GetNetworkSwitchAccessPolicies**](AccessPoliciesApi.md#GetNetworkSwitchAccessPolicies) | **Get** /networks/{networkId}/switch/accessPolicies | List the access policies for a switch network
[**GetNetworkSwitchAccessPolicy**](AccessPoliciesApi.md#GetNetworkSwitchAccessPolicy) | **Get** /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber} | Return a specific access policy for a switch network
[**UpdateNetworkSwitchAccessPolicy**](AccessPoliciesApi.md#UpdateNetworkSwitchAccessPolicy) | **Put** /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber} | Update an access policy for a switch network



## CreateNetworkSwitchAccessPolicy

> InlineResponse20068 CreateNetworkSwitchAccessPolicy(ctx, networkId).CreateNetworkSwitchAccessPolicy(createNetworkSwitchAccessPolicy).Execute()

Create an access policy for a switch network



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
    createNetworkSwitchAccessPolicy := *openapiclient.NewInlineObject113("Name_example", []openapiclient.NetworksNetworkIdSwitchAccessPoliciesRadiusServers1{*openapiclient.NewNetworksNetworkIdSwitchAccessPoliciesRadiusServers1("Host_example", int32(123), "Secret_example")}, false, false, false, "HostMode_example", false) // InlineObject113 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesApi.CreateNetworkSwitchAccessPolicy(context.Background(), networkId).CreateNetworkSwitchAccessPolicy(createNetworkSwitchAccessPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesApi.CreateNetworkSwitchAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchAccessPolicy`: InlineResponse20068
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesApi.CreateNetworkSwitchAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchAccessPolicy** | [**InlineObject113**](InlineObject113.md) |  | 

### Return type

[**InlineResponse20068**](InlineResponse20068.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchAccessPolicy

> DeleteNetworkSwitchAccessPolicy(ctx, networkId, accessPolicyNumber).Execute()

Delete an access policy for a switch network



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
    accessPolicyNumber := "accessPolicyNumber_example" // string | Access policy number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesApi.DeleteNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesApi.DeleteNetworkSwitchAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**accessPolicyNumber** | **string** | Access policy number | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchAccessPolicyRequest struct via the builder pattern


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


## GetNetworkSwitchAccessPolicies

> []InlineResponse20068 GetNetworkSwitchAccessPolicies(ctx, networkId).Execute()

List the access policies for a switch network



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesApi.GetNetworkSwitchAccessPolicies(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesApi.GetNetworkSwitchAccessPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchAccessPolicies`: []InlineResponse20068
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesApi.GetNetworkSwitchAccessPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20068**](InlineResponse20068.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchAccessPolicy

> InlineResponse20068 GetNetworkSwitchAccessPolicy(ctx, networkId, accessPolicyNumber).Execute()

Return a specific access policy for a switch network



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
    accessPolicyNumber := "accessPolicyNumber_example" // string | Access policy number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesApi.GetNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesApi.GetNetworkSwitchAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchAccessPolicy`: InlineResponse20068
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesApi.GetNetworkSwitchAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**accessPolicyNumber** | **string** | Access policy number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20068**](InlineResponse20068.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchAccessPolicy

> InlineResponse20068 UpdateNetworkSwitchAccessPolicy(ctx, networkId, accessPolicyNumber).UpdateNetworkSwitchAccessPolicy(updateNetworkSwitchAccessPolicy).Execute()

Update an access policy for a switch network



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
    accessPolicyNumber := "accessPolicyNumber_example" // string | Access policy number
    updateNetworkSwitchAccessPolicy := *openapiclient.NewInlineObject114() // InlineObject114 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessPoliciesApi.UpdateNetworkSwitchAccessPolicy(context.Background(), networkId, accessPolicyNumber).UpdateNetworkSwitchAccessPolicy(updateNetworkSwitchAccessPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessPoliciesApi.UpdateNetworkSwitchAccessPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchAccessPolicy`: InlineResponse20068
    fmt.Fprintf(os.Stdout, "Response from `AccessPoliciesApi.UpdateNetworkSwitchAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**accessPolicyNumber** | **string** | Access policy number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchAccessPolicy** | [**InlineObject114**](InlineObject114.md) |  | 

### Return type

[**InlineResponse20068**](InlineResponse20068.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

