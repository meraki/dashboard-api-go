# CreateOrganizationSamlRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | The role of the SAML administrator | 
**OrgAccess** | **string** | The privilege of the SAML administrator on the organization. Can be one of &#39;none&#39;, &#39;read-only&#39;, &#39;full&#39; or &#39;enterprise&#39; | 
**Tags** | Pointer to [**[]CreateOrganizationSamlRoleRequestTagsInner**](CreateOrganizationSamlRoleRequestTagsInner.md) | The list of tags that the SAML administrator has privleges on | [optional] 
**Networks** | Pointer to [**[]CreateOrganizationSamlRoleRequestNetworksInner**](CreateOrganizationSamlRoleRequestNetworksInner.md) | The list of networks that the SAML administrator has privileges on | [optional] 

## Methods

### NewCreateOrganizationSamlRoleRequest

`func NewCreateOrganizationSamlRoleRequest(role string, orgAccess string, ) *CreateOrganizationSamlRoleRequest`

NewCreateOrganizationSamlRoleRequest instantiates a new CreateOrganizationSamlRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrganizationSamlRoleRequestWithDefaults

`func NewCreateOrganizationSamlRoleRequestWithDefaults() *CreateOrganizationSamlRoleRequest`

NewCreateOrganizationSamlRoleRequestWithDefaults instantiates a new CreateOrganizationSamlRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *CreateOrganizationSamlRoleRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateOrganizationSamlRoleRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateOrganizationSamlRoleRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetOrgAccess

`func (o *CreateOrganizationSamlRoleRequest) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *CreateOrganizationSamlRoleRequest) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *CreateOrganizationSamlRoleRequest) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.


### GetTags

`func (o *CreateOrganizationSamlRoleRequest) GetTags() []CreateOrganizationSamlRoleRequestTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateOrganizationSamlRoleRequest) GetTagsOk() (*[]CreateOrganizationSamlRoleRequestTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateOrganizationSamlRoleRequest) SetTags(v []CreateOrganizationSamlRoleRequestTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateOrganizationSamlRoleRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworks

`func (o *CreateOrganizationSamlRoleRequest) GetNetworks() []CreateOrganizationSamlRoleRequestNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *CreateOrganizationSamlRoleRequest) GetNetworksOk() (*[]CreateOrganizationSamlRoleRequestNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *CreateOrganizationSamlRoleRequest) SetNetworks(v []CreateOrganizationSamlRoleRequestNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *CreateOrganizationSamlRoleRequest) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


