# \OpenapiSpecApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationOpenapiSpec**](OpenapiSpecApi.md#GetOrganizationOpenapiSpec) | **Get** /organizations/{organizationId}/openapiSpec | Return the OpenAPI Specification of the organization&#39;s API documentation in JSON



## GetOrganizationOpenapiSpec

> map[string]interface{} GetOrganizationOpenapiSpec(ctx, organizationId).Version(version).Execute()

Return the OpenAPI Specification of the organization's API documentation in JSON



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
    version := int32(56) // int32 | OpenAPI Specification version to return. Default is 2 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OpenapiSpecApi.GetOrganizationOpenapiSpec(context.Background(), organizationId).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenapiSpecApi.GetOrganizationOpenapiSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationOpenapiSpec`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OpenapiSpecApi.GetOrganizationOpenapiSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationOpenapiSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **int32** | OpenAPI Specification version to return. Default is 2 | 

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

