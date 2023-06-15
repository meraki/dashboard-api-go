# InlineResponse20092SixGhzSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxPower** | Pointer to **int32** | Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 30. | [optional] 
**MinPower** | Pointer to **int32** | Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 8. | [optional] 
**MinBitrate** | Pointer to **int32** | Sets min bitrate (Mbps) of 6Ghz band. Can be one of &#39;6&#39;, &#39;9&#39;, &#39;12&#39;, &#39;18&#39;, &#39;24&#39;, &#39;36&#39;, &#39;48&#39; or &#39;54&#39;. Defaults to 12. | [optional] 
**ValidAutoChannels** | Pointer to **[]int32** | Sets valid auto channels for 6Ghz band. Can be one of &#39;1&#39;, &#39;5&#39;, &#39;9&#39;, &#39;13&#39;, &#39;17&#39;, &#39;21&#39;, &#39;25&#39;, &#39;29&#39;, &#39;33&#39;, &#39;37&#39;, &#39;41&#39;, &#39;45&#39;, &#39;49&#39;, &#39;53&#39;, &#39;57&#39;, &#39;61&#39;, &#39;65&#39;, &#39;69&#39;, &#39;73&#39;, &#39;77&#39;, &#39;81&#39;, &#39;85&#39;, &#39;89&#39;, &#39;93&#39;, &#39;97&#39;, &#39;101&#39;, &#39;105&#39;, &#39;109&#39;, &#39;113&#39;, &#39;117&#39;, &#39;121&#39;, &#39;125&#39;, &#39;129&#39;, &#39;133&#39;, &#39;137&#39;, &#39;141&#39;, &#39;145&#39;, &#39;149&#39;, &#39;153&#39;, &#39;157&#39;, &#39;161&#39;, &#39;165&#39;, &#39;169&#39;, &#39;173&#39;, &#39;177&#39;, &#39;181&#39;, &#39;185&#39;, &#39;189&#39;, &#39;193&#39;, &#39;197&#39;, &#39;201&#39;, &#39;205&#39;, &#39;209&#39;, &#39;213&#39;, &#39;217&#39;, &#39;221&#39;, &#39;225&#39;, &#39;229&#39; or &#39;233&#39;. Defaults to auto. | [optional] 
**ChannelWidth** | Pointer to **string** | Sets channel width (MHz) for 6Ghz band. Can be one of &#39;0&#39;, &#39;20&#39;, &#39;40&#39;, &#39;80&#39; or &#39;160&#39;. Defaults to auto. | [optional] 
**Rxsop** | Pointer to **int32** | The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default. | [optional] 

## Methods

### NewInlineResponse20092SixGhzSettings

`func NewInlineResponse20092SixGhzSettings() *InlineResponse20092SixGhzSettings`

NewInlineResponse20092SixGhzSettings instantiates a new InlineResponse20092SixGhzSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse20092SixGhzSettingsWithDefaults

`func NewInlineResponse20092SixGhzSettingsWithDefaults() *InlineResponse20092SixGhzSettings`

NewInlineResponse20092SixGhzSettingsWithDefaults instantiates a new InlineResponse20092SixGhzSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxPower

`func (o *InlineResponse20092SixGhzSettings) GetMaxPower() int32`

GetMaxPower returns the MaxPower field if non-nil, zero value otherwise.

### GetMaxPowerOk

`func (o *InlineResponse20092SixGhzSettings) GetMaxPowerOk() (*int32, bool)`

GetMaxPowerOk returns a tuple with the MaxPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPower

`func (o *InlineResponse20092SixGhzSettings) SetMaxPower(v int32)`

SetMaxPower sets MaxPower field to given value.

### HasMaxPower

`func (o *InlineResponse20092SixGhzSettings) HasMaxPower() bool`

HasMaxPower returns a boolean if a field has been set.

### GetMinPower

`func (o *InlineResponse20092SixGhzSettings) GetMinPower() int32`

GetMinPower returns the MinPower field if non-nil, zero value otherwise.

### GetMinPowerOk

`func (o *InlineResponse20092SixGhzSettings) GetMinPowerOk() (*int32, bool)`

GetMinPowerOk returns a tuple with the MinPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPower

`func (o *InlineResponse20092SixGhzSettings) SetMinPower(v int32)`

SetMinPower sets MinPower field to given value.

### HasMinPower

`func (o *InlineResponse20092SixGhzSettings) HasMinPower() bool`

HasMinPower returns a boolean if a field has been set.

### GetMinBitrate

`func (o *InlineResponse20092SixGhzSettings) GetMinBitrate() int32`

GetMinBitrate returns the MinBitrate field if non-nil, zero value otherwise.

### GetMinBitrateOk

`func (o *InlineResponse20092SixGhzSettings) GetMinBitrateOk() (*int32, bool)`

GetMinBitrateOk returns a tuple with the MinBitrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBitrate

`func (o *InlineResponse20092SixGhzSettings) SetMinBitrate(v int32)`

SetMinBitrate sets MinBitrate field to given value.

### HasMinBitrate

`func (o *InlineResponse20092SixGhzSettings) HasMinBitrate() bool`

HasMinBitrate returns a boolean if a field has been set.

### GetValidAutoChannels

`func (o *InlineResponse20092SixGhzSettings) GetValidAutoChannels() []int32`

GetValidAutoChannels returns the ValidAutoChannels field if non-nil, zero value otherwise.

### GetValidAutoChannelsOk

`func (o *InlineResponse20092SixGhzSettings) GetValidAutoChannelsOk() (*[]int32, bool)`

GetValidAutoChannelsOk returns a tuple with the ValidAutoChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidAutoChannels

`func (o *InlineResponse20092SixGhzSettings) SetValidAutoChannels(v []int32)`

SetValidAutoChannels sets ValidAutoChannels field to given value.

### HasValidAutoChannels

`func (o *InlineResponse20092SixGhzSettings) HasValidAutoChannels() bool`

HasValidAutoChannels returns a boolean if a field has been set.

### GetChannelWidth

`func (o *InlineResponse20092SixGhzSettings) GetChannelWidth() string`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *InlineResponse20092SixGhzSettings) GetChannelWidthOk() (*string, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *InlineResponse20092SixGhzSettings) SetChannelWidth(v string)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *InlineResponse20092SixGhzSettings) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetRxsop

`func (o *InlineResponse20092SixGhzSettings) GetRxsop() int32`

GetRxsop returns the Rxsop field if non-nil, zero value otherwise.

### GetRxsopOk

`func (o *InlineResponse20092SixGhzSettings) GetRxsopOk() (*int32, bool)`

GetRxsopOk returns a tuple with the Rxsop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxsop

`func (o *InlineResponse20092SixGhzSettings) SetRxsop(v int32)`

SetRxsop sets Rxsop field to given value.

### HasRxsop

`func (o *InlineResponse20092SixGhzSettings) HasRxsop() bool`

HasRxsop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


