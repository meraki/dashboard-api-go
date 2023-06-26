# \OneToOneNatRulesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceFirewallOneToOneNatRules**](OneToOneNatRulesApi.md#GetNetworkApplianceFirewallOneToOneNatRules) | **Get** /networks/{networkId}/appliance/firewall/oneToOneNatRules | Return the 1:1 NAT mapping rules for an MX network
[**UpdateNetworkApplianceFirewallOneToOneNatRules**](OneToOneNatRulesApi.md#UpdateNetworkApplianceFirewallOneToOneNatRules) | **Put** /networks/{networkId}/appliance/firewall/oneToOneNatRules | Set the 1:1 NAT mapping rules for an MX network



## GetNetworkApplianceFirewallOneToOneNatRules

> map[string]interface{} GetNetworkApplianceFirewallOneToOneNatRules(ctx, networkId).Execute()

Return the 1:1 NAT mapping rules for an MX network



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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OneToOneNatRulesApi.GetNetworkApplianceFirewallOneToOneNatRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneToOneNatRulesApi.GetNetworkApplianceFirewallOneToOneNatRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallOneToOneNatRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OneToOneNatRulesApi.GetNetworkApplianceFirewallOneToOneNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallOneToOneNatRulesRequest struct via the builder pattern


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


## UpdateNetworkApplianceFirewallOneToOneNatRules

> map[string]interface{} UpdateNetworkApplianceFirewallOneToOneNatRules(ctx, networkId).UpdateNetworkApplianceFirewallOneToOneNatRulesRequest(updateNetworkApplianceFirewallOneToOneNatRulesRequest).Execute()

Set the 1:1 NAT mapping rules for an MX network



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
    updateNetworkApplianceFirewallOneToOneNatRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequest([]openapiclient.UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner("LanIp_example")}) // UpdateNetworkApplianceFirewallOneToOneNatRulesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OneToOneNatRulesApi.UpdateNetworkApplianceFirewallOneToOneNatRules(context.Background(), networkId).UpdateNetworkApplianceFirewallOneToOneNatRulesRequest(updateNetworkApplianceFirewallOneToOneNatRulesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneToOneNatRulesApi.UpdateNetworkApplianceFirewallOneToOneNatRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallOneToOneNatRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OneToOneNatRulesApi.UpdateNetworkApplianceFirewallOneToOneNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallOneToOneNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallOneToOneNatRulesRequest** | [**UpdateNetworkApplianceFirewallOneToOneNatRulesRequest**](UpdateNetworkApplianceFirewallOneToOneNatRulesRequest.md) |  | 

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

