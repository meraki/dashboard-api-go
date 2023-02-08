# UpdateOrganizationAdminRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the dashboard administrator | [optional] 
**OrgAccess** | Pointer to **string** | The privilege of the dashboard administrator on the organization. Can be one of &#39;full&#39;, &#39;read-only&#39;, &#39;enterprise&#39; or &#39;none&#39; | [optional] 
**Tags** | Pointer to [**[]CreateOrganizationAdminRequestTagsInner**](CreateOrganizationAdminRequestTagsInner.md) | The list of tags that the dashboard administrator has privileges on | [optional] 
**Networks** | Pointer to [**[]CreateOrganizationAdminRequestNetworksInner**](CreateOrganizationAdminRequestNetworksInner.md) | The list of networks that the dashboard administrator has privileges on | [optional] 

## Methods

### NewUpdateOrganizationAdminRequest

`func NewUpdateOrganizationAdminRequest() *UpdateOrganizationAdminRequest`

NewUpdateOrganizationAdminRequest instantiates a new UpdateOrganizationAdminRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationAdminRequestWithDefaults

`func NewUpdateOrganizationAdminRequestWithDefaults() *UpdateOrganizationAdminRequest`

NewUpdateOrganizationAdminRequestWithDefaults instantiates a new UpdateOrganizationAdminRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateOrganizationAdminRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOrganizationAdminRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOrganizationAdminRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOrganizationAdminRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgAccess

`func (o *UpdateOrganizationAdminRequest) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *UpdateOrganizationAdminRequest) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *UpdateOrganizationAdminRequest) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.

### HasOrgAccess

`func (o *UpdateOrganizationAdminRequest) HasOrgAccess() bool`

HasOrgAccess returns a boolean if a field has been set.

### GetTags

`func (o *UpdateOrganizationAdminRequest) GetTags() []CreateOrganizationAdminRequestTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateOrganizationAdminRequest) GetTagsOk() (*[]CreateOrganizationAdminRequestTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateOrganizationAdminRequest) SetTags(v []CreateOrganizationAdminRequestTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateOrganizationAdminRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworks

`func (o *UpdateOrganizationAdminRequest) GetNetworks() []CreateOrganizationAdminRequestNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *UpdateOrganizationAdminRequest) GetNetworksOk() (*[]CreateOrganizationAdminRequestNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *UpdateOrganizationAdminRequest) SetNetworks(v []CreateOrganizationAdminRequestNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *UpdateOrganizationAdminRequest) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


