# \LicensingApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BindAdministeredLicensingSubscriptionSubscription**](LicensingApi.md#BindAdministeredLicensingSubscriptionSubscription) | **Post** /administered/licensing/subscription/subscriptions/{subscriptionId}/bind | Bind networks to a subscription
[**ClaimAdministeredLicensingSubscriptionSubscriptions**](LicensingApi.md#ClaimAdministeredLicensingSubscriptionSubscriptions) | **Post** /administered/licensing/subscription/subscriptions/claim | Claim a subscription into an organization.
[**GetAdministeredLicensingSubscriptionSubscriptions**](LicensingApi.md#GetAdministeredLicensingSubscriptionSubscriptions) | **Get** /administered/licensing/subscription/subscriptions | List available subscriptions
[**GetOrganizationLicensingCotermLicenses**](LicensingApi.md#GetOrganizationLicensingCotermLicenses) | **Get** /organizations/{organizationId}/licensing/coterm/licenses | List the licenses in a coterm organization
[**MoveOrganizationLicensingCotermLicenses**](LicensingApi.md#MoveOrganizationLicensingCotermLicenses) | **Post** /organizations/{organizationId}/licensing/coterm/licenses/move | Moves a license to a different organization (coterm only)
[**ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey**](LicensingApi.md#ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey) | **Post** /administered/licensing/subscription/subscriptions/claimKey/validate | Find a subscription by claim key



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
    resp, r, err := apiClient.LicensingApi.BindAdministeredLicensingSubscriptionSubscription(context.Background(), subscriptionId).BindAdministeredLicensingSubscriptionSubscriptionRequest(bindAdministeredLicensingSubscriptionSubscriptionRequest).Validate(validate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingApi.BindAdministeredLicensingSubscriptionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BindAdministeredLicensingSubscriptionSubscription`: BindAdministeredLicensingSubscriptionSubscription200Response
    fmt.Fprintf(os.Stdout, "Response from `LicensingApi.BindAdministeredLicensingSubscriptionSubscription`: %v\n", resp)
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

> GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner ClaimAdministeredLicensingSubscriptionSubscriptions(ctx).ClaimAdministeredLicensingSubscriptionSubscriptionsRequest(claimAdministeredLicensingSubscriptionSubscriptionsRequest).Execute()

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

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingApi.ClaimAdministeredLicensingSubscriptionSubscriptions(context.Background()).ClaimAdministeredLicensingSubscriptionSubscriptionsRequest(claimAdministeredLicensingSubscriptionSubscriptionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingApi.ClaimAdministeredLicensingSubscriptionSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClaimAdministeredLicensingSubscriptionSubscriptions`: GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LicensingApi.ClaimAdministeredLicensingSubscriptionSubscriptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClaimAdministeredLicensingSubscriptionSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **claimAdministeredLicensingSubscriptionSubscriptionsRequest** | [**ClaimAdministeredLicensingSubscriptionSubscriptionsRequest**](ClaimAdministeredLicensingSubscriptionSubscriptionsRequest.md) |  | 

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
    resp, r, err := apiClient.LicensingApi.GetAdministeredLicensingSubscriptionSubscriptions(context.Background()).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SubscriptionIds(subscriptionIds).OrganizationIds(organizationIds).Statuses(statuses).ProductTypes(productTypes).StartDate(startDate).EndDate(endDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingApi.GetAdministeredLicensingSubscriptionSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdministeredLicensingSubscriptionSubscriptions`: []GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LicensingApi.GetAdministeredLicensingSubscriptionSubscriptions`: %v\n", resp)
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


## GetOrganizationLicensingCotermLicenses

> []GetOrganizationLicensingCotermLicenses200ResponseInner GetOrganizationLicensingCotermLicenses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Invalidated(invalidated).Expired(expired).Execute()

List the licenses in a coterm organization



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
    perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
    startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
    invalidated := true // bool | Filter for licenses that are invalidated (optional)
    expired := true // bool | Filter for licenses that are expired (optional)

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingApi.GetOrganizationLicensingCotermLicenses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).Invalidated(invalidated).Expired(expired).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingApi.GetOrganizationLicensingCotermLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationLicensingCotermLicenses`: []GetOrganizationLicensingCotermLicenses200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LicensingApi.GetOrganizationLicensingCotermLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationLicensingCotermLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **invalidated** | **bool** | Filter for licenses that are invalidated | 
 **expired** | **bool** | Filter for licenses that are expired | 

### Return type

[**[]GetOrganizationLicensingCotermLicenses200ResponseInner**](GetOrganizationLicensingCotermLicenses200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganizationLicensingCotermLicenses

> MoveOrganizationLicensingCotermLicenses200Response MoveOrganizationLicensingCotermLicenses(ctx, organizationId).MoveOrganizationLicensingCotermLicensesRequest(moveOrganizationLicensingCotermLicensesRequest).Execute()

Moves a license to a different organization (coterm only)



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
    moveOrganizationLicensingCotermLicensesRequest := *openapiclient.NewMoveOrganizationLicensingCotermLicensesRequest(*openapiclient.NewMoveOrganizationLicensingCotermLicensesRequestDestination(), []openapiclient.MoveOrganizationLicensingCotermLicensesRequestLicensesInner{*openapiclient.NewMoveOrganizationLicensingCotermLicensesRequestLicensesInner("Key_example", []openapiclient.MoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner{*openapiclient.NewMoveOrganizationLicensingCotermLicensesRequestLicensesInnerCountsInner("Model_example", int32(123))})}) // MoveOrganizationLicensingCotermLicensesRequest | 

    configuration := openapiclient.NewConfiguration()

    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LicensingApi.MoveOrganizationLicensingCotermLicenses(context.Background(), organizationId).MoveOrganizationLicensingCotermLicensesRequest(moveOrganizationLicensingCotermLicensesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingApi.MoveOrganizationLicensingCotermLicenses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveOrganizationLicensingCotermLicenses`: MoveOrganizationLicensingCotermLicenses200Response
    fmt.Fprintf(os.Stdout, "Response from `LicensingApi.MoveOrganizationLicensingCotermLicenses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationLicensingCotermLicensesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveOrganizationLicensingCotermLicensesRequest** | [**MoveOrganizationLicensingCotermLicensesRequest**](MoveOrganizationLicensingCotermLicensesRequest.md) |  | 

### Return type

[**MoveOrganizationLicensingCotermLicenses200Response**](MoveOrganizationLicensingCotermLicenses200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
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
    resp, r, err := apiClient.LicensingApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey(context.Background()).ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest(validateAdministeredLicensingSubscriptionSubscriptionsClaimKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicensingApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey`: GetAdministeredLicensingSubscriptionSubscriptions200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `LicensingApi.ValidateAdministeredLicensingSubscriptionSubscriptionsClaimKey`: %v\n", resp)
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

