# \GroupsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkFirmwareUpgradesStagedGroup**](GroupsApi.md#CreateNetworkFirmwareUpgradesStagedGroup) | **Post** /networks/{networkId}/firmwareUpgrades/staged/groups | Create a Staged Upgrade Group for a network
[**CreateOrganizationAdaptivePolicyGroup**](GroupsApi.md#CreateOrganizationAdaptivePolicyGroup) | **Post** /organizations/{organizationId}/adaptivePolicy/groups | Creates a new adaptive policy group
[**CreateOrganizationPolicyObjectsGroup**](GroupsApi.md#CreateOrganizationPolicyObjectsGroup) | **Post** /organizations/{organizationId}/policyObjects/groups | Creates a new Policy Object Group.
[**DeleteNetworkFirmwareUpgradesStagedGroup**](GroupsApi.md#DeleteNetworkFirmwareUpgradesStagedGroup) | **Delete** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Delete a Staged Upgrade Group
[**DeleteOrganizationAdaptivePolicyGroup**](GroupsApi.md#DeleteOrganizationAdaptivePolicyGroup) | **Delete** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Deletes the specified adaptive policy group and any associated policies and references
[**DeleteOrganizationPolicyObjectsGroup**](GroupsApi.md#DeleteOrganizationPolicyObjectsGroup) | **Delete** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Deletes a Policy Object Group.
[**GetNetworkFirmwareUpgradesStagedGroup**](GroupsApi.md#GetNetworkFirmwareUpgradesStagedGroup) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Get a Staged Upgrade Group from a network
[**GetNetworkFirmwareUpgradesStagedGroups**](GroupsApi.md#GetNetworkFirmwareUpgradesStagedGroups) | **Get** /networks/{networkId}/firmwareUpgrades/staged/groups | List of Staged Upgrade Groups in a network
[**GetOrganizationAdaptivePolicyGroup**](GroupsApi.md#GetOrganizationAdaptivePolicyGroup) | **Get** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Returns an adaptive policy group
[**GetOrganizationAdaptivePolicyGroups**](GroupsApi.md#GetOrganizationAdaptivePolicyGroups) | **Get** /organizations/{organizationId}/adaptivePolicy/groups | List adaptive policy groups in a organization
[**GetOrganizationPolicyObjectsGroup**](GroupsApi.md#GetOrganizationPolicyObjectsGroup) | **Get** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Shows details of a Policy Object Group.
[**GetOrganizationPolicyObjectsGroups**](GroupsApi.md#GetOrganizationPolicyObjectsGroups) | **Get** /organizations/{organizationId}/policyObjects/groups | Lists Policy Object Groups belonging to the organization.
[**UpdateNetworkFirmwareUpgradesStagedGroup**](GroupsApi.md#UpdateNetworkFirmwareUpgradesStagedGroup) | **Put** /networks/{networkId}/firmwareUpgrades/staged/groups/{groupId} | Update a Staged Upgrade Group for a network
[**UpdateOrganizationAdaptivePolicyGroup**](GroupsApi.md#UpdateOrganizationAdaptivePolicyGroup) | **Put** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Updates an adaptive policy group
[**UpdateOrganizationPolicyObjectsGroup**](GroupsApi.md#UpdateOrganizationPolicyObjectsGroup) | **Put** /organizations/{organizationId}/policyObjects/groups/{policyObjectGroupId} | Updates a Policy Object Group.



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
    resp, r, err := apiClient.GroupsApi.CreateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId).CreateNetworkFirmwareUpgradesStagedGroup(createNetworkFirmwareUpgradesStagedGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkFirmwareUpgradesStagedGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
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


## CreateOrganizationAdaptivePolicyGroup

> map[string]interface{} CreateOrganizationAdaptivePolicyGroup(ctx, organizationId).CreateOrganizationAdaptivePolicyGroup(createOrganizationAdaptivePolicyGroup).Execute()

Creates a new adaptive policy group



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
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationAdaptivePolicyGroup := *openapiclient.NewInlineObject173("Name_example", int32(123)) // InlineObject173 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.CreateOrganizationAdaptivePolicyGroup(context.Background(), organizationId).CreateOrganizationAdaptivePolicyGroup(createOrganizationAdaptivePolicyGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdaptivePolicyGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdaptivePolicyGroup** | [**InlineObject173**](InlineObject173.md) |  | 

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


## CreateOrganizationPolicyObjectsGroup

> map[string]interface{} CreateOrganizationPolicyObjectsGroup(ctx, organizationId).CreateOrganizationPolicyObjectsGroup(createOrganizationPolicyObjectsGroup).Execute()

Creates a new Policy Object Group.



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
    organizationId := "organizationId_example" // string | Organization ID
    createOrganizationPolicyObjectsGroup := *openapiclient.NewInlineObject214("Name_example") // InlineObject214 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.CreateOrganizationPolicyObjectsGroup(context.Background(), organizationId).CreateOrganizationPolicyObjectsGroup(createOrganizationPolicyObjectsGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.CreateOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationPolicyObjectsGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.CreateOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationPolicyObjectsGroup** | [**InlineObject214**](InlineObject214.md) |  | 

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
    resp, r, err := apiClient.GroupsApi.DeleteNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
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


## DeleteOrganizationAdaptivePolicyGroup

> DeleteOrganizationAdaptivePolicyGroup(ctx, organizationId, id).Execute()

Deletes the specified adaptive policy group and any associated policies and references



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
    organizationId := "organizationId_example" // string | Organization ID
    id := "id_example" // string | Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.DeleteOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


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


## DeleteOrganizationPolicyObjectsGroup

> DeleteOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).Execute()

Deletes a Policy Object Group.



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
    organizationId := "organizationId_example" // string | Organization ID
    policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.DeleteOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.DeleteOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPolicyObjectsGroupRequest struct via the builder pattern


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
    resp, r, err := apiClient.GroupsApi.GetNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedGroup`: InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
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
    resp, r, err := apiClient.GroupsApi.GetNetworkFirmwareUpgradesStagedGroups(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetNetworkFirmwareUpgradesStagedGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkFirmwareUpgradesStagedGroups`: []InlineResponse20030
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetNetworkFirmwareUpgradesStagedGroups`: %v\n", resp)
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


## GetOrganizationAdaptivePolicyGroup

> map[string]interface{} GetOrganizationAdaptivePolicyGroup(ctx, organizationId, id).Execute()

Returns an adaptive policy group



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
    organizationId := "organizationId_example" // string | Organization ID
    id := "id_example" // string | Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


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


## GetOrganizationAdaptivePolicyGroups

> []map[string]interface{} GetOrganizationAdaptivePolicyGroups(ctx, organizationId).Execute()

List adaptive policy groups in a organization



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
    organizationId := "organizationId_example" // string | Organization ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetOrganizationAdaptivePolicyGroups(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetOrganizationAdaptivePolicyGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyGroups`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetOrganizationAdaptivePolicyGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyGroupsRequest struct via the builder pattern


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


## GetOrganizationPolicyObjectsGroup

> map[string]interface{} GetOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).Execute()

Shows details of a Policy Object Group.



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
    organizationId := "organizationId_example" // string | Organization ID
    policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPolicyObjectsGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsGroupRequest struct via the builder pattern


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


## GetOrganizationPolicyObjectsGroups

> []map[string]interface{} GetOrganizationPolicyObjectsGroups(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

Lists Policy Object Groups belonging to the organization.



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
    organizationId := "organizationId_example" // string | Organization ID
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 10 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.GetOrganizationPolicyObjectsGroups(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.GetOrganizationPolicyObjectsGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationPolicyObjectsGroups`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.GetOrganizationPolicyObjectsGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationPolicyObjectsGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 10 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 

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
    resp, r, err := apiClient.GroupsApi.UpdateNetworkFirmwareUpgradesStagedGroup(context.Background(), networkId, groupId).UpdateNetworkFirmwareUpgradesStagedGroup(updateNetworkFirmwareUpgradesStagedGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateNetworkFirmwareUpgradesStagedGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkFirmwareUpgradesStagedGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateNetworkFirmwareUpgradesStagedGroup`: %v\n", resp)
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


## UpdateOrganizationAdaptivePolicyGroup

> map[string]interface{} UpdateOrganizationAdaptivePolicyGroup(ctx, organizationId, id).UpdateOrganizationAdaptivePolicyGroup(updateOrganizationAdaptivePolicyGroup).Execute()

Updates an adaptive policy group



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
    organizationId := "organizationId_example" // string | Organization ID
    id := "id_example" // string | Id
    updateOrganizationAdaptivePolicyGroup := *openapiclient.NewInlineObject174() // InlineObject174 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.UpdateOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).UpdateOrganizationAdaptivePolicyGroup(updateOrganizationAdaptivePolicyGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdaptivePolicyGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyGroup** | [**InlineObject174**](InlineObject174.md) |  | 

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


## UpdateOrganizationPolicyObjectsGroup

> map[string]interface{} UpdateOrganizationPolicyObjectsGroup(ctx, organizationId, policyObjectGroupId).UpdateOrganizationPolicyObjectsGroup(updateOrganizationPolicyObjectsGroup).Execute()

Updates a Policy Object Group.



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
    organizationId := "organizationId_example" // string | Organization ID
    policyObjectGroupId := "policyObjectGroupId_example" // string | Policy object group ID
    updateOrganizationPolicyObjectsGroup := *openapiclient.NewInlineObject215() // InlineObject215 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupsApi.UpdateOrganizationPolicyObjectsGroup(context.Background(), organizationId, policyObjectGroupId).UpdateOrganizationPolicyObjectsGroup(updateOrganizationPolicyObjectsGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupsApi.UpdateOrganizationPolicyObjectsGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationPolicyObjectsGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `GroupsApi.UpdateOrganizationPolicyObjectsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**policyObjectGroupId** | **string** | Policy object group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationPolicyObjectsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationPolicyObjectsGroup** | [**InlineObject215**](InlineObject215.md) |  | 

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

