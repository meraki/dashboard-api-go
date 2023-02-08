# \DhcpServerPolicyApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer**](DhcpServerPolicyApi.md#CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) | **Post** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers | Add a server to be trusted by Dynamic ARP Inspection on this network
[**DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer**](DhcpServerPolicyApi.md#DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) | **Delete** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId} | Remove a server from being trusted by Dynamic ARP Inspection on this network
[**GetNetworkSwitchDhcpServerPolicy**](DhcpServerPolicyApi.md#GetNetworkSwitchDhcpServerPolicy) | **Get** /networks/{networkId}/switch/dhcpServerPolicy | Return the DHCP server settings
[**GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers**](DhcpServerPolicyApi.md#GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers) | **Get** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers | Return the list of servers trusted by Dynamic ARP Inspection on this network
[**GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice**](DhcpServerPolicyApi.md#GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice) | **Get** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/warnings/byDevice | Return the devices that have a Dynamic ARP Inspection warning and their warnings
[**UpdateNetworkSwitchDhcpServerPolicy**](DhcpServerPolicyApi.md#UpdateNetworkSwitchDhcpServerPolicy) | **Put** /networks/{networkId}/switch/dhcpServerPolicy | Update the DHCP server settings
[**UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer**](DhcpServerPolicyApi.md#UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer) | **Put** /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers/{trustedServerId} | Update a server that is trusted by Dynamic ARP Inspection on this network



## CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer

> GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx, networkId).CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer).Execute()

Add a server to be trusted by Dynamic ARP Inspection on this network



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
    createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer := *openapiclient.NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest("Mac_example", int32(123), *openapiclient.NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequestIpv4()) // CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DhcpServerPolicyApi.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId).CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyApi.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyApi.CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer** | [**CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest**](CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest.md) |  | 

### Return type

[**GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer

> DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx, networkId, trustedServerId).Execute()

Remove a server from being trusted by Dynamic ARP Inspection on this network



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
    trustedServerId := "trustedServerId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DhcpServerPolicyApi.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId, trustedServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyApi.DeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**trustedServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchDhcpServerPolicy

> map[string]interface{} GetNetworkSwitchDhcpServerPolicy(ctx, networkId).Execute()

Return the DHCP server settings



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
    resp, r, err := apiClient.DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicy(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchDhcpServerPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpServerPolicyRequest struct via the builder pattern


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


## GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers

> []GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the list of servers trusted by Dynamic ARP Inspection on this network



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers`: []GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice

> []GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Return the devices that have a Dynamic ARP Inspection warning and their warnings



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice`: []GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyApi.GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

### Return type

[**[]GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionWarningsByDevice200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkSwitchDhcpServerPolicy

> map[string]interface{} UpdateNetworkSwitchDhcpServerPolicy(ctx, networkId).UpdateNetworkSwitchDhcpServerPolicy(updateNetworkSwitchDhcpServerPolicy).Execute()

Update the DHCP server settings



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
    updateNetworkSwitchDhcpServerPolicy := *openapiclient.NewUpdateNetworkSwitchDhcpServerPolicyRequest() // UpdateNetworkSwitchDhcpServerPolicyRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicy(context.Background(), networkId).UpdateNetworkSwitchDhcpServerPolicy(updateNetworkSwitchDhcpServerPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchDhcpServerPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchDhcpServerPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSwitchDhcpServerPolicy** | [**UpdateNetworkSwitchDhcpServerPolicyRequest**](UpdateNetworkSwitchDhcpServerPolicyRequest.md) |  | 

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


## UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer

> GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(ctx, networkId, trustedServerId).UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer).Execute()

Update a server that is trusted by Dynamic ARP Inspection on this network



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
    trustedServerId := "trustedServerId_example" // string | 
    updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer := *openapiclient.NewUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest() // UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(context.Background(), networkId, trustedServerId).UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer(updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `DhcpServerPolicyApi.UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**trustedServerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer** | [**UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest**](UpdateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerRequest.md) |  | 

### Return type

[**GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner**](GetNetworkSwitchDhcpServerPolicyArpInspectionTrustedServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

