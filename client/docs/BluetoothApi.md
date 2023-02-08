# \BluetoothApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceWirelessBluetoothSettings**](BluetoothApi.md#GetDeviceWirelessBluetoothSettings) | **Get** /devices/{serial}/wireless/bluetooth/settings | Return the bluetooth settings for a wireless device
[**GetNetworkWirelessBluetoothSettings**](BluetoothApi.md#GetNetworkWirelessBluetoothSettings) | **Get** /networks/{networkId}/wireless/bluetooth/settings | Return the Bluetooth settings for a network. &lt;a href&#x3D;\&quot;https://documentation.meraki.com/MR/Bluetooth/Bluetooth_Low_Energy_(BLE)\&quot;&gt;Bluetooth settings&lt;/a&gt; must be enabled on the network.
[**UpdateDeviceWirelessBluetoothSettings**](BluetoothApi.md#UpdateDeviceWirelessBluetoothSettings) | **Put** /devices/{serial}/wireless/bluetooth/settings | Update the bluetooth settings for a wireless device
[**UpdateNetworkWirelessBluetoothSettings**](BluetoothApi.md#UpdateNetworkWirelessBluetoothSettings) | **Put** /networks/{networkId}/wireless/bluetooth/settings | Update the Bluetooth settings for a network



## GetDeviceWirelessBluetoothSettings

> GetDeviceWirelessBluetoothSettings200Response GetDeviceWirelessBluetoothSettings(ctx, serial).Execute()

Return the bluetooth settings for a wireless device



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
    serial := "serial_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BluetoothApi.GetDeviceWirelessBluetoothSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BluetoothApi.GetDeviceWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceWirelessBluetoothSettings`: GetDeviceWirelessBluetoothSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `BluetoothApi.GetDeviceWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceWirelessBluetoothSettings200Response**](GetDeviceWirelessBluetoothSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessBluetoothSettings

> GetNetworkWirelessBluetoothSettings200Response GetNetworkWirelessBluetoothSettings(ctx, networkId).Execute()

Return the Bluetooth settings for a network. <a href=\"https://documentation.meraki.com/MR/Bluetooth/Bluetooth_Low_Energy_(BLE)\">Bluetooth settings</a> must be enabled on the network.



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
    resp, r, err := apiClient.BluetoothApi.GetNetworkWirelessBluetoothSettings(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BluetoothApi.GetNetworkWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessBluetoothSettings`: GetNetworkWirelessBluetoothSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `BluetoothApi.GetNetworkWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkWirelessBluetoothSettings200Response**](GetNetworkWirelessBluetoothSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceWirelessBluetoothSettings

> GetDeviceWirelessBluetoothSettings200Response UpdateDeviceWirelessBluetoothSettings(ctx, serial).UpdateDeviceWirelessBluetoothSettings(updateDeviceWirelessBluetoothSettings).Execute()

Update the bluetooth settings for a wireless device



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
    serial := "serial_example" // string | 
    updateDeviceWirelessBluetoothSettings := *openapiclient.NewUpdateDeviceWirelessBluetoothSettingsRequest() // UpdateDeviceWirelessBluetoothSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BluetoothApi.UpdateDeviceWirelessBluetoothSettings(context.Background(), serial).UpdateDeviceWirelessBluetoothSettings(updateDeviceWirelessBluetoothSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BluetoothApi.UpdateDeviceWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceWirelessBluetoothSettings`: GetDeviceWirelessBluetoothSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `BluetoothApi.UpdateDeviceWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceWirelessBluetoothSettings** | [**UpdateDeviceWirelessBluetoothSettingsRequest**](UpdateDeviceWirelessBluetoothSettingsRequest.md) |  | 

### Return type

[**GetDeviceWirelessBluetoothSettings200Response**](GetDeviceWirelessBluetoothSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessBluetoothSettings

> GetNetworkWirelessBluetoothSettings200Response UpdateNetworkWirelessBluetoothSettings(ctx, networkId).UpdateNetworkWirelessBluetoothSettings(updateNetworkWirelessBluetoothSettings).Execute()

Update the Bluetooth settings for a network



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
    updateNetworkWirelessBluetoothSettings := *openapiclient.NewUpdateNetworkWirelessBluetoothSettingsRequest() // UpdateNetworkWirelessBluetoothSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BluetoothApi.UpdateNetworkWirelessBluetoothSettings(context.Background(), networkId).UpdateNetworkWirelessBluetoothSettings(updateNetworkWirelessBluetoothSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BluetoothApi.UpdateNetworkWirelessBluetoothSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessBluetoothSettings`: GetNetworkWirelessBluetoothSettings200Response
    fmt.Fprintf(os.Stdout, "Response from `BluetoothApi.UpdateNetworkWirelessBluetoothSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessBluetoothSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkWirelessBluetoothSettings** | [**UpdateNetworkWirelessBluetoothSettingsRequest**](UpdateNetworkWirelessBluetoothSettingsRequest.md) |  | 

### Return type

[**GetNetworkWirelessBluetoothSettings200Response**](GetNetworkWirelessBluetoothSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

