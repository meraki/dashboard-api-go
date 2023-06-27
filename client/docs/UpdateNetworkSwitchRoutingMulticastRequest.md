# UpdateNetworkSwitchRoutingMulticastRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSettings** | Pointer to [**UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings**](UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings.md) |  | [optional] 
**Overrides** | Pointer to [**[]UpdateNetworkSwitchRoutingMulticastRequestOverridesInner**](UpdateNetworkSwitchRoutingMulticastRequestOverridesInner.md) | Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings. | [optional] 

## Methods

### NewUpdateNetworkSwitchRoutingMulticastRequest

`func NewUpdateNetworkSwitchRoutingMulticastRequest() *UpdateNetworkSwitchRoutingMulticastRequest`

NewUpdateNetworkSwitchRoutingMulticastRequest instantiates a new UpdateNetworkSwitchRoutingMulticastRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchRoutingMulticastRequestWithDefaults

`func NewUpdateNetworkSwitchRoutingMulticastRequestWithDefaults() *UpdateNetworkSwitchRoutingMulticastRequest`

NewUpdateNetworkSwitchRoutingMulticastRequestWithDefaults instantiates a new UpdateNetworkSwitchRoutingMulticastRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSettings

`func (o *UpdateNetworkSwitchRoutingMulticastRequest) GetDefaultSettings() UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings`

GetDefaultSettings returns the DefaultSettings field if non-nil, zero value otherwise.

### GetDefaultSettingsOk

`func (o *UpdateNetworkSwitchRoutingMulticastRequest) GetDefaultSettingsOk() (*UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings, bool)`

GetDefaultSettingsOk returns a tuple with the DefaultSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSettings

`func (o *UpdateNetworkSwitchRoutingMulticastRequest) SetDefaultSettings(v UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings)`

SetDefaultSettings sets DefaultSettings field to given value.

### HasDefaultSettings

`func (o *UpdateNetworkSwitchRoutingMulticastRequest) HasDefaultSettings() bool`

HasDefaultSettings returns a boolean if a field has been set.

### GetOverrides

`func (o *UpdateNetworkSwitchRoutingMulticastRequest) GetOverrides() []UpdateNetworkSwitchRoutingMulticastRequestOverridesInner`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *UpdateNetworkSwitchRoutingMulticastRequest) GetOverridesOk() (*[]UpdateNetworkSwitchRoutingMulticastRequestOverridesInner, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *UpdateNetworkSwitchRoutingMulticastRequest) SetOverrides(v []UpdateNetworkSwitchRoutingMulticastRequestOverridesInner)`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *UpdateNetworkSwitchRoutingMulticastRequest) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


