# DevicesSerialWirelessRadioSettingsTwoFourGhzSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Sets a manual channel for 2.4 GHz. Can be &#39;1&#39;, &#39;2&#39;, &#39;3&#39;, &#39;4&#39;, &#39;5&#39;, &#39;6&#39;, &#39;7&#39;, &#39;8&#39;, &#39;9&#39;, &#39;10&#39;, &#39;11&#39;, &#39;12&#39;, &#39;13&#39; or &#39;14&#39; or null for using auto channel. | [optional] 
**TargetPower** | Pointer to **int32** | Set a manual target power for 2.4 GHz. Can be between &#39;5&#39; or &#39;30&#39; or null for using auto power range. | [optional] 

## Methods

### NewDevicesSerialWirelessRadioSettingsTwoFourGhzSettings

`func NewDevicesSerialWirelessRadioSettingsTwoFourGhzSettings() *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings`

NewDevicesSerialWirelessRadioSettingsTwoFourGhzSettings instantiates a new DevicesSerialWirelessRadioSettingsTwoFourGhzSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevicesSerialWirelessRadioSettingsTwoFourGhzSettingsWithDefaults

`func NewDevicesSerialWirelessRadioSettingsTwoFourGhzSettingsWithDefaults() *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings`

NewDevicesSerialWirelessRadioSettingsTwoFourGhzSettingsWithDefaults instantiates a new DevicesSerialWirelessRadioSettingsTwoFourGhzSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTargetPower

`func (o *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings) GetTargetPower() int32`

GetTargetPower returns the TargetPower field if non-nil, zero value otherwise.

### GetTargetPowerOk

`func (o *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings) GetTargetPowerOk() (*int32, bool)`

GetTargetPowerOk returns a tuple with the TargetPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPower

`func (o *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings) SetTargetPower(v int32)`

SetTargetPower sets TargetPower field to given value.

### HasTargetPower

`func (o *DevicesSerialWirelessRadioSettingsTwoFourGhzSettings) HasTargetPower() bool`

HasTargetPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


