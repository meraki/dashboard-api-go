# UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Sets a manual channel for 5 GHz. Can be &#39;36&#39;, &#39;40&#39;, &#39;44&#39;, &#39;48&#39;, &#39;52&#39;, &#39;56&#39;, &#39;60&#39;, &#39;64&#39;, &#39;100&#39;, &#39;104&#39;, &#39;108&#39;, &#39;112&#39;, &#39;116&#39;, &#39;120&#39;, &#39;124&#39;, &#39;128&#39;, &#39;132&#39;, &#39;136&#39;, &#39;140&#39;, &#39;144&#39;, &#39;149&#39;, &#39;153&#39;, &#39;157&#39;, &#39;161&#39;, &#39;165&#39;, &#39;169&#39;, &#39;173&#39; or &#39;177&#39; or null for using auto channel. | [optional] 
**ChannelWidth** | Pointer to **int32** | Sets a manual channel for 5 GHz. Can be &#39;0&#39;, &#39;20&#39;, &#39;40&#39;, &#39;80&#39; or &#39;160&#39; or null for using auto channel width. | [optional] 
**TargetPower** | Pointer to **int32** | Set a manual target power for 5 GHz. Can be between &#39;8&#39; or &#39;30&#39; or null for using auto power range. | [optional] 

## Methods

### NewUpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings

`func NewUpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings() *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings`

NewUpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings instantiates a new UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceWirelessRadioSettingsRequestFiveGhzSettingsWithDefaults

`func NewUpdateDeviceWirelessRadioSettingsRequestFiveGhzSettingsWithDefaults() *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings`

NewUpdateDeviceWirelessRadioSettingsRequestFiveGhzSettingsWithDefaults instantiates a new UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChannelWidth

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) GetChannelWidth() int32`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) GetChannelWidthOk() (*int32, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) SetChannelWidth(v int32)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetTargetPower

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) GetTargetPower() int32`

GetTargetPower returns the TargetPower field if non-nil, zero value otherwise.

### GetTargetPowerOk

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) GetTargetPowerOk() (*int32, bool)`

GetTargetPowerOk returns a tuple with the TargetPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPower

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) SetTargetPower(v int32)`

SetTargetPower sets TargetPower field to given value.

### HasTargetPower

`func (o *UpdateDeviceWirelessRadioSettingsRequestFiveGhzSettings) HasTargetPower() bool`

HasTargetPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


