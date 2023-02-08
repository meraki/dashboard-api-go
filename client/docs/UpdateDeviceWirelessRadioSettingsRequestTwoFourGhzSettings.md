# UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Sets a manual channel for 2.4 GHz. Can be &#39;1&#39;, &#39;2&#39;, &#39;3&#39;, &#39;4&#39;, &#39;5&#39;, &#39;6&#39;, &#39;7&#39;, &#39;8&#39;, &#39;9&#39;, &#39;10&#39;, &#39;11&#39;, &#39;12&#39;, &#39;13&#39; or &#39;14&#39; or null for using auto channel. | [optional] 
**TargetPower** | Pointer to **int32** | Set a manual target power for 2.4 GHz. Can be between &#39;5&#39; or &#39;30&#39; or null for using auto power range. | [optional] 

## Methods

### NewUpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings

`func NewUpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings() *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings`

NewUpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings instantiates a new UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettingsWithDefaults

`func NewUpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettingsWithDefaults() *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings`

NewUpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettingsWithDefaults instantiates a new UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetTargetPower

`func (o *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings) GetTargetPower() int32`

GetTargetPower returns the TargetPower field if non-nil, zero value otherwise.

### GetTargetPowerOk

`func (o *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings) GetTargetPowerOk() (*int32, bool)`

GetTargetPowerOk returns a tuple with the TargetPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPower

`func (o *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings) SetTargetPower(v int32)`

SetTargetPower sets TargetPower field to given value.

### HasTargetPower

`func (o *UpdateDeviceWirelessRadioSettingsRequestTwoFourGhzSettings) HasTargetPower() bool`

HasTargetPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


