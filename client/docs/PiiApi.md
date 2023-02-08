# \PiiApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkPiiRequest**](PiiApi.md#CreateNetworkPiiRequest) | **Post** /networks/{networkId}/pii/requests | Submit a new delete or restrict processing PII request
[**DeleteNetworkPiiRequest**](PiiApi.md#DeleteNetworkPiiRequest) | **Delete** /networks/{networkId}/pii/requests/{requestId} | Delete a restrict processing PII request
[**GetNetworkPiiPiiKeys**](PiiApi.md#GetNetworkPiiPiiKeys) | **Get** /networks/{networkId}/pii/piiKeys | List the keys required to access Personally Identifiable Information (PII) for a given identifier
[**GetNetworkPiiRequest**](PiiApi.md#GetNetworkPiiRequest) | **Get** /networks/{networkId}/pii/requests/{requestId} | Return a PII request
[**GetNetworkPiiRequests**](PiiApi.md#GetNetworkPiiRequests) | **Get** /networks/{networkId}/pii/requests | List the PII requests for this network or organization
[**GetNetworkPiiSmDevicesForKey**](PiiApi.md#GetNetworkPiiSmDevicesForKey) | **Get** /networks/{networkId}/pii/smDevicesForKey | Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier
[**GetNetworkPiiSmOwnersForKey**](PiiApi.md#GetNetworkPiiSmOwnersForKey) | **Get** /networks/{networkId}/pii/smOwnersForKey | Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier



## CreateNetworkPiiRequest

> map[string]interface{} CreateNetworkPiiRequest(ctx, networkId).CreateNetworkPiiRequest(createNetworkPiiRequest).Execute()

Submit a new delete or restrict processing PII request



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
    createNetworkPiiRequest := *openapiclient.NewCreateNetworkPiiRequestRequest() // CreateNetworkPiiRequestRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiiApi.CreateNetworkPiiRequest(context.Background(), networkId).CreateNetworkPiiRequest(createNetworkPiiRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiiApi.CreateNetworkPiiRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkPiiRequest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PiiApi.CreateNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPiiRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkPiiRequest** | [**CreateNetworkPiiRequestRequest**](CreateNetworkPiiRequestRequest.md) |  | 

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


## DeleteNetworkPiiRequest

> DeleteNetworkPiiRequest(ctx, networkId, requestId).Execute()

Delete a restrict processing PII request



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
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiiApi.DeleteNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiiApi.DeleteNetworkPiiRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPiiRequestRequest struct via the builder pattern


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


## GetNetworkPiiPiiKeys

> map[string]interface{} GetNetworkPiiPiiKeys(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

List the keys required to access Personally Identifiable Information (PII) for a given identifier



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
    username := "username_example" // string | The username of a Systems Manager user (optional)
    email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
    mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
    serial := "serial_example" // string | The serial of a Systems Manager device (optional)
    imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
    bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiiApi.GetNetworkPiiPiiKeys(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiiApi.GetNetworkPiiPiiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiPiiKeys`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PiiApi.GetNetworkPiiPiiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiPiiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

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


## GetNetworkPiiRequest

> map[string]interface{} GetNetworkPiiRequest(ctx, networkId, requestId).Execute()

Return a PII request



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
    requestId := "requestId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiiApi.GetNetworkPiiRequest(context.Background(), networkId, requestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiiApi.GetNetworkPiiRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiRequest`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PiiApi.GetNetworkPiiRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**requestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestRequest struct via the builder pattern


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


## GetNetworkPiiRequests

> []map[string]interface{} GetNetworkPiiRequests(ctx, networkId).Execute()

List the PII requests for this network or organization



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
    resp, r, err := apiClient.PiiApi.GetNetworkPiiRequests(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiiApi.GetNetworkPiiRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiRequests`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PiiApi.GetNetworkPiiRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiRequestsRequest struct via the builder pattern


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


## GetNetworkPiiSmDevicesForKey

> map[string]interface{} GetNetworkPiiSmDevicesForKey(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier



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
    username := "username_example" // string | The username of a Systems Manager user (optional)
    email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
    mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
    serial := "serial_example" // string | The serial of a Systems Manager device (optional)
    imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
    bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiiApi.GetNetworkPiiSmDevicesForKey(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiiApi.GetNetworkPiiSmDevicesForKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiSmDevicesForKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PiiApi.GetNetworkPiiSmDevicesForKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiSmDevicesForKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

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


## GetNetworkPiiSmOwnersForKey

> map[string]interface{} GetNetworkPiiSmOwnersForKey(ctx, networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()

Given a piece of Personally Identifiable Information (PII), return the Systems Manager owner ID(s) associated with that identifier



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
    username := "username_example" // string | The username of a Systems Manager user (optional)
    email := "email_example" // string | The email of a network user account or a Systems Manager device (optional)
    mac := "mac_example" // string | The MAC of a network client device or a Systems Manager device (optional)
    serial := "serial_example" // string | The serial of a Systems Manager device (optional)
    imei := "imei_example" // string | The IMEI of a Systems Manager device (optional)
    bluetoothMac := "bluetoothMac_example" // string | The MAC of a Bluetooth client (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PiiApi.GetNetworkPiiSmOwnersForKey(context.Background(), networkId).Username(username).Email(email).Mac(mac).Serial(serial).Imei(imei).BluetoothMac(bluetoothMac).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PiiApi.GetNetworkPiiSmOwnersForKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPiiSmOwnersForKey`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PiiApi.GetNetworkPiiSmOwnersForKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPiiSmOwnersForKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | The username of a Systems Manager user | 
 **email** | **string** | The email of a network user account or a Systems Manager device | 
 **mac** | **string** | The MAC of a network client device or a Systems Manager device | 
 **serial** | **string** | The serial of a Systems Manager device | 
 **imei** | **string** | The IMEI of a Systems Manager device | 
 **bluetoothMac** | **string** | The MAC of a Bluetooth client | 

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

