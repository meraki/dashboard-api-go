# \InboundCellularFirewallRulesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceFirewallInboundCellularFirewallRules**](InboundCellularFirewallRulesApi.md#GetNetworkApplianceFirewallInboundCellularFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/inboundCellularFirewallRules | Return the inbound cellular firewall rules for an MX network
[**UpdateNetworkApplianceFirewallInboundCellularFirewallRules**](InboundCellularFirewallRulesApi.md#UpdateNetworkApplianceFirewallInboundCellularFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/inboundCellularFirewallRules | Update the inbound cellular firewall rules of an MX network



## GetNetworkApplianceFirewallInboundCellularFirewallRules

> []map[string]interface{} GetNetworkApplianceFirewallInboundCellularFirewallRules(ctx, networkId).Execute()

Return the inbound cellular firewall rules for an MX network



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
    resp, r, err := apiClient.InboundCellularFirewallRulesApi.GetNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboundCellularFirewallRulesApi.GetNetworkApplianceFirewallInboundCellularFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallInboundCellularFirewallRules`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InboundCellularFirewallRulesApi.GetNetworkApplianceFirewallInboundCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallInboundCellularFirewallRulesRequest struct via the builder pattern


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


## UpdateNetworkApplianceFirewallInboundCellularFirewallRules

> []map[string]interface{} UpdateNetworkApplianceFirewallInboundCellularFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallInboundCellularFirewallRules(updateNetworkApplianceFirewallInboundCellularFirewallRules).Execute()

Update the inbound cellular firewall rules of an MX network



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
    updateNetworkApplianceFirewallInboundCellularFirewallRules := *openapiclient.NewInlineObject32() // InlineObject32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InboundCellularFirewallRulesApi.UpdateNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallInboundCellularFirewallRules(updateNetworkApplianceFirewallInboundCellularFirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboundCellularFirewallRulesApi.UpdateNetworkApplianceFirewallInboundCellularFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallInboundCellularFirewallRules`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InboundCellularFirewallRulesApi.UpdateNetworkApplianceFirewallInboundCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallInboundCellularFirewallRules** | [**InlineObject32**](InlineObject32.md) |  | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

