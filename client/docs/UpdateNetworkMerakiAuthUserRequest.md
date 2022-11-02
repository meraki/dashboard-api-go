# UpdateNetworkMerakiAuthUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user. Only allowed If the user is not Dashboard administrator. | [optional] 
**Password** | Pointer to **string** | The password for this user account. Only allowed If the user is not Dashboard administrator. | [optional] 
**EmailPasswordToUser** | Pointer to **bool** | Whether or not Meraki should email the password to user. Default is false. | [optional] 
**Authorizations** | Pointer to [**[]UpdateNetworkMerakiAuthUserRequestAuthorizationsInner**](UpdateNetworkMerakiAuthUserRequestAuthorizationsInner.md) | Authorization zones and expiration dates for the user. | [optional] 

## Methods

### NewUpdateNetworkMerakiAuthUserRequest

`func NewUpdateNetworkMerakiAuthUserRequest() *UpdateNetworkMerakiAuthUserRequest`

NewUpdateNetworkMerakiAuthUserRequest instantiates a new UpdateNetworkMerakiAuthUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkMerakiAuthUserRequestWithDefaults

`func NewUpdateNetworkMerakiAuthUserRequestWithDefaults() *UpdateNetworkMerakiAuthUserRequest`

NewUpdateNetworkMerakiAuthUserRequestWithDefaults instantiates a new UpdateNetworkMerakiAuthUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkMerakiAuthUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkMerakiAuthUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkMerakiAuthUserRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkMerakiAuthUserRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateNetworkMerakiAuthUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateNetworkMerakiAuthUserRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateNetworkMerakiAuthUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateNetworkMerakiAuthUserRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmailPasswordToUser

`func (o *UpdateNetworkMerakiAuthUserRequest) GetEmailPasswordToUser() bool`

GetEmailPasswordToUser returns the EmailPasswordToUser field if non-nil, zero value otherwise.

### GetEmailPasswordToUserOk

`func (o *UpdateNetworkMerakiAuthUserRequest) GetEmailPasswordToUserOk() (*bool, bool)`

GetEmailPasswordToUserOk returns a tuple with the EmailPasswordToUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPasswordToUser

`func (o *UpdateNetworkMerakiAuthUserRequest) SetEmailPasswordToUser(v bool)`

SetEmailPasswordToUser sets EmailPasswordToUser field to given value.

### HasEmailPasswordToUser

`func (o *UpdateNetworkMerakiAuthUserRequest) HasEmailPasswordToUser() bool`

HasEmailPasswordToUser returns a boolean if a field has been set.

### GetAuthorizations

`func (o *UpdateNetworkMerakiAuthUserRequest) GetAuthorizations() []UpdateNetworkMerakiAuthUserRequestAuthorizationsInner`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *UpdateNetworkMerakiAuthUserRequest) GetAuthorizationsOk() (*[]UpdateNetworkMerakiAuthUserRequestAuthorizationsInner, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *UpdateNetworkMerakiAuthUserRequest) SetAuthorizations(v []UpdateNetworkMerakiAuthUserRequestAuthorizationsInner)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *UpdateNetworkMerakiAuthUserRequest) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


