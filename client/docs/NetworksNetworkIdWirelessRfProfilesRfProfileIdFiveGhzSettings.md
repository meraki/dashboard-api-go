# NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPower** | Pointer to **int32** | Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. | [optional] 
**MinPower** | Pointer to **int32** | Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. | [optional] 
**MinBitrate** | Pointer to **int32** | Sets min bitrate (Mbps) of 5Ghz band. Can be one of &#39;6&#39;, &#39;9&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;. | [optional] 
**ValidAutoChannels** | Pointer to **[]int32** | Sets valid auto channels for 5Ghz band. Can be one of &#39;36&#39;, &#39;40&#39;, &#39;44&#39;, &#39;48&#39;, &#39;52&#39;, &#39;56&#39;, &#39;60&#39;, &#39;64&#39;, &#39;100&#39;, &#39;104&#39;, &#39;108&#39;, &#39;112&#39;, &#39;116&#39;, &#39;120&#39;, &#39;124&#39;, &#39;128&#39;, &#39;132&#39;, &#39;136&#39;, &#39;140&#39;, &#39;144&#39;, &#39;149&#39;, &#39;153&#39;, &#39;157&#39;, &#39;161&#39; or &#39;165&#39;. | [optional] 
**ChannelWidth** | Pointer to **string** | Sets channel width (MHz) for 5Ghz band. Can be one of &#39;auto&#39;, &#39;20&#39;, &#39;40&#39; or &#39;80&#39;. | [optional] 
**Rxsop** | Pointer to **int32** | The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default. | [optional] 

## Methods

### NewNetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings

`func NewNetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings() *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings`

NewNetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings instantiates a new NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettingsWithDefaults

`func NewNetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettingsWithDefaults() *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings`

NewNetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettingsWithDefaults instantiates a new NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPower

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetMaxPower() int32`

GetMaxPower returns the MaxPower field if non-nil, zero value otherwise.

### GetMaxPowerOk

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetMaxPowerOk() (*int32, bool)`

GetMaxPowerOk returns a tuple with the MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPower

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) SetMaxPower(v int32)`

SetMaxPower sets MaxPower field to given value.

### HasMaxPower

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) HasMaxPower() bool`

HasMaxPower returns a boolean if a field has been set.

### GetMinPower

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetMinPower() int32`

GetMinPower returns the MinPower field if non-nil, zero value otherwise.

### GetMinPowerOk

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetMinPowerOk() (*int32, bool)`

GetMinPowerOk returns a tuple with the MinPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPower

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) SetMinPower(v int32)`

SetMinPower sets MinPower field to given value.

### HasMinPower

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) HasMinPower() bool`

HasMinPower returns a boolean if a field has been set.

### GetMinBitrate

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetMinBitrate() int32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetMinBitrateOk() (*int32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) SetMinBitrate(v int32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetValidAutoChannels

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetValidAutoChannels() []int32`

GetValidAutoChannels returns the ValidAutoChannels field if non-nil, zero value otherwise.

### GetValidAutoChannelsOk

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetValidAutoChannelsOk() (*[]int32, bool)`

GetValidAutoChannelsOk returns a tuple with the ValidAutoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidAutoChannels

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) SetValidAutoChannels(v []int32)`

SetValidAutoChannels sets ValidAutoChannels field to given value.

### HasValidAutoChannels

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) HasValidAutoChannels() bool`

HasValidAutoChannels returns a boolean if a field has been set.

### GetChannelWidth

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetChannelWidth() string`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetChannelWidthOk() (*string, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) SetChannelWidth(v string)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetRxsop

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetRxsop() int32`

GetRxsop returns the Rxsop field if non-nil, zero value otherwise.

### GetRxsopOk

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) GetRxsopOk() (*int32, bool)`

GetRxsopOk returns a tuple with the Rxsop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxsop

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) SetRxsop(v int32)`

SetRxsop sets Rxsop field to given value.

### HasRxsop

`func (o *NetworksNetworkIdWirelessRfProfilesRfProfileIdFiveGhzSettings) HasRxsop() bool`

HasRxsop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


