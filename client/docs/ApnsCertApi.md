# \ApnsCertApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationSmApnsCert**](ApnsCertApi.md#GetOrganizationSmApnsCert) | **Get** /organizations/{organizationId}/sm/apnsCert | Get the organization&#39;s APNS certificate



## GetOrganizationSmApnsCert

> GetOrganizationSmApnsCert200Response GetOrganizationSmApnsCert(ctx, organizationId).Execute()

Get the organization's APNS certificate



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
    resp, r, err := apiClient.ApnsCertApi.GetOrganizationSmApnsCert(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApnsCertApi.GetOrganizationSmApnsCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationSmApnsCert`: GetOrganizationSmApnsCert200Response
    fmt.Fprintf(os.Stdout, "Response from `ApnsCertApi.GetOrganizationSmApnsCert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationSmApnsCertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationSmApnsCert200Response**](GetOrganizationSmApnsCert200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

