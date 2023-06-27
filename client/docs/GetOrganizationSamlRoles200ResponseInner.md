# GetOrganizationSamlRoles200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID associated with the SAML role | [optional] 
**Role** | Pointer to **string** | The role of the SAML administrator | [optional] 
**OrgAccess** | Pointer to **string** | The privilege of the SAML administrator on the organization | [optional] 
**Networks** | Pointer to [**[]GetOrganizationSamlRoles200ResponseInnerNetworksInner**](GetOrganizationSamlRoles200ResponseInnerNetworksInner.md) | The list of networks that the SAML administrator has privileges on | [optional] 
**Tags** | Pointer to [**[]GetOrganizationSamlRoles200ResponseInnerTagsInner**](GetOrganizationSamlRoles200ResponseInnerTagsInner.md) | The list of tags that the SAML administrator has privleges on | [optional] 

## Methods

### NewGetOrganizationSamlRoles200ResponseInner

`func NewGetOrganizationSamlRoles200ResponseInner() *GetOrganizationSamlRoles200ResponseInner`

NewGetOrganizationSamlRoles200ResponseInner instantiates a new GetOrganizationSamlRoles200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationSamlRoles200ResponseInnerWithDefaults

`func NewGetOrganizationSamlRoles200ResponseInnerWithDefaults() *GetOrganizationSamlRoles200ResponseInner`

NewGetOrganizationSamlRoles200ResponseInnerWithDefaults instantiates a new GetOrganizationSamlRoles200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationSamlRoles200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationSamlRoles200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationSamlRoles200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationSamlRoles200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *GetOrganizationSamlRoles200ResponseInner) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetOrganizationSamlRoles200ResponseInner) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetOrganizationSamlRoles200ResponseInner) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetOrganizationSamlRoles200ResponseInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetOrgAccess

`func (o *GetOrganizationSamlRoles200ResponseInner) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *GetOrganizationSamlRoles200ResponseInner) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *GetOrganizationSamlRoles200ResponseInner) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.

### HasOrgAccess

`func (o *GetOrganizationSamlRoles200ResponseInner) HasOrgAccess() bool`

HasOrgAccess returns a boolean if a field has been set.

### GetNetworks

`func (o *GetOrganizationSamlRoles200ResponseInner) GetNetworks() []GetOrganizationSamlRoles200ResponseInnerNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetOrganizationSamlRoles200ResponseInner) GetNetworksOk() (*[]GetOrganizationSamlRoles200ResponseInnerNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetOrganizationSamlRoles200ResponseInner) SetNetworks(v []GetOrganizationSamlRoles200ResponseInnerNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetOrganizationSamlRoles200ResponseInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationSamlRoles200ResponseInner) GetTags() []GetOrganizationSamlRoles200ResponseInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationSamlRoles200ResponseInner) GetTagsOk() (*[]GetOrganizationSamlRoles200ResponseInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationSamlRoles200ResponseInner) SetTags(v []GetOrganizationSamlRoles200ResponseInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationSamlRoles200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


