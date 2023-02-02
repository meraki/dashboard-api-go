# InlineObject101

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WifiMacs** | Pointer to **[]string** | The wifiMacs of the devices to be locked. | [optional] 
**Ids** | Pointer to **[]string** | The ids of the devices to be locked. | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices to be locked. | [optional] 
**Scope** | Pointer to **[]string** | The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be wiped. | [optional] 
**Pin** | Pointer to **int32** | The pin number for locking macOS devices (a six digit number). Required only for macOS devices. | [optional] 

## Methods

### NewInlineObject101

`func NewInlineObject101() *InlineObject101`

NewInlineObject101 instantiates a new InlineObject101 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject101WithDefaults

`func NewInlineObject101WithDefaults() *InlineObject101`

NewInlineObject101WithDefaults instantiates a new InlineObject101 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiMacs

`func (o *InlineObject101) GetWifiMacs() []string`

GetWifiMacs returns the WifiMacs field if non-nil, zero value otherwise.

### GetWifiMacsOk

`func (o *InlineObject101) GetWifiMacsOk() (*[]string, bool)`

GetWifiMacsOk returns a tuple with the WifiMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacs

`func (o *InlineObject101) SetWifiMacs(v []string)`

SetWifiMacs sets WifiMacs field to given value.

### HasWifiMacs

`func (o *InlineObject101) HasWifiMacs() bool`

HasWifiMacs returns a boolean if a field has been set.

### GetIds

`func (o *InlineObject101) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *InlineObject101) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *InlineObject101) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *InlineObject101) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSerials

`func (o *InlineObject101) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject101) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject101) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineObject101) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetScope

`func (o *InlineObject101) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineObject101) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineObject101) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineObject101) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetPin

`func (o *InlineObject101) GetPin() int32`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *InlineObject101) GetPinOk() (*int32, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *InlineObject101) SetPin(v int32)`

SetPin sets Pin field to given value.

### HasPin

`func (o *InlineObject101) HasPin() bool`

HasPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


