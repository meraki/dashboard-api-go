# \ProfilesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignNetworkWirelessEthernetPortsProfiles**](ProfilesApi.md#AssignNetworkWirelessEthernetPortsProfiles) | **Post** /networks/{networkId}/wireless/ethernet/ports/profiles/assign | Assign AP port profile to list of APs
[**CreateNetworkSensorAlertsProfile**](ProfilesApi.md#CreateNetworkSensorAlertsProfile) | **Post** /networks/{networkId}/sensor/alerts/profiles | Creates a sensor alert profile for a network.
[**CreateNetworkWirelessEthernetPortsProfile**](ProfilesApi.md#CreateNetworkWirelessEthernetPortsProfile) | **Post** /networks/{networkId}/wireless/ethernet/ports/profiles | Create an AP port profile
[**CreateOrganizationAlertsProfile**](ProfilesApi.md#CreateOrganizationAlertsProfile) | **Post** /organizations/{organizationId}/alerts/profiles | Create an organization-wide alert configuration
[**DeleteNetworkSensorAlertsProfile**](ProfilesApi.md#DeleteNetworkSensorAlertsProfile) | **Delete** /networks/{networkId}/sensor/alerts/profiles/{id} | Deletes a sensor alert profile from a network.
[**DeleteNetworkWirelessEthernetPortsProfile**](ProfilesApi.md#DeleteNetworkWirelessEthernetPortsProfile) | **Delete** /networks/{networkId}/wireless/ethernet/ports/profiles/{profileId} | Delete an AP port profile
[**DeleteOrganizationAlertsProfile**](ProfilesApi.md#DeleteOrganizationAlertsProfile) | **Delete** /organizations/{organizationId}/alerts/profiles/{alertConfigId} | Removes an organization-wide alert config
[**GetNetworkSensorAlertsProfile**](ProfilesApi.md#GetNetworkSensorAlertsProfile) | **Get** /networks/{networkId}/sensor/alerts/profiles/{id} | Show details of a sensor alert profile for a network.
[**GetNetworkSensorAlertsProfiles**](ProfilesApi.md#GetNetworkSensorAlertsProfiles) | **Get** /networks/{networkId}/sensor/alerts/profiles | Lists all sensor alert profiles for a network.
[**GetNetworkSmProfiles**](ProfilesApi.md#GetNetworkSmProfiles) | **Get** /networks/{networkId}/sm/profiles | List all profiles in a network
[**GetNetworkWirelessEthernetPortsProfile**](ProfilesApi.md#GetNetworkWirelessEthernetPortsProfile) | **Get** /networks/{networkId}/wireless/ethernet/ports/profiles/{profileId} | Show the AP port profile by ID for this network
[**GetNetworkWirelessEthernetPortsProfiles**](ProfilesApi.md#GetNetworkWirelessEthernetPortsProfiles) | **Get** /networks/{networkId}/wireless/ethernet/ports/profiles | List the AP port profiles for this network
[**GetOrganizationAlertsProfiles**](ProfilesApi.md#GetOrganizationAlertsProfiles) | **Get** /organizations/{organizationId}/alerts/profiles | List all organization-wide alert configurations
[**GetOrganizationConfigTemplateSwitchProfilePort**](ProfilesApi.md#GetOrganizationConfigTemplateSwitchProfilePort) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId} | Return a switch template port
[**GetOrganizationConfigTemplateSwitchProfilePorts**](ProfilesApi.md#GetOrganizationConfigTemplateSwitchProfilePorts) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports | Return all the ports of a switch template
[**GetOrganizationConfigTemplateSwitchProfiles**](ProfilesApi.md#GetOrganizationConfigTemplateSwitchProfiles) | **Get** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles | List the switch templates for your switch template configuration
[**SetNetworkWirelessEthernetPortsProfilesDefault**](ProfilesApi.md#SetNetworkWirelessEthernetPortsProfilesDefault) | **Post** /networks/{networkId}/wireless/ethernet/ports/profiles/setDefault | Set the AP port profile to be default for this network
[**UpdateNetworkSensorAlertsProfile**](ProfilesApi.md#UpdateNetworkSensorAlertsProfile) | **Put** /networks/{networkId}/sensor/alerts/profiles/{id} | Updates a sensor alert profile for a network.
[**UpdateNetworkWirelessEthernetPortsProfile**](ProfilesApi.md#UpdateNetworkWirelessEthernetPortsProfile) | **Put** /networks/{networkId}/wireless/ethernet/ports/profiles/{profileId} | Update the AP port profile by ID for this network
[**UpdateOrganizationAlertsProfile**](ProfilesApi.md#UpdateOrganizationAlertsProfile) | **Put** /organizations/{organizationId}/alerts/profiles/{alertConfigId} | Update an organization-wide alert config
[**UpdateOrganizationConfigTemplateSwitchProfilePort**](ProfilesApi.md#UpdateOrganizationConfigTemplateSwitchProfilePort) | **Put** /organizations/{organizationId}/configTemplates/{configTemplateId}/switch/profiles/{profileId}/ports/{portId} | Update a switch template port



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
    resp, r, err := apiClient.ProfilesApi.AssignNetworkWirelessEthernetPortsProfiles(context.Background(), networkId).AssignNetworkWirelessEthernetPortsProfilesRequest(assignNetworkWirelessEthernetPortsProfilesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.AssignNetworkWirelessEthernetPortsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignNetworkWirelessEthernetPortsProfiles`: AssignNetworkWirelessEthernetPortsProfiles201Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.AssignNetworkWirelessEthernetPortsProfiles`: %v\n", resp)
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


## CreateNetworkSensorAlertsProfile

> GetNetworkSensorAlertsProfiles200ResponseInner CreateNetworkSensorAlertsProfile(ctx, networkId).CreateNetworkSensorAlertsProfileRequest(createNetworkSensorAlertsProfileRequest).Execute()

Creates a sensor alert profile for a network.



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
    createNetworkSensorAlertsProfileRequest := *openapiclient.NewCreateNetworkSensorAlertsProfileRequest("Name_example", []openapiclient.GetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner{*openapiclient.NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInner("Metric_example", *openapiclient.NewGetNetworkSensorAlertsProfiles200ResponseInnerConditionsInnerThreshold())}) // CreateNetworkSensorAlertsProfileRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.CreateNetworkSensorAlertsProfile(context.Background(), networkId).CreateNetworkSensorAlertsProfileRequest(createNetworkSensorAlertsProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.CreateNetworkSensorAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkSensorAlertsProfile`: GetNetworkSensorAlertsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.CreateNetworkSensorAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkSensorAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkSensorAlertsProfileRequest** | [**CreateNetworkSensorAlertsProfileRequest**](CreateNetworkSensorAlertsProfileRequest.md) |  | 

### Return type

[**GetNetworkSensorAlertsProfiles200ResponseInner**](GetNetworkSensorAlertsProfiles200ResponseInner.md)

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
    resp, r, err := apiClient.ProfilesApi.CreateNetworkWirelessEthernetPortsProfile(context.Background(), networkId).CreateNetworkWirelessEthernetPortsProfileRequest(createNetworkWirelessEthernetPortsProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.CreateNetworkWirelessEthernetPortsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkWirelessEthernetPortsProfile`: GetNetworkWirelessEthernetPortsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.CreateNetworkWirelessEthernetPortsProfile`: %v\n", resp)
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


## CreateOrganizationAlertsProfile

> map[string]interface{} CreateOrganizationAlertsProfile(ctx, organizationId).CreateOrganizationAlertsProfileRequest(createOrganizationAlertsProfileRequest).Execute()

Create an organization-wide alert configuration



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
    createOrganizationAlertsProfileRequest := *openapiclient.NewCreateOrganizationAlertsProfileRequest("Type_example", *openapiclient.NewCreateOrganizationAlertsProfileRequestAlertCondition(), *openapiclient.NewCreateOrganizationAlertsProfileRequestRecipients(), []string{"NetworkTags_example"}) // CreateOrganizationAlertsProfileRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.CreateOrganizationAlertsProfile(context.Background(), organizationId).CreateOrganizationAlertsProfileRequest(createOrganizationAlertsProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.CreateOrganizationAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAlertsProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.CreateOrganizationAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAlertsProfileRequest** | [**CreateOrganizationAlertsProfileRequest**](CreateOrganizationAlertsProfileRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkSensorAlertsProfile

> DeleteNetworkSensorAlertsProfile(ctx, networkId, id).Execute()

Deletes a sensor alert profile from a network.



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProfilesApi.DeleteNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.DeleteNetworkSensorAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkSensorAlertsProfileRequest struct via the builder pattern


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
    r, err := apiClient.ProfilesApi.DeleteNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.DeleteNetworkWirelessEthernetPortsProfile``: %v\n", err)
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


## DeleteOrganizationAlertsProfile

> DeleteOrganizationAlertsProfile(ctx, organizationId, alertConfigId).Execute()

Removes an organization-wide alert config



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
    alertConfigId := "alertConfigId_example" // string | Alert config ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProfilesApi.DeleteOrganizationAlertsProfile(context.Background(), organizationId, alertConfigId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.DeleteOrganizationAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**alertConfigId** | **string** | Alert config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAlertsProfileRequest struct via the builder pattern


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


## GetNetworkSensorAlertsProfile

> GetNetworkSensorAlertsProfiles200ResponseInner GetNetworkSensorAlertsProfile(ctx, networkId, id).Execute()

Show details of a sensor alert profile for a network.



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.GetNetworkSensorAlertsProfile(context.Background(), networkId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.GetNetworkSensorAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorAlertsProfile`: GetNetworkSensorAlertsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.GetNetworkSensorAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkSensorAlertsProfiles200ResponseInner**](GetNetworkSensorAlertsProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSensorAlertsProfiles

> []GetNetworkSensorAlertsProfiles200ResponseInner GetNetworkSensorAlertsProfiles(ctx, networkId).Execute()

Lists all sensor alert profiles for a network.



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
    resp, r, err := apiClient.ProfilesApi.GetNetworkSensorAlertsProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.GetNetworkSensorAlertsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSensorAlertsProfiles`: []GetNetworkSensorAlertsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.GetNetworkSensorAlertsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSensorAlertsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSensorAlertsProfiles200ResponseInner**](GetNetworkSensorAlertsProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkSmProfiles

> []GetNetworkSmProfiles200ResponseInner GetNetworkSmProfiles(ctx, networkId).Execute()

List all profiles in a network



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
    resp, r, err := apiClient.ProfilesApi.GetNetworkSmProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.GetNetworkSmProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkSmProfiles`: []GetNetworkSmProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.GetNetworkSmProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkSmProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkSmProfiles200ResponseInner**](GetNetworkSmProfiles200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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
    resp, r, err := apiClient.ProfilesApi.GetNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.GetNetworkWirelessEthernetPortsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessEthernetPortsProfile`: GetNetworkWirelessEthernetPortsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.GetNetworkWirelessEthernetPortsProfile`: %v\n", resp)
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
    resp, r, err := apiClient.ProfilesApi.GetNetworkWirelessEthernetPortsProfiles(context.Background(), networkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.GetNetworkWirelessEthernetPortsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkWirelessEthernetPortsProfiles`: []GetNetworkWirelessEthernetPortsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.GetNetworkWirelessEthernetPortsProfiles`: %v\n", resp)
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


## GetOrganizationAlertsProfiles

> []map[string]interface{} GetOrganizationAlertsProfiles(ctx, organizationId).Execute()

List all organization-wide alert configurations



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

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.GetOrganizationAlertsProfiles(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.GetOrganizationAlertsProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAlertsProfiles`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.GetOrganizationAlertsProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAlertsProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfilePort

> GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner GetOrganizationConfigTemplateSwitchProfilePort(ctx, organizationId, configTemplateId, profileId, portId).Execute()

Return a switch template port



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
    configTemplateId := "configTemplateId_example" // string | Config template ID
    profileId := "profileId_example" // string | Profile ID
    portId := "portId_example" // string | Port ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.GetOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.GetOrganizationConfigTemplateSwitchProfilePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigTemplateSwitchProfilePort`: GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.GetOrganizationConfigTemplateSwitchProfilePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner**](GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfilePorts

> []GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner GetOrganizationConfigTemplateSwitchProfilePorts(ctx, organizationId, configTemplateId, profileId).Execute()

Return all the ports of a switch template



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
    configTemplateId := "configTemplateId_example" // string | Config template ID
    profileId := "profileId_example" // string | Profile ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.GetOrganizationConfigTemplateSwitchProfilePorts(context.Background(), organizationId, configTemplateId, profileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.GetOrganizationConfigTemplateSwitchProfilePorts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigTemplateSwitchProfilePorts`: []GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.GetOrganizationConfigTemplateSwitchProfilePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner**](GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationConfigTemplateSwitchProfiles

> []GetOrganizationConfigTemplateSwitchProfiles200ResponseInner GetOrganizationConfigTemplateSwitchProfiles(ctx, organizationId, configTemplateId).Execute()

List the switch templates for your switch template configuration



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
    configTemplateId := "configTemplateId_example" // string | Config template ID

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.GetOrganizationConfigTemplateSwitchProfiles(context.Background(), organizationId, configTemplateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.GetOrganizationConfigTemplateSwitchProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationConfigTemplateSwitchProfiles`: []GetOrganizationConfigTemplateSwitchProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.GetOrganizationConfigTemplateSwitchProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationConfigTemplateSwitchProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]GetOrganizationConfigTemplateSwitchProfiles200ResponseInner**](GetOrganizationConfigTemplateSwitchProfiles200ResponseInner.md)

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
    resp, r, err := apiClient.ProfilesApi.SetNetworkWirelessEthernetPortsProfilesDefault(context.Background(), networkId).SetNetworkWirelessEthernetPortsProfilesDefaultRequest(setNetworkWirelessEthernetPortsProfilesDefaultRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.SetNetworkWirelessEthernetPortsProfilesDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetNetworkWirelessEthernetPortsProfilesDefault`: SetNetworkWirelessEthernetPortsProfilesDefault200Response
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.SetNetworkWirelessEthernetPortsProfilesDefault`: %v\n", resp)
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


## UpdateNetworkSensorAlertsProfile

> GetNetworkSensorAlertsProfiles200ResponseInner UpdateNetworkSensorAlertsProfile(ctx, networkId, id).UpdateNetworkSensorAlertsProfileRequest(updateNetworkSensorAlertsProfileRequest).Execute()

Updates a sensor alert profile for a network.



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
    id := "id_example" // string | ID
    updateNetworkSensorAlertsProfileRequest := *openapiclient.NewUpdateNetworkSensorAlertsProfileRequest() // UpdateNetworkSensorAlertsProfileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.UpdateNetworkSensorAlertsProfile(context.Background(), networkId, id).UpdateNetworkSensorAlertsProfileRequest(updateNetworkSensorAlertsProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.UpdateNetworkSensorAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkSensorAlertsProfile`: GetNetworkSensorAlertsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.UpdateNetworkSensorAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkSensorAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkSensorAlertsProfileRequest** | [**UpdateNetworkSensorAlertsProfileRequest**](UpdateNetworkSensorAlertsProfileRequest.md) |  | 

### Return type

[**GetNetworkSensorAlertsProfiles200ResponseInner**](GetNetworkSensorAlertsProfiles200ResponseInner.md)

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
    resp, r, err := apiClient.ProfilesApi.UpdateNetworkWirelessEthernetPortsProfile(context.Background(), networkId, profileId).UpdateNetworkWirelessEthernetPortsProfileRequest(updateNetworkWirelessEthernetPortsProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.UpdateNetworkWirelessEthernetPortsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkWirelessEthernetPortsProfile`: GetNetworkWirelessEthernetPortsProfiles200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.UpdateNetworkWirelessEthernetPortsProfile`: %v\n", resp)
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


## UpdateOrganizationAlertsProfile

> map[string]interface{} UpdateOrganizationAlertsProfile(ctx, organizationId, alertConfigId).UpdateOrganizationAlertsProfileRequest(updateOrganizationAlertsProfileRequest).Execute()

Update an organization-wide alert config



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
    alertConfigId := "alertConfigId_example" // string | Alert config ID
    updateOrganizationAlertsProfileRequest := *openapiclient.NewUpdateOrganizationAlertsProfileRequest() // UpdateOrganizationAlertsProfileRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.UpdateOrganizationAlertsProfile(context.Background(), organizationId, alertConfigId).UpdateOrganizationAlertsProfileRequest(updateOrganizationAlertsProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.UpdateOrganizationAlertsProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAlertsProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.UpdateOrganizationAlertsProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**alertConfigId** | **string** | Alert config ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAlertsProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAlertsProfileRequest** | [**UpdateOrganizationAlertsProfileRequest**](UpdateOrganizationAlertsProfileRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationConfigTemplateSwitchProfilePort

> GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner UpdateOrganizationConfigTemplateSwitchProfilePort(ctx, organizationId, configTemplateId, profileId, portId).UpdateOrganizationConfigTemplateSwitchProfilePortRequest(updateOrganizationConfigTemplateSwitchProfilePortRequest).Execute()

Update a switch template port



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
    configTemplateId := "configTemplateId_example" // string | Config template ID
    profileId := "profileId_example" // string | Profile ID
    portId := "portId_example" // string | Port ID
    updateOrganizationConfigTemplateSwitchProfilePortRequest := *openapiclient.NewUpdateOrganizationConfigTemplateSwitchProfilePortRequest() // UpdateOrganizationConfigTemplateSwitchProfilePortRequest |  (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfilesApi.UpdateOrganizationConfigTemplateSwitchProfilePort(context.Background(), organizationId, configTemplateId, profileId, portId).UpdateOrganizationConfigTemplateSwitchProfilePortRequest(updateOrganizationConfigTemplateSwitchProfilePortRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfilesApi.UpdateOrganizationConfigTemplateSwitchProfilePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationConfigTemplateSwitchProfilePort`: GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `ProfilesApi.UpdateOrganizationConfigTemplateSwitchProfilePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**configTemplateId** | **string** | Config template ID | 
**profileId** | **string** | Profile ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationConfigTemplateSwitchProfilePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **updateOrganizationConfigTemplateSwitchProfilePortRequest** | [**UpdateOrganizationConfigTemplateSwitchProfilePortRequest**](UpdateOrganizationConfigTemplateSwitchProfilePortRequest.md) |  | 

### Return type

[**GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner**](GetOrganizationConfigTemplateSwitchProfilePorts200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

