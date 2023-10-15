# \VlanProfilesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkVlanProfile**](VlanProfilesApi.md#CreateNetworkVlanProfile) | **Post** /networks/{networkId}/vlanProfiles | Create a VLAN profile for a network
[**DeleteNetworkVlanProfile**](VlanProfilesApi.md#DeleteNetworkVlanProfile) | **Delete** /networks/{networkId}/vlanProfiles/{iname} | Delete a VLAN profile of a network
[**GetNetworkVlanProfile**](VlanProfilesApi.md#GetNetworkVlanProfile) | **Get** /networks/{networkId}/vlanProfiles/{iname} | Get an existing VLAN profile of a network
[**GetNetworkVlanProfiles**](VlanProfilesApi.md#GetNetworkVlanProfiles) | **Get** /networks/{networkId}/vlanProfiles | List VLAN profiles for a network
[**GetNetworkVlanProfilesAssignmentsByDevice**](VlanProfilesApi.md#GetNetworkVlanProfilesAssignmentsByDevice) | **Get** /networks/{networkId}/vlanProfiles/assignments/byDevice | Get the assigned VLAN Profiles for devices in a network
[**ReassignNetworkVlanProfilesAssignments**](VlanProfilesApi.md#ReassignNetworkVlanProfilesAssignments) | **Post** /networks/{networkId}/vlanProfiles/assignments/reassign | Update the assigned VLAN Profile for devices in a network
[**UpdateNetworkVlanProfile**](VlanProfilesApi.md#UpdateNetworkVlanProfile) | **Put** /networks/{networkId}/vlanProfiles/{iname} | Update an existing VLAN profile of a network



## CreateNetworkVlanProfile

> GetNetworkVlanProfiles200ResponseInner CreateNetworkVlanProfile(ctx, networkId).CreateNetworkVlanProfileRequest(createNetworkVlanProfileRequest).Execute()

Create a VLAN profile for a network



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
    createNetworkVlanProfileRequest := *openapiclient.NewCreateNetworkVlanProfileRequest("Name_example", []openapiclient.CreateNetworkVlanProfileRequestVlanNamesInner{*openapiclient.NewCreateNetworkVlanProfileRequestVlanNamesInner("Name_example", "VlanId_example")}, []openapiclient.CreateNetworkVlanProfileRequestVlanGroupsInner{*openapiclient.NewCreateNetworkVlanProfileRequestVlanGroupsInner("Name_example", "VlanIds_example")}, "Iname_example") // CreateNetworkVlanProfileRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlanProfilesApi.CreateNetworkVlanProfile(context.Background(), networkId).CreateNetworkVlanProfileRequest(createNetworkVlanProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanProfilesApi.CreateNetworkVlanProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkVlanProfile`: GetNetworkVlanProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `VlanProfilesApi.CreateNetworkVlanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkVlanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkVlanProfileRequest** | [**CreateNetworkVlanProfileRequest**](CreateNetworkVlanProfileRequest.md) |  | 

### Return type

[**GetNetworkVlanProfiles200ResponseInner**](GetNetworkVlanProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkVlanProfile

> DeleteNetworkVlanProfile(ctx, networkId, iname).Execute()

Delete a VLAN profile of a network



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
    iname := "iname_example" // string | Iname

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VlanProfilesApi.DeleteNetworkVlanProfile(context.Background(), networkId, iname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanProfilesApi.DeleteNetworkVlanProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**iname** | **string** | Iname | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkVlanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkVlanProfile

> GetNetworkVlanProfiles200ResponseInner GetNetworkVlanProfile(ctx, networkId, iname).Execute()

Get an existing VLAN profile of a network



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
    iname := "iname_example" // string | Iname

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlanProfilesApi.GetNetworkVlanProfile(context.Background(), networkId, iname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanProfilesApi.GetNetworkVlanProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkVlanProfile`: GetNetworkVlanProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `VlanProfilesApi.GetNetworkVlanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**iname** | **string** | Iname | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkVlanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkVlanProfiles200ResponseInner**](GetNetworkVlanProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkVlanProfiles

> []GetNetworkVlanProfiles200ResponseInner GetNetworkVlanProfiles(ctx, networkId).Execute()

List VLAN profiles for a network



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

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlanProfilesApi.GetNetworkVlanProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanProfilesApi.GetNetworkVlanProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkVlanProfiles`: []GetNetworkVlanProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `VlanProfilesApi.GetNetworkVlanProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkVlanProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkVlanProfiles200ResponseInner**](GetNetworkVlanProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkVlanProfilesAssignmentsByDevice

> []GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner GetNetworkVlanProfilesAssignmentsByDevice(ctx, networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Serials(serials).ProductTypes(productTypes).StackIds(stackIds).Execute()

Get the assigned VLAN Profiles for devices in a network



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    serials := []string{"Inner_example"} // []string | Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match. (optional)
    productTypes := []string{"ProductTypes_example"} // []string | Optional parameter to filter devices by product types. (optional)
    stackIds := []string{"Inner_example"} // []string | Optional parameter to filter devices by Switch Stack ids. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlanProfilesApi.GetNetworkVlanProfilesAssignmentsByDevice(context.Background(), networkId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Serials(serials).ProductTypes(productTypes).StackIds(stackIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanProfilesApi.GetNetworkVlanProfilesAssignmentsByDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkVlanProfilesAssignmentsByDevice`: []GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `VlanProfilesApi.GetNetworkVlanProfilesAssignmentsByDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkVlanProfilesAssignmentsByDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **serials** | **[]string** | Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match. | 
 **productTypes** | **[]string** | Optional parameter to filter devices by product types. | 
 **stackIds** | **[]string** | Optional parameter to filter devices by Switch Stack ids. | 

### Return type

[**[]GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner**](GetNetworkVlanProfilesAssignmentsByDevice200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReassignNetworkVlanProfilesAssignments

> ReassignNetworkVlanProfilesAssignments200Response ReassignNetworkVlanProfilesAssignments(ctx, networkId).ReassignNetworkVlanProfilesAssignmentsRequest(reassignNetworkVlanProfilesAssignmentsRequest).Execute()

Update the assigned VLAN Profile for devices in a network



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
    reassignNetworkVlanProfilesAssignmentsRequest := *openapiclient.NewReassignNetworkVlanProfilesAssignmentsRequest([]string{"Serials_example"}, []string{"StackIds_example"}) // ReassignNetworkVlanProfilesAssignmentsRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlanProfilesApi.ReassignNetworkVlanProfilesAssignments(context.Background(), networkId).ReassignNetworkVlanProfilesAssignmentsRequest(reassignNetworkVlanProfilesAssignmentsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanProfilesApi.ReassignNetworkVlanProfilesAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReassignNetworkVlanProfilesAssignments`: ReassignNetworkVlanProfilesAssignments200Response
    fmt.Fprintf(os.Stdout, "Response from `VlanProfilesApi.ReassignNetworkVlanProfilesAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReassignNetworkVlanProfilesAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reassignNetworkVlanProfilesAssignmentsRequest** | [**ReassignNetworkVlanProfilesAssignmentsRequest**](ReassignNetworkVlanProfilesAssignmentsRequest.md) |  | 

### Return type

[**ReassignNetworkVlanProfilesAssignments200Response**](ReassignNetworkVlanProfilesAssignments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkVlanProfile

> GetNetworkVlanProfiles200ResponseInner UpdateNetworkVlanProfile(ctx, networkId, iname).UpdateNetworkVlanProfileRequest(updateNetworkVlanProfileRequest).Execute()

Update an existing VLAN profile of a network



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
    iname := "iname_example" // string | Iname
    updateNetworkVlanProfileRequest := *openapiclient.NewUpdateNetworkVlanProfileRequest("Name_example", []openapiclient.CreateNetworkVlanProfileRequestVlanNamesInner{*openapiclient.NewCreateNetworkVlanProfileRequestVlanNamesInner("Name_example", "VlanId_example")}, []openapiclient.CreateNetworkVlanProfileRequestVlanGroupsInner{*openapiclient.NewCreateNetworkVlanProfileRequestVlanGroupsInner("Name_example", "VlanIds_example")}) // UpdateNetworkVlanProfileRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VlanProfilesApi.UpdateNetworkVlanProfile(context.Background(), networkId, iname).UpdateNetworkVlanProfileRequest(updateNetworkVlanProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanProfilesApi.UpdateNetworkVlanProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkVlanProfile`: GetNetworkVlanProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `VlanProfilesApi.UpdateNetworkVlanProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**iname** | **string** | Iname | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkVlanProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkVlanProfileRequest** | [**UpdateNetworkVlanProfileRequest**](UpdateNetworkVlanProfileRequest.md) |  | 

### Return type

[**GetNetworkVlanProfiles200ResponseInner**](GetNetworkVlanProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

