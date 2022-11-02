# \SubnetPoolApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkCellularGatewaySubnetPool**](SubnetPoolApi.md#GetNetworkCellularGatewaySubnetPool) | **Get** /networks/{networkId}/cellularGateway/subnetPool | Return the subnet pool and mask configured for MGs in the network.
[**UpdateNetworkCellularGatewaySubnetPool**](SubnetPoolApi.md#UpdateNetworkCellularGatewaySubnetPool) | **Put** /networks/{networkId}/cellularGateway/subnetPool | Update the subnet pool and mask configuration for MGs in the network.



## GetNetworkCellularGatewaySubnetPool

> map[string]interface{} GetNetworkCellularGatewaySubnetPool(ctx, networkId).Execute()

Return the subnet pool and mask configured for MGs in the network.



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
    resp, r, err := apiClient.SubnetPoolApi.GetNetworkCellularGatewaySubnetPool(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubnetPoolApi.GetNetworkCellularGatewaySubnetPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCellularGatewaySubnetPool`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SubnetPoolApi.GetNetworkCellularGatewaySubnetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCellularGatewaySubnetPoolRequest struct via the builder pattern


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


## UpdateNetworkCellularGatewaySubnetPool

> map[string]interface{} UpdateNetworkCellularGatewaySubnetPool(ctx, networkId).UpdateNetworkCellularGatewaySubnetPool(updateNetworkCellularGatewaySubnetPool).Execute()

Update the subnet pool and mask configuration for MGs in the network.



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
    updateNetworkCellularGatewaySubnetPool := *openapiclient.NewUpdateNetworkCellularGatewaySubnetPoolRequest() // UpdateNetworkCellularGatewaySubnetPoolRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubnetPoolApi.UpdateNetworkCellularGatewaySubnetPool(context.Background(), networkId).UpdateNetworkCellularGatewaySubnetPool(updateNetworkCellularGatewaySubnetPool).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubnetPoolApi.UpdateNetworkCellularGatewaySubnetPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCellularGatewaySubnetPool`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SubnetPoolApi.UpdateNetworkCellularGatewaySubnetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCellularGatewaySubnetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkCellularGatewaySubnetPool** | [**UpdateNetworkCellularGatewaySubnetPoolRequest**](UpdateNetworkCellularGatewaySubnetPoolRequest.md) |  | 

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

