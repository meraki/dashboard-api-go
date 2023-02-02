# InlineObject138

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RstpEnabled** | Pointer to **bool** | The spanning tree protocol status in network | [optional] 
**StpBridgePriority** | Pointer to [**[]NetworksNetworkIdSwitchStpStpBridgePriority**](NetworksNetworkIdSwitchStpStpBridgePriority.md) | STP bridge priority for switches/stacks or switch profiles. An empty array will clear the STP bridge priority settings. | [optional] 

## Methods

### NewInlineObject138

`func NewInlineObject138() *InlineObject138`

NewInlineObject138 instantiates a new InlineObject138 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject138WithDefaults

`func NewInlineObject138WithDefaults() *InlineObject138`

NewInlineObject138WithDefaults instantiates a new InlineObject138 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRstpEnabled

`func (o *InlineObject138) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *InlineObject138) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *InlineObject138) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *InlineObject138) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpBridgePriority

`func (o *InlineObject138) GetStpBridgePriority() []NetworksNetworkIdSwitchStpStpBridgePriority`

GetStpBridgePriority returns the StpBridgePriority field if non-nil, zero value otherwise.

### GetStpBridgePriorityOk

`func (o *InlineObject138) GetStpBridgePriorityOk() (*[]NetworksNetworkIdSwitchStpStpBridgePriority, bool)`

GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpBridgePriority

`func (o *InlineObject138) SetStpBridgePriority(v []NetworksNetworkIdSwitchStpStpBridgePriority)`

SetStpBridgePriority sets StpBridgePriority field to given value.

### HasStpBridgePriority

`func (o *InlineObject138) HasStpBridgePriority() bool`

HasStpBridgePriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


