# CheckinNetworkSmDevicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WifiMacs** | Pointer to **[]string** | The wifiMacs of the devices to be checked-in. | [optional] 
**Ids** | Pointer to **[]string** | The ids of the devices to be checked-in. | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices to be checked-in. | [optional] 
**Scope** | Pointer to **[]string** | The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be checked-in. | [optional] 

## Methods

### NewCheckinNetworkSmDevicesRequest

`func NewCheckinNetworkSmDevicesRequest() *CheckinNetworkSmDevicesRequest`

NewCheckinNetworkSmDevicesRequest instantiates a new CheckinNetworkSmDevicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckinNetworkSmDevicesRequestWithDefaults

`func NewCheckinNetworkSmDevicesRequestWithDefaults() *CheckinNetworkSmDevicesRequest`

NewCheckinNetworkSmDevicesRequestWithDefaults instantiates a new CheckinNetworkSmDevicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiMacs

`func (o *CheckinNetworkSmDevicesRequest) GetWifiMacs() []string`

GetWifiMacs returns the WifiMacs field if non-nil, zero value otherwise.

### GetWifiMacsOk

`func (o *CheckinNetworkSmDevicesRequest) GetWifiMacsOk() (*[]string, bool)`

GetWifiMacsOk returns a tuple with the WifiMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacs

`func (o *CheckinNetworkSmDevicesRequest) SetWifiMacs(v []string)`

SetWifiMacs sets WifiMacs field to given value.

### HasWifiMacs

`func (o *CheckinNetworkSmDevicesRequest) HasWifiMacs() bool`

HasWifiMacs returns a boolean if a field has been set.

### GetIds

`func (o *CheckinNetworkSmDevicesRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *CheckinNetworkSmDevicesRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *CheckinNetworkSmDevicesRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *CheckinNetworkSmDevicesRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSerials

`func (o *CheckinNetworkSmDevicesRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *CheckinNetworkSmDevicesRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *CheckinNetworkSmDevicesRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *CheckinNetworkSmDevicesRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetScope

`func (o *CheckinNetworkSmDevicesRequest) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CheckinNetworkSmDevicesRequest) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CheckinNetworkSmDevicesRequest) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CheckinNetworkSmDevicesRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


