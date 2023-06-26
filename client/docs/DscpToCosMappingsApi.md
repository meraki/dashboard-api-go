# \DscpToCosMappingsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSwitchDscpToCosMappings**](DscpToCosMappingsApi.md#GetNetworkSwitchDscpToCosMappings) | **Get** /networks/{networkId}/switch/dscpToCosMappings | Return the DSCP to CoS mappings
[**UpdateNetworkSwitchDscpToCosMappings**](DscpToCosMappingsApi.md#UpdateNetworkSwitchDscpToCosMappings) | **Put** /networks/{networkId}/switch/dscpToCosMappings | Update the DSCP to CoS mappings



## GetNetworkSwitchDscpToCosMappings

> InlineResponse20074 GetNetworkSwitchDscpToCosMappings(ctx, networkId).Execute()

Return the DSCP to CoS mappings



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
    resp, r, err := apiClient.DscpToCosMappingsApi.GetNetworkSwitchDscpToCosMappings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DscpToCosMappingsApi.GetNetworkSwitchDscpToCosMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchDscpToCosMappings`: InlineResponse20074
    fmt.Fprintf(os.Stdout, "Response from `DscpToCosMappingsApi.GetNetworkSwitchDscpToCosMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDscpToCosMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20074**](InlineResponse20074.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchDscpToCosMappings

> map[string]interface{} UpdateNetworkSwitchDscpToCosMappings(ctx, networkId).UpdateNetworkSwitchDscpToCosMappings(updateNetworkSwitchDscpToCosMappings).Execute()

Update the DSCP to CoS mappings



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
    updateNetworkSwitchDscpToCosMappings := *openapiclient.NewInlineObject119([]openapiclient.NetworksNetworkIdSwitchDscpToCosMappingsMappings{*openapiclient.NewNetworksNetworkIdSwitchDscpToCosMappingsMappings(int32(123), int32(123))}) // InlineObject119 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DscpToCosMappingsApi.UpdateNetworkSwitchDscpToCosMappings(context.Background(), networkId).UpdateNetworkSwitchDscpToCosMappings(updateNetworkSwitchDscpToCosMappings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DscpToCosMappingsApi.UpdateNetworkSwitchDscpToCosMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchDscpToCosMappings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DscpToCosMappingsApi.UpdateNetworkSwitchDscpToCosMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchDscpToCosMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchDscpToCosMappings** | [**InlineObject119**](InlineObject119.md) |  | 

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

