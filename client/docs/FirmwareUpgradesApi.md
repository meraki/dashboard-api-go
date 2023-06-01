# \FirmwareUpgradesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkFirmwareUpgradesRollback**](FirmwareUpgradesApi.md#CreateNetworkFirmwareUpgradesRollback) | **Post** /networks/{networkId}/firmwareUpgrades/rollbacks | Rollback a Firmware Upgrade For A Network
[**CreateNetworkFirmwareUpgradesStagedEvent**](FirmwareUpgradesApi.md#CreateNetworkFirmwareUpgradesStagedEvent) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events | Create a Staged Upgrade Event for a network
[**CreateNetworkFirmwareUpgradesStagedGroup**](FirmwareUpgradesApi.md#CreateNetworkFirmwareUpgradesStagedGroup) | **Post** /networks/{networkId}/firmwareUpgrades/staged/groups | Create a Staged Upgrade Group for a network
[**DeferNetworkFirmwareUpgradesStagedEvents**](FirmwareUpgradesApi.md#DeferNetworkFirmwareUpgradesStagedEvents) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events/defer | Postpone by 1 week all pending staged upgrade stages for a network
[**DeleteNetworkFirmwareUpgradesStagedGroup**](FirmwareUpgradesApi.md#DeleteNetworkFirmwareUpgradesStagedGroup) | **Delete** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Delete a Staged Upgrade Group
[**GetNetworkFirmwareUpgrades**](FirmwareUpgradesApi.md#GetNetworkFirmwareUpgrades) | **Get** /networks/{networkId}/firmwareUpgrades | Get firmware upgrade information for a network
[**GetNetworkFirmwareUpgradesStagedEvents**](FirmwareUpgradesApi.md#GetNetworkFirmwareUpgradesStagedEvents) | **Get** /networks/{networkId}/firmwareUpgrades/staged/events | Get the Staged Upgrade Event from a network
[**GetNetworkFirmwareUpgradesStagedGroup**](FirmwareUpgradesApi.md#GetNetworkFirmwareUpgradesStagedGroup) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Get a Staged Upgrade Group from a network
[**GetNetworkFirmwareUpgradesStagedGroups**](FirmwareUpgradesApi.md#GetNetworkFirmwareUpgradesStagedGroups) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups | List of Staged Upgrade Groups in a network
[**GetNetworkFirmwareUpgradesStagedStages**](FirmwareUpgradesApi.md#GetNetworkFirmwareUpgradesStagedStages) | **Get** /networks/{networkId}/firmwareUpgrades/staged/stages | Order of Staged Upgrade Groups in a network
[**RollbacksNetworkFirmwareUpgradesStagedEvents**](FirmwareUpgradesApi.md#RollbacksNetworkFirmwareUpgradesStagedEvents) | **Post** /networks/{networkId}/firmwareUpgrades/staged/events/rollbacks | Rollback a Staged Upgrade Event for a network
[**UpdateNetworkFirmwareUpgrades**](FirmwareUpgradesApi.md#UpdateNetworkFirmwareUpgrades) | **Put** /networks/{networkId}/firmwareUpgrades | Update firmware upgrade information for a network
[**UpdateNetworkFirmwareUpgradesStagedEvents**](FirmwareUpgradesApi.md#UpdateNetworkFirmwareUpgradesStagedEvents) | **Put** /networks/{networkId}/firmwareUpgrades/staged/events | Update the Staged Upgrade Event for a network
[**UpdateNetworkFirmwareUpgradesStagedGroup**](FirmwareUpgradesApi.md#UpdateNetworkFirmwareUpgradesStagedGroup) | **Put** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Update a Staged Upgrade Group for a network
[**UpdateNetworkFirmwareUpgradesStagedStages**](FirmwareUpgradesApi.md#UpdateNetworkFirmwareUpgradesStagedStages) | **Put** /networks/{networkId}/firmwareUpgrades/staged/stages | Assign Staged Upgrade Group order in the sequence.



## CreateNetworkFirmwareUpgradesRollback

> InlineResponse20028 CreateNetworkFirmwareUpgradesRollback(ctx, networkId).CreateNetworkFirmwareUpgradesRollback(createNetworkFirmwareUpgradesRollback).Execute()

Rollback a Firmware Upgrade For A Network



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
    createNetworkFirmwareUpgradesRollback := *openapiclient.NewInlineObject78([]openapiclient.NetworksNetworkIdFirmwareUpgradesRollbacksReasons{*openapiclient.NewNetworksNetworkIdFirmwareUpgradesRollbacksReasons("Category_example", "Comment_example")}) // InlineObject78 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesRollback(context.Background(), networkId).CreateNetworkFirmwareUpgradesRollback(createNetworkFirmwareUpgradesRollback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesRollback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesRollback`: InlineResponse20028
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesRollback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkFirmwareUpgradesRollbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkFirmwareUpgradesRollback** | [**InlineObject78**](InlineObject78.md) |  | 

### Return type

[**InlineResponse20028**](InlineResponse20028.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkFirmwareUpgradesStagedEvent

> InlineResponse20029 CreateNetworkFirmwareUpgradesStagedEvent(ctx, networkId).CreateNetworkFirmwareUpgradesStagedEvent(createNetworkFirmwareUpgradesStagedEvent).Execute()

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
    createNetworkFirmwareUpgradesStagedEvent := *openapiclient.NewInlineObject80([]openapiclient.NetworksNetworkIdFirmwareUpgradesStagedEventsStages{*openapiclient.NewNetworksNetworkIdFirmwareUpgradesStagedEventsStages()}) // InlineObject80 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesStagedEvent(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedEvent(createNetworkFirmwareUpgradesStagedEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesStagedEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesStagedEvent`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesStagedEvent`: %v\n", resp)
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

 **createNetworkFirmwareUpgradesStagedEvent** | [**InlineObject80**](InlineObject80.md) |  | 

### Return type

[**InlineResponse20029**](InlineResponse20029.md)

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
    createNetworkFirmwareUpgradesStagedGroup := *openapiclient.NewInlineObject82("Name_example", false) // InlineObject82 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedGroup(createNetworkFirmwareUpgradesStagedGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesStagedGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.CreateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
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

 **createNetworkFirmwareUpgradesStagedGroup** | [**InlineObject82**](InlineObject82.md) |  | 

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

> InlineResponse20029 DeferNetworkFirmwareUpgradesStagedEvents(ctx, networkId).Execute()

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
    resp, r, err := apiClient.FirmwareUpgradesApi.DeferNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.DeferNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeferNetworkFirmwareUpgradesStagedEvents`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.DeferNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
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

[**InlineResponse20029**](InlineResponse20029.md)

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
    resp, r, err := apiClient.FirmwareUpgradesApi.DeleteNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.DeleteNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
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


## GetNetworkFirmwareUpgrades

> InlineResponse20027 GetNetworkFirmwareUpgrades(ctx, networkId).Execute()

Get firmware upgrade information for a network



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
    resp, r, err := apiClient.FirmwareUpgradesApi.GetNetworkFirmwareUpgrades(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.GetNetworkFirmwareUpgrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgrades`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.GetNetworkFirmwareUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkFirmwareUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedEvents

> InlineResponse20029 GetNetworkFirmwareUpgradesStagedEvents(ctx, networkId).Execute()

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
    resp, r, err := apiClient.FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedEvents`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
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

[**InlineResponse20029**](InlineResponse20029.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedGroup

> InlineResponse20030 GetNetworkFirmwareUpgradesStagedGroup(ctx, networkId, groupId).Execute()

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
    resp, r, err := apiClient.FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedGroup`: InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
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

[**InlineResponse20030**](InlineResponse20030.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedGroups

> []InlineResponse20030 GetNetworkFirmwareUpgradesStagedGroups(ctx, networkId).Execute()

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
    resp, r, err := apiClient.FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedGroups(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedGroups`: []InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedGroups`: %v\n", resp)
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

[**[]InlineResponse20030**](InlineResponse20030.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkFirmwareUpgradesStagedStages

> []InlineResponse20031 GetNetworkFirmwareUpgradesStagedStages(ctx, networkId).Execute()

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
    resp, r, err := apiClient.FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedStages`: []InlineResponse20031
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.GetNetworkFirmwareUpgradesStagedStages`: %v\n", resp)
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

[**[]InlineResponse20031**](InlineResponse20031.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbacksNetworkFirmwareUpgradesStagedEvents

> InlineResponse20029 RollbacksNetworkFirmwareUpgradesStagedEvents(ctx, networkId).RollbacksNetworkFirmwareUpgradesStagedEvents(rollbacksNetworkFirmwareUpgradesStagedEvents).Execute()

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
    rollbacksNetworkFirmwareUpgradesStagedEvents := *openapiclient.NewInlineObject81([]openapiclient.NetworksNetworkIdFirmwareUpgradesStagedEventsStages{*openapiclient.NewNetworksNetworkIdFirmwareUpgradesStagedEventsStages()}) // InlineObject81 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.RollbacksNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).RollbacksNetworkFirmwareUpgradesStagedEvents(rollbacksNetworkFirmwareUpgradesStagedEvents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.RollbacksNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbacksNetworkFirmwareUpgradesStagedEvents`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.RollbacksNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
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

 **rollbacksNetworkFirmwareUpgradesStagedEvents** | [**InlineObject81**](InlineObject81.md) |  | 

### Return type

[**InlineResponse20029**](InlineResponse20029.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgrades

> InlineResponse20027 UpdateNetworkFirmwareUpgrades(ctx, networkId).UpdateNetworkFirmwareUpgrades(updateNetworkFirmwareUpgrades).Execute()

Update firmware upgrade information for a network



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
    updateNetworkFirmwareUpgrades := *openapiclient.NewInlineObject77() // InlineObject77 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.UpdateNetworkFirmwareUpgrades(context.Background(), networkId).UpdateNetworkFirmwareUpgrades(updateNetworkFirmwareUpgrades).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.UpdateNetworkFirmwareUpgrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgrades`: InlineResponse20027
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.UpdateNetworkFirmwareUpgrades`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkFirmwareUpgradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkFirmwareUpgrades** | [**InlineObject77**](InlineObject77.md) |  | 

### Return type

[**InlineResponse20027**](InlineResponse20027.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkFirmwareUpgradesStagedEvents

> InlineResponse20029 UpdateNetworkFirmwareUpgradesStagedEvents(ctx, networkId).UpdateNetworkFirmwareUpgradesStagedEvents(updateNetworkFirmwareUpgradesStagedEvents).Execute()

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
    updateNetworkFirmwareUpgradesStagedEvents := *openapiclient.NewInlineObject79([]openapiclient.NetworksNetworkIdFirmwareUpgradesStagedEventsStages{*openapiclient.NewNetworksNetworkIdFirmwareUpgradesStagedEventsStages()}) // InlineObject79 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.UpdateNetworkFirmwareUpgradesStagedEvents(context.Background(), networkId).UpdateNetworkFirmwareUpgradesStagedEvents(updateNetworkFirmwareUpgradesStagedEvents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.UpdateNetworkFirmwareUpgradesStagedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgradesStagedEvents`: InlineResponse20029
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.UpdateNetworkFirmwareUpgradesStagedEvents`: %v\n", resp)
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

 **updateNetworkFirmwareUpgradesStagedEvents** | [**InlineObject79**](InlineObject79.md) |  | 

### Return type

[**InlineResponse20029**](InlineResponse20029.md)

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
    updateNetworkFirmwareUpgradesStagedGroup := *openapiclient.NewInlineObject83("Name_example", false) // InlineObject83 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.UpdateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).UpdateNetworkFirmwareUpgradesStagedGroup(updateNetworkFirmwareUpgradesStagedGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.UpdateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgradesStagedGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.UpdateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
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


 **updateNetworkFirmwareUpgradesStagedGroup** | [**InlineObject83**](InlineObject83.md) |  | 

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

> []InlineResponse20031 UpdateNetworkFirmwareUpgradesStagedStages(ctx, networkId).UpdateNetworkFirmwareUpgradesStagedStages(updateNetworkFirmwareUpgradesStagedStages).Execute()

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
    updateNetworkFirmwareUpgradesStagedStages := *openapiclient.NewInlineObject84() // InlineObject84 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirmwareUpgradesApi.UpdateNetworkFirmwareUpgradesStagedStages(context.Background(), networkId).UpdateNetworkFirmwareUpgradesStagedStages(updateNetworkFirmwareUpgradesStagedStages).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirmwareUpgradesApi.UpdateNetworkFirmwareUpgradesStagedStages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgradesStagedStages`: []InlineResponse20031
    fmt.Fprintf(os.Stdout, "Response from `FirmwareUpgradesApi.UpdateNetworkFirmwareUpgradesStagedStages`: %v\n", resp)
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

 **updateNetworkFirmwareUpgradesStagedStages** | [**InlineObject84**](InlineObject84.md) |  | 

### Return type

[**[]InlineResponse20031**](InlineResponse20031.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

