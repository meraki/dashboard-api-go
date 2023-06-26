# GetOrganizationAdmins200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Admin&#39;s ID | [optional] 
**Name** | Pointer to **string** | Admin&#39;s username | [optional] 
**Email** | Pointer to **string** | Admin&#39;s email address | [optional] 
**OrgAccess** | Pointer to **string** | Admin&#39;s level of access to the organization | [optional] 
**AccountStatus** | Pointer to **string** | Status of the admin&#39;s account | [optional] 
**TwoFactorAuthEnabled** | Pointer to **bool** | Indicates whether two-factor authentication is enabled | [optional] 
**HasApiKey** | Pointer to **bool** | Indicates whether the admin has an API key | [optional] 
**LastActive** | Pointer to **time.Time** | Time when the admin was last active | [optional] 
**Tags** | Pointer to [**[]GetOrganizationAdmins200ResponseInnerTagsInner**](GetOrganizationAdmins200ResponseInnerTagsInner.md) | Admin tag information | [optional] 
**Networks** | Pointer to [**[]GetOrganizationAdmins200ResponseInnerNetworksInner**](GetOrganizationAdmins200ResponseInnerNetworksInner.md) | Admin network access information | [optional] 
**AuthenticationMethod** | Pointer to **string** | Admin&#39;s authentication method | [optional] 

## Methods

### NewGetOrganizationAdmins200ResponseInner

`func NewGetOrganizationAdmins200ResponseInner() *GetOrganizationAdmins200ResponseInner`

NewGetOrganizationAdmins200ResponseInner instantiates a new GetOrganizationAdmins200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationAdmins200ResponseInnerWithDefaults

`func NewGetOrganizationAdmins200ResponseInnerWithDefaults() *GetOrganizationAdmins200ResponseInner`

NewGetOrganizationAdmins200ResponseInnerWithDefaults instantiates a new GetOrganizationAdmins200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrganizationAdmins200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrganizationAdmins200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrganizationAdmins200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetOrganizationAdmins200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetOrganizationAdmins200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOrganizationAdmins200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOrganizationAdmins200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOrganizationAdmins200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *GetOrganizationAdmins200ResponseInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetOrganizationAdmins200ResponseInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetOrganizationAdmins200ResponseInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetOrganizationAdmins200ResponseInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetOrgAccess

`func (o *GetOrganizationAdmins200ResponseInner) GetOrgAccess() string`

GetOrgAccess returns the OrgAccess field if non-nil, zero value otherwise.

### GetOrgAccessOk

`func (o *GetOrganizationAdmins200ResponseInner) GetOrgAccessOk() (*string, bool)`

GetOrgAccessOk returns a tuple with the OrgAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgAccess

`func (o *GetOrganizationAdmins200ResponseInner) SetOrgAccess(v string)`

SetOrgAccess sets OrgAccess field to given value.

### HasOrgAccess

`func (o *GetOrganizationAdmins200ResponseInner) HasOrgAccess() bool`

HasOrgAccess returns a boolean if a field has been set.

### GetAccountStatus

`func (o *GetOrganizationAdmins200ResponseInner) GetAccountStatus() string`

GetAccountStatus returns the AccountStatus field if non-nil, zero value otherwise.

### GetAccountStatusOk

`func (o *GetOrganizationAdmins200ResponseInner) GetAccountStatusOk() (*string, bool)`

GetAccountStatusOk returns a tuple with the AccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStatus

`func (o *GetOrganizationAdmins200ResponseInner) SetAccountStatus(v string)`

SetAccountStatus sets AccountStatus field to given value.

### HasAccountStatus

`func (o *GetOrganizationAdmins200ResponseInner) HasAccountStatus() bool`

HasAccountStatus returns a boolean if a field has been set.

### GetTwoFactorAuthEnabled

`func (o *GetOrganizationAdmins200ResponseInner) GetTwoFactorAuthEnabled() bool`

GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field if non-nil, zero value otherwise.

### GetTwoFactorAuthEnabledOk

`func (o *GetOrganizationAdmins200ResponseInner) GetTwoFactorAuthEnabledOk() (*bool, bool)`

GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthEnabled

`func (o *GetOrganizationAdmins200ResponseInner) SetTwoFactorAuthEnabled(v bool)`

SetTwoFactorAuthEnabled sets TwoFactorAuthEnabled field to given value.

### HasTwoFactorAuthEnabled

`func (o *GetOrganizationAdmins200ResponseInner) HasTwoFactorAuthEnabled() bool`

HasTwoFactorAuthEnabled returns a boolean if a field has been set.

### GetHasApiKey

`func (o *GetOrganizationAdmins200ResponseInner) GetHasApiKey() bool`

GetHasApiKey returns the HasApiKey field if non-nil, zero value otherwise.

### GetHasApiKeyOk

`func (o *GetOrganizationAdmins200ResponseInner) GetHasApiKeyOk() (*bool, bool)`

GetHasApiKeyOk returns a tuple with the HasApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasApiKey

`func (o *GetOrganizationAdmins200ResponseInner) SetHasApiKey(v bool)`

SetHasApiKey sets HasApiKey field to given value.

### HasHasApiKey

`func (o *GetOrganizationAdmins200ResponseInner) HasHasApiKey() bool`

HasHasApiKey returns a boolean if a field has been set.

### GetLastActive

`func (o *GetOrganizationAdmins200ResponseInner) GetLastActive() time.Time`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *GetOrganizationAdmins200ResponseInner) GetLastActiveOk() (*time.Time, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *GetOrganizationAdmins200ResponseInner) SetLastActive(v time.Time)`

SetLastActive sets LastActive field to given value.

### HasLastActive

`func (o *GetOrganizationAdmins200ResponseInner) HasLastActive() bool`

HasLastActive returns a boolean if a field has been set.

### GetTags

`func (o *GetOrganizationAdmins200ResponseInner) GetTags() []GetOrganizationAdmins200ResponseInnerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetOrganizationAdmins200ResponseInner) GetTagsOk() (*[]GetOrganizationAdmins200ResponseInnerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetOrganizationAdmins200ResponseInner) SetTags(v []GetOrganizationAdmins200ResponseInnerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetOrganizationAdmins200ResponseInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetNetworks

`func (o *GetOrganizationAdmins200ResponseInner) GetNetworks() []GetOrganizationAdmins200ResponseInnerNetworksInner`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetOrganizationAdmins200ResponseInner) GetNetworksOk() (*[]GetOrganizationAdmins200ResponseInnerNetworksInner, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetOrganizationAdmins200ResponseInner) SetNetworks(v []GetOrganizationAdmins200ResponseInnerNetworksInner)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetOrganizationAdmins200ResponseInner) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetAuthenticationMethod

`func (o *GetOrganizationAdmins200ResponseInner) GetAuthenticationMethod() string`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *GetOrganizationAdmins200ResponseInner) GetAuthenticationMethodOk() (*string, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *GetOrganizationAdmins200ResponseInner) SetAuthenticationMethod(v string)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.

### HasAuthenticationMethod

`func (o *GetOrganizationAdmins200ResponseInner) HasAuthenticationMethod() bool`

HasAuthenticationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


