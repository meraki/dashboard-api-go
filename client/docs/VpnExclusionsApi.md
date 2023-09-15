# \VpnExclusionsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork**](VpnExclusionsApi.md#GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork) | **Get** /organizations/{organizationId}/appliance/trafficShaping/vpnExclusions/byNetwork | Display VPN exclusion rules for MX networks.
[**UpdateNetworkApplianceTrafficShapingVpnExclusions**](VpnExclusionsApi.md#UpdateNetworkApplianceTrafficShapingVpnExclusions) | **Put** /networks/{networkId}/appliance/trafficShaping/vpnExclusions | Update VPN exclusion rules for an MX network.



## GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork

> []UpdateNetworkApplianceTrafficShapingVpnExclusions200Response GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Display VPN exclusion rules for MX networks.



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
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VpnExclusionsApi.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VpnExclusionsApi.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork`: []UpdateNetworkApplianceTrafficShapingVpnExclusions200Response
    fmt.Fprintf(os.Stdout, "Response from `VpnExclusionsApi.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 

### Return type

[**[]UpdateNetworkApplianceTrafficShapingVpnExclusions200Response**](UpdateNetworkApplianceTrafficShapingVpnExclusions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingVpnExclusions

> UpdateNetworkApplianceTrafficShapingVpnExclusions200Response UpdateNetworkApplianceTrafficShapingVpnExclusions(ctx, networkId).UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest(updateNetworkApplianceTrafficShapingVpnExclusionsRequest).Execute()

Update VPN exclusion rules for an MX network.



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
    updateNetworkApplianceTrafficShapingVpnExclusionsRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest() // UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VpnExclusionsApi.UpdateNetworkApplianceTrafficShapingVpnExclusions(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest(updateNetworkApplianceTrafficShapingVpnExclusionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VpnExclusionsApi.UpdateNetworkApplianceTrafficShapingVpnExclusions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingVpnExclusions`: UpdateNetworkApplianceTrafficShapingVpnExclusions200Response
    fmt.Fprintf(os.Stdout, "Response from `VpnExclusionsApi.UpdateNetworkApplianceTrafficShapingVpnExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingVpnExclusionsRequest** | [**UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest**](UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest.md) |  | 

### Return type

[**UpdateNetworkApplianceTrafficShapingVpnExclusions200Response**](UpdateNetworkApplianceTrafficShapingVpnExclusions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

