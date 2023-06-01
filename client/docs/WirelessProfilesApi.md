# \WirelessProfilesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkCameraWirelessProfile**](WirelessProfilesApi.md#CreateNetworkCameraWirelessProfile) | **Post** /networks/{networkId}/camera/wirelessProfiles | Creates a new camera wireless profile for this network.
[**DeleteNetworkCameraWirelessProfile**](WirelessProfilesApi.md#DeleteNetworkCameraWirelessProfile) | **Delete** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Delete an existing camera wireless profile for this network.
[**GetDeviceCameraWirelessProfiles**](WirelessProfilesApi.md#GetDeviceCameraWirelessProfiles) | **Get** /devices/{serial}/camera/wirelessProfiles | Returns wireless profile assigned to the given camera
[**GetNetworkCameraWirelessProfile**](WirelessProfilesApi.md#GetNetworkCameraWirelessProfile) | **Get** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Retrieve a single camera wireless profile.
[**GetNetworkCameraWirelessProfiles**](WirelessProfilesApi.md#GetNetworkCameraWirelessProfiles) | **Get** /networks/{networkId}/camera/wirelessProfiles | List the camera wireless profiles for this network.
[**UpdateDeviceCameraWirelessProfiles**](WirelessProfilesApi.md#UpdateDeviceCameraWirelessProfiles) | **Put** /devices/{serial}/camera/wirelessProfiles | Assign wireless profiles to the given camera
[**UpdateNetworkCameraWirelessProfile**](WirelessProfilesApi.md#UpdateNetworkCameraWirelessProfile) | **Put** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Update an existing camera wireless profile in this network.



## CreateNetworkCameraWirelessProfile

> map[string]interface{} CreateNetworkCameraWirelessProfile(ctx, networkId).CreateNetworkCameraWirelessProfile(createNetworkCameraWirelessProfile).Execute()

Creates a new camera wireless profile for this network.



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
    createNetworkCameraWirelessProfile := *openapiclient.NewInlineObject65("Name_example", *openapiclient.NewNetworksNetworkIdCameraWirelessProfilesSsid()) // InlineObject65 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessProfilesApi.CreateNetworkCameraWirelessProfile(context.Background(), networkId).CreateNetworkCameraWirelessProfile(createNetworkCameraWirelessProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesApi.CreateNetworkCameraWirelessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkCameraWirelessProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesApi.CreateNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkCameraWirelessProfile** | [**InlineObject65**](InlineObject65.md) |  | 

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


## DeleteNetworkCameraWirelessProfile

> DeleteNetworkCameraWirelessProfile(ctx, networkId, wirelessProfileId).Execute()

Delete an existing camera wireless profile for this network.



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
    wirelessProfileId := "wirelessProfileId_example" // string | Wireless profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessProfilesApi.DeleteNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesApi.DeleteNetworkCameraWirelessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**wirelessProfileId** | **string** | Wireless profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkCameraWirelessProfileRequest struct via the builder pattern


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


## GetDeviceCameraWirelessProfiles

> map[string]interface{} GetDeviceCameraWirelessProfiles(ctx, serial).Execute()

Returns wireless profile assigned to the given camera



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
    serial := "serial_example" // string | Serial

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessProfilesApi.GetDeviceCameraWirelessProfiles(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesApi.GetDeviceCameraWirelessProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraWirelessProfiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesApi.GetDeviceCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraWirelessProfilesRequest struct via the builder pattern


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


## GetNetworkCameraWirelessProfile

> map[string]interface{} GetNetworkCameraWirelessProfile(ctx, networkId, wirelessProfileId).Execute()

Retrieve a single camera wireless profile.



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
    wirelessProfileId := "wirelessProfileId_example" // string | Wireless profile ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessProfilesApi.GetNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesApi.GetNetworkCameraWirelessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCameraWirelessProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesApi.GetNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**wirelessProfileId** | **string** | Wireless profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraWirelessProfileRequest struct via the builder pattern


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


## GetNetworkCameraWirelessProfiles

> []map[string]interface{} GetNetworkCameraWirelessProfiles(ctx, networkId).Execute()

List the camera wireless profiles for this network.



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
    resp, r, err := apiClient.WirelessProfilesApi.GetNetworkCameraWirelessProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesApi.GetNetworkCameraWirelessProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCameraWirelessProfiles`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesApi.GetNetworkCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraWirelessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceCameraWirelessProfiles

> map[string]interface{} UpdateDeviceCameraWirelessProfiles(ctx, serial).UpdateDeviceCameraWirelessProfiles(updateDeviceCameraWirelessProfiles).Execute()

Assign wireless profiles to the given camera



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
    serial := "serial_example" // string | Serial
    updateDeviceCameraWirelessProfiles := *openapiclient.NewInlineObject8(*openapiclient.NewDevicesSerialCameraWirelessProfilesIds()) // InlineObject8 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessProfilesApi.UpdateDeviceCameraWirelessProfiles(context.Background(), serial).UpdateDeviceCameraWirelessProfiles(updateDeviceCameraWirelessProfiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesApi.UpdateDeviceCameraWirelessProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCameraWirelessProfiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesApi.UpdateDeviceCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraWirelessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraWirelessProfiles** | [**InlineObject8**](InlineObject8.md) |  | 

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


## UpdateNetworkCameraWirelessProfile

> map[string]interface{} UpdateNetworkCameraWirelessProfile(ctx, networkId, wirelessProfileId).UpdateNetworkCameraWirelessProfile(updateNetworkCameraWirelessProfile).Execute()

Update an existing camera wireless profile in this network.



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
    wirelessProfileId := "wirelessProfileId_example" // string | Wireless profile ID
    updateNetworkCameraWirelessProfile := *openapiclient.NewInlineObject66() // InlineObject66 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WirelessProfilesApi.UpdateNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).UpdateNetworkCameraWirelessProfile(updateNetworkCameraWirelessProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WirelessProfilesApi.UpdateNetworkCameraWirelessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCameraWirelessProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `WirelessProfilesApi.UpdateNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**wirelessProfileId** | **string** | Wireless profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkCameraWirelessProfile** | [**InlineObject66**](InlineObject66.md) |  | 

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

