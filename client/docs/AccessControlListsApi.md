# \AccessControlListsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchAccessControlLists**](AccessControlListsApi.md#GetNetworkSwitchAccessControlLists) | **Get** /networks/{networkId}/switch/accessControlLists | Return the access control lists for a MS network
[**UpdateNetworkSwitchAccessControlLists**](AccessControlListsApi.md#UpdateNetworkSwitchAccessControlLists) | **Put** /networks/{networkId}/switch/accessControlLists | Update the access control lists for a MS network



## GetNetworkSwitchAccessControlLists

> InlineResponse20067 GetNetworkSwitchAccessControlLists(ctx, networkId).Execute()

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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessControlListsApi.GetNetworkSwitchAccessControlLists(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessControlListsApi.GetNetworkSwitchAccessControlLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchAccessControlLists`: InlineResponse20067
    fmt.Fprintf(os.Stdout, "Response from `AccessControlListsApi.GetNetworkSwitchAccessControlLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchAccessControlListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20067**](InlineResponse20067.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchAccessControlLists

> InlineResponse20067 UpdateNetworkSwitchAccessControlLists(ctx, networkId).UpdateNetworkSwitchAccessControlLists(updateNetworkSwitchAccessControlLists).Execute()

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
    networkId := "networkId_example" // string | Network ID
    updateNetworkSwitchAccessControlLists := *openapiclient.NewInlineObject112([]openapiclient.NetworksNetworkIdSwitchAccessControlListsRules{*openapiclient.NewNetworksNetworkIdSwitchAccessControlListsRules("Policy_example", "Protocol_example", "SrcCidr_example", "DstCidr_example")}) // InlineObject112 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessControlListsApi.UpdateNetworkSwitchAccessControlLists(context.Background(), networkId).UpdateNetworkSwitchAccessControlLists(updateNetworkSwitchAccessControlLists).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessControlListsApi.UpdateNetworkSwitchAccessControlLists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchAccessControlLists`: InlineResponse20067
    fmt.Fprintf(os.Stdout, "Response from `AccessControlListsApi.UpdateNetworkSwitchAccessControlLists`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchAccessControlListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchAccessControlLists** | [**InlineObject112**](InlineObject112.md) |  | 

### Return type

[**InlineResponse20067**](InlineResponse20067.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

