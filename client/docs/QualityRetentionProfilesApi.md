# \QualityRetentionProfilesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkCameraQualityRetentionProfile**](QualityRetentionProfilesApi.md#CreateNetworkCameraQualityRetentionProfile) | **Post** /networks/{networkId}/camera/qualityRetentionProfiles | Creates new quality retention profile for this network.
[**DeleteNetworkCameraQualityRetentionProfile**](QualityRetentionProfilesApi.md#DeleteNetworkCameraQualityRetentionProfile) | **Delete** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Delete an existing quality retention profile for this network.
[**GetNetworkCameraQualityRetentionProfile**](QualityRetentionProfilesApi.md#GetNetworkCameraQualityRetentionProfile) | **Get** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Retrieve a single quality retention profile
[**GetNetworkCameraQualityRetentionProfiles**](QualityRetentionProfilesApi.md#GetNetworkCameraQualityRetentionProfiles) | **Get** /networks/{networkId}/camera/qualityRetentionProfiles | List the quality retention profiles for this network
[**UpdateNetworkCameraQualityRetentionProfile**](QualityRetentionProfilesApi.md#UpdateNetworkCameraQualityRetentionProfile) | **Put** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Update an existing quality retention profile for this network.



## CreateNetworkCameraQualityRetentionProfile

> map[string]interface{} CreateNetworkCameraQualityRetentionProfile(ctx, networkId).CreateNetworkCameraQualityRetentionProfile(createNetworkCameraQualityRetentionProfile).Execute()

Creates new quality retention profile for this network.



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
    createNetworkCameraQualityRetentionProfile := *openapiclient.NewCreateNetworkCameraQualityRetentionProfileRequest("Name_example") // CreateNetworkCameraQualityRetentionProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityRetentionProfilesApi.CreateNetworkCameraQualityRetentionProfile(context.Background(), networkId).CreateNetworkCameraQualityRetentionProfile(createNetworkCameraQualityRetentionProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityRetentionProfilesApi.CreateNetworkCameraQualityRetentionProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkCameraQualityRetentionProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QualityRetentionProfilesApi.CreateNetworkCameraQualityRetentionProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkCameraQualityRetentionProfile** | [**CreateNetworkCameraQualityRetentionProfileRequest**](CreateNetworkCameraQualityRetentionProfileRequest.md) |  | 

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


## DeleteNetworkCameraQualityRetentionProfile

> DeleteNetworkCameraQualityRetentionProfile(ctx, networkId, qualityRetentionProfileId).Execute()

Delete an existing quality retention profile for this network.



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
    qualityRetentionProfileId := "qualityRetentionProfileId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityRetentionProfilesApi.DeleteNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityRetentionProfilesApi.DeleteNetworkCameraQualityRetentionProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**qualityRetentionProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


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


## GetNetworkCameraQualityRetentionProfile

> map[string]interface{} GetNetworkCameraQualityRetentionProfile(ctx, networkId, qualityRetentionProfileId).Execute()

Retrieve a single quality retention profile



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
    qualityRetentionProfileId := "qualityRetentionProfileId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityRetentionProfilesApi.GetNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityRetentionProfilesApi.GetNetworkCameraQualityRetentionProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCameraQualityRetentionProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QualityRetentionProfilesApi.GetNetworkCameraQualityRetentionProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**qualityRetentionProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


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


## GetNetworkCameraQualityRetentionProfiles

> []map[string]interface{} GetNetworkCameraQualityRetentionProfiles(ctx, networkId).Execute()

List the quality retention profiles for this network



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
    resp, r, err := apiClient.QualityRetentionProfilesApi.GetNetworkCameraQualityRetentionProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityRetentionProfilesApi.GetNetworkCameraQualityRetentionProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCameraQualityRetentionProfiles`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QualityRetentionProfilesApi.GetNetworkCameraQualityRetentionProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraQualityRetentionProfilesRequest struct via the builder pattern


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


## UpdateNetworkCameraQualityRetentionProfile

> map[string]interface{} UpdateNetworkCameraQualityRetentionProfile(ctx, networkId, qualityRetentionProfileId).UpdateNetworkCameraQualityRetentionProfile(updateNetworkCameraQualityRetentionProfile).Execute()

Update an existing quality retention profile for this network.



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
    qualityRetentionProfileId := "qualityRetentionProfileId_example" // string | 
    updateNetworkCameraQualityRetentionProfile := *openapiclient.NewUpdateNetworkCameraQualityRetentionProfileRequest() // UpdateNetworkCameraQualityRetentionProfileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QualityRetentionProfilesApi.UpdateNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).UpdateNetworkCameraQualityRetentionProfile(updateNetworkCameraQualityRetentionProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QualityRetentionProfilesApi.UpdateNetworkCameraQualityRetentionProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCameraQualityRetentionProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `QualityRetentionProfilesApi.UpdateNetworkCameraQualityRetentionProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**qualityRetentionProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCameraQualityRetentionProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkCameraQualityRetentionProfile** | [**UpdateNetworkCameraQualityRetentionProfileRequest**](UpdateNetworkCameraQualityRetentionProfileRequest.md) |  | 

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

