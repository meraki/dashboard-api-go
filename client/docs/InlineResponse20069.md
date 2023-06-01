# InlineResponse20069

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMtuSize** | Pointer to **int32** | MTU size for the entire network. Default value is 9578. | [optional] 
**Overrides** | Pointer to [**[]InlineResponse20069Overrides**](InlineResponse20069Overrides.md) | Override MTU size for individual switches or switch profiles.       An empty array will clear overrides. | [optional] 

## Methods

### NewInlineResponse20069

`func NewInlineResponse20069() *InlineResponse20069`

NewInlineResponse20069 instantiates a new InlineResponse20069 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20069WithDefaults

`func NewInlineResponse20069WithDefaults() *InlineResponse20069`

NewInlineResponse20069WithDefaults instantiates a new InlineResponse20069 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMtuSize

`func (o *InlineResponse20069) GetDefaultMtuSize() int32`

GetDefaultMtuSize returns the DefaultMtuSize field if non-nil, zero value otherwise.

### GetDefaultMtuSizeOk

`func (o *InlineResponse20069) GetDefaultMtuSizeOk() (*int32, bool)`

GetDefaultMtuSizeOk returns a tuple with the DefaultMtuSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMtuSize

`func (o *InlineResponse20069) SetDefaultMtuSize(v int32)`

SetDefaultMtuSize sets DefaultMtuSize field to given value.

### HasDefaultMtuSize

`func (o *InlineResponse20069) HasDefaultMtuSize() bool`

HasDefaultMtuSize returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineResponse20069) GetOverrides() []InlineResponse20069Overrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineResponse20069) GetOverridesOk() (*[]InlineResponse20069Overrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineResponse20069) SetOverrides(v []InlineResponse20069Overrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineResponse20069) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


