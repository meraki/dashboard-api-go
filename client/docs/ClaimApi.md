# \ClaimApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VmxNetworkDevicesClaim**](ClaimApi.md#VmxNetworkDevicesClaim) | **Post** /networks/{networkId}/devices/claim/vmx | Claim a vMX into a network



## VmxNetworkDevicesClaim

> map[string]interface{} VmxNetworkDevicesClaim(ctx, networkId).VmxNetworkDevicesClaim(vmxNetworkDevicesClaim).Execute()

Claim a vMX into a network



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
    vmxNetworkDevicesClaim := *openapiclient.NewVmxNetworkDevicesClaimRequest("Size_example") // VmxNetworkDevicesClaimRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClaimApi.VmxNetworkDevicesClaim(context.Background(), networkId).VmxNetworkDevicesClaim(vmxNetworkDevicesClaim).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClaimApi.VmxNetworkDevicesClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VmxNetworkDevicesClaim`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ClaimApi.VmxNetworkDevicesClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVmxNetworkDevicesClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vmxNetworkDevicesClaim** | [**VmxNetworkDevicesClaimRequest**](VmxNetworkDevicesClaimRequest.md) |  | 

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

