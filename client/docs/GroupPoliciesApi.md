# \GroupPoliciesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkGroupPolicy**](GroupPoliciesApi.md#CreateNetworkGroupPolicy) | **Post** /networks/{networkId}/groupPolicies | Create a group policy
[**DeleteNetworkGroupPolicy**](GroupPoliciesApi.md#DeleteNetworkGroupPolicy) | **Delete** /networks/{networkId}/groupPolicies/{groupPolicyId} | Delete a group policy
[**GetNetworkGroupPolicies**](GroupPoliciesApi.md#GetNetworkGroupPolicies) | **Get** /networks/{networkId}/groupPolicies | List the group policies in a network
[**GetNetworkGroupPolicy**](GroupPoliciesApi.md#GetNetworkGroupPolicy) | **Get** /networks/{networkId}/groupPolicies/{groupPolicyId} | Display a group policy
[**UpdateNetworkGroupPolicy**](GroupPoliciesApi.md#UpdateNetworkGroupPolicy) | **Put** /networks/{networkId}/groupPolicies/{groupPolicyId} | Update a group policy



## CreateNetworkGroupPolicy

> map[string]interface{} CreateNetworkGroupPolicy(ctx, networkId).CreateNetworkGroupPolicy(createNetworkGroupPolicy).Execute()

Create a group policy



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
    networkId := "networkId_example" // string | 
    createNetworkGroupPolicy := *openapiclient.NewCreateNetworkGroupPolicyRequest("Name_example") // CreateNetworkGroupPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPoliciesApi.CreateNetworkGroupPolicy(context.Background(), networkId).CreateNetworkGroupPolicy(createNetworkGroupPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesApi.CreateNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkGroupPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesApi.CreateNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkGroupPolicy** | [**CreateNetworkGroupPolicyRequest**](CreateNetworkGroupPolicyRequest.md) |  | 

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


## DeleteNetworkGroupPolicy

> DeleteNetworkGroupPolicy(ctx, networkId, groupPolicyId).Execute()

Delete a group policy



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
    networkId := "networkId_example" // string | 
    groupPolicyId := "groupPolicyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPoliciesApi.DeleteNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesApi.DeleteNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**groupPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkGroupPolicyRequest struct via the builder pattern


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


## GetNetworkGroupPolicies

> []map[string]interface{} GetNetworkGroupPolicies(ctx, networkId).Execute()

List the group policies in a network



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
    networkId := "networkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPoliciesApi.GetNetworkGroupPolicies(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesApi.GetNetworkGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkGroupPolicies`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesApi.GetNetworkGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupPoliciesRequest struct via the builder pattern


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


## GetNetworkGroupPolicy

> map[string]interface{} GetNetworkGroupPolicy(ctx, networkId, groupPolicyId).Execute()

Display a group policy



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
    networkId := "networkId_example" // string | 
    groupPolicyId := "groupPolicyId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPoliciesApi.GetNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesApi.GetNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkGroupPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesApi.GetNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**groupPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkGroupPolicy

> map[string]interface{} UpdateNetworkGroupPolicy(ctx, networkId, groupPolicyId).UpdateNetworkGroupPolicy(updateNetworkGroupPolicy).Execute()

Update a group policy



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
    networkId := "networkId_example" // string | 
    groupPolicyId := "groupPolicyId_example" // string | 
    updateNetworkGroupPolicy := *openapiclient.NewUpdateNetworkGroupPolicyRequest() // UpdateNetworkGroupPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupPoliciesApi.UpdateNetworkGroupPolicy(context.Background(), networkId, groupPolicyId).UpdateNetworkGroupPolicy(updateNetworkGroupPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupPoliciesApi.UpdateNetworkGroupPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkGroupPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupPoliciesApi.UpdateNetworkGroupPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**groupPolicyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkGroupPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkGroupPolicy** | [**UpdateNetworkGroupPolicyRequest**](UpdateNetworkGroupPolicyRequest.md) |  | 

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

