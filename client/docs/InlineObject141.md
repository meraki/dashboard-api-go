# InlineObject141

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BroadcastThreshold** | Pointer to **int32** | Percentage (1 to 99) of total available port bandwidth for broadcast traffic type. Default value 100 percent rate is to clear the configuration. | [optional] 
**MulticastThreshold** | Pointer to **int32** | Percentage (1 to 99) of total available port bandwidth for multicast traffic type. Default value 100 percent rate is to clear the configuration. | [optional] 
**UnknownUnicastThreshold** | Pointer to **int32** | Percentage (1 to 99) of total available port bandwidth for unknown unicast (dlf-destination lookup failure) traffic type. Default value 100 percent rate is to clear the configuration. | [optional] 

## Methods

### NewInlineObject141

`func NewInlineObject141() *InlineObject141`

NewInlineObject141 instantiates a new InlineObject141 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject141WithDefaults

`func NewInlineObject141WithDefaults() *InlineObject141`

NewInlineObject141WithDefaults instantiates a new InlineObject141 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroadcastThreshold

`func (o *InlineObject141) GetBroadcastThreshold() int32`

GetBroadcastThreshold returns the BroadcastThreshold field if non-nil, zero value otherwise.

### GetBroadcastThresholdOk

`func (o *InlineObject141) GetBroadcastThresholdOk() (*int32, bool)`

GetBroadcastThresholdOk returns a tuple with the BroadcastThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastThreshold

`func (o *InlineObject141) SetBroadcastThreshold(v int32)`

SetBroadcastThreshold sets BroadcastThreshold field to given value.

### HasBroadcastThreshold

`func (o *InlineObject141) HasBroadcastThreshold() bool`

HasBroadcastThreshold returns a boolean if a field has been set.

### GetMulticastThreshold

`func (o *InlineObject141) GetMulticastThreshold() int32`

GetMulticastThreshold returns the MulticastThreshold field if non-nil, zero value otherwise.

### GetMulticastThresholdOk

`func (o *InlineObject141) GetMulticastThresholdOk() (*int32, bool)`

GetMulticastThresholdOk returns a tuple with the MulticastThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastThreshold

`func (o *InlineObject141) SetMulticastThreshold(v int32)`

SetMulticastThreshold sets MulticastThreshold field to given value.

### HasMulticastThreshold

`func (o *InlineObject141) HasMulticastThreshold() bool`

HasMulticastThreshold returns a boolean if a field has been set.

### GetUnknownUnicastThreshold

`func (o *InlineObject141) GetUnknownUnicastThreshold() int32`

GetUnknownUnicastThreshold returns the UnknownUnicastThreshold field if non-nil, zero value otherwise.

### GetUnknownUnicastThresholdOk

`func (o *InlineObject141) GetUnknownUnicastThresholdOk() (*int32, bool)`

GetUnknownUnicastThresholdOk returns a tuple with the UnknownUnicastThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUnicastThreshold

`func (o *InlineObject141) SetUnknownUnicastThreshold(v int32)`

SetUnknownUnicastThreshold sets UnknownUnicastThreshold field to given value.

### HasUnknownUnicastThreshold

`func (o *InlineObject141) HasUnknownUnicastThreshold() bool`

HasUnknownUnicastThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


