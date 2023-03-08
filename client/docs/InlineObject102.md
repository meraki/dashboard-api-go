# InlineObject102

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

### NewInlineObject102

`func NewInlineObject102(tags []string, updateAction string, ) *InlineObject102`

NewInlineObject102 instantiates a new InlineObject102 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject102WithDefaults

`func NewInlineObject102WithDefaults() *InlineObject102`

NewInlineObject102WithDefaults instantiates a new InlineObject102 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWifiMacs

`func (o *InlineObject102) GetWifiMacs() []string`

GetWifiMacs returns the WifiMacs field if non-nil, zero value otherwise.

### GetWifiMacsOk

`func (o *InlineObject102) GetWifiMacsOk() (*[]string, bool)`

GetWifiMacsOk returns a tuple with the WifiMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiMacs

`func (o *InlineObject102) SetWifiMacs(v []string)`

SetWifiMacs sets WifiMacs field to given value.

### HasWifiMacs

`func (o *InlineObject102) HasWifiMacs() bool`

HasWifiMacs returns a boolean if a field has been set.

### GetIds

`func (o *InlineObject102) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *InlineObject102) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *InlineObject102) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *InlineObject102) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSerials

`func (o *InlineObject102) GetSerials() []string`

GetSerials returns the Serials field if non-nil, zero value otherwise.

### GetSerialsOk

`func (o *InlineObject102) GetSerialsOk() (*[]string, bool)`

GetSerialsOk returns a tuple with the Serials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerials

`func (o *InlineObject102) SetSerials(v []string)`

SetSerials sets Serials field to given value.

### HasSerials

`func (o *InlineObject102) HasSerials() bool`

HasSerials returns a boolean if a field has been set.

### GetScope

`func (o *InlineObject102) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *InlineObject102) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *InlineObject102) SetScope(v []string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *InlineObject102) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTags

`func (o *InlineObject102) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InlineObject102) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InlineObject102) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUpdateAction

`func (o *InlineObject102) GetUpdateAction() string`

GetUpdateAction returns the UpdateAction field if non-nil, zero value otherwise.

### GetUpdateActionOk

`func (o *InlineObject102) GetUpdateActionOk() (*string, bool)`

GetUpdateActionOk returns a tuple with the UpdateAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateAction

`func (o *InlineObject102) SetUpdateAction(v string)`

SetUpdateAction sets UpdateAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


