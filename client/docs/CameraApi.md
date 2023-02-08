# \CameraApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkCameraQualityRetentionProfile**](CameraApi.md#CreateNetworkCameraQualityRetentionProfile) | **Post** /networks/{networkId}/camera/qualityRetentionProfiles | Creates new quality retention profile for this network.
[**CreateNetworkCameraWirelessProfile**](CameraApi.md#CreateNetworkCameraWirelessProfile) | **Post** /networks/{networkId}/camera/wirelessProfiles | Creates a new camera wireless profile for this network.
[**CreateOrganizationCameraCustomAnalyticsArtifact**](CameraApi.md#CreateOrganizationCameraCustomAnalyticsArtifact) | **Post** /organizations/{organizationId}/camera/customAnalytics/artifacts | Create custom analytics artifact
[**DeleteNetworkCameraQualityRetentionProfile**](CameraApi.md#DeleteNetworkCameraQualityRetentionProfile) | **Delete** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Delete an existing quality retention profile for this network.
[**DeleteNetworkCameraWirelessProfile**](CameraApi.md#DeleteNetworkCameraWirelessProfile) | **Delete** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Delete an existing camera wireless profile for this network.
[**DeleteOrganizationCameraCustomAnalyticsArtifact**](CameraApi.md#DeleteOrganizationCameraCustomAnalyticsArtifact) | **Delete** /organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId} | Delete Custom Analytics Artifact
[**GenerateDeviceCameraSnapshot**](CameraApi.md#GenerateDeviceCameraSnapshot) | **Post** /devices/{serial}/camera/generateSnapshot | Generate a snapshot of what the camera sees at the specified time and return a link to that image.
[**GetDeviceCameraAnalyticsLive**](CameraApi.md#GetDeviceCameraAnalyticsLive) | **Get** /devices/{serial}/camera/analytics/live | Returns live state from camera of analytics zones
[**GetDeviceCameraAnalyticsOverview**](CameraApi.md#GetDeviceCameraAnalyticsOverview) | **Get** /devices/{serial}/camera/analytics/overview | Returns an overview of aggregate analytics data for a timespan
[**GetDeviceCameraAnalyticsRecent**](CameraApi.md#GetDeviceCameraAnalyticsRecent) | **Get** /devices/{serial}/camera/analytics/recent | Returns most recent record for analytics zones
[**GetDeviceCameraAnalyticsZoneHistory**](CameraApi.md#GetDeviceCameraAnalyticsZoneHistory) | **Get** /devices/{serial}/camera/analytics/zones/{zoneId}/history | Return historical records for analytic zones
[**GetDeviceCameraAnalyticsZones**](CameraApi.md#GetDeviceCameraAnalyticsZones) | **Get** /devices/{serial}/camera/analytics/zones | Returns all configured analytic zones for this camera
[**GetDeviceCameraCustomAnalytics**](CameraApi.md#GetDeviceCameraCustomAnalytics) | **Get** /devices/{serial}/camera/customAnalytics | Return custom analytics settings for a camera
[**GetDeviceCameraQualityAndRetention**](CameraApi.md#GetDeviceCameraQualityAndRetention) | **Get** /devices/{serial}/camera/qualityAndRetention | Returns quality and retention settings for the given camera
[**GetDeviceCameraSense**](CameraApi.md#GetDeviceCameraSense) | **Get** /devices/{serial}/camera/sense | Returns sense settings for a given camera
[**GetDeviceCameraSenseObjectDetectionModels**](CameraApi.md#GetDeviceCameraSenseObjectDetectionModels) | **Get** /devices/{serial}/camera/sense/objectDetectionModels | Returns the MV Sense object detection model list for the given camera
[**GetDeviceCameraVideoLink**](CameraApi.md#GetDeviceCameraVideoLink) | **Get** /devices/{serial}/camera/videoLink | Returns video link to the specified camera
[**GetDeviceCameraVideoSettings**](CameraApi.md#GetDeviceCameraVideoSettings) | **Get** /devices/{serial}/camera/video/settings | Returns video settings for the given camera
[**GetDeviceCameraWirelessProfiles**](CameraApi.md#GetDeviceCameraWirelessProfiles) | **Get** /devices/{serial}/camera/wirelessProfiles | Returns wireless profile assigned to the given camera
[**GetNetworkCameraQualityRetentionProfile**](CameraApi.md#GetNetworkCameraQualityRetentionProfile) | **Get** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Retrieve a single quality retention profile
[**GetNetworkCameraQualityRetentionProfiles**](CameraApi.md#GetNetworkCameraQualityRetentionProfiles) | **Get** /networks/{networkId}/camera/qualityRetentionProfiles | List the quality retention profiles for this network
[**GetNetworkCameraSchedules**](CameraApi.md#GetNetworkCameraSchedules) | **Get** /networks/{networkId}/camera/schedules | Returns a list of all camera recording schedules.
[**GetNetworkCameraWirelessProfile**](CameraApi.md#GetNetworkCameraWirelessProfile) | **Get** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Retrieve a single camera wireless profile.
[**GetNetworkCameraWirelessProfiles**](CameraApi.md#GetNetworkCameraWirelessProfiles) | **Get** /networks/{networkId}/camera/wirelessProfiles | List the camera wireless profiles for this network.
[**GetOrganizationCameraCustomAnalyticsArtifact**](CameraApi.md#GetOrganizationCameraCustomAnalyticsArtifact) | **Get** /organizations/{organizationId}/camera/customAnalytics/artifacts/{artifactId} | Get Custom Analytics Artifact
[**GetOrganizationCameraCustomAnalyticsArtifacts**](CameraApi.md#GetOrganizationCameraCustomAnalyticsArtifacts) | **Get** /organizations/{organizationId}/camera/customAnalytics/artifacts | List Custom Analytics Artifacts
[**GetOrganizationCameraOnboardingStatuses**](CameraApi.md#GetOrganizationCameraOnboardingStatuses) | **Get** /organizations/{organizationId}/camera/onboarding/statuses | Fetch onboarding status of cameras
[**UpdateDeviceCameraCustomAnalytics**](CameraApi.md#UpdateDeviceCameraCustomAnalytics) | **Put** /devices/{serial}/camera/customAnalytics | Update custom analytics settings for a camera
[**UpdateDeviceCameraQualityAndRetention**](CameraApi.md#UpdateDeviceCameraQualityAndRetention) | **Put** /devices/{serial}/camera/qualityAndRetention | Update quality and retention settings for the given camera
[**UpdateDeviceCameraSense**](CameraApi.md#UpdateDeviceCameraSense) | **Put** /devices/{serial}/camera/sense | Update sense settings for the given camera
[**UpdateDeviceCameraVideoSettings**](CameraApi.md#UpdateDeviceCameraVideoSettings) | **Put** /devices/{serial}/camera/video/settings | Update video settings for the given camera
[**UpdateDeviceCameraWirelessProfiles**](CameraApi.md#UpdateDeviceCameraWirelessProfiles) | **Put** /devices/{serial}/camera/wirelessProfiles | Assign wireless profiles to the given camera
[**UpdateNetworkCameraQualityRetentionProfile**](CameraApi.md#UpdateNetworkCameraQualityRetentionProfile) | **Put** /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId} | Update an existing quality retention profile for this network.
[**UpdateNetworkCameraWirelessProfile**](CameraApi.md#UpdateNetworkCameraWirelessProfile) | **Put** /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId} | Update an existing camera wireless profile in this network.
[**UpdateOrganizationCameraOnboardingStatuses**](CameraApi.md#UpdateOrganizationCameraOnboardingStatuses) | **Put** /organizations/{organizationId}/camera/onboarding/statuses | Notify that credential handoff to camera has completed



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
    resp, r, err := apiClient.CameraApi.CreateNetworkCameraQualityRetentionProfile(context.Background(), networkId).CreateNetworkCameraQualityRetentionProfile(createNetworkCameraQualityRetentionProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.CreateNetworkCameraQualityRetentionProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkCameraQualityRetentionProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.CreateNetworkCameraQualityRetentionProfile`: %v\n", resp)
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
    networkId := "networkId_example" // string | 
    createNetworkCameraWirelessProfile := *openapiclient.NewCreateNetworkCameraWirelessProfileRequest("Name_example", *openapiclient.NewCreateNetworkCameraWirelessProfileRequestSsid()) // CreateNetworkCameraWirelessProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.CreateNetworkCameraWirelessProfile(context.Background(), networkId).CreateNetworkCameraWirelessProfile(createNetworkCameraWirelessProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.CreateNetworkCameraWirelessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkCameraWirelessProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.CreateNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkCameraWirelessProfile** | [**CreateNetworkCameraWirelessProfileRequest**](CreateNetworkCameraWirelessProfileRequest.md) |  | 

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


## CreateOrganizationCameraCustomAnalyticsArtifact

> map[string]interface{} CreateOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId).CreateOrganizationCameraCustomAnalyticsArtifact(createOrganizationCameraCustomAnalyticsArtifact).Execute()

Create custom analytics artifact



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
    createOrganizationCameraCustomAnalyticsArtifact := *openapiclient.NewCreateOrganizationCameraCustomAnalyticsArtifactRequest() // CreateOrganizationCameraCustomAnalyticsArtifactRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.CreateOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId).CreateOrganizationCameraCustomAnalyticsArtifact(createOrganizationCameraCustomAnalyticsArtifact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.CreateOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationCameraCustomAnalyticsArtifact`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.CreateOrganizationCameraCustomAnalyticsArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationCameraCustomAnalyticsArtifact** | [**CreateOrganizationCameraCustomAnalyticsArtifactRequest**](CreateOrganizationCameraCustomAnalyticsArtifactRequest.md) |  | 

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
    resp, r, err := apiClient.CameraApi.DeleteNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.DeleteNetworkCameraQualityRetentionProfile``: %v\n", err)
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
    networkId := "networkId_example" // string | 
    wirelessProfileId := "wirelessProfileId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.DeleteNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.DeleteNetworkCameraWirelessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**wirelessProfileId** | **string** |  | 

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


## DeleteOrganizationCameraCustomAnalyticsArtifact

> DeleteOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId, artifactId).Execute()

Delete Custom Analytics Artifact



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
    artifactId := "artifactId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.DeleteOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.DeleteOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**artifactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


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


## GenerateDeviceCameraSnapshot

> map[string]interface{} GenerateDeviceCameraSnapshot(ctx, serial).GenerateDeviceCameraSnapshot(generateDeviceCameraSnapshot).Execute()

Generate a snapshot of what the camera sees at the specified time and return a link to that image.



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
    generateDeviceCameraSnapshot := *openapiclient.NewGenerateDeviceCameraSnapshotRequest() // GenerateDeviceCameraSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.GenerateDeviceCameraSnapshot(context.Background(), serial).GenerateDeviceCameraSnapshot(generateDeviceCameraSnapshot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GenerateDeviceCameraSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateDeviceCameraSnapshot`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GenerateDeviceCameraSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateDeviceCameraSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateDeviceCameraSnapshot** | [**GenerateDeviceCameraSnapshotRequest**](GenerateDeviceCameraSnapshotRequest.md) |  | 

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


## GetDeviceCameraAnalyticsLive

> map[string]interface{} GetDeviceCameraAnalyticsLive(ctx, serial).Execute()

Returns live state from camera of analytics zones



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
    resp, r, err := apiClient.CameraApi.GetDeviceCameraAnalyticsLive(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraAnalyticsLive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraAnalyticsLive`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraAnalyticsLive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsLiveRequest struct via the builder pattern


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


## GetDeviceCameraAnalyticsOverview

> []map[string]interface{} GetDeviceCameraAnalyticsOverview(ctx, serial).T0(t0).T1(t1).Timespan(timespan).ObjectType(objectType).Execute()

Returns an overview of aggregate analytics data for a timespan



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
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour. (optional)
    objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.GetDeviceCameraAnalyticsOverview(context.Background(), serial).T0(t0).T1(t1).Timespan(timespan).ObjectType(objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraAnalyticsOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraAnalyticsOverview`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraAnalyticsOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 1 hour. | 
 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

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


## GetDeviceCameraAnalyticsRecent

> []map[string]interface{} GetDeviceCameraAnalyticsRecent(ctx, serial).ObjectType(objectType).Execute()

Returns most recent record for analytics zones



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
    objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.GetDeviceCameraAnalyticsRecent(context.Background(), serial).ObjectType(objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraAnalyticsRecent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraAnalyticsRecent`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraAnalyticsRecent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

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


## GetDeviceCameraAnalyticsZoneHistory

> []map[string]interface{} GetDeviceCameraAnalyticsZoneHistory(ctx, serial, zoneId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).ObjectType(objectType).Execute()

Return historical records for analytic zones



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
    zoneId := "zoneId_example" // string | 
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 hours after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60. (optional)
    objectType := "objectType_example" // string | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.GetDeviceCameraAnalyticsZoneHistory(context.Background(), serial, zoneId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).ObjectType(objectType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraAnalyticsZoneHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraAnalyticsZoneHistory`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraAnalyticsZoneHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsZoneHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 hours after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 hours. The default is 1 hour. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 60. The default is 60. | 
 **objectType** | **string** | [optional] The object type for which analytics will be retrieved. The default object type is person. The available types are [person, vehicle]. | 

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


## GetDeviceCameraAnalyticsZones

> []map[string]interface{} GetDeviceCameraAnalyticsZones(ctx, serial).Execute()

Returns all configured analytic zones for this camera



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
    resp, r, err := apiClient.CameraApi.GetDeviceCameraAnalyticsZones(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraAnalyticsZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraAnalyticsZones`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraAnalyticsZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraAnalyticsZonesRequest struct via the builder pattern


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


## GetDeviceCameraCustomAnalytics

> map[string]interface{} GetDeviceCameraCustomAnalytics(ctx, serial).Execute()

Return custom analytics settings for a camera



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
    resp, r, err := apiClient.CameraApi.GetDeviceCameraCustomAnalytics(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraCustomAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraCustomAnalytics`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraCustomAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraCustomAnalyticsRequest struct via the builder pattern


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
    resp, r, err := apiClient.CameraApi.GetDeviceCameraQualityAndRetention(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraQualityAndRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraQualityAndRetention`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraQualityAndRetention`: %v\n", resp)
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


## GetDeviceCameraSense

> map[string]interface{} GetDeviceCameraSense(ctx, serial).Execute()

Returns sense settings for a given camera



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
    resp, r, err := apiClient.CameraApi.GetDeviceCameraSense(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraSense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraSense`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraSense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraSenseRequest struct via the builder pattern


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


## GetDeviceCameraSenseObjectDetectionModels

> []map[string]interface{} GetDeviceCameraSenseObjectDetectionModels(ctx, serial).Execute()

Returns the MV Sense object detection model list for the given camera



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
    resp, r, err := apiClient.CameraApi.GetDeviceCameraSenseObjectDetectionModels(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraSenseObjectDetectionModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraSenseObjectDetectionModels`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraSenseObjectDetectionModels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraSenseObjectDetectionModelsRequest struct via the builder pattern


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


## GetDeviceCameraVideoLink

> map[string]interface{} GetDeviceCameraVideoLink(ctx, serial).Timestamp(timestamp).Execute()

Returns video link to the specified camera



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    serial := "serial_example" // string | 
    timestamp := time.Now() // time.Time | [optional] The video link will start at this time. The timestamp should be a string in ISO8601 format. If no timestamp is specified, we will assume current time. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.GetDeviceCameraVideoLink(context.Background(), serial).Timestamp(timestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraVideoLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraVideoLink`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraVideoLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraVideoLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timestamp** | **time.Time** | [optional] The video link will start at this time. The timestamp should be a string in ISO8601 format. If no timestamp is specified, we will assume current time. | 

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


## GetDeviceCameraVideoSettings

> map[string]interface{} GetDeviceCameraVideoSettings(ctx, serial).Execute()

Returns video settings for the given camera



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
    resp, r, err := apiClient.CameraApi.GetDeviceCameraVideoSettings(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraVideoSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraVideoSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraVideoSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceCameraVideoSettingsRequest struct via the builder pattern


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
    serial := "serial_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.GetDeviceCameraWirelessProfiles(context.Background(), serial).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetDeviceCameraWirelessProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceCameraWirelessProfiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetDeviceCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

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
    resp, r, err := apiClient.CameraApi.GetNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetNetworkCameraQualityRetentionProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCameraQualityRetentionProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetNetworkCameraQualityRetentionProfile`: %v\n", resp)
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
    resp, r, err := apiClient.CameraApi.GetNetworkCameraQualityRetentionProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetNetworkCameraQualityRetentionProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCameraQualityRetentionProfiles`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetNetworkCameraQualityRetentionProfiles`: %v\n", resp)
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


## GetNetworkCameraSchedules

> []map[string]interface{} GetNetworkCameraSchedules(ctx, networkId).Execute()

Returns a list of all camera recording schedules.



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
    resp, r, err := apiClient.CameraApi.GetNetworkCameraSchedules(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetNetworkCameraSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCameraSchedules`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetNetworkCameraSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkCameraSchedulesRequest struct via the builder pattern


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
    networkId := "networkId_example" // string | 
    wirelessProfileId := "wirelessProfileId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.GetNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetNetworkCameraWirelessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCameraWirelessProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**wirelessProfileId** | **string** |  | 

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
    networkId := "networkId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.GetNetworkCameraWirelessProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetNetworkCameraWirelessProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkCameraWirelessProfiles`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetNetworkCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

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


## GetOrganizationCameraCustomAnalyticsArtifact

> map[string]interface{} GetOrganizationCameraCustomAnalyticsArtifact(ctx, organizationId, artifactId).Execute()

Get Custom Analytics Artifact



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
    artifactId := "artifactId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.GetOrganizationCameraCustomAnalyticsArtifact(context.Background(), organizationId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetOrganizationCameraCustomAnalyticsArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraCustomAnalyticsArtifact`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetOrganizationCameraCustomAnalyticsArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**artifactId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraCustomAnalyticsArtifactRequest struct via the builder pattern


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


## GetOrganizationCameraCustomAnalyticsArtifacts

> []map[string]interface{} GetOrganizationCameraCustomAnalyticsArtifacts(ctx, organizationId).Execute()

List Custom Analytics Artifacts



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
    resp, r, err := apiClient.CameraApi.GetOrganizationCameraCustomAnalyticsArtifacts(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetOrganizationCameraCustomAnalyticsArtifacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraCustomAnalyticsArtifacts`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetOrganizationCameraCustomAnalyticsArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraCustomAnalyticsArtifactsRequest struct via the builder pattern


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


## GetOrganizationCameraOnboardingStatuses

> []map[string]interface{} GetOrganizationCameraOnboardingStatuses(ctx, organizationId).Serials(serials).NetworkIds(networkIds).Execute()

Fetch onboarding status of cameras



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
    serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned cameras will be filtered to only include these serials. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned cameras will be filtered to only include these networks. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.GetOrganizationCameraOnboardingStatuses(context.Background(), organizationId).Serials(serials).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.GetOrganizationCameraOnboardingStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationCameraOnboardingStatuses`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.GetOrganizationCameraOnboardingStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationCameraOnboardingStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serials** | **[]string** | A list of serial numbers. The returned cameras will be filtered to only include these serials. | 
 **networkIds** | **[]string** | A list of network IDs. The returned cameras will be filtered to only include these networks. | 

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


## UpdateDeviceCameraCustomAnalytics

> map[string]interface{} UpdateDeviceCameraCustomAnalytics(ctx, serial).UpdateDeviceCameraCustomAnalytics(updateDeviceCameraCustomAnalytics).Execute()

Update custom analytics settings for a camera



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
    updateDeviceCameraCustomAnalytics := *openapiclient.NewUpdateDeviceCameraCustomAnalyticsRequest() // UpdateDeviceCameraCustomAnalyticsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.UpdateDeviceCameraCustomAnalytics(context.Background(), serial).UpdateDeviceCameraCustomAnalytics(updateDeviceCameraCustomAnalytics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.UpdateDeviceCameraCustomAnalytics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCameraCustomAnalytics`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.UpdateDeviceCameraCustomAnalytics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraCustomAnalyticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraCustomAnalytics** | [**UpdateDeviceCameraCustomAnalyticsRequest**](UpdateDeviceCameraCustomAnalyticsRequest.md) |  | 

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
    resp, r, err := apiClient.CameraApi.UpdateDeviceCameraQualityAndRetention(context.Background(), serial).UpdateDeviceCameraQualityAndRetention(updateDeviceCameraQualityAndRetention).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.UpdateDeviceCameraQualityAndRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCameraQualityAndRetention`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.UpdateDeviceCameraQualityAndRetention`: %v\n", resp)
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


## UpdateDeviceCameraSense

> map[string]interface{} UpdateDeviceCameraSense(ctx, serial).UpdateDeviceCameraSense(updateDeviceCameraSense).Execute()

Update sense settings for the given camera



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
    updateDeviceCameraSense := *openapiclient.NewUpdateDeviceCameraSenseRequest() // UpdateDeviceCameraSenseRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.UpdateDeviceCameraSense(context.Background(), serial).UpdateDeviceCameraSense(updateDeviceCameraSense).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.UpdateDeviceCameraSense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCameraSense`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.UpdateDeviceCameraSense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraSenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraSense** | [**UpdateDeviceCameraSenseRequest**](UpdateDeviceCameraSenseRequest.md) |  | 

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


## UpdateDeviceCameraVideoSettings

> map[string]interface{} UpdateDeviceCameraVideoSettings(ctx, serial).UpdateDeviceCameraVideoSettings(updateDeviceCameraVideoSettings).Execute()

Update video settings for the given camera



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
    updateDeviceCameraVideoSettings := *openapiclient.NewUpdateDeviceCameraVideoSettingsRequest() // UpdateDeviceCameraVideoSettingsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.UpdateDeviceCameraVideoSettings(context.Background(), serial).UpdateDeviceCameraVideoSettings(updateDeviceCameraVideoSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.UpdateDeviceCameraVideoSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCameraVideoSettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.UpdateDeviceCameraVideoSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraVideoSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraVideoSettings** | [**UpdateDeviceCameraVideoSettingsRequest**](UpdateDeviceCameraVideoSettingsRequest.md) |  | 

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
    serial := "serial_example" // string | 
    updateDeviceCameraWirelessProfiles := *openapiclient.NewUpdateDeviceCameraWirelessProfilesRequest(*openapiclient.NewUpdateDeviceCameraWirelessProfilesRequestIds()) // UpdateDeviceCameraWirelessProfilesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.UpdateDeviceCameraWirelessProfiles(context.Background(), serial).UpdateDeviceCameraWirelessProfiles(updateDeviceCameraWirelessProfiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.UpdateDeviceCameraWirelessProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceCameraWirelessProfiles`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.UpdateDeviceCameraWirelessProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceCameraWirelessProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceCameraWirelessProfiles** | [**UpdateDeviceCameraWirelessProfilesRequest**](UpdateDeviceCameraWirelessProfilesRequest.md) |  | 

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
    resp, r, err := apiClient.CameraApi.UpdateNetworkCameraQualityRetentionProfile(context.Background(), networkId, qualityRetentionProfileId).UpdateNetworkCameraQualityRetentionProfile(updateNetworkCameraQualityRetentionProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.UpdateNetworkCameraQualityRetentionProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCameraQualityRetentionProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.UpdateNetworkCameraQualityRetentionProfile`: %v\n", resp)
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
    networkId := "networkId_example" // string | 
    wirelessProfileId := "wirelessProfileId_example" // string | 
    updateNetworkCameraWirelessProfile := *openapiclient.NewUpdateNetworkCameraWirelessProfileRequest() // UpdateNetworkCameraWirelessProfileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.UpdateNetworkCameraWirelessProfile(context.Background(), networkId, wirelessProfileId).UpdateNetworkCameraWirelessProfile(updateNetworkCameraWirelessProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.UpdateNetworkCameraWirelessProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkCameraWirelessProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.UpdateNetworkCameraWirelessProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**wirelessProfileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkCameraWirelessProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkCameraWirelessProfile** | [**UpdateNetworkCameraWirelessProfileRequest**](UpdateNetworkCameraWirelessProfileRequest.md) |  | 

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


## UpdateOrganizationCameraOnboardingStatuses

> map[string]interface{} UpdateOrganizationCameraOnboardingStatuses(ctx, organizationId).UpdateOrganizationCameraOnboardingStatuses(updateOrganizationCameraOnboardingStatuses).Execute()

Notify that credential handoff to camera has completed



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
    updateOrganizationCameraOnboardingStatuses := *openapiclient.NewUpdateOrganizationCameraOnboardingStatusesRequest() // UpdateOrganizationCameraOnboardingStatusesRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CameraApi.UpdateOrganizationCameraOnboardingStatuses(context.Background(), organizationId).UpdateOrganizationCameraOnboardingStatuses(updateOrganizationCameraOnboardingStatuses).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CameraApi.UpdateOrganizationCameraOnboardingStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationCameraOnboardingStatuses`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CameraApi.UpdateOrganizationCameraOnboardingStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationCameraOnboardingStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationCameraOnboardingStatuses** | [**UpdateOrganizationCameraOnboardingStatusesRequest**](UpdateOrganizationCameraOnboardingStatusesRequest.md) |  | 

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

