# \CellularUsageHistoryApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSmDeviceCellularUsageHistory**](CellularUsageHistoryApi.md#GetNetworkSmDeviceCellularUsageHistory) | **Get** /networks/{networkId}/sm/devices/{deviceId}/cellularUsageHistory | Return the client&#39;s daily cellular data usage history



## GetNetworkSmDeviceCellularUsageHistory

> []GetNetworkSmDeviceCellularUsageHistory200ResponseInner GetNetworkSmDeviceCellularUsageHistory(ctx, networkId, deviceId).Execute()

Return the client's daily cellular data usage history



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
    deviceId := "deviceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellularUsageHistoryApi.GetNetworkSmDeviceCellularUsageHistory(context.Background(), networkId, deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellularUsageHistoryApi.GetNetworkSmDeviceCellularUsageHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmDeviceCellularUsageHistory`: []GetNetworkSmDeviceCellularUsageHistory200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `CellularUsageHistoryApi.GetNetworkSmDeviceCellularUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**deviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmDeviceCellularUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetNetworkSmDeviceCellularUsageHistory200ResponseInner**](GetNetworkSmDeviceCellularUsageHistory200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

