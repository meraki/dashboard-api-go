# InlineResponse20065

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

### NewInlineResponse20065

`func NewInlineResponse20065() *InlineResponse20065`

NewInlineResponse20065 instantiates a new InlineResponse20065 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20065WithDefaults

`func NewInlineResponse20065WithDefaults() *InlineResponse20065`

NewInlineResponse20065WithDefaults instantiates a new InlineResponse20065 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InlineResponse20065) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InlineResponse20065) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InlineResponse20065) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InlineResponse20065) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *InlineResponse20065) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineResponse20065) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineResponse20065) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *InlineResponse20065) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFullName

`func (o *InlineResponse20065) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *InlineResponse20065) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *InlineResponse20065) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *InlineResponse20065) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetUsername

`func (o *InlineResponse20065) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InlineResponse20065) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InlineResponse20065) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InlineResponse20065) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetHasPassword

`func (o *InlineResponse20065) GetHasPassword() bool`

GetHasPassword returns the HasPassword field if non-nil, zero value otherwise.

### GetHasPasswordOk

`func (o *InlineResponse20065) GetHasPasswordOk() (*bool, bool)`

GetHasPasswordOk returns a tuple with the HasPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPassword

`func (o *InlineResponse20065) SetHasPassword(v bool)`

SetHasPassword sets HasPassword field to given value.

### HasHasPassword

`func (o *InlineResponse20065) HasHasPassword() bool`

HasHasPassword returns a boolean if a field has been set.

### GetTags

`func (o *InlineResponse20065) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineResponse20065) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineResponse20065) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InlineResponse20065) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAdGroups

`func (o *InlineResponse20065) GetAdGroups() []string`

GetAdGroups returns the AdGroups field if non-nil, zero value otherwise.

### GetAdGroupsOk

`func (o *InlineResponse20065) GetAdGroupsOk() (*[]string, bool)`

GetAdGroupsOk returns a tuple with the AdGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdGroups

`func (o *InlineResponse20065) SetAdGroups(v []string)`

SetAdGroups sets AdGroups field to given value.

### HasAdGroups

`func (o *InlineResponse20065) HasAdGroups() bool`

HasAdGroups returns a boolean if a field has been set.

### GetAzureAdGroups

`func (o *InlineResponse20065) GetAzureAdGroups() []string`

GetAzureAdGroups returns the AzureAdGroups field if non-nil, zero value otherwise.

### GetAzureAdGroupsOk

`func (o *InlineResponse20065) GetAzureAdGroupsOk() (*[]string, bool)`

GetAzureAdGroupsOk returns a tuple with the AzureAdGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdGroups

`func (o *InlineResponse20065) SetAzureAdGroups(v []string)`

SetAzureAdGroups sets AzureAdGroups field to given value.

### HasAzureAdGroups

`func (o *InlineResponse20065) HasAzureAdGroups() bool`

HasAzureAdGroups returns a boolean if a field has been set.

### GetSamlGroups

`func (o *InlineResponse20065) GetSamlGroups() []string`

GetSamlGroups returns the SamlGroups field if non-nil, zero value otherwise.

### GetSamlGroupsOk

`func (o *InlineResponse20065) GetSamlGroupsOk() (*[]string, bool)`

GetSamlGroupsOk returns a tuple with the SamlGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlGroups

`func (o *InlineResponse20065) SetSamlGroups(v []string)`

SetSamlGroups sets SamlGroups field to given value.

### HasSamlGroups

`func (o *InlineResponse20065) HasSamlGroups() bool`

HasSamlGroups returns a boolean if a field has been set.

### GetAsmGroups

`func (o *InlineResponse20065) GetAsmGroups() []string`

GetAsmGroups returns the AsmGroups field if non-nil, zero value otherwise.

### GetAsmGroupsOk

`func (o *InlineResponse20065) GetAsmGroupsOk() (*[]string, bool)`

GetAsmGroupsOk returns a tuple with the AsmGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsmGroups

`func (o *InlineResponse20065) SetAsmGroups(v []string)`

SetAsmGroups sets AsmGroups field to given value.

### HasAsmGroups

`func (o *InlineResponse20065) HasAsmGroups() bool`

HasAsmGroups returns a boolean if a field has been set.

### GetIsExternal

`func (o *InlineResponse20065) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *InlineResponse20065) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *InlineResponse20065) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *InlineResponse20065) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetDisplayName

`func (o *InlineResponse20065) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InlineResponse20065) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InlineResponse20065) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InlineResponse20065) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHasIdentityCertificate

`func (o *InlineResponse20065) GetHasIdentityCertificate() bool`

GetHasIdentityCertificate returns the HasIdentityCertificate field if non-nil, zero value otherwise.

### GetHasIdentityCertificateOk

`func (o *InlineResponse20065) GetHasIdentityCertificateOk() (*bool, bool)`

GetHasIdentityCertificateOk returns a tuple with the HasIdentityCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasIdentityCertificate

`func (o *InlineResponse20065) SetHasIdentityCertificate(v bool)`

SetHasIdentityCertificate sets HasIdentityCertificate field to given value.

### HasHasIdentityCertificate

`func (o *InlineResponse20065) HasHasIdentityCertificate() bool`

HasHasIdentityCertificate returns a boolean if a field has been set.

### GetUserThumbnail

`func (o *InlineResponse20065) GetUserThumbnail() string`

GetUserThumbnail returns the UserThumbnail field if non-nil, zero value otherwise.

### GetUserThumbnailOk

`func (o *InlineResponse20065) GetUserThumbnailOk() (*string, bool)`

GetUserThumbnailOk returns a tuple with the UserThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserThumbnail

`func (o *InlineResponse20065) SetUserThumbnail(v string)`

SetUserThumbnail sets UserThumbnail field to given value.

### HasUserThumbnail

`func (o *InlineResponse20065) HasUserThumbnail() bool`

HasUserThumbnail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


