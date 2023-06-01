# InlineObject139

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RstpEnabled** | Pointer to **bool** | The spanning tree protocol status in network | [optional] 
**StpBridgePriority** | Pointer to [**[]NetworksNetworkIdSwitchStpStpBridgePriority**](NetworksNetworkIdSwitchStpStpBridgePriority.md) | STP bridge priority for switches/stacks or switch profiles. An empty array will clear the STP bridge priority settings. | [optional] 

## Methods

### NewInlineObject139

`func NewInlineObject139() *InlineObject139`

NewInlineObject139 instantiates a new InlineObject139 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject139WithDefaults

`func NewInlineObject139WithDefaults() *InlineObject139`

NewInlineObject139WithDefaults instantiates a new InlineObject139 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRstpEnabled

`func (o *InlineObject139) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *InlineObject139) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *InlineObject139) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *InlineObject139) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpBridgePriority

`func (o *InlineObject139) GetStpBridgePriority() []NetworksNetworkIdSwitchStpStpBridgePriority`

GetStpBridgePriority returns the StpBridgePriority field if non-nil, zero value otherwise.

### GetStpBridgePriorityOk

`func (o *InlineObject139) GetStpBridgePriorityOk() (*[]NetworksNetworkIdSwitchStpStpBridgePriority, bool)`

GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpBridgePriority

`func (o *InlineObject139) SetStpBridgePriority(v []NetworksNetworkIdSwitchStpStpBridgePriority)`

SetStpBridgePriority sets StpBridgePriority field to given value.

### HasStpBridgePriority

`func (o *InlineObject139) HasStpBridgePriority() bool`

HasStpBridgePriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


