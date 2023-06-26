# \BrandingPoliciesApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationBrandingPolicy**](BrandingPoliciesApi.md#CreateOrganizationBrandingPolicy) | **Post** /organizations/{organizationId}/brandingPolicies | Add a new branding policy to an organization
[**DeleteOrganizationBrandingPolicy**](BrandingPoliciesApi.md#DeleteOrganizationBrandingPolicy) | **Delete** /organizations/{organizationId}/brandingPolicies/{brandingPolicyId} | Delete a branding policy
[**GetOrganizationBrandingPolicies**](BrandingPoliciesApi.md#GetOrganizationBrandingPolicies) | **Get** /organizations/{organizationId}/brandingPolicies | List the branding policies of an organization
[**GetOrganizationBrandingPoliciesPriorities**](BrandingPoliciesApi.md#GetOrganizationBrandingPoliciesPriorities) | **Get** /organizations/{organizationId}/brandingPolicies/priorities | Return the branding policy IDs of an organization in priority order
[**GetOrganizationBrandingPolicy**](BrandingPoliciesApi.md#GetOrganizationBrandingPolicy) | **Get** /organizations/{organizationId}/brandingPolicies/{brandingPolicyId} | Return a branding policy
[**UpdateOrganizationBrandingPoliciesPriorities**](BrandingPoliciesApi.md#UpdateOrganizationBrandingPoliciesPriorities) | **Put** /organizations/{organizationId}/brandingPolicies/priorities | Update the priority ordering of an organization&#39;s branding policies.
[**UpdateOrganizationBrandingPolicy**](BrandingPoliciesApi.md#UpdateOrganizationBrandingPolicy) | **Put** /organizations/{organizationId}/brandingPolicies/{brandingPolicyId} | Update a branding policy



## CreateOrganizationBrandingPolicy

> InlineResponse2015 CreateOrganizationBrandingPolicy(ctx, organizationId).CreateOrganizationBrandingPolicy(createOrganizationBrandingPolicy).Execute()

Add a new branding policy to an organization



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
    createOrganizationBrandingPolicy := *openapiclient.NewInlineObject188() // InlineObject188 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingPoliciesApi.CreateOrganizationBrandingPolicy(context.Background(), organizationId).CreateOrganizationBrandingPolicy(createOrganizationBrandingPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesApi.CreateOrganizationBrandingPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationBrandingPolicy`: InlineResponse2015
    fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesApi.CreateOrganizationBrandingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationBrandingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationBrandingPolicy** | [**InlineObject188**](InlineObject188.md) |  | 

### Return type

[**InlineResponse2015**](InlineResponse2015.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationBrandingPolicy

> DeleteOrganizationBrandingPolicy(ctx, organizationId, brandingPolicyId).Execute()

Delete a branding policy



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
    brandingPolicyId := "brandingPolicyId_example" // string | Branding policy ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingPoliciesApi.DeleteOrganizationBrandingPolicy(context.Background(), organizationId, brandingPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesApi.DeleteOrganizationBrandingPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**brandingPolicyId** | **string** | Branding policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationBrandingPolicyRequest struct via the builder pattern


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


## GetOrganizationBrandingPolicies

> []InlineResponse200108 GetOrganizationBrandingPolicies(ctx, organizationId).Execute()

List the branding policies of an organization



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
    resp, r, err := apiClient.BrandingPoliciesApi.GetOrganizationBrandingPolicies(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesApi.GetOrganizationBrandingPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBrandingPolicies`: []InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesApi.GetOrganizationBrandingPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBrandingPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200108**](InlineResponse200108.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBrandingPoliciesPriorities

> InlineResponse200109 GetOrganizationBrandingPoliciesPriorities(ctx, organizationId).Execute()

Return the branding policy IDs of an organization in priority order



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
    resp, r, err := apiClient.BrandingPoliciesApi.GetOrganizationBrandingPoliciesPriorities(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesApi.GetOrganizationBrandingPoliciesPriorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBrandingPoliciesPriorities`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesApi.GetOrganizationBrandingPoliciesPriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBrandingPoliciesPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationBrandingPolicy

> InlineResponse200108 GetOrganizationBrandingPolicy(ctx, organizationId, brandingPolicyId).Execute()

Return a branding policy



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
    brandingPolicyId := "brandingPolicyId_example" // string | Branding policy ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingPoliciesApi.GetOrganizationBrandingPolicy(context.Background(), organizationId, brandingPolicyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesApi.GetOrganizationBrandingPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationBrandingPolicy`: InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesApi.GetOrganizationBrandingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**brandingPolicyId** | **string** | Branding policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationBrandingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200108**](InlineResponse200108.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationBrandingPoliciesPriorities

> InlineResponse200109 UpdateOrganizationBrandingPoliciesPriorities(ctx, organizationId).UpdateOrganizationBrandingPoliciesPriorities(updateOrganizationBrandingPoliciesPriorities).Execute()

Update the priority ordering of an organization's branding policies.



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
    updateOrganizationBrandingPoliciesPriorities := *openapiclient.NewInlineObject189() // InlineObject189 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingPoliciesApi.UpdateOrganizationBrandingPoliciesPriorities(context.Background(), organizationId).UpdateOrganizationBrandingPoliciesPriorities(updateOrganizationBrandingPoliciesPriorities).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesApi.UpdateOrganizationBrandingPoliciesPriorities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationBrandingPoliciesPriorities`: InlineResponse200109
    fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesApi.UpdateOrganizationBrandingPoliciesPriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationBrandingPoliciesPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationBrandingPoliciesPriorities** | [**InlineObject189**](InlineObject189.md) |  | 

### Return type

[**InlineResponse200109**](InlineResponse200109.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationBrandingPolicy

> InlineResponse200108 UpdateOrganizationBrandingPolicy(ctx, organizationId, brandingPolicyId).UpdateOrganizationBrandingPolicy(updateOrganizationBrandingPolicy).Execute()

Update a branding policy



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
    brandingPolicyId := "brandingPolicyId_example" // string | Branding policy ID
    updateOrganizationBrandingPolicy := *openapiclient.NewInlineObject190() // InlineObject190 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingPoliciesApi.UpdateOrganizationBrandingPolicy(context.Background(), organizationId, brandingPolicyId).UpdateOrganizationBrandingPolicy(updateOrganizationBrandingPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingPoliciesApi.UpdateOrganizationBrandingPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationBrandingPolicy`: InlineResponse200108
    fmt.Fprintf(os.Stdout, "Response from `BrandingPoliciesApi.UpdateOrganizationBrandingPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**brandingPolicyId** | **string** | Branding policy ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationBrandingPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationBrandingPolicy** | [**InlineObject190**](InlineObject190.md) |  | 

### Return type

[**InlineResponse200108**](InlineResponse200108.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

