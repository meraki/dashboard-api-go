# LockNetworkSmDevicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WifiMacs** | Pointer to **[]string** | The wifiMacs of the devices to be locked. | [optional] 
**Ids** | Pointer to **[]string** | The ids of the devices to be locked. | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices to be locked. | [optional] 
**Scope** | Pointer to **[]string** | The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be wiped. | [optional] 
**Pin** | Pointer to **int32** | The pin number for locking macOS devices (a six digit number). Required only for macOS devices. | [optional] 

## Methods

### NewLockNetworkSmDevicesRequest

`func NewLockNetworkSmDevicesRequest() *LockNetworkSmDevicesRequest`

NewLockNetworkSmDevicesRequest instantiates a new LockNetworkSmDevicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockNetworkSmDevicesRequestWithDefaults

`func NewLockNetworkSmDevicesRequestWithDefaults() *LockNetworkSmDevicesRequest`

NewLockNetworkSmDevicesRequestWithDefaults instantiates a new LockNetworkSmDevicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiMacs

`func (o *LockNetworkSmDevicesRequest) GetWifiMacs() []string`

GetWifiMacs returns the WifiMacs field if non-nil, zero value otherwise.

### GetWifiMacsOk

`func (o *LockNetworkSmDevicesRequest) GetWifiMacsOk() (*[]string, bool)`

GetWifiMacsOk returns a tuple with the WifiMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacs

`func (o *LockNetworkSmDevicesRequest) SetWifiMacs(v []string)`

SetWifiMacs sets WifiMacs field to given value.

### HasWifiMacs

`func (o *LockNetworkSmDevicesRequest) HasWifiMacs() bool`

HasWifiMacs returns a boolean if a field has been set.

### GetIds

`func (o *LockNetworkSmDevicesRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *LockNetworkSmDevicesRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *LockNetworkSmDevicesRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *LockNetworkSmDevicesRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSerials

`func (o *LockNetworkSmDevicesRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *LockNetworkSmDevicesRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *LockNetworkSmDevicesRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *LockNetworkSmDevicesRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetScope

`func (o *LockNetworkSmDevicesRequest) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *LockNetworkSmDevicesRequest) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *LockNetworkSmDevicesRequest) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *LockNetworkSmDevicesRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetPin

`func (o *LockNetworkSmDevicesRequest) GetPin() int32`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *LockNetworkSmDevicesRequest) GetPinOk() (*int32, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *LockNetworkSmDevicesRequest) SetPin(v int32)`

SetPin sets Pin field to given value.

### HasPin

`func (o *LockNetworkSmDevicesRequest) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


