# UpdateNetworkMerakiAuthUserRequestAuthorizationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidNumber** | **int32** | SSID for which the user is being authorized | 
**ExpiresAt** | Pointer to **string** | Date for authorization to expire. Default is for authorization to not expire. | [optional] [default to "Never"]

## Methods

### NewUpdateNetworkMerakiAuthUserRequestAuthorizationsInner

`func NewUpdateNetworkMerakiAuthUserRequestAuthorizationsInner(ssidNumber int32, ) *UpdateNetworkMerakiAuthUserRequestAuthorizationsInner`

NewUpdateNetworkMerakiAuthUserRequestAuthorizationsInner instantiates a new UpdateNetworkMerakiAuthUserRequestAuthorizationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkMerakiAuthUserRequestAuthorizationsInnerWithDefaults

`func NewUpdateNetworkMerakiAuthUserRequestAuthorizationsInnerWithDefaults() *UpdateNetworkMerakiAuthUserRequestAuthorizationsInner`

NewUpdateNetworkMerakiAuthUserRequestAuthorizationsInnerWithDefaults instantiates a new UpdateNetworkMerakiAuthUserRequestAuthorizationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *UpdateNetworkMerakiAuthUserRequestAuthorizationsInner) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *UpdateNetworkMerakiAuthUserRequestAuthorizationsInner) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *UpdateNetworkMerakiAuthUserRequestAuthorizationsInner) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.


### GetExpiresAt

`func (o *UpdateNetworkMerakiAuthUserRequestAuthorizationsInner) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *UpdateNetworkMerakiAuthUserRequestAuthorizationsInner) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *UpdateNetworkMerakiAuthUserRequestAuthorizationsInner) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *UpdateNetworkMerakiAuthUserRequestAuthorizationsInner) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


