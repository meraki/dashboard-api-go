# InlineObject89

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email address of the user | 
**Name** | Pointer to **string** | Name of the user. Only required If the user is not a Dashboard administrator. | [optional] 
**Password** | Pointer to **string** | The password for this user account. Only required If the user is not a Dashboard administrator. | [optional] 
**AccountType** | Pointer to **string** | Authorization type for user. Can be &#39;Guest&#39; or &#39;802.1X&#39; for wireless networks, or &#39;Client VPN&#39; for wired networks. Defaults to &#39;802.1X&#39;. | [optional] [default to "802.1X"]
**EmailPasswordToUser** | Pointer to **bool** | Whether or not Meraki should email the password to user. Default is false. | [optional] 
**IsAdmin** | Pointer to **bool** | Whether or not the user is a Dashboard administrator. | [optional] 
**Authorizations** | [**[]NetworksNetworkIdMerakiAuthUsersAuthorizations1**](NetworksNetworkIdMerakiAuthUsersAuthorizations1.md) | Authorization zones and expiration dates for the user. | 

## Methods

### NewInlineObject89

`func NewInlineObject89(email string, authorizations []NetworksNetworkIdMerakiAuthUsersAuthorizations1, ) *InlineObject89`

NewInlineObject89 instantiates a new InlineObject89 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject89WithDefaults

`func NewInlineObject89WithDefaults() *InlineObject89`

NewInlineObject89WithDefaults instantiates a new InlineObject89 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *InlineObject89) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *InlineObject89) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *InlineObject89) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *InlineObject89) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InlineObject89) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InlineObject89) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InlineObject89) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *InlineObject89) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InlineObject89) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InlineObject89) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *InlineObject89) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAccountType

`func (o *InlineObject89) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *InlineObject89) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *InlineObject89) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *InlineObject89) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetEmailPasswordToUser

`func (o *InlineObject89) GetEmailPasswordToUser() bool`

GetEmailPasswordToUser returns the EmailPasswordToUser field if non-nil, zero value otherwise.

### GetEmailPasswordToUserOk

`func (o *InlineObject89) GetEmailPasswordToUserOk() (*bool, bool)`

GetEmailPasswordToUserOk returns a tuple with the EmailPasswordToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPasswordToUser

`func (o *InlineObject89) SetEmailPasswordToUser(v bool)`

SetEmailPasswordToUser sets EmailPasswordToUser field to given value.

### HasEmailPasswordToUser

`func (o *InlineObject89) HasEmailPasswordToUser() bool`

HasEmailPasswordToUser returns a boolean if a field has been set.

### GetIsAdmin

`func (o *InlineObject89) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *InlineObject89) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *InlineObject89) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *InlineObject89) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetAuthorizations

`func (o *InlineObject89) GetAuthorizations() []NetworksNetworkIdMerakiAuthUsersAuthorizations1`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *InlineObject89) GetAuthorizationsOk() (*[]NetworksNetworkIdMerakiAuthUsersAuthorizations1, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *InlineObject89) SetAuthorizations(v []NetworksNetworkIdMerakiAuthUsersAuthorizations1)`

SetAuthorizations sets Authorizations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


