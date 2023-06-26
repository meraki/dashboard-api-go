# InlineObject182

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the dashboard administrator | [optional] 
**OrgAccess** | Pointer to **string** | The privilege of the dashboard administrator on the organization. Can be one of &#39;full&#39;, &#39;read-only&#39;, &#39;enterprise&#39; or &#39;none&#39; | [optional] 
**Tags** | Pointer to [**[]OrganizationsOrganizationIdAdminsTags1**](OrganizationsOrganizationIdAdminsTags1.md) | The list of tags that the dashboard administrator has privileges on | [optional] 
**Networks** | Pointer to [**[]OrganizationsOrganizationIdAdminsNetworks1**](OrganizationsOrganizationIdAdminsNetworks1.md) | The list of networks that the dashboard administrator has privileges on | [optional] 

## Methods

### NewInlineObject182

`func NewInlineObject182() *InlineObject182`

NewInlineObject182 instantiates a new InlineObject182 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject182WithDefaults

`func NewInlineObject182WithDefaults() *InlineObject182`

NewInlineObject182WithDefaults instantiates a new InlineObject182 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InlineObject182) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject182) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject182) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject182) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgAccess

`func (o *InlineObject182) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *InlineObject182) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *InlineObject182) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.

### HasOrgAccess

`func (o *InlineObject182) HasOrgAccess() bool`

HasOrgAccess returns a boolean if a field has been set.

### GetTags

`func (o *InlineObject182) GetTags() []OrganizationsOrganizationIdAdminsTags1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject182) GetTagsOk() (*[]OrganizationsOrganizationIdAdminsTags1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject182) SetTags(v []OrganizationsOrganizationIdAdminsTags1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject182) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineObject182) GetNetworks() []OrganizationsOrganizationIdAdminsNetworks1`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineObject182) GetNetworksOk() (*[]OrganizationsOrganizationIdAdminsNetworks1, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineObject182) SetNetworks(v []OrganizationsOrganizationIdAdminsNetworks1)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineObject182) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


