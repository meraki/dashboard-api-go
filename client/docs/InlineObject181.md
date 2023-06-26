# InlineObject181

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the dashboard administrator. This attribute can not be updated. | 
**Name** | **string** | The name of the dashboard administrator | 
**OrgAccess** | **string** | The privilege of the dashboard administrator on the organization. Can be one of &#39;full&#39;, &#39;read-only&#39;, &#39;enterprise&#39; or &#39;none&#39; | 
**Tags** | Pointer to [**[]OrganizationsOrganizationIdAdminsTags1**](OrganizationsOrganizationIdAdminsTags1.md) | The list of tags that the dashboard administrator has privileges on | [optional] 
**Networks** | Pointer to [**[]OrganizationsOrganizationIdAdminsNetworks1**](OrganizationsOrganizationIdAdminsNetworks1.md) | The list of networks that the dashboard administrator has privileges on | [optional] 
**AuthenticationMethod** | Pointer to **string** | The method of authentication the user will use to sign in to the Meraki dashboard. Can be one of &#39;Email&#39; or &#39;Cisco SecureX Sign-On&#39;. The default is Email authentication | [optional] 

## Methods

### NewInlineObject181

`func NewInlineObject181(email string, name string, orgAccess string, ) *InlineObject181`

NewInlineObject181 instantiates a new InlineObject181 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject181WithDefaults

`func NewInlineObject181WithDefaults() *InlineObject181`

NewInlineObject181WithDefaults instantiates a new InlineObject181 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InlineObject181) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineObject181) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineObject181) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *InlineObject181) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject181) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject181) SetName(v string)`

SetName sets Name field to given value.


### GetOrgAccess

`func (o *InlineObject181) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *InlineObject181) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *InlineObject181) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.


### GetTags

`func (o *InlineObject181) GetTags() []OrganizationsOrganizationIdAdminsTags1`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject181) GetTagsOk() (*[]OrganizationsOrganizationIdAdminsTags1, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject181) SetTags(v []OrganizationsOrganizationIdAdminsTags1)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineObject181) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworks

`func (o *InlineObject181) GetNetworks() []OrganizationsOrganizationIdAdminsNetworks1`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *InlineObject181) GetNetworksOk() (*[]OrganizationsOrganizationIdAdminsNetworks1, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *InlineObject181) SetNetworks(v []OrganizationsOrganizationIdAdminsNetworks1)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *InlineObject181) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *InlineObject181) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *InlineObject181) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *InlineObject181) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *InlineObject181) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


