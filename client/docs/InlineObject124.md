# InlineObject124

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSettings** | Pointer to [**NetworksNetworkIdSwitchRoutingMulticastDefaultSettings**](NetworksNetworkIdSwitchRoutingMulticastDefaultSettings.md) |  | [optional] 
**Overrides** | Pointer to [**[]NetworksNetworkIdSwitchRoutingMulticastOverrides**](NetworksNetworkIdSwitchRoutingMulticastOverrides.md) | Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings. | [optional] 

## Methods

### NewInlineObject124

`func NewInlineObject124() *InlineObject124`

NewInlineObject124 instantiates a new InlineObject124 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject124WithDefaults

`func NewInlineObject124WithDefaults() *InlineObject124`

NewInlineObject124WithDefaults instantiates a new InlineObject124 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSettings

`func (o *InlineObject124) GetDefaultSettings() NetworksNetworkIdSwitchRoutingMulticastDefaultSettings`

GetDefaultSettings returns the DefaultSettings field if non-nil, zero value otherwise.

### GetDefaultSettingsOk

`func (o *InlineObject124) GetDefaultSettingsOk() (*NetworksNetworkIdSwitchRoutingMulticastDefaultSettings, bool)`

GetDefaultSettingsOk returns a tuple with the DefaultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSettings

`func (o *InlineObject124) SetDefaultSettings(v NetworksNetworkIdSwitchRoutingMulticastDefaultSettings)`

SetDefaultSettings sets DefaultSettings field to given value.

### HasDefaultSettings

`func (o *InlineObject124) HasDefaultSettings() bool`

HasDefaultSettings returns a boolean if a field has been set.

### GetOverrides

`func (o *InlineObject124) GetOverrides() []NetworksNetworkIdSwitchRoutingMulticastOverrides`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *InlineObject124) GetOverridesOk() (*[]NetworksNetworkIdSwitchRoutingMulticastOverrides, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *InlineObject124) SetOverrides(v []NetworksNetworkIdSwitchRoutingMulticastOverrides)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *InlineObject124) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


