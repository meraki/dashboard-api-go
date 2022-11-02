# UpdateNetworkSwitchStpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RstpEnabled** | Pointer to **bool** | The spanning tree protocol status in network | [optional] 
**StpBridgePriority** | Pointer to [**[]UpdateNetworkSwitchStpRequestStpBridgePriorityInner**](UpdateNetworkSwitchStpRequestStpBridgePriorityInner.md) | STP bridge priority for switches/stacks or switch profiles. An empty array will clear the STP bridge priority settings. | [optional] 

## Methods

### NewUpdateNetworkSwitchStpRequest

`func NewUpdateNetworkSwitchStpRequest() *UpdateNetworkSwitchStpRequest`

NewUpdateNetworkSwitchStpRequest instantiates a new UpdateNetworkSwitchStpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchStpRequestWithDefaults

`func NewUpdateNetworkSwitchStpRequestWithDefaults() *UpdateNetworkSwitchStpRequest`

NewUpdateNetworkSwitchStpRequestWithDefaults instantiates a new UpdateNetworkSwitchStpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRstpEnabled

`func (o *UpdateNetworkSwitchStpRequest) GetRstpEnabled() bool`

GetRstpEnabled returns the RstpEnabled field if non-nil, zero value otherwise.

### GetRstpEnabledOk

`func (o *UpdateNetworkSwitchStpRequest) GetRstpEnabledOk() (*bool, bool)`

GetRstpEnabledOk returns a tuple with the RstpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRstpEnabled

`func (o *UpdateNetworkSwitchStpRequest) SetRstpEnabled(v bool)`

SetRstpEnabled sets RstpEnabled field to given value.

### HasRstpEnabled

`func (o *UpdateNetworkSwitchStpRequest) HasRstpEnabled() bool`

HasRstpEnabled returns a boolean if a field has been set.

### GetStpBridgePriority

`func (o *UpdateNetworkSwitchStpRequest) GetStpBridgePriority() []UpdateNetworkSwitchStpRequestStpBridgePriorityInner`

GetStpBridgePriority returns the StpBridgePriority field if non-nil, zero value otherwise.

### GetStpBridgePriorityOk

`func (o *UpdateNetworkSwitchStpRequest) GetStpBridgePriorityOk() (*[]UpdateNetworkSwitchStpRequestStpBridgePriorityInner, bool)`

GetStpBridgePriorityOk returns a tuple with the StpBridgePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpBridgePriority

`func (o *UpdateNetworkSwitchStpRequest) SetStpBridgePriority(v []UpdateNetworkSwitchStpRequestStpBridgePriorityInner)`

SetStpBridgePriority sets StpBridgePriority field to given value.

### HasStpBridgePriority

`func (o *UpdateNetworkSwitchStpRequest) HasStpBridgePriority() bool`

HasStpBridgePriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


