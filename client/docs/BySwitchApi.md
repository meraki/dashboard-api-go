# \BySwitchApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSwitchPortsBySwitch**](BySwitchApi.md#GetOrganizationSwitchPortsBySwitch) | **Get** /organizations/{organizationId}/switch/ports/bySwitch | List the switchports in an organization by switch



## GetOrganizationSwitchPortsBySwitch

> []InlineResponse200140 GetOrganizationSwitchPortsBySwitch(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).PortProfileIds(portProfileIds).Name(name).Mac(mac).Macs(macs).Serial(serial).Serials(serials).ConfigurationUpdatedAfter(configurationUpdatedAfter).Execute()

List the switchports in an organization by switch



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
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 50. Default is 50. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | Optional parameter to filter switchports by network. (optional)
    portProfileIds := []string{"Inner_example"} // []string | Optional parameter to filter switchports belonging to the specified port profiles. (optional)
    name := "name_example" // string | Optional parameter to filter switchports belonging to switches by name. All returned switches will have a name that contains the search term or is an exact match. (optional)
    mac := "mac_example" // string | Optional parameter to filter switchports belonging to switches by MAC address. All returned switches will have a MAC address that contains the search term or is an exact match. (optional)
    macs := []string{"Inner_example"} // []string | Optional parameter to filter switchports by one or more MAC addresses belonging to devices. All switchports returned belong to MAC addresses of switches that are an exact match. (optional)
    serial := "serial_example" // string | Optional parameter to filter switchports belonging to switches by serial number. All returned switches will have a serial number that contains the search term or is an exact match. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter switchports belonging to switches with one or more serial numbers. All switchports returned belong to serial numbers of switches that are an exact match. (optional)
    configurationUpdatedAfter := "configurationUpdatedAfter_example" // string | Optional parameter to filter results by switches where the configuration has been updated after the given timestamp. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BySwitchApi.GetOrganizationSwitchPortsBySwitch(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).PortProfileIds(portProfileIds).Name(name).Mac(mac).Macs(macs).Serial(serial).Serials(serials).ConfigurationUpdatedAfter(configurationUpdatedAfter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BySwitchApi.GetOrganizationSwitchPortsBySwitch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSwitchPortsBySwitch`: []InlineResponse200140
    fmt.Fprintf(os.Stdout, "Response from `BySwitchApi.GetOrganizationSwitchPortsBySwitch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSwitchPortsBySwitchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 50. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter switchports by network. | 
 **portProfileIds** | **[]string** | Optional parameter to filter switchports belonging to the specified port profiles. | 
 **name** | **string** | Optional parameter to filter switchports belonging to switches by name. All returned switches will have a name that contains the search term or is an exact match. | 
 **mac** | **string** | Optional parameter to filter switchports belonging to switches by MAC address. All returned switches will have a MAC address that contains the search term or is an exact match. | 
 **macs** | **[]string** | Optional parameter to filter switchports by one or more MAC addresses belonging to devices. All switchports returned belong to MAC addresses of switches that are an exact match. | 
 **serial** | **string** | Optional parameter to filter switchports belonging to switches by serial number. All returned switches will have a serial number that contains the search term or is an exact match. | 
 **serials** | **[]string** | Optional parameter to filter switchports belonging to switches with one or more serial numbers. All switchports returned belong to serial numbers of switches that are an exact match. | 
 **configurationUpdatedAfter** | **string** | Optional parameter to filter results by switches where the configuration has been updated after the given timestamp. | 

### Return type

[**[]InlineResponse200140**](InlineResponse200140.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

