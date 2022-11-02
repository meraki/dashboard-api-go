# \QualityAndRetentionApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceCameraQualityAndRetention**](QualityAndRetentionApi.md#GetDeviceCameraQualityAndRetention) | **Get** /devices/{serial}/camera/qualityAndRetention | Returns quality and retention settings for the given camera
[**UpdateDeviceCameraQualityAndRetention**](QualityAndRetentionApi.md#UpdateDeviceCameraQualityAndRetention) | **Put** /devices/{serial}/camera/qualityAndRetention | Update quality and retention settings for the given camera



## GetDeviceCameraQualityAndRetention

> map[string]interface{} GetDeviceCameraQualityAndRetention(ctx, serial).Execute()

Returns quality and retention settings for the given camera



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
    resp, r, err := apiClient.QualityAndRetentionApi.GetDeviceCameraQualityAndRetention(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityAndRetentionApi.GetDeviceCameraQualityAndRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraQualityAndRetention`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QualityAndRetentionApi.GetDeviceCameraQualityAndRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraQualityAndRetentionRequest struct via the builder pattern


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


## UpdateDeviceCameraQualityAndRetention

> map[string]interface{} UpdateDeviceCameraQualityAndRetention(ctx, serial).UpdateDeviceCameraQualityAndRetention(updateDeviceCameraQualityAndRetention).Execute()

Update quality and retention settings for the given camera



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
    updateDeviceCameraQualityAndRetention := *openapiclient.NewUpdateDeviceCameraQualityAndRetentionRequest() // UpdateDeviceCameraQualityAndRetentionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityAndRetentionApi.UpdateDeviceCameraQualityAndRetention(context.Background(), serial).UpdateDeviceCameraQualityAndRetention(updateDeviceCameraQualityAndRetention).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityAndRetentionApi.UpdateDeviceCameraQualityAndRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCameraQualityAndRetention`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QualityAndRetentionApi.UpdateDeviceCameraQualityAndRetention`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraQualityAndRetentionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraQualityAndRetention** | [**UpdateDeviceCameraQualityAndRetentionRequest**](UpdateDeviceCameraQualityAndRetentionRequest.md) |  | 

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

