# UpdateNetworkSwitchStpRequestStpBridgePriorityInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwitchProfiles** | Pointer to **[]string** | List of switch profile IDs | [optional] 
**Switches** | Pointer to **[]string** | List of switch serial numbers | [optional] 
**Stacks** | Pointer to **[]string** | List of stack IDs | [optional] 
**StpPriority** | **int32** | STP priority for switch, stacks, or switch profiles | 

## Methods

### NewUpdateNetworkSwitchStpRequestStpBridgePriorityInner

`func NewUpdateNetworkSwitchStpRequestStpBridgePriorityInner(stpPriority int32, ) *UpdateNetworkSwitchStpRequestStpBridgePriorityInner`

NewUpdateNetworkSwitchStpRequestStpBridgePriorityInner instantiates a new UpdateNetworkSwitchStpRequestStpBridgePriorityInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchStpRequestStpBridgePriorityInnerWithDefaults

`func NewUpdateNetworkSwitchStpRequestStpBridgePriorityInnerWithDefaults() *UpdateNetworkSwitchStpRequestStpBridgePriorityInner`

NewUpdateNetworkSwitchStpRequestStpBridgePriorityInnerWithDefaults instantiates a new UpdateNetworkSwitchStpRequestStpBridgePriorityInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchProfiles

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetSwitchProfiles() []string`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetSwitchProfilesOk() (*[]string, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) SetSwitchProfiles(v []string)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### GetSwitches

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetSwitches() []string`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetSwitchesOk() (*[]string, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) SetSwitches(v []string)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetStacks

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetStacks() []string`

GetStacks returns the Stacks field if non-nil, zero value otherwise.

### GetStacksOk

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetStacksOk() (*[]string, bool)`

GetStacksOk returns a tuple with the Stacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacks

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) SetStacks(v []string)`

SetStacks sets Stacks field to given value.

### HasStacks

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) HasStacks() bool`

HasStacks returns a boolean if a field has been set.

### GetStpPriority

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetStpPriority() int32`

GetStpPriority returns the StpPriority field if non-nil, zero value otherwise.

### GetStpPriorityOk

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) GetStpPriorityOk() (*int32, bool)`

GetStpPriorityOk returns a tuple with the StpPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpPriority

`func (o *UpdateNetworkSwitchStpRequestStpBridgePriorityInner) SetStpPriority(v int32)`

SetStpPriority sets StpPriority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


