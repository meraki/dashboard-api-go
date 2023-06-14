# UpdateOrganizationSamlRole200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID associated with the SAML role | [optional] 
**Role** | Pointer to **string** | The role of the SAML administrator | [optional] 
**OrgAccess** | Pointer to **string** | The privilege of the SAML administrator on the organization | [optional] 
**Networks** | Pointer to [**[]UpdateOrganizationSamlRole200ResponseNetworksInner**](UpdateOrganizationSamlRole200ResponseNetworksInner.md) | The list of networks that the SAML administrator has privileges on | [optional] 
**Tags** | Pointer to [**[]UpdateOrganizationSamlRole200ResponseTagsInner**](UpdateOrganizationSamlRole200ResponseTagsInner.md) | The list of tags that the SAML administrator has privleges on | [optional] 

## Methods

### NewUpdateOrganizationSamlRole200Response

`func NewUpdateOrganizationSamlRole200Response() *UpdateOrganizationSamlRole200Response`

NewUpdateOrganizationSamlRole200Response instantiates a new UpdateOrganizationSamlRole200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrganizationSamlRole200ResponseWithDefaults

`func NewUpdateOrganizationSamlRole200ResponseWithDefaults() *UpdateOrganizationSamlRole200Response`

NewUpdateOrganizationSamlRole200ResponseWithDefaults instantiates a new UpdateOrganizationSamlRole200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateOrganizationSamlRole200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrganizationSamlRole200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrganizationSamlRole200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateOrganizationSamlRole200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *UpdateOrganizationSamlRole200Response) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateOrganizationSamlRole200Response) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateOrganizationSamlRole200Response) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateOrganizationSamlRole200Response) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetOrgAccess

`func (o *UpdateOrganizationSamlRole200Response) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *UpdateOrganizationSamlRole200Response) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *UpdateOrganizationSamlRole200Response) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.

### HasOrgAccess

`func (o *UpdateOrganizationSamlRole200Response) HasOrgAccess() bool`

HasOrgAccess returns a boolean if a field has been set.

### GetNetworks

`func (o *UpdateOrganizationSamlRole200Response) GetNetworks() []UpdateOrganizationSamlRole200ResponseNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *UpdateOrganizationSamlRole200Response) GetNetworksOk() (*[]UpdateOrganizationSamlRole200ResponseNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *UpdateOrganizationSamlRole200Response) SetNetworks(v []UpdateOrganizationSamlRole200ResponseNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *UpdateOrganizationSamlRole200Response) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetTags

`func (o *UpdateOrganizationSamlRole200Response) GetTags() []UpdateOrganizationSamlRole200ResponseTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateOrganizationSamlRole200Response) GetTagsOk() (*[]UpdateOrganizationSamlRole200ResponseTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateOrganizationSamlRole200Response) SetTags(v []UpdateOrganizationSamlRole200ResponseTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateOrganizationSamlRole200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


