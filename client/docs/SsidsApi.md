# \SsidsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkWirelessSsidIdentityPsk**](SsidsApi.md#CreateNetworkWirelessSsidIdentityPsk) | **Post** /networks/{networkId}/wireless/ssids/{number}/identityPsks | Create an Identity PSK
[**DeleteNetworkWirelessSsidIdentityPsk**](SsidsApi.md#DeleteNetworkWirelessSsidIdentityPsk) | **Delete** /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} | Delete an Identity PSK
[**GetNetworkApplianceSsid**](SsidsApi.md#GetNetworkApplianceSsid) | **Get** /networks/{networkId}/appliance/ssids/{number} | Return a single MX SSID
[**GetNetworkApplianceSsids**](SsidsApi.md#GetNetworkApplianceSsids) | **Get** /networks/{networkId}/appliance/ssids | List the MX SSIDs in a network
[**GetNetworkWirelessSsid**](SsidsApi.md#GetNetworkWirelessSsid) | **Get** /networks/{networkId}/wireless/ssids/{number} | Return a single MR SSID
[**GetNetworkWirelessSsidBonjourForwarding**](SsidsApi.md#GetNetworkWirelessSsidBonjourForwarding) | **Get** /networks/{networkId}/wireless/ssids/{number}/bonjourForwarding | List the Bonjour forwarding setting and rules for the SSID
[**GetNetworkWirelessSsidDeviceTypeGroupPolicies**](SsidsApi.md#GetNetworkWirelessSsidDeviceTypeGroupPolicies) | **Get** /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies | List the device type group policies for the SSID
[**GetNetworkWirelessSsidEapOverride**](SsidsApi.md#GetNetworkWirelessSsidEapOverride) | **Get** /networks/{networkId}/wireless/ssids/{number}/eapOverride | Return the EAP overridden parameters for an SSID
[**GetNetworkWirelessSsidFirewallL3FirewallRules**](SsidsApi.md#GetNetworkWirelessSsidFirewallL3FirewallRules) | **Get** /networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules | Return the L3 firewall rules for an SSID on an MR network
[**GetNetworkWirelessSsidFirewallL7FirewallRules**](SsidsApi.md#GetNetworkWirelessSsidFirewallL7FirewallRules) | **Get** /networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules | Return the L7 firewall rules for an SSID on an MR network
[**GetNetworkWirelessSsidHotspot20**](SsidsApi.md#GetNetworkWirelessSsidHotspot20) | **Get** /networks/{networkId}/wireless/ssids/{number}/hotspot20 | Return the Hotspot 2.0 settings for an SSID
[**GetNetworkWirelessSsidIdentityPsk**](SsidsApi.md#GetNetworkWirelessSsidIdentityPsk) | **Get** /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} | Return an Identity PSK
[**GetNetworkWirelessSsidIdentityPsks**](SsidsApi.md#GetNetworkWirelessSsidIdentityPsks) | **Get** /networks/{networkId}/wireless/ssids/{number}/identityPsks | List all Identity PSKs in a wireless network
[**GetNetworkWirelessSsidSchedules**](SsidsApi.md#GetNetworkWirelessSsidSchedules) | **Get** /networks/{networkId}/wireless/ssids/{number}/schedules | List the outage schedule for the SSID
[**GetNetworkWirelessSsidSplashSettings**](SsidsApi.md#GetNetworkWirelessSsidSplashSettings) | **Get** /networks/{networkId}/wireless/ssids/{number}/splash/settings | Display the splash page settings for the given SSID
[**GetNetworkWirelessSsidTrafficShapingRules**](SsidsApi.md#GetNetworkWirelessSsidTrafficShapingRules) | **Get** /networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules | Display the traffic shaping settings for a SSID on an MR network
[**GetNetworkWirelessSsidVpn**](SsidsApi.md#GetNetworkWirelessSsidVpn) | **Get** /networks/{networkId}/wireless/ssids/{number}/vpn | List the VPN settings for the SSID.
[**GetNetworkWirelessSsids**](SsidsApi.md#GetNetworkWirelessSsids) | **Get** /networks/{networkId}/wireless/ssids | List the MR SSIDs in a network
[**GetOrganizationSummaryTopSsidsByUsage**](SsidsApi.md#GetOrganizationSummaryTopSsidsByUsage) | **Get** /organizations/{organizationId}/summary/top/ssids/byUsage | Return metrics for organization&#39;s top 10 ssids by data usage over given time range
[**UpdateNetworkApplianceSsid**](SsidsApi.md#UpdateNetworkApplianceSsid) | **Put** /networks/{networkId}/appliance/ssids/{number} | Update the attributes of an MX SSID
[**UpdateNetworkWirelessSsid**](SsidsApi.md#UpdateNetworkWirelessSsid) | **Put** /networks/{networkId}/wireless/ssids/{number} | Update the attributes of an MR SSID
[**UpdateNetworkWirelessSsidBonjourForwarding**](SsidsApi.md#UpdateNetworkWirelessSsidBonjourForwarding) | **Put** /networks/{networkId}/wireless/ssids/{number}/bonjourForwarding | Update the bonjour forwarding setting and rules for the SSID
[**UpdateNetworkWirelessSsidDeviceTypeGroupPolicies**](SsidsApi.md#UpdateNetworkWirelessSsidDeviceTypeGroupPolicies) | **Put** /networks/{networkId}/wireless/ssids/{number}/deviceTypeGroupPolicies | Update the device type group policies for the SSID
[**UpdateNetworkWirelessSsidEapOverride**](SsidsApi.md#UpdateNetworkWirelessSsidEapOverride) | **Put** /networks/{networkId}/wireless/ssids/{number}/eapOverride | Update the EAP overridden parameters for an SSID.
[**UpdateNetworkWirelessSsidFirewallL3FirewallRules**](SsidsApi.md#UpdateNetworkWirelessSsidFirewallL3FirewallRules) | **Put** /networks/{networkId}/wireless/ssids/{number}/firewall/l3FirewallRules | Update the L3 firewall rules of an SSID on an MR network
[**UpdateNetworkWirelessSsidFirewallL7FirewallRules**](SsidsApi.md#UpdateNetworkWirelessSsidFirewallL7FirewallRules) | **Put** /networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules | Update the L7 firewall rules of an SSID on an MR network
[**UpdateNetworkWirelessSsidHotspot20**](SsidsApi.md#UpdateNetworkWirelessSsidHotspot20) | **Put** /networks/{networkId}/wireless/ssids/{number}/hotspot20 | Update the Hotspot 2.0 settings of an SSID
[**UpdateNetworkWirelessSsidIdentityPsk**](SsidsApi.md#UpdateNetworkWirelessSsidIdentityPsk) | **Put** /networks/{networkId}/wireless/ssids/{number}/identityPsks/{identityPskId} | Update an Identity PSK
[**UpdateNetworkWirelessSsidSchedules**](SsidsApi.md#UpdateNetworkWirelessSsidSchedules) | **Put** /networks/{networkId}/wireless/ssids/{number}/schedules | Update the outage schedule for the SSID
[**UpdateNetworkWirelessSsidSplashSettings**](SsidsApi.md#UpdateNetworkWirelessSsidSplashSettings) | **Put** /networks/{networkId}/wireless/ssids/{number}/splash/settings | Modify the splash page settings for the given SSID
[**UpdateNetworkWirelessSsidTrafficShapingRules**](SsidsApi.md#UpdateNetworkWirelessSsidTrafficShapingRules) | **Put** /networks/{networkId}/wireless/ssids/{number}/trafficShaping/rules | Update the traffic shaping settings for an SSID on an MR network
[**UpdateNetworkWirelessSsidVpn**](SsidsApi.md#UpdateNetworkWirelessSsidVpn) | **Put** /networks/{networkId}/wireless/ssids/{number}/vpn | Update the VPN settings for the SSID



## CreateNetworkWirelessSsidIdentityPsk

> map[string]interface{} CreateNetworkWirelessSsidIdentityPsk(ctx, networkId, number).CreateNetworkWirelessSsidIdentityPsk(createNetworkWirelessSsidIdentityPsk).Execute()

Create an Identity PSK



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
    number := "number_example" // string | Number
    createNetworkWirelessSsidIdentityPsk := *openapiclient.NewInlineObject161("Name_example", "GroupPolicyId_example") // InlineObject161 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.CreateNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number).CreateNetworkWirelessSsidIdentityPsk(createNetworkWirelessSsidIdentityPsk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.CreateNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWirelessSsidIdentityPsk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.CreateNetworkWirelessSsidIdentityPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNetworkWirelessSsidIdentityPsk** | [**InlineObject161**](InlineObject161.md) |  | 

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


## DeleteNetworkWirelessSsidIdentityPsk

> DeleteNetworkWirelessSsidIdentityPsk(ctx, networkId, number, identityPskId).Execute()

Delete an Identity PSK



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
    number := "number_example" // string | Number
    identityPskId := "identityPskId_example" // string | Identity psk ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.DeleteNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.DeleteNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 
**identityPskId** | **string** | Identity psk ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


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


## GetNetworkApplianceSsid

> InlineResponse20017 GetNetworkApplianceSsid(ctx, networkId, number).Execute()

Return a single MX SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkApplianceSsid(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkApplianceSsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSsid`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkApplianceSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSsids

> []InlineResponse20017 GetNetworkApplianceSsids(ctx, networkId).Execute()

List the MX SSIDs in a network



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
    resp, r, err := apiClient.SsidsApi.GetNetworkApplianceSsids(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkApplianceSsids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceSsids`: []InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkApplianceSsids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSsidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20017**](InlineResponse20017.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsid

> map[string]interface{} GetNetworkWirelessSsid(ctx, networkId, number).Execute()

Return a single MR SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsid(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsid`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidRequest struct via the builder pattern


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


## GetNetworkWirelessSsidBonjourForwarding

> map[string]interface{} GetNetworkWirelessSsidBonjourForwarding(ctx, networkId, number).Execute()

List the Bonjour forwarding setting and rules for the SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidBonjourForwarding(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidBonjourForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidBonjourForwarding`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidBonjourForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidBonjourForwardingRequest struct via the builder pattern


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


## GetNetworkWirelessSsidDeviceTypeGroupPolicies

> map[string]interface{} GetNetworkWirelessSsidDeviceTypeGroupPolicies(ctx, networkId, number).Execute()

List the device type group policies for the SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidDeviceTypeGroupPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidDeviceTypeGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct via the builder pattern


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


## GetNetworkWirelessSsidEapOverride

> InlineResponse20085 GetNetworkWirelessSsidEapOverride(ctx, networkId, number).Execute()

Return the EAP overridden parameters for an SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidEapOverride(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidEapOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidEapOverride`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidEapOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidEapOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidFirewallL3FirewallRules

> map[string]interface{} GetNetworkWirelessSsidFirewallL3FirewallRules(ctx, networkId, number).Execute()

Return the L3 firewall rules for an SSID on an MR network



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidFirewallL3FirewallRules(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidFirewallL3FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidFirewallL3FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidFirewallL3FirewallRulesRequest struct via the builder pattern


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


## GetNetworkWirelessSsidFirewallL7FirewallRules

> map[string]interface{} GetNetworkWirelessSsidFirewallL7FirewallRules(ctx, networkId, number).Execute()

Return the L7 firewall rules for an SSID on an MR network



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidFirewallL7FirewallRules(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidFirewallL7FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidFirewallL7FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidFirewallL7FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidFirewallL7FirewallRulesRequest struct via the builder pattern


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


## GetNetworkWirelessSsidHotspot20

> map[string]interface{} GetNetworkWirelessSsidHotspot20(ctx, networkId, number).Execute()

Return the Hotspot 2.0 settings for an SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidHotspot20(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidHotspot20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidHotspot20`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidHotspot20`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidHotspot20Request struct via the builder pattern


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


## GetNetworkWirelessSsidIdentityPsk

> InlineResponse20086 GetNetworkWirelessSsidIdentityPsk(ctx, networkId, number, identityPskId).Execute()

Return an Identity PSK



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
    number := "number_example" // string | Number
    identityPskId := "identityPskId_example" // string | Identity psk ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidIdentityPsk`: InlineResponse20086
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidIdentityPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 
**identityPskId** | **string** | Identity psk ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InlineResponse20086**](InlineResponse20086.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidIdentityPsks

> []InlineResponse20086 GetNetworkWirelessSsidIdentityPsks(ctx, networkId, number).Execute()

List all Identity PSKs in a wireless network



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidIdentityPsks(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidIdentityPsks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidIdentityPsks`: []InlineResponse20086
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidIdentityPsks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidIdentityPsksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]InlineResponse20086**](InlineResponse20086.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidSchedules

> map[string]interface{} GetNetworkWirelessSsidSchedules(ctx, networkId, number).Execute()

List the outage schedule for the SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidSchedules(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidSchedulesRequest struct via the builder pattern


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


## GetNetworkWirelessSsidSplashSettings

> InlineResponse20087 GetNetworkWirelessSsidSplashSettings(ctx, networkId, number).Execute()

Display the splash page settings for the given SSID



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidSplashSettings(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidSplashSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidSplashSettings`: InlineResponse20087
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidSplashSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidSplashSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20087**](InlineResponse20087.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessSsidTrafficShapingRules

> map[string]interface{} GetNetworkWirelessSsidTrafficShapingRules(ctx, networkId, number).Execute()

Display the traffic shaping settings for a SSID on an MR network



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidTrafficShapingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidTrafficShapingRulesRequest struct via the builder pattern


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


## GetNetworkWirelessSsidVpn

> map[string]interface{} GetNetworkWirelessSsidVpn(ctx, networkId, number).Execute()

List the VPN settings for the SSID.



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
    number := "number_example" // string | Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsidVpn(context.Background(), networkId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsidVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsidVpn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsidVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidVpnRequest struct via the builder pattern


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


## GetNetworkWirelessSsids

> []map[string]interface{} GetNetworkWirelessSsids(ctx, networkId).Execute()

List the MR SSIDs in a network



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
    resp, r, err := apiClient.SsidsApi.GetNetworkWirelessSsids(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetNetworkWirelessSsids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSsids`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetNetworkWirelessSsids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSsidsRequest struct via the builder pattern


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


## GetOrganizationSummaryTopSsidsByUsage

> []InlineResponse200138 GetOrganizationSummaryTopSsidsByUsage(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Return metrics for organization's top 10 ssids by data usage over given time range



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.GetOrganizationSummaryTopSsidsByUsage(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.GetOrganizationSummaryTopSsidsByUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSummaryTopSsidsByUsage`: []InlineResponse200138
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.GetOrganizationSummaryTopSsidsByUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSummaryTopSsidsByUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

[**[]InlineResponse200138**](InlineResponse200138.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSsid

> InlineResponse20017 UpdateNetworkApplianceSsid(ctx, networkId, number).UpdateNetworkApplianceSsid(updateNetworkApplianceSsid).Execute()

Update the attributes of an MX SSID



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
    number := "number_example" // string | Number
    updateNetworkApplianceSsid := *openapiclient.NewInlineObject47() // InlineObject47 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkApplianceSsid(context.Background(), networkId, number).UpdateNetworkApplianceSsid(updateNetworkApplianceSsid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkApplianceSsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceSsid`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkApplianceSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceSsid** | [**InlineObject47**](InlineObject47.md) |  | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsid

> map[string]interface{} UpdateNetworkWirelessSsid(ctx, networkId, number).UpdateNetworkWirelessSsid(updateNetworkWirelessSsid).Execute()

Update the attributes of an MR SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsid := *openapiclient.NewInlineObject154() // InlineObject154 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsid(context.Background(), networkId, number).UpdateNetworkWirelessSsid(updateNetworkWirelessSsid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsid`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsid** | [**InlineObject154**](InlineObject154.md) |  | 

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


## UpdateNetworkWirelessSsidBonjourForwarding

> map[string]interface{} UpdateNetworkWirelessSsidBonjourForwarding(ctx, networkId, number).UpdateNetworkWirelessSsidBonjourForwarding(updateNetworkWirelessSsidBonjourForwarding).Execute()

Update the bonjour forwarding setting and rules for the SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidBonjourForwarding := *openapiclient.NewInlineObject155() // InlineObject155 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidBonjourForwarding(context.Background(), networkId, number).UpdateNetworkWirelessSsidBonjourForwarding(updateNetworkWirelessSsidBonjourForwarding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidBonjourForwarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidBonjourForwarding`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidBonjourForwarding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidBonjourForwardingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidBonjourForwarding** | [**InlineObject155**](InlineObject155.md) |  | 

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


## UpdateNetworkWirelessSsidDeviceTypeGroupPolicies

> map[string]interface{} UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(ctx, networkId, number).UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies).Execute()

Update the device type group policies for the SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidDeviceTypeGroupPolicies := *openapiclient.NewInlineObject156() // InlineObject156 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(context.Background(), networkId, number).UpdateNetworkWirelessSsidDeviceTypeGroupPolicies(updateNetworkWirelessSsidDeviceTypeGroupPolicies).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidDeviceTypeGroupPolicies`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidDeviceTypeGroupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidDeviceTypeGroupPolicies** | [**InlineObject156**](InlineObject156.md) |  | 

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


## UpdateNetworkWirelessSsidEapOverride

> InlineResponse20085 UpdateNetworkWirelessSsidEapOverride(ctx, networkId, number).UpdateNetworkWirelessSsidEapOverride(updateNetworkWirelessSsidEapOverride).Execute()

Update the EAP overridden parameters for an SSID.



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidEapOverride := *openapiclient.NewInlineObject157() // InlineObject157 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidEapOverride(context.Background(), networkId, number).UpdateNetworkWirelessSsidEapOverride(updateNetworkWirelessSsidEapOverride).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidEapOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidEapOverride`: InlineResponse20085
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidEapOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidEapOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidEapOverride** | [**InlineObject157**](InlineObject157.md) |  | 

### Return type

[**InlineResponse20085**](InlineResponse20085.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidFirewallL3FirewallRules

> map[string]interface{} UpdateNetworkWirelessSsidFirewallL3FirewallRules(ctx, networkId, number).UpdateNetworkWirelessSsidFirewallL3FirewallRules(updateNetworkWirelessSsidFirewallL3FirewallRules).Execute()

Update the L3 firewall rules of an SSID on an MR network



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidFirewallL3FirewallRules := *openapiclient.NewInlineObject158() // InlineObject158 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidFirewallL3FirewallRules(context.Background(), networkId, number).UpdateNetworkWirelessSsidFirewallL3FirewallRules(updateNetworkWirelessSsidFirewallL3FirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidFirewallL3FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidFirewallL3FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidFirewallL3FirewallRules** | [**InlineObject158**](InlineObject158.md) |  | 

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


## UpdateNetworkWirelessSsidFirewallL7FirewallRules

> map[string]interface{} UpdateNetworkWirelessSsidFirewallL7FirewallRules(ctx, networkId, number).UpdateNetworkWirelessSsidFirewallL7FirewallRules(updateNetworkWirelessSsidFirewallL7FirewallRules).Execute()

Update the L7 firewall rules of an SSID on an MR network



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidFirewallL7FirewallRules := *openapiclient.NewInlineObject159() // InlineObject159 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidFirewallL7FirewallRules(context.Background(), networkId, number).UpdateNetworkWirelessSsidFirewallL7FirewallRules(updateNetworkWirelessSsidFirewallL7FirewallRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidFirewallL7FirewallRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidFirewallL7FirewallRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidFirewallL7FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidFirewallL7FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidFirewallL7FirewallRules** | [**InlineObject159**](InlineObject159.md) |  | 

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


## UpdateNetworkWirelessSsidHotspot20

> map[string]interface{} UpdateNetworkWirelessSsidHotspot20(ctx, networkId, number).UpdateNetworkWirelessSsidHotspot20(updateNetworkWirelessSsidHotspot20).Execute()

Update the Hotspot 2.0 settings of an SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidHotspot20 := *openapiclient.NewInlineObject160() // InlineObject160 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidHotspot20(context.Background(), networkId, number).UpdateNetworkWirelessSsidHotspot20(updateNetworkWirelessSsidHotspot20).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidHotspot20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidHotspot20`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidHotspot20`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidHotspot20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidHotspot20** | [**InlineObject160**](InlineObject160.md) |  | 

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


## UpdateNetworkWirelessSsidIdentityPsk

> map[string]interface{} UpdateNetworkWirelessSsidIdentityPsk(ctx, networkId, number, identityPskId).UpdateNetworkWirelessSsidIdentityPsk(updateNetworkWirelessSsidIdentityPsk).Execute()

Update an Identity PSK



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
    number := "number_example" // string | Number
    identityPskId := "identityPskId_example" // string | Identity psk ID
    updateNetworkWirelessSsidIdentityPsk := *openapiclient.NewInlineObject162() // InlineObject162 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidIdentityPsk(context.Background(), networkId, number, identityPskId).UpdateNetworkWirelessSsidIdentityPsk(updateNetworkWirelessSsidIdentityPsk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidIdentityPsk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidIdentityPsk`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidIdentityPsk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 
**identityPskId** | **string** | Identity psk ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidIdentityPskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateNetworkWirelessSsidIdentityPsk** | [**InlineObject162**](InlineObject162.md) |  | 

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


## UpdateNetworkWirelessSsidSchedules

> map[string]interface{} UpdateNetworkWirelessSsidSchedules(ctx, networkId, number).UpdateNetworkWirelessSsidSchedules(updateNetworkWirelessSsidSchedules).Execute()

Update the outage schedule for the SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidSchedules := *openapiclient.NewInlineObject163() // InlineObject163 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidSchedules(context.Background(), networkId, number).UpdateNetworkWirelessSsidSchedules(updateNetworkWirelessSsidSchedules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidSchedules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidSchedules** | [**InlineObject163**](InlineObject163.md) |  | 

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


## UpdateNetworkWirelessSsidSplashSettings

> InlineResponse20087 UpdateNetworkWirelessSsidSplashSettings(ctx, networkId, number).UpdateNetworkWirelessSsidSplashSettings(updateNetworkWirelessSsidSplashSettings).Execute()

Modify the splash page settings for the given SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidSplashSettings := *openapiclient.NewInlineObject164() // InlineObject164 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidSplashSettings(context.Background(), networkId, number).UpdateNetworkWirelessSsidSplashSettings(updateNetworkWirelessSsidSplashSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidSplashSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidSplashSettings`: InlineResponse20087
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidSplashSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidSplashSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidSplashSettings** | [**InlineObject164**](InlineObject164.md) |  | 

### Return type

[**InlineResponse20087**](InlineResponse20087.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessSsidTrafficShapingRules

> map[string]interface{} UpdateNetworkWirelessSsidTrafficShapingRules(ctx, networkId, number).UpdateNetworkWirelessSsidTrafficShapingRules(updateNetworkWirelessSsidTrafficShapingRules).Execute()

Update the traffic shaping settings for an SSID on an MR network



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidTrafficShapingRules := *openapiclient.NewInlineObject165() // InlineObject165 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidTrafficShapingRules(context.Background(), networkId, number).UpdateNetworkWirelessSsidTrafficShapingRules(updateNetworkWirelessSsidTrafficShapingRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidTrafficShapingRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidTrafficShapingRules`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidTrafficShapingRules** | [**InlineObject165**](InlineObject165.md) |  | 

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


## UpdateNetworkWirelessSsidVpn

> map[string]interface{} UpdateNetworkWirelessSsidVpn(ctx, networkId, number).UpdateNetworkWirelessSsidVpn(updateNetworkWirelessSsidVpn).Execute()

Update the VPN settings for the SSID



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
    number := "number_example" // string | Number
    updateNetworkWirelessSsidVpn := *openapiclient.NewInlineObject166() // InlineObject166 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SsidsApi.UpdateNetworkWirelessSsidVpn(context.Background(), networkId, number).UpdateNetworkWirelessSsidVpn(updateNetworkWirelessSsidVpn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SsidsApi.UpdateNetworkWirelessSsidVpn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessSsidVpn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SsidsApi.UpdateNetworkWirelessSsidVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessSsidVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessSsidVpn** | [**InlineObject166**](InlineObject166.md) |  | 

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

