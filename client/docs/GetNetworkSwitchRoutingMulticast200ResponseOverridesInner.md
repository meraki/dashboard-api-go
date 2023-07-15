# GetNetworkSwitchRoutingMulticast200ResponseOverridesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Switches** | Pointer to **[]string** | (optional) List of switch serials for non-template network | [optional] 
**Stacks** | Pointer to **[]string** | (optional) List of switch stack ids for non-template network | [optional] 
**SwitchProfiles** | Pointer to **[]string** | (optional) List of switch templates ids for template network | [optional] 
**IgmpSnoopingEnabled** | Pointer to **bool** | IGMP snooping enabled for switches, switch stacks or switch templates | [optional] 
**FloodUnknownMulticastTrafficEnabled** | Pointer to **bool** | Flood unknown multicast traffic enabled for switches, switch stacks or switch templates | [optional] 

## Methods

### NewGetNetworkSwitchRoutingMulticast200ResponseOverridesInner

`func NewGetNetworkSwitchRoutingMulticast200ResponseOverridesInner() *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner`

NewGetNetworkSwitchRoutingMulticast200ResponseOverridesInner instantiates a new GetNetworkSwitchRoutingMulticast200ResponseOverridesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkSwitchRoutingMulticast200ResponseOverridesInnerWithDefaults

`func NewGetNetworkSwitchRoutingMulticast200ResponseOverridesInnerWithDefaults() *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner`

NewGetNetworkSwitchRoutingMulticast200ResponseOverridesInnerWithDefaults instantiates a new GetNetworkSwitchRoutingMulticast200ResponseOverridesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwitches

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetSwitches() []string`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetSwitchesOk() (*[]string, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) SetSwitches(v []string)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetStacks

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetStacks() []string`

GetStacks returns the Stacks field if non-nil, zero value otherwise.

### GetStacksOk

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetStacksOk() (*[]string, bool)`

GetStacksOk returns a tuple with the Stacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacks

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) SetStacks(v []string)`

SetStacks sets Stacks field to given value.

### HasStacks

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) HasStacks() bool`

HasStacks returns a boolean if a field has been set.

### GetSwitchProfiles

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetSwitchProfiles() []string`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetSwitchProfilesOk() (*[]string, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) SetSwitchProfiles(v []string)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### GetIgmpSnoopingEnabled

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetIgmpSnoopingEnabled() bool`

GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field if non-nil, zero value otherwise.

### GetIgmpSnoopingEnabledOk

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetIgmpSnoopingEnabledOk() (*bool, bool)`

GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgmpSnoopingEnabled

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) SetIgmpSnoopingEnabled(v bool)`

SetIgmpSnoopingEnabled sets IgmpSnoopingEnabled field to given value.

### HasIgmpSnoopingEnabled

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) HasIgmpSnoopingEnabled() bool`

HasIgmpSnoopingEnabled returns a boolean if a field has been set.

### GetFloodUnknownMulticastTrafficEnabled

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetFloodUnknownMulticastTrafficEnabled() bool`

GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field if non-nil, zero value otherwise.

### GetFloodUnknownMulticastTrafficEnabledOk

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool)`

GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloodUnknownMulticastTrafficEnabled

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) SetFloodUnknownMulticastTrafficEnabled(v bool)`

SetFloodUnknownMulticastTrafficEnabled sets FloodUnknownMulticastTrafficEnabled field to given value.

### HasFloodUnknownMulticastTrafficEnabled

`func (o *GetNetworkSwitchRoutingMulticast200ResponseOverridesInner) HasFloodUnknownMulticastTrafficEnabled() bool`

HasFloodUnknownMulticastTrafficEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


