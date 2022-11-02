# \UpgradesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationFirmwareUpgrades**](UpgradesApi.md#GetOrganizationFirmwareUpgrades) | **Get** /organizations/{organizationId}/firmware/upgrades | Get firmware upgrade information for an organization



## GetOrganizationFirmwareUpgrades

> []GetOrganizationFirmwareUpgrades200ResponseInner GetOrganizationFirmwareUpgrades(ctx, organizationId).Status(status).ProductType(productType).Execute()

Get firmware upgrade information for an organization



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
    organizationId := "organizationId_example" // string | 
    status := []string{"Inner_example"} // []string | The status of an upgrade  (optional)
    productType := []string{"Inner_example"} // []string | The product type in a given upgrade ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UpgradesApi.GetOrganizationFirmwareUpgrades(context.Background(), organizationId).Status(status).ProductType(productType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UpgradesApi.GetOrganizationFirmwareUpgrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationFirmwareUpgrades`: []GetOrganizationFirmwareUpgrades200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UpgradesApi.GetOrganizationFirmwareUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationFirmwareUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **[]string** | The status of an upgrade  | 
 **productType** | **[]string** | The product type in a given upgrade ID | 

### Return type

[**[]GetOrganizationFirmwareUpgrades200ResponseInner**](GetOrganizationFirmwareUpgrades200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

