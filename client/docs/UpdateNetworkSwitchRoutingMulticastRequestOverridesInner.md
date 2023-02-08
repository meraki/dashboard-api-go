# UpdateNetworkSwitchRoutingMulticastRequestOverridesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwitchProfiles** | Pointer to **[]string** | List of switch profiles ids for template network | [optional] 
**Switches** | Pointer to **[]string** | List of switch serials for non-template network | [optional] 
**Stacks** | Pointer to **[]string** | List of switch stack ids for non-template network | [optional] 
**IgmpSnoopingEnabled** | **bool** | IGMP snooping setting for switches, switch stacks or switch profiles | 
**FloodUnknownMulticastTrafficEnabled** | **bool** | Flood unknown multicast traffic setting for switches, switch stacks or switch profiles | 

## Methods

### NewUpdateNetworkSwitchRoutingMulticastRequestOverridesInner

`func NewUpdateNetworkSwitchRoutingMulticastRequestOverridesInner(igmpSnoopingEnabled bool, floodUnknownMulticastTrafficEnabled bool, ) *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner`

NewUpdateNetworkSwitchRoutingMulticastRequestOverridesInner instantiates a new UpdateNetworkSwitchRoutingMulticastRequestOverridesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchRoutingMulticastRequestOverridesInnerWithDefaults

`func NewUpdateNetworkSwitchRoutingMulticastRequestOverridesInnerWithDefaults() *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner`

NewUpdateNetworkSwitchRoutingMulticastRequestOverridesInnerWithDefaults instantiates a new UpdateNetworkSwitchRoutingMulticastRequestOverridesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchProfiles

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetSwitchProfiles() []string`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetSwitchProfilesOk() (*[]string, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) SetSwitchProfiles(v []string)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### GetSwitches

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetSwitches() []string`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetSwitchesOk() (*[]string, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) SetSwitches(v []string)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetStacks

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetStacks() []string`

GetStacks returns the Stacks field if non-nil, zero value otherwise.

### GetStacksOk

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetStacksOk() (*[]string, bool)`

GetStacksOk returns a tuple with the Stacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacks

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) SetStacks(v []string)`

SetStacks sets Stacks field to given value.

### HasStacks

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) HasStacks() bool`

HasStacks returns a boolean if a field has been set.

### GetIgmpSnoopingEnabled

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetIgmpSnoopingEnabled() bool`

GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field if non-nil, zero value otherwise.

### GetIgmpSnoopingEnabledOk

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetIgmpSnoopingEnabledOk() (*bool, bool)`

GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingEnabled

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) SetIgmpSnoopingEnabled(v bool)`

SetIgmpSnoopingEnabled sets IgmpSnoopingEnabled field to given value.


### GetFloodUnknownMulticastTrafficEnabled

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetFloodUnknownMulticastTrafficEnabled() bool`

GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field if non-nil, zero value otherwise.

### GetFloodUnknownMulticastTrafficEnabledOk

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool)`

GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloodUnknownMulticastTrafficEnabled

`func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) SetFloodUnknownMulticastTrafficEnabled(v bool)`

SetFloodUnknownMulticastTrafficEnabled sets FloodUnknownMulticastTrafficEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


