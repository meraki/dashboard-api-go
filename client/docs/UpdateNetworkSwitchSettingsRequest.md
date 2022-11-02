# UpdateNetworkSwitchSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **int32** | Management VLAN | [optional] 
**UseCombinedPower** | Pointer to **bool** | The use Combined Power as the default behavior of secondary power supplies on supported devices. | [optional] 
**PowerExceptions** | Pointer to [**[]UpdateNetworkSwitchSettingsRequestPowerExceptionsInner**](UpdateNetworkSwitchSettingsRequestPowerExceptionsInner.md) | Exceptions on a per switch basis to \&quot;useCombinedPower\&quot; | [optional] 

## Methods

### NewUpdateNetworkSwitchSettingsRequest

`func NewUpdateNetworkSwitchSettingsRequest() *UpdateNetworkSwitchSettingsRequest`

NewUpdateNetworkSwitchSettingsRequest instantiates a new UpdateNetworkSwitchSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkSwitchSettingsRequestWithDefaults

`func NewUpdateNetworkSwitchSettingsRequestWithDefaults() *UpdateNetworkSwitchSettingsRequest`

NewUpdateNetworkSwitchSettingsRequestWithDefaults instantiates a new UpdateNetworkSwitchSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *UpdateNetworkSwitchSettingsRequest) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *UpdateNetworkSwitchSettingsRequest) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *UpdateNetworkSwitchSettingsRequest) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *UpdateNetworkSwitchSettingsRequest) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetUseCombinedPower

`func (o *UpdateNetworkSwitchSettingsRequest) GetUseCombinedPower() bool`

GetUseCombinedPower returns the UseCombinedPower field if non-nil, zero value otherwise.

### GetUseCombinedPowerOk

`func (o *UpdateNetworkSwitchSettingsRequest) GetUseCombinedPowerOk() (*bool, bool)`

GetUseCombinedPowerOk returns a tuple with the UseCombinedPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCombinedPower

`func (o *UpdateNetworkSwitchSettingsRequest) SetUseCombinedPower(v bool)`

SetUseCombinedPower sets UseCombinedPower field to given value.

### HasUseCombinedPower

`func (o *UpdateNetworkSwitchSettingsRequest) HasUseCombinedPower() bool`

HasUseCombinedPower returns a boolean if a field has been set.

### GetPowerExceptions

`func (o *UpdateNetworkSwitchSettingsRequest) GetPowerExceptions() []UpdateNetworkSwitchSettingsRequestPowerExceptionsInner`

GetPowerExceptions returns the PowerExceptions field if non-nil, zero value otherwise.

### GetPowerExceptionsOk

`func (o *UpdateNetworkSwitchSettingsRequest) GetPowerExceptionsOk() (*[]UpdateNetworkSwitchSettingsRequestPowerExceptionsInner, bool)`

GetPowerExceptionsOk returns a tuple with the PowerExceptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerExceptions

`func (o *UpdateNetworkSwitchSettingsRequest) SetPowerExceptions(v []UpdateNetworkSwitchSettingsRequestPowerExceptionsInner)`

SetPowerExceptions sets PowerExceptions field to given value.

### HasPowerExceptions

`func (o *UpdateNetworkSwitchSettingsRequest) HasPowerExceptions() bool`

HasPowerExceptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


