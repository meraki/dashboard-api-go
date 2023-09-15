# \Ipv6Api

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateDeviceWirelessAlternateManagementInterfaceIpv6**](Ipv6Api.md#UpdateDeviceWirelessAlternateManagementInterfaceIpv6) | **Put** /devices/{serial}/wireless/alternateManagementInterface/ipv6 | Update alternate management interface IPv6 address



## UpdateDeviceWirelessAlternateManagementInterfaceIpv6

> UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response UpdateDeviceWirelessAlternateManagementInterfaceIpv6(ctx, serial).UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request(updateDeviceWirelessAlternateManagementInterfaceIpv6Request).Execute()

Update alternate management interface IPv6 address



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
    serial := "serial_example" // string | Serial
    updateDeviceWirelessAlternateManagementInterfaceIpv6Request := *openapiclient.NewUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request() // UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Ipv6Api.UpdateDeviceWirelessAlternateManagementInterfaceIpv6(context.Background(), serial).UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request(updateDeviceWirelessAlternateManagementInterfaceIpv6Request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Ipv6Api.UpdateDeviceWirelessAlternateManagementInterfaceIpv6``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceWirelessAlternateManagementInterfaceIpv6`: UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response
    fmt.Fprintf(os.Stdout, "Response from `Ipv6Api.UpdateDeviceWirelessAlternateManagementInterfaceIpv6`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessAlternateManagementInterfaceIpv6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessAlternateManagementInterfaceIpv6Request** | [**UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request**](UpdateDeviceWirelessAlternateManagementInterfaceIpv6Request.md) |  | 

### Return type

[**UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response**](UpdateDeviceWirelessAlternateManagementInterfaceIpv6200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

