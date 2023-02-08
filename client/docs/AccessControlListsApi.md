# \AccessControlListsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchAccessControlLists**](AccessControlListsApi.md#GetNetworkSwitchAccessControlLists) | **Get** /networks/{networkId}/switch/accessControlLists | Return the access control lists for a MS network
[**UpdateNetworkSwitchAccessControlLists**](AccessControlListsApi.md#UpdateNetworkSwitchAccessControlLists) | **Put** /networks/{networkId}/switch/accessControlLists | Update the access control lists for a MS network



## GetNetworkSwitchAccessControlLists

> map[string]interface{} GetNetworkSwitchAccessControlLists(ctx, networkId).Execute()

Return the access control lists for a MS network



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
    resp, r, err := apiClient.AccessControlListsApi.GetNetworkSwitchAccessControlLists(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessControlListsApi.GetNetworkSwitchAccessControlLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchAccessControlLists`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccessControlListsApi.GetNetworkSwitchAccessControlLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAccessControlListsRequest struct via the builder pattern


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


## UpdateNetworkSwitchAccessControlLists

> map[string]interface{} UpdateNetworkSwitchAccessControlLists(ctx, networkId).UpdateNetworkSwitchAccessControlLists(updateNetworkSwitchAccessControlLists).Execute()

Update the access control lists for a MS network



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
    updateNetworkSwitchAccessControlLists := *openapiclient.NewUpdateNetworkSwitchAccessControlListsRequest([]openapiclient.UpdateNetworkSwitchAccessControlListsRequestRulesInner{*openapiclient.NewUpdateNetworkSwitchAccessControlListsRequestRulesInner("Policy_example", "Protocol_example", "SrcCidr_example", "DstCidr_example")}) // UpdateNetworkSwitchAccessControlListsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessControlListsApi.UpdateNetworkSwitchAccessControlLists(context.Background(), networkId).UpdateNetworkSwitchAccessControlLists(updateNetworkSwitchAccessControlLists).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessControlListsApi.UpdateNetworkSwitchAccessControlLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchAccessControlLists`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AccessControlListsApi.UpdateNetworkSwitchAccessControlLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchAccessControlListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchAccessControlLists** | [**UpdateNetworkSwitchAccessControlListsRequest**](UpdateNetworkSwitchAccessControlListsRequest.md) |  | 

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

