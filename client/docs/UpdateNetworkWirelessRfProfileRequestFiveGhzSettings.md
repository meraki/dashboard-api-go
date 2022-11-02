# UpdateNetworkWirelessRfProfileRequestFiveGhzSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPower** | Pointer to **int32** | Sets max power (dBm) of 5Ghz band. Can be integer between 8 and 30. | [optional] 
**MinPower** | Pointer to **int32** | Sets min power (dBm) of 5Ghz band. Can be integer between 8 and 30. | [optional] 
**MinBitrate** | Pointer to **int32** | Sets min bitrate (Mbps) of 5Ghz band. Can be one of &#39;6&#39;, &#39;9&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;. | [optional] 
**ValidAutoChannels** | Pointer to **[]int32** | Sets valid auto channels for 5Ghz band. Can be one of &#39;36&#39;, &#39;40&#39;, &#39;44&#39;, &#39;48&#39;, &#39;52&#39;, &#39;56&#39;, &#39;60&#39;, &#39;64&#39;, &#39;100&#39;, &#39;104&#39;, &#39;108&#39;, &#39;112&#39;, &#39;116&#39;, &#39;120&#39;, &#39;124&#39;, &#39;128&#39;, &#39;132&#39;, &#39;136&#39;, &#39;140&#39;, &#39;144&#39;, &#39;149&#39;, &#39;153&#39;, &#39;157&#39;, &#39;161&#39; or &#39;165&#39;. | [optional] 
**ChannelWidth** | Pointer to **string** | Sets channel width (MHz) for 5Ghz band. Can be one of &#39;auto&#39;, &#39;20&#39;, &#39;40&#39; or &#39;80&#39;. | [optional] 
**Rxsop** | Pointer to **int32** | The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default. | [optional] 

## Methods

### NewUpdateNetworkWirelessRfProfileRequestFiveGhzSettings

`func NewUpdateNetworkWirelessRfProfileRequestFiveGhzSettings() *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings`

NewUpdateNetworkWirelessRfProfileRequestFiveGhzSettings instantiates a new UpdateNetworkWirelessRfProfileRequestFiveGhzSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessRfProfileRequestFiveGhzSettingsWithDefaults

`func NewUpdateNetworkWirelessRfProfileRequestFiveGhzSettingsWithDefaults() *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings`

NewUpdateNetworkWirelessRfProfileRequestFiveGhzSettingsWithDefaults instantiates a new UpdateNetworkWirelessRfProfileRequestFiveGhzSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPower

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetMaxPower() int32`

GetMaxPower returns the MaxPower field if non-nil, zero value otherwise.

### GetMaxPowerOk

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetMaxPowerOk() (*int32, bool)`

GetMaxPowerOk returns a tuple with the MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPower

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) SetMaxPower(v int32)`

SetMaxPower sets MaxPower field to given value.

### HasMaxPower

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) HasMaxPower() bool`

HasMaxPower returns a boolean if a field has been set.

### GetMinPower

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetMinPower() int32`

GetMinPower returns the MinPower field if non-nil, zero value otherwise.

### GetMinPowerOk

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetMinPowerOk() (*int32, bool)`

GetMinPowerOk returns a tuple with the MinPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPower

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) SetMinPower(v int32)`

SetMinPower sets MinPower field to given value.

### HasMinPower

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) HasMinPower() bool`

HasMinPower returns a boolean if a field has been set.

### GetMinBitrate

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetMinBitrate() int32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetMinBitrateOk() (*int32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) SetMinBitrate(v int32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetValidAutoChannels

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetValidAutoChannels() []int32`

GetValidAutoChannels returns the ValidAutoChannels field if non-nil, zero value otherwise.

### GetValidAutoChannelsOk

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetValidAutoChannelsOk() (*[]int32, bool)`

GetValidAutoChannelsOk returns a tuple with the ValidAutoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidAutoChannels

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) SetValidAutoChannels(v []int32)`

SetValidAutoChannels sets ValidAutoChannels field to given value.

### HasValidAutoChannels

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) HasValidAutoChannels() bool`

HasValidAutoChannels returns a boolean if a field has been set.

### GetChannelWidth

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetChannelWidth() string`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetChannelWidthOk() (*string, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) SetChannelWidth(v string)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetRxsop

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetRxsop() int32`

GetRxsop returns the Rxsop field if non-nil, zero value otherwise.

### GetRxsopOk

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) GetRxsopOk() (*int32, bool)`

GetRxsopOk returns a tuple with the Rxsop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxsop

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) SetRxsop(v int32)`

SetRxsop sets Rxsop field to given value.

### HasRxsop

`func (o *UpdateNetworkWirelessRfProfileRequestFiveGhzSettings) HasRxsop() bool`

HasRxsop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


