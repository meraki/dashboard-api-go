# ShutdownNetworkSmDevicesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WifiMacs** | Pointer to **[]string** | The wifiMacs of the endpoints to be shutdown. | [optional] 
**Ids** | Pointer to **[]string** | The ids of the endpoints to be shutdown. | [optional] 
**Serials** | Pointer to **[]string** | The serials of the endpoints to be shutdown. | [optional] 
**Scope** | Pointer to **[]string** | The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the endpoints to be shutdown. | [optional] 

## Methods

### NewShutdownNetworkSmDevicesRequest

`func NewShutdownNetworkSmDevicesRequest() *ShutdownNetworkSmDevicesRequest`

NewShutdownNetworkSmDevicesRequest instantiates a new ShutdownNetworkSmDevicesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShutdownNetworkSmDevicesRequestWithDefaults

`func NewShutdownNetworkSmDevicesRequestWithDefaults() *ShutdownNetworkSmDevicesRequest`

NewShutdownNetworkSmDevicesRequestWithDefaults instantiates a new ShutdownNetworkSmDevicesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiMacs

`func (o *ShutdownNetworkSmDevicesRequest) GetWifiMacs() []string`

GetWifiMacs returns the WifiMacs field if non-nil, zero value otherwise.

### GetWifiMacsOk

`func (o *ShutdownNetworkSmDevicesRequest) GetWifiMacsOk() (*[]string, bool)`

GetWifiMacsOk returns a tuple with the WifiMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacs

`func (o *ShutdownNetworkSmDevicesRequest) SetWifiMacs(v []string)`

SetWifiMacs sets WifiMacs field to given value.

### HasWifiMacs

`func (o *ShutdownNetworkSmDevicesRequest) HasWifiMacs() bool`

HasWifiMacs returns a boolean if a field has been set.

### GetIds

`func (o *ShutdownNetworkSmDevicesRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ShutdownNetworkSmDevicesRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ShutdownNetworkSmDevicesRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ShutdownNetworkSmDevicesRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSerials

`func (o *ShutdownNetworkSmDevicesRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *ShutdownNetworkSmDevicesRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *ShutdownNetworkSmDevicesRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *ShutdownNetworkSmDevicesRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetScope

`func (o *ShutdownNetworkSmDevicesRequest) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ShutdownNetworkSmDevicesRequest) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ShutdownNetworkSmDevicesRequest) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ShutdownNetworkSmDevicesRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


