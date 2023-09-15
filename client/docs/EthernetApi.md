# \EthernetApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignNetworkWirelessEthernetPortsProfiles**](EthernetApi.md#AssignNetworkWirelessEthernetPortsProfiles) | **Post** /networks/{networkId}/wireless/ethernet/ports/profiles/assign | Assign AP port profile to list of APs
[**CreateNetworkWirelessEthernetPortsProfile**](EthernetApi.md#CreateNetworkWirelessEthernetPortsProfile) | **Post** /networks/{networkId}/wireless/ethernet/ports/profiles | Create an AP port profile
[**DeleteNetworkWirelessEthernetPortsProfile**](EthernetApi.md#DeleteNetworkWirelessEthernetPortsProfile) | **Delete** /networks/{networkId}/wireless/ethernet/ports/profiles/{profileId} | Delete an AP port profile
[**GetNetworkWirelessEthernetPortsProfile**](EthernetApi.md#GetNetworkWirelessEthernetPortsProfile) | **Get** /networks/{networkId}/wireless/ethernet/ports/profiles/{profileId} | Show the AP port profile by ID for this network
[**GetNetworkWirelessEthernetPortsProfiles**](EthernetApi.md#GetNetworkWirelessEthernetPortsProfiles) | **Get** /networks/{networkId}/wireless/ethernet/ports/profiles | List the AP port profiles for this network
[**GetOrganizationWirelessDevicesEthernetStatuses**](EthernetApi.md#GetOrganizationWirelessDevicesEthernetStatuses) | **Get** /organizations/{organizationId}/wireless/devices/ethernet/statuses | List the most recent Ethernet link speed, duplex, aggregation and power mode and status information for wireless devices.
[**SetNetworkWirelessEthernetPortsProfilesDefault**](EthernetApi.md#SetNetworkWirelessEthernetPortsProfilesDefault) | **Post** /networks/{networkId}/wireless/ethernet/ports/profiles/setDefault | Set the AP port profile to be default for this network
[**UpdateNetworkWirelessEthernetPortsProfile**](EthernetApi.md#UpdateNetworkWirelessEthernetPortsProfile) | **Put** /networks/{networkId}/wireless/ethernet/ports/profiles/{profileId} | Update the AP port profile by ID for this network



## AssignNetworkWirelessEthernetPortsProfiles

> AssignNetworkWirelessEthernetPortsProfiles201Response AssignNetworkWirelessEthernetPortsProfiles(ctx, networkId).AssignNetworkWirelessEthernetPortsProfilesRequest(assignNetworkWirelessEthernetPortsProfilesRequest).Execute()

Assign AP port profile to list of APs



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
    assignNetworkWirelessEthernetPortsProfilesRequest := *openapiclient.NewAssignNetworkWirelessEthernetPortsProfilesRequest([]string{"Serials_example"}, "ProfileId_example") // AssignNetworkWirelessEthernetPortsProfilesRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EthernetApi.AssignNetworkWirelessEthernetPortsProfiles(context.Background(), networkId).AssignNetworkWirelessEthernetPortsProfilesRequest(assignNetworkWirelessEthernetPortsProfilesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthernetApi.AssignNetworkWirelessEthernetPortsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignNetworkWirelessEthernetPortsProfiles`: AssignNetworkWirelessEthernetPortsProfiles201Response
    fmt.Fprintf(os.Stdout, "Response from `EthernetApi.AssignNetworkWirelessEthernetPortsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignNetworkWirelessEthernetPortsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignNetworkWirelessEthernetPortsProfilesRequest** | [**AssignNetworkWirelessEthernetPortsProfilesRequest**](AssignNetworkWirelessEthernetPortsProfilesRequest.md) |  | 

### Return type

[**AssignNetworkWirelessEthernetPortsProfiles201Response**](AssignNetworkWirelessEthernetPortsProfiles201Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWirelessEthernetPortsProfile

> GetNetworkWirelessEthernetPortsProfiles200ResponseInner CreateNetworkWirelessEthernetPortsProfile(ctx, networkId).CreateNetworkWirelessEthernetPortsProfileRequest(createNetworkWirelessEthernetPortsProfileRequest).Execute()

Create an AP port profile



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
    createNetworkWirelessEthernetPortsProfileRequest := *openapiclient.NewCreateNetworkWirelessEthernetPortsProfileRequest("Name_example", []openapiclient.CreateNetworkWirelessEthernetPortsProfileRequestPortsInner{*openapiclient.NewCreateNetworkWirelessEthernetPortsProfileRequestPortsInner("Name_example")}) // CreateNetworkWirelessEthernetPortsProfileRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EthernetApi.CreateNetworkWirelessEthernetPortsProfile(context.Background(), networkId).CreateNetworkWirelessEthernetPortsProfileRequest(createNetworkWirelessEthernetPortsProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthernetApi.CreateNetworkWirelessEthernetPortsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWirelessEthernetPortsProfile`: GetNetworkWirelessEthernetPortsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EthernetApi.CreateNetworkWirelessEthernetPortsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWirelessEthernetPortsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWirelessEthernetPortsProfileRequest** | [**CreateNetworkWirelessEthernetPortsProfileRequest**](CreateNetworkWirelessEthernetPortsProfileRequest.md) |  | 

### Return type

[**GetNetworkWirelessEthernetPortsProfiles200ResponseInner**](GetNetworkWirelessEthernetPortsProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkWirelessEthernetPortsProfile

> DeleteNetworkWirelessEthernetPortsProfile(ctx, networkId, profileId).Execute()

Delete an AP port profile



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
    profileId := "profileId_example" // string | Profile ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EthernetApi.DeleteNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthernetApi.DeleteNetworkWirelessEthernetPortsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWirelessEthernetPortsProfileRequest struct via the builder pattern


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


## GetNetworkWirelessEthernetPortsProfile

> GetNetworkWirelessEthernetPortsProfiles200ResponseInner GetNetworkWirelessEthernetPortsProfile(ctx, networkId, profileId).Execute()

Show the AP port profile by ID for this network



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
    profileId := "profileId_example" // string | Profile ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EthernetApi.GetNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthernetApi.GetNetworkWirelessEthernetPortsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessEthernetPortsProfile`: GetNetworkWirelessEthernetPortsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EthernetApi.GetNetworkWirelessEthernetPortsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessEthernetPortsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWirelessEthernetPortsProfiles200ResponseInner**](GetNetworkWirelessEthernetPortsProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessEthernetPortsProfiles

> []GetNetworkWirelessEthernetPortsProfiles200ResponseInner GetNetworkWirelessEthernetPortsProfiles(ctx, networkId).Execute()

List the AP port profiles for this network



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
    resp, r, err := apiClient.EthernetApi.GetNetworkWirelessEthernetPortsProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthernetApi.GetNetworkWirelessEthernetPortsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessEthernetPortsProfiles`: []GetNetworkWirelessEthernetPortsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EthernetApi.GetNetworkWirelessEthernetPortsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessEthernetPortsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkWirelessEthernetPortsProfiles200ResponseInner**](GetNetworkWirelessEthernetPortsProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationWirelessDevicesEthernetStatuses

> []GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner GetOrganizationWirelessDevicesEthernetStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

List the most recent Ethernet link speed, duplex, aggregation and power mode and status information for wireless devices.



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    networkIds := []string{"Inner_example"} // []string | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456 (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EthernetApi.GetOrganizationWirelessDevicesEthernetStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthernetApi.GetOrganizationWirelessDevicesEthernetStatuses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationWirelessDevicesEthernetStatuses`: []GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EthernetApi.GetOrganizationWirelessDevicesEthernetStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationWirelessDevicesEthernetStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456 | 

### Return type

[**[]GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner**](GetOrganizationWirelessDevicesEthernetStatuses200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetNetworkWirelessEthernetPortsProfilesDefault

> SetNetworkWirelessEthernetPortsProfilesDefault200Response SetNetworkWirelessEthernetPortsProfilesDefault(ctx, networkId).SetNetworkWirelessEthernetPortsProfilesDefaultRequest(setNetworkWirelessEthernetPortsProfilesDefaultRequest).Execute()

Set the AP port profile to be default for this network



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
    setNetworkWirelessEthernetPortsProfilesDefaultRequest := *openapiclient.NewSetNetworkWirelessEthernetPortsProfilesDefaultRequest("ProfileId_example") // SetNetworkWirelessEthernetPortsProfilesDefaultRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EthernetApi.SetNetworkWirelessEthernetPortsProfilesDefault(context.Background(), networkId).SetNetworkWirelessEthernetPortsProfilesDefaultRequest(setNetworkWirelessEthernetPortsProfilesDefaultRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthernetApi.SetNetworkWirelessEthernetPortsProfilesDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetNetworkWirelessEthernetPortsProfilesDefault`: SetNetworkWirelessEthernetPortsProfilesDefault200Response
    fmt.Fprintf(os.Stdout, "Response from `EthernetApi.SetNetworkWirelessEthernetPortsProfilesDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetNetworkWirelessEthernetPortsProfilesDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setNetworkWirelessEthernetPortsProfilesDefaultRequest** | [**SetNetworkWirelessEthernetPortsProfilesDefaultRequest**](SetNetworkWirelessEthernetPortsProfilesDefaultRequest.md) |  | 

### Return type

[**SetNetworkWirelessEthernetPortsProfilesDefault200Response**](SetNetworkWirelessEthernetPortsProfilesDefault200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessEthernetPortsProfile

> GetNetworkWirelessEthernetPortsProfiles200ResponseInner UpdateNetworkWirelessEthernetPortsProfile(ctx, networkId, profileId).UpdateNetworkWirelessEthernetPortsProfileRequest(updateNetworkWirelessEthernetPortsProfileRequest).Execute()

Update the AP port profile by ID for this network



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
    profileId := "profileId_example" // string | Profile ID
    updateNetworkWirelessEthernetPortsProfileRequest := *openapiclient.NewUpdateNetworkWirelessEthernetPortsProfileRequest() // UpdateNetworkWirelessEthernetPortsProfileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EthernetApi.UpdateNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).UpdateNetworkWirelessEthernetPortsProfileRequest(updateNetworkWirelessEthernetPortsProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EthernetApi.UpdateNetworkWirelessEthernetPortsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessEthernetPortsProfile`: GetNetworkWirelessEthernetPortsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `EthernetApi.UpdateNetworkWirelessEthernetPortsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessEthernetPortsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessEthernetPortsProfileRequest** | [**UpdateNetworkWirelessEthernetPortsProfileRequest**](UpdateNetworkWirelessEthernetPortsProfileRequest.md) |  | 

### Return type

[**GetNetworkWirelessEthernetPortsProfiles200ResponseInner**](GetNetworkWirelessEthernetPortsProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

