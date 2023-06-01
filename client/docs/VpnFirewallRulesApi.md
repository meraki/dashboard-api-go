# \VpnFirewallRulesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApplianceVpnVpnFirewallRules**](VpnFirewallRulesApi.md#GetOrganizationApplianceVpnVpnFirewallRules) | **Get** /organizations/{organizationId}/appliance/vpn/vpnFirewallRules | Return the firewall rules for an organization&#39;s site-to-site VPN
[**UpdateOrganizationApplianceVpnVpnFirewallRules**](VpnFirewallRulesApi.md#UpdateOrganizationApplianceVpnVpnFirewallRules) | **Put** /organizations/{organizationId}/appliance/vpn/vpnFirewallRules | Update the firewall rules of an organization&#39;s site-to-site VPN



## GetOrganizationApplianceVpnVpnFirewallRules

> map[string]interface{} GetOrganizationApplianceVpnVpnFirewallRules(ctx, organizationId).Execute()

Return the firewall rules for an organization's site-to-site VPN



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VpnFirewallRulesApi.GetOrganizationApplianceVpnVpnFirewallRules(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VpnFirewallRulesApi.GetOrganizationApplianceVpnVpnFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceVpnVpnFirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `VpnFirewallRulesApi.GetOrganizationApplianceVpnVpnFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnVpnFirewallRulesRequest struct via the builder pattern


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


## UpdateOrganizationApplianceVpnVpnFirewallRules

> InlineResponse20095 UpdateOrganizationApplianceVpnVpnFirewallRules(ctx, organizationId).UpdateOrganizationApplianceVpnVpnFirewallRules(updateOrganizationApplianceVpnVpnFirewallRules).Execute()

Update the firewall rules of an organization's site-to-site VPN



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
    updateOrganizationApplianceVpnVpnFirewallRules := *openapiclient.NewInlineObject184() // InlineObject184 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VpnFirewallRulesApi.UpdateOrganizationApplianceVpnVpnFirewallRules(context.Background(), organizationId).UpdateOrganizationApplianceVpnVpnFirewallRules(updateOrganizationApplianceVpnVpnFirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VpnFirewallRulesApi.UpdateOrganizationApplianceVpnVpnFirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationApplianceVpnVpnFirewallRules`: InlineResponse20095
    fmt.Fprintf(os.Stdout, "Response from `VpnFirewallRulesApi.UpdateOrganizationApplianceVpnVpnFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnVpnFirewallRules** | [**InlineObject184**](InlineObject184.md) |  | 

### Return type

[**InlineResponse20095**](InlineResponse20095.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

