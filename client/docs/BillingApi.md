# \BillingApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkWirelessBilling**](BillingApi.md#GetNetworkWirelessBilling) | **Get** /networks/{networkId}/wireless/billing | Return the billing settings of this network
[**UpdateNetworkWirelessBilling**](BillingApi.md#UpdateNetworkWirelessBilling) | **Put** /networks/{networkId}/wireless/billing | Update the billing settings



## GetNetworkWirelessBilling

> map[string]interface{} GetNetworkWirelessBilling(ctx, networkId).Execute()

Return the billing settings of this network



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
    resp, r, err := apiClient.BillingApi.GetNetworkWirelessBilling(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetNetworkWirelessBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessBilling`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetNetworkWirelessBilling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessBillingRequest struct via the builder pattern


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


## UpdateNetworkWirelessBilling

> map[string]interface{} UpdateNetworkWirelessBilling(ctx, networkId).UpdateNetworkWirelessBilling(updateNetworkWirelessBilling).Execute()

Update the billing settings



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
    updateNetworkWirelessBilling := *openapiclient.NewUpdateNetworkWirelessBillingRequest() // UpdateNetworkWirelessBillingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.UpdateNetworkWirelessBilling(context.Background(), networkId).UpdateNetworkWirelessBilling(updateNetworkWirelessBilling).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.UpdateNetworkWirelessBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessBilling`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.UpdateNetworkWirelessBilling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessBilling** | [**UpdateNetworkWirelessBillingRequest**](UpdateNetworkWirelessBillingRequest.md) |  | 

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

