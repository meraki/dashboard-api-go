# \EarlyAccessApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationEarlyAccessFeaturesOptIn**](EarlyAccessApi.md#CreateOrganizationEarlyAccessFeaturesOptIn) | **Post** /organizations/{organizationId}/earlyAccess/features/optIns | Create a new early access feature opt-in for an organization
[**DeleteOrganizationEarlyAccessFeaturesOptIn**](EarlyAccessApi.md#DeleteOrganizationEarlyAccessFeaturesOptIn) | **Delete** /organizations/{organizationId}/earlyAccess/features/optIns/{optInId} | Delete an early access feature opt-in
[**GetOrganizationEarlyAccessFeatures**](EarlyAccessApi.md#GetOrganizationEarlyAccessFeatures) | **Get** /organizations/{organizationId}/earlyAccess/features | List the available early access features for organization
[**GetOrganizationEarlyAccessFeaturesOptIn**](EarlyAccessApi.md#GetOrganizationEarlyAccessFeaturesOptIn) | **Get** /organizations/{organizationId}/earlyAccess/features/optIns/{optInId} | Show an early access feature opt-in for an organization
[**GetOrganizationEarlyAccessFeaturesOptIns**](EarlyAccessApi.md#GetOrganizationEarlyAccessFeaturesOptIns) | **Get** /organizations/{organizationId}/earlyAccess/features/optIns | List the early access feature opt-ins for an organization
[**UpdateOrganizationEarlyAccessFeaturesOptIn**](EarlyAccessApi.md#UpdateOrganizationEarlyAccessFeaturesOptIn) | **Put** /organizations/{organizationId}/earlyAccess/features/optIns/{optInId} | Update an early access feature opt-in for an organization



## CreateOrganizationEarlyAccessFeaturesOptIn

> map[string]interface{} CreateOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId).CreateOrganizationEarlyAccessFeaturesOptIn(createOrganizationEarlyAccessFeaturesOptIn).Execute()

Create a new early access feature opt-in for an organization



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
    createOrganizationEarlyAccessFeaturesOptIn := *openapiclient.NewCreateOrganizationEarlyAccessFeaturesOptInRequest("ShortName_example") // CreateOrganizationEarlyAccessFeaturesOptInRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EarlyAccessApi.CreateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId).CreateOrganizationEarlyAccessFeaturesOptIn(createOrganizationEarlyAccessFeaturesOptIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessApi.CreateOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationEarlyAccessFeaturesOptIn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EarlyAccessApi.CreateOrganizationEarlyAccessFeaturesOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationEarlyAccessFeaturesOptIn** | [**CreateOrganizationEarlyAccessFeaturesOptInRequest**](CreateOrganizationEarlyAccessFeaturesOptInRequest.md) |  | 

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


## DeleteOrganizationEarlyAccessFeaturesOptIn

> DeleteOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId, optInId).Execute()

Delete an early access feature opt-in



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
    optInId := "optInId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EarlyAccessApi.DeleteOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessApi.DeleteOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**optInId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


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


## GetOrganizationEarlyAccessFeatures

> []map[string]interface{} GetOrganizationEarlyAccessFeatures(ctx, organizationId).Execute()

List the available early access features for organization



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
    resp, r, err := apiClient.EarlyAccessApi.GetOrganizationEarlyAccessFeatures(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessApi.GetOrganizationEarlyAccessFeatures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationEarlyAccessFeatures`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EarlyAccessApi.GetOrganizationEarlyAccessFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEarlyAccessFeaturesRequest struct via the builder pattern


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


## GetOrganizationEarlyAccessFeaturesOptIn

> map[string]interface{} GetOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId, optInId).Execute()

Show an early access feature opt-in for an organization



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
    optInId := "optInId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EarlyAccessApi.GetOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessApi.GetOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationEarlyAccessFeaturesOptIn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EarlyAccessApi.GetOrganizationEarlyAccessFeaturesOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**optInId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


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


## GetOrganizationEarlyAccessFeaturesOptIns

> []map[string]interface{} GetOrganizationEarlyAccessFeaturesOptIns(ctx, organizationId).Execute()

List the early access feature opt-ins for an organization



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
    resp, r, err := apiClient.EarlyAccessApi.GetOrganizationEarlyAccessFeaturesOptIns(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessApi.GetOrganizationEarlyAccessFeaturesOptIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationEarlyAccessFeaturesOptIns`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EarlyAccessApi.GetOrganizationEarlyAccessFeaturesOptIns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationEarlyAccessFeaturesOptInsRequest struct via the builder pattern


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


## UpdateOrganizationEarlyAccessFeaturesOptIn

> map[string]interface{} UpdateOrganizationEarlyAccessFeaturesOptIn(ctx, organizationId, optInId).UpdateOrganizationEarlyAccessFeaturesOptIn(updateOrganizationEarlyAccessFeaturesOptIn).Execute()

Update an early access feature opt-in for an organization



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
    optInId := "optInId_example" // string | 
    updateOrganizationEarlyAccessFeaturesOptIn := *openapiclient.NewUpdateOrganizationEarlyAccessFeaturesOptInRequest() // UpdateOrganizationEarlyAccessFeaturesOptInRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EarlyAccessApi.UpdateOrganizationEarlyAccessFeaturesOptIn(context.Background(), organizationId, optInId).UpdateOrganizationEarlyAccessFeaturesOptIn(updateOrganizationEarlyAccessFeaturesOptIn).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EarlyAccessApi.UpdateOrganizationEarlyAccessFeaturesOptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationEarlyAccessFeaturesOptIn`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EarlyAccessApi.UpdateOrganizationEarlyAccessFeaturesOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**optInId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationEarlyAccessFeaturesOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationEarlyAccessFeaturesOptIn** | [**UpdateOrganizationEarlyAccessFeaturesOptInRequest**](UpdateOrganizationEarlyAccessFeaturesOptInRequest.md) |  | 

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

