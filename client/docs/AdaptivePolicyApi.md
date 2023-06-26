# \AdaptivePolicyApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationAdaptivePolicyAcl**](AdaptivePolicyApi.md#CreateOrganizationAdaptivePolicyAcl) | **Post** /organizations/{organizationId}/adaptivePolicy/acls | Creates new adaptive policy ACL
[**CreateOrganizationAdaptivePolicyGroup**](AdaptivePolicyApi.md#CreateOrganizationAdaptivePolicyGroup) | **Post** /organizations/{organizationId}/adaptivePolicy/groups | Creates a new adaptive policy group
[**CreateOrganizationAdaptivePolicyPolicy**](AdaptivePolicyApi.md#CreateOrganizationAdaptivePolicyPolicy) | **Post** /organizations/{organizationId}/adaptivePolicy/policies | Add an Adaptive Policy
[**DeleteOrganizationAdaptivePolicyAcl**](AdaptivePolicyApi.md#DeleteOrganizationAdaptivePolicyAcl) | **Delete** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Deletes the specified adaptive policy ACL
[**DeleteOrganizationAdaptivePolicyGroup**](AdaptivePolicyApi.md#DeleteOrganizationAdaptivePolicyGroup) | **Delete** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Deletes the specified adaptive policy group and any associated policies and references
[**DeleteOrganizationAdaptivePolicyPolicy**](AdaptivePolicyApi.md#DeleteOrganizationAdaptivePolicyPolicy) | **Delete** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Delete an Adaptive Policy
[**GetOrganizationAdaptivePolicyAcl**](AdaptivePolicyApi.md#GetOrganizationAdaptivePolicyAcl) | **Get** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Returns the adaptive policy ACL information
[**GetOrganizationAdaptivePolicyAcls**](AdaptivePolicyApi.md#GetOrganizationAdaptivePolicyAcls) | **Get** /organizations/{organizationId}/adaptivePolicy/acls | List adaptive policy ACLs in a organization
[**GetOrganizationAdaptivePolicyGroup**](AdaptivePolicyApi.md#GetOrganizationAdaptivePolicyGroup) | **Get** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Returns an adaptive policy group
[**GetOrganizationAdaptivePolicyGroups**](AdaptivePolicyApi.md#GetOrganizationAdaptivePolicyGroups) | **Get** /organizations/{organizationId}/adaptivePolicy/groups | List adaptive policy groups in a organization
[**GetOrganizationAdaptivePolicyOverview**](AdaptivePolicyApi.md#GetOrganizationAdaptivePolicyOverview) | **Get** /organizations/{organizationId}/adaptivePolicy/overview | Returns adaptive policy aggregate statistics for an organization
[**GetOrganizationAdaptivePolicyPolicies**](AdaptivePolicyApi.md#GetOrganizationAdaptivePolicyPolicies) | **Get** /organizations/{organizationId}/adaptivePolicy/policies | List adaptive policies in an organization
[**GetOrganizationAdaptivePolicyPolicy**](AdaptivePolicyApi.md#GetOrganizationAdaptivePolicyPolicy) | **Get** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Return an adaptive policy
[**GetOrganizationAdaptivePolicySettings**](AdaptivePolicyApi.md#GetOrganizationAdaptivePolicySettings) | **Get** /organizations/{organizationId}/adaptivePolicy/settings | Returns global adaptive policy settings in an organization
[**UpdateOrganizationAdaptivePolicyAcl**](AdaptivePolicyApi.md#UpdateOrganizationAdaptivePolicyAcl) | **Put** /organizations/{organizationId}/adaptivePolicy/acls/{aclId} | Updates an adaptive policy ACL
[**UpdateOrganizationAdaptivePolicyGroup**](AdaptivePolicyApi.md#UpdateOrganizationAdaptivePolicyGroup) | **Put** /organizations/{organizationId}/adaptivePolicy/groups/{id} | Updates an adaptive policy group
[**UpdateOrganizationAdaptivePolicyPolicy**](AdaptivePolicyApi.md#UpdateOrganizationAdaptivePolicyPolicy) | **Put** /organizations/{organizationId}/adaptivePolicy/policies/{id} | Update an Adaptive Policy
[**UpdateOrganizationAdaptivePolicySettings**](AdaptivePolicyApi.md#UpdateOrganizationAdaptivePolicySettings) | **Put** /organizations/{organizationId}/adaptivePolicy/settings | Update global adaptive policy settings



## CreateOrganizationAdaptivePolicyAcl

> InlineResponse200100 CreateOrganizationAdaptivePolicyAcl(ctx, organizationId).CreateOrganizationAdaptivePolicyAcl(createOrganizationAdaptivePolicyAcl).Execute()

Creates new adaptive policy ACL



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
    createOrganizationAdaptivePolicyAcl := *openapiclient.NewInlineObject174("Name_example", []openapiclient.OrganizationsOrganizationIdAdaptivePolicyAclsRules1{*openapiclient.NewOrganizationsOrganizationIdAdaptivePolicyAclsRules1("Policy_example", "Protocol_example")}, "IpVersion_example") // InlineObject174 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.CreateOrganizationAdaptivePolicyAcl(context.Background(), organizationId).CreateOrganizationAdaptivePolicyAcl(createOrganizationAdaptivePolicyAcl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.CreateOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdaptivePolicyAcl`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.CreateOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdaptivePolicyAcl** | [**InlineObject174**](InlineObject174.md) |  | 

### Return type

[**InlineResponse200100**](InlineResponse200100.md)

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
    createOrganizationAdaptivePolicyGroup := *openapiclient.NewInlineObject176("Name_example", int32(123)) // InlineObject176 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.CreateOrganizationAdaptivePolicyGroup(context.Background(), organizationId).CreateOrganizationAdaptivePolicyGroup(createOrganizationAdaptivePolicyGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.CreateOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdaptivePolicyGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.CreateOrganizationAdaptivePolicyGroup`: %v\n", resp)
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

 **createOrganizationAdaptivePolicyGroup** | [**InlineObject176**](InlineObject176.md) |  | 

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


## CreateOrganizationAdaptivePolicyPolicy

> map[string]interface{} CreateOrganizationAdaptivePolicyPolicy(ctx, organizationId).CreateOrganizationAdaptivePolicyPolicy(createOrganizationAdaptivePolicyPolicy).Execute()

Add an Adaptive Policy



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
    createOrganizationAdaptivePolicyPolicy := *openapiclient.NewInlineObject178(*openapiclient.NewOrganizationsOrganizationIdAdaptivePolicyPoliciesSourceGroup(), *openapiclient.NewOrganizationsOrganizationIdAdaptivePolicyPoliciesDestinationGroup()) // InlineObject178 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.CreateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId).CreateOrganizationAdaptivePolicyPolicy(createOrganizationAdaptivePolicyPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.CreateOrganizationAdaptivePolicyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganizationAdaptivePolicyPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.CreateOrganizationAdaptivePolicyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrganizationAdaptivePolicyPolicy** | [**InlineObject178**](InlineObject178.md) |  | 

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


## DeleteOrganizationAdaptivePolicyAcl

> DeleteOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).Execute()

Deletes the specified adaptive policy ACL



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
    aclId := "aclId_example" // string | Acl ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.DeleteOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.DeleteOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdaptivePolicyAclRequest struct via the builder pattern


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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.DeleteOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.DeleteOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

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


## DeleteOrganizationAdaptivePolicyPolicy

> DeleteOrganizationAdaptivePolicyPolicy(ctx, organizationId, id).Execute()

Delete an Adaptive Policy



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.DeleteOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.DeleteOrganizationAdaptivePolicyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


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


## GetOrganizationAdaptivePolicyAcl

> InlineResponse200100 GetOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).Execute()

Returns the adaptive policy ACL information



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
    aclId := "aclId_example" // string | Acl ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.GetOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyAcl`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.GetOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InlineResponse200100**](InlineResponse200100.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyAcls

> []InlineResponse200100 GetOrganizationAdaptivePolicyAcls(ctx, organizationId).Execute()

List adaptive policy ACLs in a organization



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
    resp, r, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyAcls(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.GetOrganizationAdaptivePolicyAcls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyAcls`: []InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.GetOrganizationAdaptivePolicyAcls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyAclsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InlineResponse200100**](InlineResponse200100.md)

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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.GetOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.GetOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

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
    resp, r, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyGroups(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.GetOrganizationAdaptivePolicyGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyGroups`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.GetOrganizationAdaptivePolicyGroups`: %v\n", resp)
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


## GetOrganizationAdaptivePolicyOverview

> InlineResponse200101 GetOrganizationAdaptivePolicyOverview(ctx, organizationId).Execute()

Returns adaptive policy aggregate statistics for an organization



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
    resp, r, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyOverview(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.GetOrganizationAdaptivePolicyOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyOverview`: InlineResponse200101
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.GetOrganizationAdaptivePolicyOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InlineResponse200101**](InlineResponse200101.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationAdaptivePolicyPolicies

> []map[string]interface{} GetOrganizationAdaptivePolicyPolicies(ctx, organizationId).Execute()

List adaptive policies in an organization



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
    resp, r, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyPolicies(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.GetOrganizationAdaptivePolicyPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyPolicies`: []map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.GetOrganizationAdaptivePolicyPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyPoliciesRequest struct via the builder pattern


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


## GetOrganizationAdaptivePolicyPolicy

> map[string]interface{} GetOrganizationAdaptivePolicyPolicy(ctx, organizationId, id).Execute()

Return an adaptive policy



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
    id := "id_example" // string | ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.GetOrganizationAdaptivePolicyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicyPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.GetOrganizationAdaptivePolicyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


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


## GetOrganizationAdaptivePolicySettings

> map[string]interface{} GetOrganizationAdaptivePolicySettings(ctx, organizationId).Execute()

Returns global adaptive policy settings in an organization



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
    resp, r, err := apiClient.AdaptivePolicyApi.GetOrganizationAdaptivePolicySettings(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.GetOrganizationAdaptivePolicySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationAdaptivePolicySettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.GetOrganizationAdaptivePolicySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationAdaptivePolicySettingsRequest struct via the builder pattern


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


## UpdateOrganizationAdaptivePolicyAcl

> InlineResponse200100 UpdateOrganizationAdaptivePolicyAcl(ctx, organizationId, aclId).UpdateOrganizationAdaptivePolicyAcl(updateOrganizationAdaptivePolicyAcl).Execute()

Updates an adaptive policy ACL



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
    aclId := "aclId_example" // string | Acl ID
    updateOrganizationAdaptivePolicyAcl := *openapiclient.NewInlineObject175() // InlineObject175 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyAcl(context.Background(), organizationId, aclId).UpdateOrganizationAdaptivePolicyAcl(updateOrganizationAdaptivePolicyAcl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdaptivePolicyAcl`: InlineResponse200100
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**aclId** | **string** | Acl ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyAcl** | [**InlineObject175**](InlineObject175.md) |  | 

### Return type

[**InlineResponse200100**](InlineResponse200100.md)

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
    id := "id_example" // string | ID
    updateOrganizationAdaptivePolicyGroup := *openapiclient.NewInlineObject177() // InlineObject177 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyGroup(context.Background(), organizationId, id).UpdateOrganizationAdaptivePolicyGroup(updateOrganizationAdaptivePolicyGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdaptivePolicyGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyGroup** | [**InlineObject177**](InlineObject177.md) |  | 

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


## UpdateOrganizationAdaptivePolicyPolicy

> map[string]interface{} UpdateOrganizationAdaptivePolicyPolicy(ctx, organizationId, id).UpdateOrganizationAdaptivePolicyPolicy(updateOrganizationAdaptivePolicyPolicy).Execute()

Update an Adaptive Policy



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
    id := "id_example" // string | ID
    updateOrganizationAdaptivePolicyPolicy := *openapiclient.NewInlineObject179() // InlineObject179 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyPolicy(context.Background(), organizationId, id).UpdateOrganizationAdaptivePolicyPolicy(updateOrganizationAdaptivePolicyPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdaptivePolicyPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.UpdateOrganizationAdaptivePolicyPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 
**id** | **string** | ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicyPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateOrganizationAdaptivePolicyPolicy** | [**InlineObject179**](InlineObject179.md) |  | 

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


## UpdateOrganizationAdaptivePolicySettings

> map[string]interface{} UpdateOrganizationAdaptivePolicySettings(ctx, organizationId).UpdateOrganizationAdaptivePolicySettings(updateOrganizationAdaptivePolicySettings).Execute()

Update global adaptive policy settings



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
    updateOrganizationAdaptivePolicySettings := *openapiclient.NewInlineObject180() // InlineObject180 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdaptivePolicyApi.UpdateOrganizationAdaptivePolicySettings(context.Background(), organizationId).UpdateOrganizationAdaptivePolicySettings(updateOrganizationAdaptivePolicySettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdaptivePolicyApi.UpdateOrganizationAdaptivePolicySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrganizationAdaptivePolicySettings`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `AdaptivePolicyApi.UpdateOrganizationAdaptivePolicySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationAdaptivePolicySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationAdaptivePolicySettings** | [**InlineObject180**](InlineObject180.md) |  | 

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

