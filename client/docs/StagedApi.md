# \StagedApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkFirmwareUpgradesStagedEvent**](StagedApi.md#CreateNetworkFirmwareUpgradesStagedEvent) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events | Create a Staged Upgrade Event for a network
[**CreateNetworkFirmwareUpgradesStagedGroup**](StagedApi.md#CreateNetworkFirmwareUpgradesStagedGroup) | **Post** /networks/{networkId}/firmwareUpgrades/staged/groups | Create a Staged Upgrade Group for a network
[**DeferNetworkFirmwareUpgradesStagedEvents**](StagedApi.md#DeferNetworkFirmwareUpgradesStagedEvents) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events/defer | Postpone by 1 week all pending staged upgrade stages for a network
[**DeleteNetworkFirmwareUpgradesStagedGroup**](StagedApi.md#DeleteNetworkFirmwareUpgradesStagedGroup) | **Delete** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Delete a Staged Upgrade Group
[**GetNetworkFirmwareUpgradesStagedEvents**](StagedApi.md#GetNetworkFirmwareUpgradesStagedEvents) | **Get** /networks/{networkId}/firmwareUpgrades/staged/events | Get the Staged Upgrade Event from a network
[**GetNetworkFirmwareUpgradesStagedGroup**](StagedApi.md#GetNetworkFirmwareUpgradesStagedGroup) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Get a Staged Upgrade Group from a network
[**GetNetworkFirmwareUpgradesStagedGroups**](StagedApi.md#GetNetworkFirmwareUpgradesStagedGroups) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups | List of Staged Upgrade Groups in a network
[**GetNetworkFirmwareUpgradesStagedStages**](StagedApi.md#GetNetworkFirmwareUpgradesStagedStages) | **Get** /networks/{networkId}/firmwareUpgrades/staged/stages | Order of Staged Upgrade Groups in a network
[**RollbacksNetworkFirmwareUpgradesStagedEvents**](StagedApi.md#RollbacksNetworkFirmwareUpgradesStagedEvents) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events/rollbacks | Rollback a Staged Upgrade Event for a network
[**UpdateNetworkFirmwareUpgradesStagedEvents**](StagedApi.md#UpdateNetworkFirmwareUpgradesStagedEvents) | **Put** /networks/{networkId}/firmwareUpgrades/staged/events | Update the Staged Upgrade Event for a network
[**UpdateNetworkFirmwareUpgradesStagedGroup**](StagedApi.md#UpdateNetworkFirmwareUpgradesStagedGroup) | **Put** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Update a Staged Upgrade Group for a network
[**UpdateNetworkFirmwareUpgradesStagedStages**](StagedApi.md#UpdateNetworkFirmwareUpgradesStagedStages) | **Put** /networks/{networkId}/firmwareUpgrades/staged/stages | Assign Staged Upgrade Group order in the sequence.



## CreateNetworkFirmwareUpgradesStagedEvent

> InlineResponse20032 CreateNetworkFirmwareUpgradesStagedEvent(ctx, networkId).CreateNetworkFirmwareUpgradesStagedEvent(createNetworkFirmwareUpgradesStagedEvent).Execute()

Create a Staged Upgrade Event for a network



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
    createNetworkFirmwareUpgradesStagedEvent := *openapiclient.NewInlineObject83([]openapiclient.NetworksNetworkIdFirmwareUpgradesStagedEventsStages{*openapiclient.NewNetworksNetworkIdFirmwareUpgradesStagedEventsStages()}) // InlineObject83 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagedApi.CreateNetworkFirmwareUpgradesStagedEvent(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedEvent(createNetworkFirmwareUpgradesStagedEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.CreateNetworkFirmwareUpgradesStagedEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesStagedEvent`: InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.CreateNetworkFirmwareUpgradesStagedEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesStagedEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesStagedEvent** | [**InlineObject83**](InlineObject83.md) |  | 

### Return type

[**InlineResponse20032**](InlineResponse20032.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirmwareUpgradesStagedGroup

> map[string]interface{} CreateNetworkFirmwareUpgradesStagedGroup(ctx, networkId).CreateNetworkFirmwareUpgradesStagedGroup(createNetworkFirmwareUpgradesStagedGroup).Execute()

Create a Staged Upgrade Group for a network



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
    createNetworkFirmwareUpgradesStagedGroup := *openapiclient.NewInlineObject85("Name_example", false) // InlineObject85 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagedApi.CreateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedGroup(createNetworkFirmwareUpgradesStagedGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.CreateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesStagedGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.CreateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesStagedGroup** | [**InlineObject85**](InlineObject85.md) |  | 

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


## DeferNetworkFirmwareUpgradesStagedEvents

> InlineResponse20032 DeferNetworkFirmwareUpgradesStagedEvents(ctx, networkId).Execute()

Postpone by 1 week all pending staged upgrade stages for a network



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
    resp, r, err := apiClient.StagedApi.DeferNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.DeferNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeferNetworkFirmwareUpgradesStagedEvents`: InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.DeferNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeferNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20032**](InlineResponse20032.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkFirmwareUpgradesStagedGroup

> DeleteNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).Execute()

Delete a Staged Upgrade Group



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
    groupId := "groupId_example" // string | Group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagedApi.DeleteNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.DeleteNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


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


## GetNetworkFirmwareUpgradesStagedEvents

> InlineResponse20032 GetNetworkFirmwareUpgradesStagedEvents(ctx, networkId).Execute()

Get the Staged Upgrade Event from a network



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
    resp, r, err := apiClient.StagedApi.GetNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.GetNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedEvents`: InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.GetNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20032**](InlineResponse20032.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedGroup

> InlineResponse20033 GetNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).Execute()

Get a Staged Upgrade Group from a network



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
    groupId := "groupId_example" // string | Group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagedApi.GetNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.GetNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedGroup`: InlineResponse20033
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.GetNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse20033**](InlineResponse20033.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedGroups

> []InlineResponse20033 GetNetworkFirmwareUpgradesStagedGroups(ctx, networkId).Execute()

List of Staged Upgrade Groups in a network



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
    resp, r, err := apiClient.StagedApi.GetNetworkFirmwareUpgradesStagedGroups(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.GetNetworkFirmwareUpgradesStagedGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedGroups`: []InlineResponse20033
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.GetNetworkFirmwareUpgradesStagedGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20033**](InlineResponse20033.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedStages

> []InlineResponse20034 GetNetworkFirmwareUpgradesStagedStages(ctx, networkId).Execute()

Order of Staged Upgrade Groups in a network



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
    resp, r, err := apiClient.StagedApi.GetNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.GetNetworkFirmwareUpgradesStagedStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedStages`: []InlineResponse20034
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.GetNetworkFirmwareUpgradesStagedStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesStagedStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse20034**](InlineResponse20034.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbacksNetworkFirmwareUpgradesStagedEvents

> InlineResponse20032 RollbacksNetworkFirmwareUpgradesStagedEvents(ctx, networkId).RollbacksNetworkFirmwareUpgradesStagedEvents(rollbacksNetworkFirmwareUpgradesStagedEvents).Execute()

Rollback a Staged Upgrade Event for a network



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
    rollbacksNetworkFirmwareUpgradesStagedEvents := *openapiclient.NewInlineObject84([]openapiclient.NetworksNetworkIdFirmwareUpgradesStagedEventsStages{*openapiclient.NewNetworksNetworkIdFirmwareUpgradesStagedEventsStages()}) // InlineObject84 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagedApi.RollbacksNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).RollbacksNetworkFirmwareUpgradesStagedEvents(rollbacksNetworkFirmwareUpgradesStagedEvents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.RollbacksNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbacksNetworkFirmwareUpgradesStagedEvents`: InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.RollbacksNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbacksNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rollbacksNetworkFirmwareUpgradesStagedEvents** | [**InlineObject84**](InlineObject84.md) |  | 

### Return type

[**InlineResponse20032**](InlineResponse20032.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedEvents

> InlineResponse20032 UpdateNetworkFirmwareUpgradesStagedEvents(ctx, networkId).UpdateNetworkFirmwareUpgradesStagedEvents(updateNetworkFirmwareUpgradesStagedEvents).Execute()

Update the Staged Upgrade Event for a network



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
    updateNetworkFirmwareUpgradesStagedEvents := *openapiclient.NewInlineObject82([]openapiclient.NetworksNetworkIdFirmwareUpgradesStagedEventsStages{*openapiclient.NewNetworksNetworkIdFirmwareUpgradesStagedEventsStages()}) // InlineObject82 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagedApi.UpdateNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).UpdateNetworkFirmwareUpgradesStagedEvents(updateNetworkFirmwareUpgradesStagedEvents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.UpdateNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgradesStagedEvents`: InlineResponse20032
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.UpdateNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgradesStagedEvents** | [**InlineObject82**](InlineObject82.md) |  | 

### Return type

[**InlineResponse20032**](InlineResponse20032.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedGroup

> map[string]interface{} UpdateNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).UpdateNetworkFirmwareUpgradesStagedGroup(updateNetworkFirmwareUpgradesStagedGroup).Execute()

Update a Staged Upgrade Group for a network



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
    groupId := "groupId_example" // string | Group ID
    updateNetworkFirmwareUpgradesStagedGroup := *openapiclient.NewInlineObject86("Name_example", false) // InlineObject86 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagedApi.UpdateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).UpdateNetworkFirmwareUpgradesStagedGroup(updateNetworkFirmwareUpgradesStagedGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.UpdateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgradesStagedGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.UpdateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**groupId** | **string** | Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkFirmwareUpgradesStagedGroup** | [**InlineObject86**](InlineObject86.md) |  | 

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


## UpdateNetworkFirmwareUpgradesStagedStages

> []InlineResponse20034 UpdateNetworkFirmwareUpgradesStagedStages(ctx, networkId).UpdateNetworkFirmwareUpgradesStagedStages(updateNetworkFirmwareUpgradesStagedStages).Execute()

Assign Staged Upgrade Group order in the sequence.



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
    updateNetworkFirmwareUpgradesStagedStages := *openapiclient.NewInlineObject87() // InlineObject87 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StagedApi.UpdateNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).UpdateNetworkFirmwareUpgradesStagedStages(updateNetworkFirmwareUpgradesStagedStages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StagedApi.UpdateNetworkFirmwareUpgradesStagedStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgradesStagedStages`: []InlineResponse20034
    fmt.Fprintf(os.Stdout, "Response from `StagedApi.UpdateNetworkFirmwareUpgradesStagedStages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesStagedStagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgradesStagedStages** | [**InlineObject87**](InlineObject87.md) |  | 

### Return type

[**[]InlineResponse20034**](InlineResponse20034.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

