# InlineResponse20062

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Meraki managed Id of the user record. | [optional] 
**Email** | Pointer to **string** | User email. | [optional] 
**FullName** | Pointer to **string** | User full name. | [optional] 
**Username** | Pointer to **string** | The users username. | [optional] 
**HasPassword** | Pointer to **bool** | A boolean denoting if the user has a password associated with the record. | [optional] 
**Tags** | Pointer to **string** | The set of tags the user is scoped to. | [optional] 
**AdGroups** | Pointer to **[]string** | Active Directory Groups the user belongs to. | [optional] 
**AzureAdGroups** | Pointer to **[]string** | Azure Active Directory Groups the user belongs to. | [optional] 
**SamlGroups** | Pointer to **[]string** | SAML Groups the user belongs to. | [optional] 
**AsmGroups** | Pointer to **[]string** | Apple School Manager Groups the user belongs to. | [optional] 
**IsExternal** | Pointer to **bool** | Whether the user was created using an external integration, or via the Meraki Dashboard. | [optional] 
**DisplayName** | Pointer to **string** | The user display name. | [optional] 
**HasIdentityCertificate** | Pointer to **bool** | A boolean indicating if the user has an associated identity certificate.. | [optional] 
**UserThumbnail** | Pointer to **string** | The url where the users thumbnail is hosted. | [optional] 

## Methods

### NewInlineResponse20062

`func NewInlineResponse20062() *InlineResponse20062`

NewInlineResponse20062 instantiates a new InlineResponse20062 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20062WithDefaults

`func NewInlineResponse20062WithDefaults() *InlineResponse20062`

NewInlineResponse20062WithDefaults instantiates a new InlineResponse20062 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20062) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20062) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20062) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20062) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse20062) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse20062) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse20062) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse20062) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullName

`func (o *InlineResponse20062) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *InlineResponse20062) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *InlineResponse20062) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *InlineResponse20062) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetUsername

`func (o *InlineResponse20062) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InlineResponse20062) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InlineResponse20062) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InlineResponse20062) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetHasPassword

`func (o *InlineResponse20062) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *InlineResponse20062) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *InlineResponse20062) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *InlineResponse20062) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20062) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20062) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20062) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20062) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAdGroups

`func (o *InlineResponse20062) GetAdGroups() []string`

GetAdGroups returns the AdGroups field if non-nil, zero value otherwise.

### GetAdGroupsOk

`func (o *InlineResponse20062) GetAdGroupsOk() (*[]string, bool)`

GetAdGroupsOk returns a tuple with the AdGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdGroups

`func (o *InlineResponse20062) SetAdGroups(v []string)`

SetAdGroups sets AdGroups field to given value.

### HasAdGroups

`func (o *InlineResponse20062) HasAdGroups() bool`

HasAdGroups returns a boolean if a field has been set.

### GetAzureAdGroups

`func (o *InlineResponse20062) GetAzureAdGroups() []string`

GetAzureAdGroups returns the AzureAdGroups field if non-nil, zero value otherwise.

### GetAzureAdGroupsOk

`func (o *InlineResponse20062) GetAzureAdGroupsOk() (*[]string, bool)`

GetAzureAdGroupsOk returns a tuple with the AzureAdGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdGroups

`func (o *InlineResponse20062) SetAzureAdGroups(v []string)`

SetAzureAdGroups sets AzureAdGroups field to given value.

### HasAzureAdGroups

`func (o *InlineResponse20062) HasAzureAdGroups() bool`

HasAzureAdGroups returns a boolean if a field has been set.

### GetSamlGroups

`func (o *InlineResponse20062) GetSamlGroups() []string`

GetSamlGroups returns the SamlGroups field if non-nil, zero value otherwise.

### GetSamlGroupsOk

`func (o *InlineResponse20062) GetSamlGroupsOk() (*[]string, bool)`

GetSamlGroupsOk returns a tuple with the SamlGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlGroups

`func (o *InlineResponse20062) SetSamlGroups(v []string)`

SetSamlGroups sets SamlGroups field to given value.

### HasSamlGroups

`func (o *InlineResponse20062) HasSamlGroups() bool`

HasSamlGroups returns a boolean if a field has been set.

### GetAsmGroups

`func (o *InlineResponse20062) GetAsmGroups() []string`

GetAsmGroups returns the AsmGroups field if non-nil, zero value otherwise.

### GetAsmGroupsOk

`func (o *InlineResponse20062) GetAsmGroupsOk() (*[]string, bool)`

GetAsmGroupsOk returns a tuple with the AsmGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmGroups

`func (o *InlineResponse20062) SetAsmGroups(v []string)`

SetAsmGroups sets AsmGroups field to given value.

### HasAsmGroups

`func (o *InlineResponse20062) HasAsmGroups() bool`

HasAsmGroups returns a boolean if a field has been set.

### GetIsExternal

`func (o *InlineResponse20062) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *InlineResponse20062) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *InlineResponse20062) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *InlineResponse20062) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetDisplayName

`func (o *InlineResponse20062) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InlineResponse20062) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InlineResponse20062) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InlineResponse20062) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHasIdentityCertificate

`func (o *InlineResponse20062) GetHasIdentityCertificate() bool`

GetHasIdentityCertificate returns the HasIdentityCertificate field if non-nil, zero value otherwise.

### GetHasIdentityCertificateOk

`func (o *InlineResponse20062) GetHasIdentityCertificateOk() (*bool, bool)`

GetHasIdentityCertificateOk returns a tuple with the HasIdentityCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIdentityCertificate

`func (o *InlineResponse20062) SetHasIdentityCertificate(v bool)`

SetHasIdentityCertificate sets HasIdentityCertificate field to given value.

### HasHasIdentityCertificate

`func (o *InlineResponse20062) HasHasIdentityCertificate() bool`

HasHasIdentityCertificate returns a boolean if a field has been set.

### GetUserThumbnail

`func (o *InlineResponse20062) GetUserThumbnail() string`

GetUserThumbnail returns the UserThumbnail field if non-nil, zero value otherwise.

### GetUserThumbnailOk

`func (o *InlineResponse20062) GetUserThumbnailOk() (*string, bool)`

GetUserThumbnailOk returns a tuple with the UserThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserThumbnail

`func (o *InlineResponse20062) SetUserThumbnail(v string)`

SetUserThumbnail sets UserThumbnail field to given value.

### HasUserThumbnail

`func (o *InlineResponse20062) HasUserThumbnail() bool`

HasUserThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


