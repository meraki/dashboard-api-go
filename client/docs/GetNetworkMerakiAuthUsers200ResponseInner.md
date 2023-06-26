# GetNetworkMerakiAuthUsers200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Meraki auth user id | [optional] 
**Email** | Pointer to **string** | Email address of the user | [optional] 
**Name** | Pointer to **string** | Name of the user | [optional] 
**CreatedAt** | Pointer to **time.Time** | Creation time of the user | [optional] 
**AccountType** | Pointer to **string** | Authorization type for user. | [optional] 
**IsAdmin** | Pointer to **bool** | Whether or not the user is a Dashboard administrator | [optional] 
**Authorizations** | Pointer to [**[]GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner**](GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner.md) | User authorization info | [optional] 

## Methods

### NewGetNetworkMerakiAuthUsers200ResponseInner

`func NewGetNetworkMerakiAuthUsers200ResponseInner() *GetNetworkMerakiAuthUsers200ResponseInner`

NewGetNetworkMerakiAuthUsers200ResponseInner instantiates a new GetNetworkMerakiAuthUsers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkMerakiAuthUsers200ResponseInnerWithDefaults

`func NewGetNetworkMerakiAuthUsers200ResponseInnerWithDefaults() *GetNetworkMerakiAuthUsers200ResponseInner`

NewGetNetworkMerakiAuthUsers200ResponseInnerWithDefaults instantiates a new GetNetworkMerakiAuthUsers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEmail

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAccountType

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetIsAdmin

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetIsAdminOk() (*bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdmin

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) SetIsAdmin(v bool)`

SetIsAdmin sets IsAdmin field to given value.

### HasIsAdmin

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### GetAuthorizations

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetAuthorizations() []GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner`

GetAuthorizations returns the Authorizations field if non-nil, zero value otherwise.

### GetAuthorizationsOk

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) GetAuthorizationsOk() (*[]GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner, bool)`

GetAuthorizationsOk returns a tuple with the Authorizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizations

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) SetAuthorizations(v []GetNetworkMerakiAuthUsers200ResponseInnerAuthorizationsInner)`

SetAuthorizations sets Authorizations field to given value.

### HasAuthorizations

`func (o *GetNetworkMerakiAuthUsers200ResponseInner) HasAuthorizations() bool`

HasAuthorizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


