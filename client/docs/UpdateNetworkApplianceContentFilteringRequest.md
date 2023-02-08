# UpdateNetworkApplianceContentFilteringRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedUrlPatterns** | Pointer to **[]string** | A list of URL patterns that are allowed | [optional] 
**BlockedUrlPatterns** | Pointer to **[]string** | A list of URL patterns that are blocked | [optional] 
**BlockedUrlCategories** | Pointer to **[]string** | A list of URL categories to block | [optional] 
**UrlCategoryListSize** | Pointer to **string** | URL category list size which is either &#39;topSites&#39; or &#39;fullList&#39; | [optional] 

## Methods

### NewUpdateNetworkApplianceContentFilteringRequest

`func NewUpdateNetworkApplianceContentFilteringRequest() *UpdateNetworkApplianceContentFilteringRequest`

NewUpdateNetworkApplianceContentFilteringRequest instantiates a new UpdateNetworkApplianceContentFilteringRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkApplianceContentFilteringRequestWithDefaults

`func NewUpdateNetworkApplianceContentFilteringRequestWithDefaults() *UpdateNetworkApplianceContentFilteringRequest`

NewUpdateNetworkApplianceContentFilteringRequestWithDefaults instantiates a new UpdateNetworkApplianceContentFilteringRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedUrlPatterns

`func (o *UpdateNetworkApplianceContentFilteringRequest) GetAllowedUrlPatterns() []string`

GetAllowedUrlPatterns returns the AllowedUrlPatterns field if non-nil, zero value otherwise.

### GetAllowedUrlPatternsOk

`func (o *UpdateNetworkApplianceContentFilteringRequest) GetAllowedUrlPatternsOk() (*[]string, bool)`

GetAllowedUrlPatternsOk returns a tuple with the AllowedUrlPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrlPatterns

`func (o *UpdateNetworkApplianceContentFilteringRequest) SetAllowedUrlPatterns(v []string)`

SetAllowedUrlPatterns sets AllowedUrlPatterns field to given value.

### HasAllowedUrlPatterns

`func (o *UpdateNetworkApplianceContentFilteringRequest) HasAllowedUrlPatterns() bool`

HasAllowedUrlPatterns returns a boolean if a field has been set.

### GetBlockedUrlPatterns

`func (o *UpdateNetworkApplianceContentFilteringRequest) GetBlockedUrlPatterns() []string`

GetBlockedUrlPatterns returns the BlockedUrlPatterns field if non-nil, zero value otherwise.

### GetBlockedUrlPatternsOk

`func (o *UpdateNetworkApplianceContentFilteringRequest) GetBlockedUrlPatternsOk() (*[]string, bool)`

GetBlockedUrlPatternsOk returns a tuple with the BlockedUrlPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUrlPatterns

`func (o *UpdateNetworkApplianceContentFilteringRequest) SetBlockedUrlPatterns(v []string)`

SetBlockedUrlPatterns sets BlockedUrlPatterns field to given value.

### HasBlockedUrlPatterns

`func (o *UpdateNetworkApplianceContentFilteringRequest) HasBlockedUrlPatterns() bool`

HasBlockedUrlPatterns returns a boolean if a field has been set.

### GetBlockedUrlCategories

`func (o *UpdateNetworkApplianceContentFilteringRequest) GetBlockedUrlCategories() []string`

GetBlockedUrlCategories returns the BlockedUrlCategories field if non-nil, zero value otherwise.

### GetBlockedUrlCategoriesOk

`func (o *UpdateNetworkApplianceContentFilteringRequest) GetBlockedUrlCategoriesOk() (*[]string, bool)`

GetBlockedUrlCategoriesOk returns a tuple with the BlockedUrlCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUrlCategories

`func (o *UpdateNetworkApplianceContentFilteringRequest) SetBlockedUrlCategories(v []string)`

SetBlockedUrlCategories sets BlockedUrlCategories field to given value.

### HasBlockedUrlCategories

`func (o *UpdateNetworkApplianceContentFilteringRequest) HasBlockedUrlCategories() bool`

HasBlockedUrlCategories returns a boolean if a field has been set.

### GetUrlCategoryListSize

`func (o *UpdateNetworkApplianceContentFilteringRequest) GetUrlCategoryListSize() string`

GetUrlCategoryListSize returns the UrlCategoryListSize field if non-nil, zero value otherwise.

### GetUrlCategoryListSizeOk

`func (o *UpdateNetworkApplianceContentFilteringRequest) GetUrlCategoryListSizeOk() (*string, bool)`

GetUrlCategoryListSizeOk returns a tuple with the UrlCategoryListSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlCategoryListSize

`func (o *UpdateNetworkApplianceContentFilteringRequest) SetUrlCategoryListSize(v string)`

SetUrlCategoryListSize sets UrlCategoryListSize field to given value.

### HasUrlCategoryListSize

`func (o *UpdateNetworkApplianceContentFilteringRequest) HasUrlCategoryListSize() bool`

HasUrlCategoryListSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


