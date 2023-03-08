# NetworksNetworkIdSwitchRoutingMulticastOverrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SwitchProfiles** | Pointer to **[]string** | List of switch profiles ids for template network | [optional] 
**Switches** | Pointer to **[]string** | List of switch serials for non-template network | [optional] 
**Stacks** | Pointer to **[]string** | List of switch stack ids for non-template network | [optional] 
**IgmpSnoopingEnabled** | **bool** | IGMP snooping setting for switches, switch stacks or switch profiles | 
**FloodUnknownMulticastTrafficEnabled** | **bool** | Flood unknown multicast traffic setting for switches, switch stacks or switch profiles | 

## Methods

### NewNetworksNetworkIdSwitchRoutingMulticastOverrides

`func NewNetworksNetworkIdSwitchRoutingMulticastOverrides(igmpSnoopingEnabled bool, floodUnknownMulticastTrafficEnabled bool, ) *NetworksNetworkIdSwitchRoutingMulticastOverrides`

NewNetworksNetworkIdSwitchRoutingMulticastOverrides instantiates a new NetworksNetworkIdSwitchRoutingMulticastOverrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdSwitchRoutingMulticastOverridesWithDefaults

`func NewNetworksNetworkIdSwitchRoutingMulticastOverridesWithDefaults() *NetworksNetworkIdSwitchRoutingMulticastOverrides`

NewNetworksNetworkIdSwitchRoutingMulticastOverridesWithDefaults instantiates a new NetworksNetworkIdSwitchRoutingMulticastOverrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitchProfiles

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetSwitchProfiles() []string`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetSwitchProfilesOk() (*[]string, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) SetSwitchProfiles(v []string)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### GetSwitches

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetSwitches() []string`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetSwitchesOk() (*[]string, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) SetSwitches(v []string)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetStacks

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetStacks() []string`

GetStacks returns the Stacks field if non-nil, zero value otherwise.

### GetStacksOk

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetStacksOk() (*[]string, bool)`

GetStacksOk returns a tuple with the Stacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacks

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) SetStacks(v []string)`

SetStacks sets Stacks field to given value.

### HasStacks

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) HasStacks() bool`

HasStacks returns a boolean if a field has been set.

### GetIgmpSnoopingEnabled

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetIgmpSnoopingEnabled() bool`

GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field if non-nil, zero value otherwise.

### GetIgmpSnoopingEnabledOk

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetIgmpSnoopingEnabledOk() (*bool, bool)`

GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingEnabled

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) SetIgmpSnoopingEnabled(v bool)`

SetIgmpSnoopingEnabled sets IgmpSnoopingEnabled field to given value.


### GetFloodUnknownMulticastTrafficEnabled

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetFloodUnknownMulticastTrafficEnabled() bool`

GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field if non-nil, zero value otherwise.

### GetFloodUnknownMulticastTrafficEnabledOk

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool)`

GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloodUnknownMulticastTrafficEnabled

`func (o *NetworksNetworkIdSwitchRoutingMulticastOverrides) SetFloodUnknownMulticastTrafficEnabled(v bool)`

SetFloodUnknownMulticastTrafficEnabled sets FloodUnknownMulticastTrafficEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


