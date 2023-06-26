# \LicensingApi

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationLicensingCotermLicenses**](LicensingApi.md#GetOrganizationLicensingCotermLicenses) | **Get** /organizations/{organizationId}/licensing/coterm/licenses | List the licenses in a coterm organization
[**MoveOrganizationLicensingCotermLicenses**](LicensingApi.md#MoveOrganizationLicensingCotermLicenses) | **Post** /organizations/{organizationId}/licensing/coterm/licenses/move | Moves a license to a different organization (coterm only)



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

[meraki_api_key](../README.md#meraki_api_key)

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

[meraki_api_key](../README.md#meraki_api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

