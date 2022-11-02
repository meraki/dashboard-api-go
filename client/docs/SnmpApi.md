# \SnmpApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNetworkSnmp**](SnmpApi.md#GetNetworkSnmp) | **Get** /networks/{networkId}/snmp | Return the SNMP settings for a network
[**GetOrganizationSnmp**](SnmpApi.md#GetOrganizationSnmp) | **Get** /organizations/{organizationId}/snmp | Return the SNMP settings for an organization
[**UpdateNetworkSnmp**](SnmpApi.md#UpdateNetworkSnmp) | **Put** /networks/{networkId}/snmp | Update the SNMP settings for a network
[**UpdateOrganizationSnmp**](SnmpApi.md#UpdateOrganizationSnmp) | **Put** /organizations/{organizationId}/snmp | Update the SNMP settings for an organization



## GetNetworkSnmp

> map[string]interface{} GetNetworkSnmp(ctx, networkId).Execute()

Return the SNMP settings for a network



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
    resp, r, err := apiClient.SnmpApi.GetNetworkSnmp(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnmpApi.GetNetworkSnmp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSnmp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SnmpApi.GetNetworkSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSnmpRequest struct via the builder pattern


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


## GetOrganizationSnmp

> map[string]interface{} GetOrganizationSnmp(ctx, organizationId).Execute()

Return the SNMP settings for an organization



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnmpApi.GetOrganizationSnmp(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnmpApi.GetOrganizationSnmp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSnmp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SnmpApi.GetOrganizationSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSnmpRequest struct via the builder pattern


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


## UpdateNetworkSnmp

> map[string]interface{} UpdateNetworkSnmp(ctx, networkId).UpdateNetworkSnmp(updateNetworkSnmp).Execute()

Update the SNMP settings for a network



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
    updateNetworkSnmp := *openapiclient.NewUpdateNetworkSnmpRequest() // UpdateNetworkSnmpRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnmpApi.UpdateNetworkSnmp(context.Background(), networkId).UpdateNetworkSnmp(updateNetworkSnmp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnmpApi.UpdateNetworkSnmp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSnmp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SnmpApi.UpdateNetworkSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkSnmp** | [**UpdateNetworkSnmpRequest**](UpdateNetworkSnmpRequest.md) |  | 

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


## UpdateOrganizationSnmp

> map[string]interface{} UpdateOrganizationSnmp(ctx, organizationId).UpdateOrganizationSnmp(updateOrganizationSnmp).Execute()

Update the SNMP settings for an organization



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
    updateOrganizationSnmp := *openapiclient.NewUpdateOrganizationSnmpRequest() // UpdateOrganizationSnmpRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnmpApi.UpdateOrganizationSnmp(context.Background(), organizationId).UpdateOrganizationSnmp(updateOrganizationSnmp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnmpApi.UpdateOrganizationSnmp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationSnmp`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SnmpApi.UpdateOrganizationSnmp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationSnmpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationSnmp** | [**UpdateOrganizationSnmpRequest**](UpdateOrganizationSnmpRequest.md) |  | 

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

