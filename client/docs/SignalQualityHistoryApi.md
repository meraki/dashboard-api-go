# \SignalQualityHistoryApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessSignalQualityHistory**](SignalQualityHistoryApi.md#GetNetworkWirelessSignalQualityHistory) | **Get** /networks/{networkId}/wireless/signalQualityHistory | Return signal quality (SNR/RSSI) over time for a device or network client



## GetNetworkWirelessSignalQualityHistory

> []GetNetworkWirelessSignalQualityHistory200ResponseInner GetNetworkWirelessSignalQualityHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).Execute()

Return signal quality (SNR/RSSI) over time for a device or network client



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. (optional)
    autoResolution := true // bool | Automatically select a data resolution based on the given timespan; this overrides the value specified by the 'resolution' parameter. The default setting is false. (optional)
    clientId := "clientId_example" // string | Filter results by network client. (optional)
    deviceSerial := "deviceSerial_example" // string | Filter results by device. (optional)
    apTag := "apTag_example" // string | Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified. (optional)
    band := "band_example" // string | Filter results by band (either '2.4', '5' or '6'). (optional)
    ssid := int32(56) // int32 | Filter results by SSID number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignalQualityHistoryApi.GetNetworkWirelessSignalQualityHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).AutoResolution(autoResolution).ClientId(clientId).DeviceSerial(deviceSerial).ApTag(apTag).Band(band).Ssid(ssid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignalQualityHistoryApi.GetNetworkWirelessSignalQualityHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessSignalQualityHistory`: []GetNetworkWirelessSignalQualityHistory200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SignalQualityHistoryApi.GetNetworkWirelessSignalQualityHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessSignalQualityHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 7 days. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 300, 600, 1200, 3600, 14400, 86400. The default is 86400. | 
 **autoResolution** | **bool** | Automatically select a data resolution based on the given timespan; this overrides the value specified by the &#39;resolution&#39; parameter. The default setting is false. | 
 **clientId** | **string** | Filter results by network client. | 
 **deviceSerial** | **string** | Filter results by device. | 
 **apTag** | **string** | Filter results by AP tag; either :clientId or :deviceSerial must be jointly specified. | 
 **band** | **string** | Filter results by band (either &#39;2.4&#39;, &#39;5&#39; or &#39;6&#39;). | 
 **ssid** | **int32** | Filter results by SSID number. | 

### Return type

[**[]GetNetworkWirelessSignalQualityHistory200ResponseInner**](GetNetworkWirelessSignalQualityHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

