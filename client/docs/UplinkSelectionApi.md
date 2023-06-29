# \UplinkSelectionApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkApplianceTrafficShapingUplinkSelection**](UplinkSelectionApi.md#GetNetworkApplianceTrafficShapingUplinkSelection) | **Get** /networks/{networkId}/appliance/trafficShaping/uplinkSelection | Show uplink selection settings for an MX network
[**UpdateNetworkApplianceTrafficShapingUplinkSelection**](UplinkSelectionApi.md#UpdateNetworkApplianceTrafficShapingUplinkSelection) | **Put** /networks/{networkId}/appliance/trafficShaping/uplinkSelection | Update uplink selection settings for an MX network



## GetNetworkApplianceTrafficShapingUplinkSelection

> GetNetworkApplianceTrafficShapingUplinkSelection200Response GetNetworkApplianceTrafficShapingUplinkSelection(ctx, networkId).Execute()

Show uplink selection settings for an MX network



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
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UplinkSelectionApi.GetNetworkApplianceTrafficShapingUplinkSelection(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UplinkSelectionApi.GetNetworkApplianceTrafficShapingUplinkSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceTrafficShapingUplinkSelection`: GetNetworkApplianceTrafficShapingUplinkSelection200Response
    fmt.Fprintf(os.Stdout, "Response from `UplinkSelectionApi.GetNetworkApplianceTrafficShapingUplinkSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingUplinkSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceTrafficShapingUplinkSelection200Response**](GetNetworkApplianceTrafficShapingUplinkSelection200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingUplinkSelection

> GetNetworkApplianceTrafficShapingUplinkSelection200Response UpdateNetworkApplianceTrafficShapingUplinkSelection(ctx, networkId).UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest(updateNetworkApplianceTrafficShapingUplinkSelectionRequest).Execute()

Update uplink selection settings for an MX network



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
    updateNetworkApplianceTrafficShapingUplinkSelectionRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest() // UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UplinkSelectionApi.UpdateNetworkApplianceTrafficShapingUplinkSelection(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest(updateNetworkApplianceTrafficShapingUplinkSelectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UplinkSelectionApi.UpdateNetworkApplianceTrafficShapingUplinkSelection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceTrafficShapingUplinkSelection`: GetNetworkApplianceTrafficShapingUplinkSelection200Response
    fmt.Fprintf(os.Stdout, "Response from `UplinkSelectionApi.UpdateNetworkApplianceTrafficShapingUplinkSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingUplinkSelectionRequest** | [**UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest**](UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest.md) |  | 

### Return type

[**GetNetworkApplianceTrafficShapingUplinkSelection200Response**](GetNetworkApplianceTrafficShapingUplinkSelection200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

