# GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overall** | Pointer to **string** | Overall status of network | [optional] 
**ByProductType** | Pointer to [**[]GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner**](GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner.md) | List of status counts by product type | [optional] 

## Methods

### NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses

`func NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses() *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses`

NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses instantiates a new GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesWithDefaults

`func NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesWithDefaults() *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses`

NewGetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesWithDefaults instantiates a new GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverall

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses) GetOverall() string`

GetOverall returns the Overall field if non-nil, zero value otherwise.

### GetOverallOk

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses) GetOverallOk() (*string, bool)`

GetOverallOk returns a tuple with the Overall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverall

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses) SetOverall(v string)`

SetOverall sets Overall field to given value.

### HasOverall

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses) HasOverall() bool`

HasOverall returns a boolean if a field has been set.

### GetByProductType

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses) GetByProductType() []GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner`

GetByProductType returns the ByProductType field if non-nil, zero value otherwise.

### GetByProductTypeOk

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses) GetByProductTypeOk() (*[]GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner, bool)`

GetByProductTypeOk returns a tuple with the ByProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByProductType

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses) SetByProductType(v []GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatusesByProductTypeInner)`

SetByProductType sets ByProductType field to given value.

### HasByProductType

`func (o *GetOrganizationSummaryTopNetworksByStatus200ResponseInnerStatuses) HasByProductType() bool`

HasByProductType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


