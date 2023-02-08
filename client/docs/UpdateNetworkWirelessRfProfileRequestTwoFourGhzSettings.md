# UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPower** | Pointer to **int32** | Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. | [optional] 
**MinPower** | Pointer to **int32** | Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. | [optional] 
**MinBitrate** | Pointer to **float32** | Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of &#39;1&#39;, &#39;2&#39;, &#39;5.5&#39;, &#39;6&#39;, &#39;9&#39;, &#39;11&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;. | [optional] 
**ValidAutoChannels** | Pointer to **[]int32** | Sets valid auto channels for 2.4Ghz band. Can be one of &#39;1&#39;, &#39;6&#39; or &#39;11&#39;. | [optional] 
**AxEnabled** | Pointer to **bool** | Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. | [optional] 
**Rxsop** | Pointer to **int32** | The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default. | [optional] 

## Methods

### NewUpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings

`func NewUpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings() *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings`

NewUpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings instantiates a new UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkWirelessRfProfileRequestTwoFourGhzSettingsWithDefaults

`func NewUpdateNetworkWirelessRfProfileRequestTwoFourGhzSettingsWithDefaults() *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings`

NewUpdateNetworkWirelessRfProfileRequestTwoFourGhzSettingsWithDefaults instantiates a new UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPower

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetMaxPower() int32`

GetMaxPower returns the MaxPower field if non-nil, zero value otherwise.

### GetMaxPowerOk

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetMaxPowerOk() (*int32, bool)`

GetMaxPowerOk returns a tuple with the MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPower

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) SetMaxPower(v int32)`

SetMaxPower sets MaxPower field to given value.

### HasMaxPower

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) HasMaxPower() bool`

HasMaxPower returns a boolean if a field has been set.

### GetMinPower

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetMinPower() int32`

GetMinPower returns the MinPower field if non-nil, zero value otherwise.

### GetMinPowerOk

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetMinPowerOk() (*int32, bool)`

GetMinPowerOk returns a tuple with the MinPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPower

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) SetMinPower(v int32)`

SetMinPower sets MinPower field to given value.

### HasMinPower

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) HasMinPower() bool`

HasMinPower returns a boolean if a field has been set.

### GetMinBitrate

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetMinBitrate() float32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetMinBitrateOk() (*float32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) SetMinBitrate(v float32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetValidAutoChannels

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetValidAutoChannels() []int32`

GetValidAutoChannels returns the ValidAutoChannels field if non-nil, zero value otherwise.

### GetValidAutoChannelsOk

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetValidAutoChannelsOk() (*[]int32, bool)`

GetValidAutoChannelsOk returns a tuple with the ValidAutoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidAutoChannels

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) SetValidAutoChannels(v []int32)`

SetValidAutoChannels sets ValidAutoChannels field to given value.

### HasValidAutoChannels

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) HasValidAutoChannels() bool`

HasValidAutoChannels returns a boolean if a field has been set.

### GetAxEnabled

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetAxEnabled() bool`

GetAxEnabled returns the AxEnabled field if non-nil, zero value otherwise.

### GetAxEnabledOk

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetAxEnabledOk() (*bool, bool)`

GetAxEnabledOk returns a tuple with the AxEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxEnabled

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) SetAxEnabled(v bool)`

SetAxEnabled sets AxEnabled field to given value.

### HasAxEnabled

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) HasAxEnabled() bool`

HasAxEnabled returns a boolean if a field has been set.

### GetRxsop

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetRxsop() int32`

GetRxsop returns the Rxsop field if non-nil, zero value otherwise.

### GetRxsopOk

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) GetRxsopOk() (*int32, bool)`

GetRxsopOk returns a tuple with the Rxsop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxsop

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) SetRxsop(v int32)`

SetRxsop sets Rxsop field to given value.

### HasRxsop

`func (o *UpdateNetworkWirelessRfProfileRequestTwoFourGhzSettings) HasRxsop() bool`

HasRxsop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


