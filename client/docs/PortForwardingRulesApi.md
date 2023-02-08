# \PortForwardingRulesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCellularGatewayPortForwardingRules**](PortForwardingRulesApi.md#GetDeviceCellularGatewayPortForwardingRules) | **Get** /devices/{serial}/cellularGateway/portForwardingRules | Returns the port forwarding rules for a single MG.
[**GetNetworkApplianceFirewallPortForwardingRules**](PortForwardingRulesApi.md#GetNetworkApplianceFirewallPortForwardingRules) | **Get** /networks/{networkId}/appliance/firewall/portForwardingRules | Return the port forwarding rules for an MX network
[**UpdateDeviceCellularGatewayPortForwardingRules**](PortForwardingRulesApi.md#UpdateDeviceCellularGatewayPortForwardingRules) | **Put** /devices/{serial}/cellularGateway/portForwardingRules | Updates the port forwarding rules for a single MG.
[**UpdateNetworkApplianceFirewallPortForwardingRules**](PortForwardingRulesApi.md#UpdateNetworkApplianceFirewallPortForwardingRules) | **Put** /networks/{networkId}/appliance/firewall/portForwardingRules | Update the port forwarding rules for an MX network



## GetDeviceCellularGatewayPortForwardingRules

> map[string]interface{} GetDeviceCellularGatewayPortForwardingRules(ctx, serial).Execute()

Returns the port forwarding rules for a single MG.



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
    serial := "serial_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortForwardingRulesApi.GetDeviceCellularGatewayPortForwardingRules(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesApi.GetDeviceCellularGatewayPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCellularGatewayPortForwardingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesApi.GetDeviceCellularGatewayPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCellularGatewayPortForwardingRulesRequest struct via the builder pattern


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


## GetNetworkApplianceFirewallPortForwardingRules

> map[string]interface{} GetNetworkApplianceFirewallPortForwardingRules(ctx, networkId).Execute()

Return the port forwarding rules for an MX network



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
    resp, r, err := apiClient.PortForwardingRulesApi.GetNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesApi.GetNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceFirewallPortForwardingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesApi.GetNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


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


## UpdateDeviceCellularGatewayPortForwardingRules

> map[string]interface{} UpdateDeviceCellularGatewayPortForwardingRules(ctx, serial).UpdateDeviceCellularGatewayPortForwardingRules(updateDeviceCellularGatewayPortForwardingRules).Execute()

Updates the port forwarding rules for a single MG.



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
    serial := "serial_example" // string | 
    updateDeviceCellularGatewayPortForwardingRules := *openapiclient.NewUpdateDeviceCellularGatewayPortForwardingRulesRequest() // UpdateDeviceCellularGatewayPortForwardingRulesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortForwardingRulesApi.UpdateDeviceCellularGatewayPortForwardingRules(context.Background(), serial).UpdateDeviceCellularGatewayPortForwardingRules(updateDeviceCellularGatewayPortForwardingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesApi.UpdateDeviceCellularGatewayPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCellularGatewayPortForwardingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesApi.UpdateDeviceCellularGatewayPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCellularGatewayPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCellularGatewayPortForwardingRules** | [**UpdateDeviceCellularGatewayPortForwardingRulesRequest**](UpdateDeviceCellularGatewayPortForwardingRulesRequest.md) |  | 

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


## UpdateNetworkApplianceFirewallPortForwardingRules

> map[string]interface{} UpdateNetworkApplianceFirewallPortForwardingRules(ctx, networkId).UpdateNetworkApplianceFirewallPortForwardingRules(updateNetworkApplianceFirewallPortForwardingRules).Execute()

Update the port forwarding rules for an MX network



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
    updateNetworkApplianceFirewallPortForwardingRules := *openapiclient.NewUpdateNetworkApplianceFirewallPortForwardingRulesRequest([]openapiclient.UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner("LanIp_example", "PublicPort_example", "LocalPort_example", []string{"AllowedIps_example"}, "Protocol_example")}) // UpdateNetworkApplianceFirewallPortForwardingRulesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortForwardingRulesApi.UpdateNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).UpdateNetworkApplianceFirewallPortForwardingRules(updateNetworkApplianceFirewallPortForwardingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortForwardingRulesApi.UpdateNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceFirewallPortForwardingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PortForwardingRulesApi.UpdateNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallPortForwardingRules** | [**UpdateNetworkApplianceFirewallPortForwardingRulesRequest**](UpdateNetworkApplianceFirewallPortForwardingRulesRequest.md) |  | 

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

