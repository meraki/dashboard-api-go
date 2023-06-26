# MoveNetworkSmDevicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WifiMacs** | Pointer to **[]string** | The wifiMacs of the devices to be moved. | [optional] 
**Ids** | Pointer to **[]string** | The ids of the devices to be moved. | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices to be moved. | [optional] 
**Scope** | Pointer to **[]string** | The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be moved. | [optional] 
**NewNetwork** | **string** | The new network to which the devices will be moved. | 

## Methods

### NewMoveNetworkSmDevicesRequest

`func NewMoveNetworkSmDevicesRequest(newNetwork string, ) *MoveNetworkSmDevicesRequest`

NewMoveNetworkSmDevicesRequest instantiates a new MoveNetworkSmDevicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveNetworkSmDevicesRequestWithDefaults

`func NewMoveNetworkSmDevicesRequestWithDefaults() *MoveNetworkSmDevicesRequest`

NewMoveNetworkSmDevicesRequestWithDefaults instantiates a new MoveNetworkSmDevicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiMacs

`func (o *MoveNetworkSmDevicesRequest) GetWifiMacs() []string`

GetWifiMacs returns the WifiMacs field if non-nil, zero value otherwise.

### GetWifiMacsOk

`func (o *MoveNetworkSmDevicesRequest) GetWifiMacsOk() (*[]string, bool)`

GetWifiMacsOk returns a tuple with the WifiMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacs

`func (o *MoveNetworkSmDevicesRequest) SetWifiMacs(v []string)`

SetWifiMacs sets WifiMacs field to given value.

### HasWifiMacs

`func (o *MoveNetworkSmDevicesRequest) HasWifiMacs() bool`

HasWifiMacs returns a boolean if a field has been set.

### GetIds

`func (o *MoveNetworkSmDevicesRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *MoveNetworkSmDevicesRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *MoveNetworkSmDevicesRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *MoveNetworkSmDevicesRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSerials

`func (o *MoveNetworkSmDevicesRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *MoveNetworkSmDevicesRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *MoveNetworkSmDevicesRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *MoveNetworkSmDevicesRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetScope

`func (o *MoveNetworkSmDevicesRequest) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *MoveNetworkSmDevicesRequest) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *MoveNetworkSmDevicesRequest) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *MoveNetworkSmDevicesRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetNewNetwork

`func (o *MoveNetworkSmDevicesRequest) GetNewNetwork() string`

GetNewNetwork returns the NewNetwork field if non-nil, zero value otherwise.

### GetNewNetworkOk

`func (o *MoveNetworkSmDevicesRequest) GetNewNetworkOk() (*string, bool)`

GetNewNetworkOk returns a tuple with the NewNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewNetwork

`func (o *MoveNetworkSmDevicesRequest) SetNewNetwork(v string)`

SetNewNetwork sets NewNetwork field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


