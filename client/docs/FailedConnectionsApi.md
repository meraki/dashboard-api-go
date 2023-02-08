# \FailedConnectionsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessFailedConnections**](FailedConnectionsApi.md#GetNetworkWirelessFailedConnections) | **Get** /networks/{networkId}/wireless/failedConnections | List of all failed client connection events on this network in a given time range



## GetNetworkWirelessFailedConnections

> []GetNetworkWirelessFailedConnections200ResponseInner GetNetworkWirelessFailedConnections(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Serial(serial).ClientId(clientId).Execute()

List of all failed client connection events on this network in a given time range



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information. (optional)
    ssid := int32(56) // int32 | Filter results by SSID (optional)
    vlan := int32(56) // int32 | Filter results by VLAN (optional)
    apTag := "apTag_example" // string | Filter results by AP Tag (optional)
    serial := "serial_example" // string | Filter by AP (optional)
    clientId := "clientId_example" // string | Filter by client MAC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FailedConnectionsApi.GetNetworkWirelessFailedConnections(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Band(band).Ssid(ssid).Vlan(vlan).ApTag(apTag).Serial(serial).ClientId(clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FailedConnectionsApi.GetNetworkWirelessFailedConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessFailedConnections`: []GetNetworkWirelessFailedConnections200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `FailedConnectionsApi.GetNetworkWirelessFailedConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessFailedConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 180 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). Note that data prior to February 2020 will not have band information. | 
 **ssid** | **int32** | Filter results by SSID | 
 **vlan** | **int32** | Filter results by VLAN | 
 **apTag** | **string** | Filter results by AP Tag | 
 **serial** | **string** | Filter by AP | 
 **clientId** | **string** | Filter by client MAC | 

### Return type

[**[]GetNetworkWirelessFailedConnections200ResponseInner**](GetNetworkWirelessFailedConnections200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

