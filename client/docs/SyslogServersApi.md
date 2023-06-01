# \SyslogServersApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSyslogServers**](SyslogServersApi.md#GetNetworkSyslogServers) | **Get** /networks/{networkId}/syslogServers | List the syslog servers for a network
[**UpdateNetworkSyslogServers**](SyslogServersApi.md#UpdateNetworkSyslogServers) | **Put** /networks/{networkId}/syslogServers | Update the syslog servers for a network



## GetNetworkSyslogServers

> InlineResponse20073 GetNetworkSyslogServers(ctx, networkId).Execute()

List the syslog servers for a network



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
    networkId := "networkId_example" // string | Network ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyslogServersApi.GetNetworkSyslogServers(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyslogServersApi.GetNetworkSyslogServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSyslogServers`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `SyslogServersApi.GetNetworkSyslogServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSyslogServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSyslogServers

> InlineResponse20073 UpdateNetworkSyslogServers(ctx, networkId).UpdateNetworkSyslogServers(updateNetworkSyslogServers).Execute()

Update the syslog servers for a network



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
    networkId := "networkId_example" // string | Network ID
    updateNetworkSyslogServers := *openapiclient.NewInlineObject140([]openapiclient.NetworksNetworkIdSyslogServersServers{*openapiclient.NewNetworksNetworkIdSyslogServersServers("Host_example", int32(123), []string{"Roles_example"})}) // InlineObject140 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyslogServersApi.UpdateNetworkSyslogServers(context.Background(), networkId).UpdateNetworkSyslogServers(updateNetworkSyslogServers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyslogServersApi.UpdateNetworkSyslogServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSyslogServers`: InlineResponse20073
    fmt.Fprintf(os.Stdout, "Response from `SyslogServersApi.UpdateNetworkSyslogServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSyslogServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSyslogServers** | [**InlineObject140**](InlineObject140.md) |  | 

### Return type

[**InlineResponse20073**](InlineResponse20073.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

