# \SubscriptionsApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindAdministeredLicensingSubscriptionSubscription**](SubscriptionsApi.md#BindAdministeredLicensingSubscriptionSubscription) | **Post** /administered/licensing/subscription/subscriptions/{subscriptionId}/bind | Bind networks to a subscription
[**ClaimAdministeredLicensingSubscriptionSubscriptions**](SubscriptionsApi.md#ClaimAdministeredLicensingSubscriptionSubscriptions) | **Post** /administered/licensing/subscription/subscriptions/claim | Claim a subscription into an organization.
[**GetAdministeredLicensingSubscriptionSubscriptions**](SubscriptionsApi.md#GetAdministeredLicensingSubscriptionSubscriptions) | **Get** /administered/licensing/subscription/subscriptions | List available subscriptions
[**ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey**](SubscriptionsApi.md#ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey) | **Post** /administered/licensing/subscription/subscriptions/claimKey/validate | Find a subscription by claim key



## BindAdministeredLicensingSubscriptionSubscription

> BindAdministeredLicensingSubscriptionSubscription200Response BindAdministeredLicensingSubscriptionSubscription(ctx, subscriptionId).BindAdministeredLicensingSubscriptionSubscriptionRequest(bindAdministeredLicensingSubscriptionSubscriptionRequest).Validate(validate).Execute()

Bind networks to a subscription



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
    subscriptionId := "subscriptionId_example" // string | Subscription ID
    bindAdministeredLicensingSubscriptionSubscriptionRequest := *openapiclient.NewBindAdministeredLicensingSubscriptionSubscriptionRequest([]string{"NetworkIds_example"}) // BindAdministeredLicensingSubscriptionSubscriptionRequest | 
    validate := true // bool | Check if the provided networks can be bound to the subscription. Returns any licensing problems and does not commit the results. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsApi.BindAdministeredLicensingSubscriptionSubscription(context.Background(), subscriptionId).BindAdministeredLicensingSubscriptionSubscriptionRequest(bindAdministeredLicensingSubscriptionSubscriptionRequest).Validate(validate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.BindAdministeredLicensingSubscriptionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BindAdministeredLicensingSubscriptionSubscription`: BindAdministeredLicensingSubscriptionSubscription200Response
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.BindAdministeredLicensingSubscriptionSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Subscription ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBindAdministeredLicensingSubscriptionSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bindAdministeredLicensingSubscriptionSubscriptionRequest** | [**BindAdministeredLicensingSubscriptionSubscriptionRequest**](BindAdministeredLicensingSubscriptionSubscriptionRequest.md) |  | 
 **validate** | **bool** | Check if the provided networks can be bound to the subscription. Returns any licensing problems and does not commit the results. | 

### Return type

[**BindAdministeredLicensingSubscriptionSubscription200Response**](BindAdministeredLicensingSubscriptionSubscription200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClaimAdministeredLicensingSubscriptionSubscriptions

> GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner ClaimAdministeredLicensingSubscriptionSubscriptions(ctx).ClaimAdministeredLicensingSubscriptionSubscriptionsRequest(claimAdministeredLicensingSubscriptionSubscriptionsRequest).Validate(validate).Execute()

Claim a subscription into an organization.



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
    claimAdministeredLicensingSubscriptionSubscriptionsRequest := *openapiclient.NewClaimAdministeredLicensingSubscriptionSubscriptionsRequest("ClaimKey_example", "OrganizationId_example") // ClaimAdministeredLicensingSubscriptionSubscriptionsRequest | 
    validate := true // bool | Check if the provided claim key is valid and can be claimed into the organization. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsApi.ClaimAdministeredLicensingSubscriptionSubscriptions(context.Background()).ClaimAdministeredLicensingSubscriptionSubscriptionsRequest(claimAdministeredLicensingSubscriptionSubscriptionsRequest).Validate(validate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.ClaimAdministeredLicensingSubscriptionSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimAdministeredLicensingSubscriptionSubscriptions`: GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.ClaimAdministeredLicensingSubscriptionSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClaimAdministeredLicensingSubscriptionSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **claimAdministeredLicensingSubscriptionSubscriptionsRequest** | [**ClaimAdministeredLicensingSubscriptionSubscriptionsRequest**](ClaimAdministeredLicensingSubscriptionSubscriptionsRequest.md) |  | 
 **validate** | **bool** | Check if the provided claim key is valid and can be claimed into the organization. | 

### Return type

[**GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner**](GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdministeredLicensingSubscriptionSubscriptions

> []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner GetAdministeredLicensingSubscriptionSubscriptions(ctx).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SubscriptionIds(subscriptionIds).OrganizationIds(organizationIds).Statuses(statuses).ProductTypes(productTypes).StartDate(startDate).EndDate(endDate).Execute()

List available subscriptions



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    subscriptionIds := []string{"Inner_example"} // []string | List of subscription ids to fetch (optional)
    organizationIds := []string{"Inner_example"} // []string | Organizations to get associated subscriptions for (optional)
    statuses := []string{"Statuses_example"} // []string | List of statuses that returned subscriptions can have (optional)
    productTypes := []string{"ProductTypes_example"} // []string | List of product types that returned subscriptions need to have entitlements for. (optional)
    startDate := openapiclient.getAdministeredLicensingSubscriptionSubscriptions_startDate_parameter{GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf: openapiclient.NewGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf()} // GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameter | Filter subscriptions by start date, ISO 8601 format. To filter with a range of dates, use 'startDate[<option>]=?' in the request. Accepted options include lt, gt, lte, gte. (optional)
    endDate := openapiclient.getAdministeredLicensingSubscriptionSubscriptions_startDate_parameter{GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf: openapiclient.NewGetAdministeredLicensingSubscriptionSubscriptionsStartDateParameterOneOf()} // GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameter | Filter subscriptions by end date, ISO 8601 format. To filter with a range of dates, use 'endDate[<option>]=?' in the request. Accepted options include lt, gt, lte, gte. (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsApi.GetAdministeredLicensingSubscriptionSubscriptions(context.Background()).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SubscriptionIds(subscriptionIds).OrganizationIds(organizationIds).Statuses(statuses).ProductTypes(productTypes).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.GetAdministeredLicensingSubscriptionSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdministeredLicensingSubscriptionSubscriptions`: []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.GetAdministeredLicensingSubscriptionSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdministeredLicensingSubscriptionSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **subscriptionIds** | **[]string** | List of subscription ids to fetch | 
 **organizationIds** | **[]string** | Organizations to get associated subscriptions for | 
 **statuses** | **[]string** | List of statuses that returned subscriptions can have | 
 **productTypes** | **[]string** | List of product types that returned subscriptions need to have entitlements for. | 
 **startDate** | [**GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameter**](GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameter.md) | Filter subscriptions by start date, ISO 8601 format. To filter with a range of dates, use &#39;startDate[&lt;option&gt;]&#x3D;?&#39; in the request. Accepted options include lt, gt, lte, gte. | 
 **endDate** | [**GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameter**](GetAdministeredLicensingSubscriptionSubscriptionsStartDateParameter.md) | Filter subscriptions by end date, ISO 8601 format. To filter with a range of dates, use &#39;endDate[&lt;option&gt;]&#x3D;?&#39; in the request. Accepted options include lt, gt, lte, gte. | 

### Return type

[**[]GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner**](GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey

> GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(ctx).ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest(validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest).Execute()

Find a subscription by claim key



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
    validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest := *openapiclient.NewValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest("ClaimKey_example") // ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionsApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(context.Background()).ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest(validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey`: GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionsApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest** | [**ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest**](ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest.md) |  | 

### Return type

[**GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner**](GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

