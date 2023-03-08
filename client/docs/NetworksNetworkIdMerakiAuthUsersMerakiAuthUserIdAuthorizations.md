# NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidNumber** | **int32** | SSID for which the user is being authorized | 
**ExpiresAt** | Pointer to **string** | Date for authorization to expire. Default is for authorization to not expire. | [optional] [default to "Never"]

## Methods

### NewNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations

`func NewNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations(ssidNumber int32, ) *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations`

NewNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations instantiates a new NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizationsWithDefaults

`func NewNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizationsWithDefaults() *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations`

NewNetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizationsWithDefaults instantiates a new NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.


### GetExpiresAt

`func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *NetworksNetworkIdMerakiAuthUsersMerakiAuthUserIdAuthorizations) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


