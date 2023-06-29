# \InsightApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationInsightMonitoredMediaServer**](InsightApi.md#CreateOrganizationInsightMonitoredMediaServer) | **Post** /organizations/{organizationId}/insight/monitoredMediaServers | Add a media server to be monitored for this organization
[**DeleteOrganizationInsightMonitoredMediaServer**](InsightApi.md#DeleteOrganizationInsightMonitoredMediaServer) | **Delete** /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} | Delete a monitored media server from this organization
[**GetNetworkInsightApplicationHealthByTime**](InsightApi.md#GetNetworkInsightApplicationHealthByTime) | **Get** /networks/{networkId}/insight/applications/{applicationId}/healthByTime | Get application health by time
[**GetOrganizationInsightApplications**](InsightApi.md#GetOrganizationInsightApplications) | **Get** /organizations/{organizationId}/insight/applications | List all Insight tracked applications
[**GetOrganizationInsightMonitoredMediaServer**](InsightApi.md#GetOrganizationInsightMonitoredMediaServer) | **Get** /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} | Return a monitored media server for this organization
[**GetOrganizationInsightMonitoredMediaServers**](InsightApi.md#GetOrganizationInsightMonitoredMediaServers) | **Get** /organizations/{organizationId}/insight/monitoredMediaServers | List the monitored media servers for this organization
[**UpdateOrganizationInsightMonitoredMediaServer**](InsightApi.md#UpdateOrganizationInsightMonitoredMediaServer) | **Put** /organizations/{organizationId}/insight/monitoredMediaServers/{monitoredMediaServerId} | Update a monitored media server for this organization



## CreateOrganizationInsightMonitoredMediaServer

> map[string]interface{} CreateOrganizationInsightMonitoredMediaServer(ctx, organizationId).CreateOrganizationInsightMonitoredMediaServerRequest(createOrganizationInsightMonitoredMediaServerRequest).Execute()

Add a media server to be monitored for this organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationInsightMonitoredMediaServerRequest := *openapiclient.NewCreateOrganizationInsightMonitoredMediaServerRequest("Name_example", "Address_example") // CreateOrganizationInsightMonitoredMediaServerRequest | 

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightApi.CreateOrganizationInsightMonitoredMediaServer(context.Background(), organizationId).CreateOrganizationInsightMonitoredMediaServerRequest(createOrganizationInsightMonitoredMediaServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.CreateOrganizationInsightMonitoredMediaServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationInsightMonitoredMediaServer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.CreateOrganizationInsightMonitoredMediaServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationInsightMonitoredMediaServerRequest** | [**CreateOrganizationInsightMonitoredMediaServerRequest**](CreateOrganizationInsightMonitoredMediaServerRequest.md) |  | 

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


## DeleteOrganizationInsightMonitoredMediaServer

> DeleteOrganizationInsightMonitoredMediaServer(ctx, organizationId, monitoredMediaServerId).Execute()

Delete a monitored media server from this organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    monitoredMediaServerId := "monitoredMediaServerId_example" // string | Monitored media server ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.InsightApi.DeleteOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.DeleteOrganizationInsightMonitoredMediaServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**monitoredMediaServerId** | **string** | Monitored media server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


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


## GetNetworkInsightApplicationHealthByTime

> []GetNetworkInsightApplicationHealthByTime200ResponseInner GetNetworkInsightApplicationHealthByTime(ctx, networkId, applicationId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()

Get application health by time



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    networkId := "networkId_example" // string | Network ID
    applicationId := "applicationId_example" // string | Application ID
    t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 7 days from today. (optional)
    t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. (optional)
    timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours. (optional)
    resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 3600, 86400. The default is 300. (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightApi.GetNetworkInsightApplicationHealthByTime(context.Background(), networkId, applicationId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.GetNetworkInsightApplicationHealthByTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkInsightApplicationHealthByTime`: []GetNetworkInsightApplicationHealthByTime200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.GetNetworkInsightApplicationHealthByTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**applicationId** | **string** | Application ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInsightApplicationHealthByTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 7 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 7 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days. The default is 2 hours. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 3600, 86400. The default is 300. | 

### Return type

[**[]GetNetworkInsightApplicationHealthByTime200ResponseInner**](GetNetworkInsightApplicationHealthByTime200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInsightApplications

> []GetOrganizationInsightApplications200ResponseInner GetOrganizationInsightApplications(ctx, organizationId).Execute()

List all Insight tracked applications



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightApi.GetOrganizationInsightApplications(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.GetOrganizationInsightApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInsightApplications`: []GetOrganizationInsightApplications200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.GetOrganizationInsightApplications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInsightApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationInsightApplications200ResponseInner**](GetOrganizationInsightApplications200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInsightMonitoredMediaServer

> map[string]interface{} GetOrganizationInsightMonitoredMediaServer(ctx, organizationId, monitoredMediaServerId).Execute()

Return a monitored media server for this organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    monitoredMediaServerId := "monitoredMediaServerId_example" // string | Monitored media server ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightApi.GetOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.GetOrganizationInsightMonitoredMediaServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInsightMonitoredMediaServer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.GetOrganizationInsightMonitoredMediaServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**monitoredMediaServerId** | **string** | Monitored media server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


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


## GetOrganizationInsightMonitoredMediaServers

> []GetOrganizationInsightMonitoredMediaServers200ResponseInner GetOrganizationInsightMonitoredMediaServers(ctx, organizationId).Execute()

List the monitored media servers for this organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightApi.GetOrganizationInsightMonitoredMediaServers(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.GetOrganizationInsightMonitoredMediaServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInsightMonitoredMediaServers`: []GetOrganizationInsightMonitoredMediaServers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.GetOrganizationInsightMonitoredMediaServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInsightMonitoredMediaServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetOrganizationInsightMonitoredMediaServers200ResponseInner**](GetOrganizationInsightMonitoredMediaServers200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationInsightMonitoredMediaServer

> map[string]interface{} UpdateOrganizationInsightMonitoredMediaServer(ctx, organizationId, monitoredMediaServerId).UpdateOrganizationInsightMonitoredMediaServerRequest(updateOrganizationInsightMonitoredMediaServerRequest).Execute()

Update a monitored media server for this organization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/meraki/dashboard-api-go/client"
)

func main() {
    organizationId := "organizationId_example" // string | Organization ID
    monitoredMediaServerId := "monitoredMediaServerId_example" // string | Monitored media server ID
    updateOrganizationInsightMonitoredMediaServerRequest := *openapiclient.NewUpdateOrganizationInsightMonitoredMediaServerRequest() // UpdateOrganizationInsightMonitoredMediaServerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    configuration.AddDefaultHeader("Authorization", "Bearer "+os.Getenv("MERAKI_DASHBOARD_API_KEY"))
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InsightApi.UpdateOrganizationInsightMonitoredMediaServer(context.Background(), organizationId, monitoredMediaServerId).UpdateOrganizationInsightMonitoredMediaServerRequest(updateOrganizationInsightMonitoredMediaServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InsightApi.UpdateOrganizationInsightMonitoredMediaServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationInsightMonitoredMediaServer`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `InsightApi.UpdateOrganizationInsightMonitoredMediaServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**monitoredMediaServerId** | **string** | Monitored media server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationInsightMonitoredMediaServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationInsightMonitoredMediaServerRequest** | [**UpdateOrganizationInsightMonitoredMediaServerRequest**](UpdateOrganizationInsightMonitoredMediaServerRequest.md) |  | 

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

