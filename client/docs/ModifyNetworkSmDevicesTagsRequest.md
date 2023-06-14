# ModifyNetworkSmDevicesTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WifiMacs** | Pointer to **[]string** | The wifiMacs of the devices to be modified. | [optional] 
**Ids** | Pointer to **[]string** | The ids of the devices to be modified. | [optional] 
**Serials** | Pointer to **[]string** | The serials of the devices to be modified. | [optional] 
**Scope** | Pointer to **[]string** | The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be modified. | [optional] 
**Tags** | **[]string** | The tags to be added, deleted, or updated. | 
**UpdateAction** | **string** | One of add, delete, or update. Only devices that have been modified will be returned. | 

## Methods

### NewModifyNetworkSmDevicesTagsRequest

`func NewModifyNetworkSmDevicesTagsRequest(tags []string, updateAction string, ) *ModifyNetworkSmDevicesTagsRequest`

NewModifyNetworkSmDevicesTagsRequest instantiates a new ModifyNetworkSmDevicesTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyNetworkSmDevicesTagsRequestWithDefaults

`func NewModifyNetworkSmDevicesTagsRequestWithDefaults() *ModifyNetworkSmDevicesTagsRequest`

NewModifyNetworkSmDevicesTagsRequestWithDefaults instantiates a new ModifyNetworkSmDevicesTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiMacs

`func (o *ModifyNetworkSmDevicesTagsRequest) GetWifiMacs() []string`

GetWifiMacs returns the WifiMacs field if non-nil, zero value otherwise.

### GetWifiMacsOk

`func (o *ModifyNetworkSmDevicesTagsRequest) GetWifiMacsOk() (*[]string, bool)`

GetWifiMacsOk returns a tuple with the WifiMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacs

`func (o *ModifyNetworkSmDevicesTagsRequest) SetWifiMacs(v []string)`

SetWifiMacs sets WifiMacs field to given value.

### HasWifiMacs

`func (o *ModifyNetworkSmDevicesTagsRequest) HasWifiMacs() bool`

HasWifiMacs returns a boolean if a field has been set.

### GetIds

`func (o *ModifyNetworkSmDevicesTagsRequest) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ModifyNetworkSmDevicesTagsRequest) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ModifyNetworkSmDevicesTagsRequest) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *ModifyNetworkSmDevicesTagsRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSerials

`func (o *ModifyNetworkSmDevicesTagsRequest) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *ModifyNetworkSmDevicesTagsRequest) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *ModifyNetworkSmDevicesTagsRequest) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *ModifyNetworkSmDevicesTagsRequest) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetScope

`func (o *ModifyNetworkSmDevicesTagsRequest) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ModifyNetworkSmDevicesTagsRequest) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ModifyNetworkSmDevicesTagsRequest) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ModifyNetworkSmDevicesTagsRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTags

`func (o *ModifyNetworkSmDevicesTagsRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModifyNetworkSmDevicesTagsRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModifyNetworkSmDevicesTagsRequest) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUpdateAction

`func (o *ModifyNetworkSmDevicesTagsRequest) GetUpdateAction() string`

GetUpdateAction returns the UpdateAction field if non-nil, zero value otherwise.

### GetUpdateActionOk

`func (o *ModifyNetworkSmDevicesTagsRequest) GetUpdateActionOk() (*string, bool)`

GetUpdateActionOk returns a tuple with the UpdateAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAction

`func (o *ModifyNetworkSmDevicesTagsRequest) SetUpdateAction(v string)`

SetUpdateAction sets UpdateAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


