# \RfProfilesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkApplianceRfProfile**](RfProfilesApi.md#CreateNetworkApplianceRfProfile) | **Post** /networks/{networkId}/appliance/rfProfiles | Creates new RF profile for this network
[**CreateNetworkWirelessRfProfile**](RfProfilesApi.md#CreateNetworkWirelessRfProfile) | **Post** /networks/{networkId}/wireless/rfProfiles | Creates new RF profile for this network
[**DeleteNetworkApplianceRfProfile**](RfProfilesApi.md#DeleteNetworkApplianceRfProfile) | **Delete** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Delete a RF Profile
[**DeleteNetworkWirelessRfProfile**](RfProfilesApi.md#DeleteNetworkWirelessRfProfile) | **Delete** /networks/{networkId}/wireless/rfProfiles/{rfProfileId} | Delete a RF Profile
[**GetNetworkApplianceRfProfile**](RfProfilesApi.md#GetNetworkApplianceRfProfile) | **Get** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Return a RF profile
[**GetNetworkApplianceRfProfiles**](RfProfilesApi.md#GetNetworkApplianceRfProfiles) | **Get** /networks/{networkId}/appliance/rfProfiles | List the RF profiles for this network
[**GetNetworkWirelessRfProfile**](RfProfilesApi.md#GetNetworkWirelessRfProfile) | **Get** /networks/{networkId}/wireless/rfProfiles/{rfProfileId} | Return a RF profile
[**GetNetworkWirelessRfProfiles**](RfProfilesApi.md#GetNetworkWirelessRfProfiles) | **Get** /networks/{networkId}/wireless/rfProfiles | List RF profiles for this network
[**UpdateNetworkApplianceRfProfile**](RfProfilesApi.md#UpdateNetworkApplianceRfProfile) | **Put** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Updates specified RF profile for this network
[**UpdateNetworkWirelessRfProfile**](RfProfilesApi.md#UpdateNetworkWirelessRfProfile) | **Put** /networks/{networkId}/wireless/rfProfiles/{rfProfileId} | Updates specified RF profile for this network



## CreateNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner CreateNetworkApplianceRfProfile(ctx, networkId).CreateNetworkApplianceRfProfileRequest(createNetworkApplianceRfProfileRequest).Execute()

Creates new RF profile for this network



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
    createNetworkApplianceRfProfileRequest := *openapiclient.NewCreateNetworkApplianceRfProfileRequest("Name_example") // CreateNetworkApplianceRfProfileRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RfProfilesApi.CreateNetworkApplianceRfProfile(context.Background(), networkId).CreateNetworkApplianceRfProfileRequest(createNetworkApplianceRfProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesApi.CreateNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
    fmt.Fprintf(os.Stdout, "Response from `RfProfilesApi.CreateNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceRfProfileRequest** | [**CreateNetworkApplianceRfProfileRequest**](CreateNetworkApplianceRfProfileRequest.md) |  | 

### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkWirelessRfProfile

> GetNetworkWirelessRfProfiles200Response CreateNetworkWirelessRfProfile(ctx, networkId).CreateNetworkWirelessRfProfileRequest(createNetworkWirelessRfProfileRequest).Execute()

Creates new RF profile for this network



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
    createNetworkWirelessRfProfileRequest := *openapiclient.NewCreateNetworkWirelessRfProfileRequest("Name_example", "BandSelectionType_example") // CreateNetworkWirelessRfProfileRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RfProfilesApi.CreateNetworkWirelessRfProfile(context.Background(), networkId).CreateNetworkWirelessRfProfileRequest(createNetworkWirelessRfProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesApi.CreateNetworkWirelessRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWirelessRfProfile`: GetNetworkWirelessRfProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `RfProfilesApi.CreateNetworkWirelessRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkWirelessRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkWirelessRfProfileRequest** | [**CreateNetworkWirelessRfProfileRequest**](CreateNetworkWirelessRfProfileRequest.md) |  | 

### Return type

[**GetNetworkWirelessRfProfiles200Response**](GetNetworkWirelessRfProfiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceRfProfile

> DeleteNetworkApplianceRfProfile(ctx, networkId, rfProfileId).Execute()

Delete a RF Profile



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RfProfilesApi.DeleteNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesApi.DeleteNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceRfProfileRequest struct via the builder pattern


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


## DeleteNetworkWirelessRfProfile

> DeleteNetworkWirelessRfProfile(ctx, networkId, rfProfileId).Execute()

Delete a RF Profile



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RfProfilesApi.DeleteNetworkWirelessRfProfile(context.Background(), networkId, rfProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesApi.DeleteNetworkWirelessRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkWirelessRfProfileRequest struct via the builder pattern


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


## GetNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner GetNetworkApplianceRfProfile(ctx, networkId, rfProfileId).Execute()

Return a RF profile



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RfProfilesApi.GetNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesApi.GetNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
    fmt.Fprintf(os.Stdout, "Response from `RfProfilesApi.GetNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceRfProfiles

> GetNetworkApplianceRfProfiles200Response GetNetworkApplianceRfProfiles(ctx, networkId).Execute()

List the RF profiles for this network



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
    resp, r, err := apiClient.RfProfilesApi.GetNetworkApplianceRfProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesApi.GetNetworkApplianceRfProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkApplianceRfProfiles`: GetNetworkApplianceRfProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `RfProfilesApi.GetNetworkApplianceRfProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceRfProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceRfProfiles200Response**](GetNetworkApplianceRfProfiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessRfProfile

> GetNetworkWirelessRfProfiles200Response GetNetworkWirelessRfProfile(ctx, networkId, rfProfileId).Execute()

Return a RF profile



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RfProfilesApi.GetNetworkWirelessRfProfile(context.Background(), networkId, rfProfileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesApi.GetNetworkWirelessRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessRfProfile`: GetNetworkWirelessRfProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `RfProfilesApi.GetNetworkWirelessRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkWirelessRfProfiles200Response**](GetNetworkWirelessRfProfiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkWirelessRfProfiles

> GetNetworkWirelessRfProfiles200Response GetNetworkWirelessRfProfiles(ctx, networkId).IncludeTemplateProfiles(includeTemplateProfiles).Execute()

List RF profiles for this network



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
    includeTemplateProfiles := true // bool | If the network is bound to a template, this parameter controls whether or not the non-basic RF profiles defined on the template should be included in the response alongside the non-basic profiles defined on the bound network. Defaults to false. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RfProfilesApi.GetNetworkWirelessRfProfiles(context.Background(), networkId).IncludeTemplateProfiles(includeTemplateProfiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesApi.GetNetworkWirelessRfProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessRfProfiles`: GetNetworkWirelessRfProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `RfProfilesApi.GetNetworkWirelessRfProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkWirelessRfProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTemplateProfiles** | **bool** | If the network is bound to a template, this parameter controls whether or not the non-basic RF profiles defined on the template should be included in the response alongside the non-basic profiles defined on the bound network. Defaults to false. | 

### Return type

[**GetNetworkWirelessRfProfiles200Response**](GetNetworkWirelessRfProfiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner UpdateNetworkApplianceRfProfile(ctx, networkId, rfProfileId).UpdateNetworkApplianceRfProfileRequest(updateNetworkApplianceRfProfileRequest).Execute()

Updates specified RF profile for this network



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID
    updateNetworkApplianceRfProfileRequest := *openapiclient.NewUpdateNetworkApplianceRfProfileRequest() // UpdateNetworkApplianceRfProfileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RfProfilesApi.UpdateNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).UpdateNetworkApplianceRfProfileRequest(updateNetworkApplianceRfProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesApi.UpdateNetworkApplianceRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
    fmt.Fprintf(os.Stdout, "Response from `RfProfilesApi.UpdateNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceRfProfileRequest** | [**UpdateNetworkApplianceRfProfileRequest**](UpdateNetworkApplianceRfProfileRequest.md) |  | 

### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkWirelessRfProfile

> GetNetworkWirelessRfProfiles200Response UpdateNetworkWirelessRfProfile(ctx, networkId, rfProfileId).UpdateNetworkWirelessRfProfileRequest(updateNetworkWirelessRfProfileRequest).Execute()

Updates specified RF profile for this network



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
    rfProfileId := "rfProfileId_example" // string | Rf profile ID
    updateNetworkWirelessRfProfileRequest := *openapiclient.NewUpdateNetworkWirelessRfProfileRequest() // UpdateNetworkWirelessRfProfileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RfProfilesApi.UpdateNetworkWirelessRfProfile(context.Background(), networkId, rfProfileId).UpdateNetworkWirelessRfProfileRequest(updateNetworkWirelessRfProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RfProfilesApi.UpdateNetworkWirelessRfProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessRfProfile`: GetNetworkWirelessRfProfiles200Response
    fmt.Fprintf(os.Stdout, "Response from `RfProfilesApi.UpdateNetworkWirelessRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkWirelessRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkWirelessRfProfileRequest** | [**UpdateNetworkWirelessRfProfileRequest**](UpdateNetworkWirelessRfProfileRequest.md) |  | 

### Return type

[**GetNetworkWirelessRfProfiles200Response**](GetNetworkWirelessRfProfiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

