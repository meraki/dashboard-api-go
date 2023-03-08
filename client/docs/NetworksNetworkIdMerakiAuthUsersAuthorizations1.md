# NetworksNetworkIdMerakiAuthUsersAuthorizations1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidNumber** | Pointer to **int32** | Required for wireless networks. The SSID for which the user is being authorized, which must be configured for the user&#39;s given accountType. | [optional] 
**ExpiresAt** | Pointer to **string** | Date for authorization to expire. Set to &#39;Never&#39; for the authorization to not expire, which is the default. | [optional] [default to "Never"]

## Methods

### NewNetworksNetworkIdMerakiAuthUsersAuthorizations1

`func NewNetworksNetworkIdMerakiAuthUsersAuthorizations1() *NetworksNetworkIdMerakiAuthUsersAuthorizations1`

NewNetworksNetworkIdMerakiAuthUsersAuthorizations1 instantiates a new NetworksNetworkIdMerakiAuthUsersAuthorizations1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdMerakiAuthUsersAuthorizations1WithDefaults

`func NewNetworksNetworkIdMerakiAuthUsersAuthorizations1WithDefaults() *NetworksNetworkIdMerakiAuthUsersAuthorizations1`

NewNetworksNetworkIdMerakiAuthUsersAuthorizations1WithDefaults instantiates a new NetworksNetworkIdMerakiAuthUsersAuthorizations1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidNumber

`func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations1) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations1) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations1) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations1) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetExpiresAt

`func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations1) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations1) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations1) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *NetworksNetworkIdMerakiAuthUsersAuthorizations1) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


