# GetDeviceWirelessStatus200ResponseBasicServiceSetsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidName** | Pointer to **string** | Name of wireless network | [optional] 
**SsidNumber** | Pointer to **int32** | Unique identifier of wireless network | [optional] 
**Enabled** | Pointer to **bool** | Status of wireless network | [optional] 
**Band** | Pointer to **string** | Frequency range used by wireless network | [optional] 
**Bssid** | Pointer to **string** | Unique identifier of wireless access point | [optional] 
**Channel** | Pointer to **int32** | Frequency channel used by wireless network | [optional] 
**ChannelWidth** | Pointer to **string** | Width of frequency channel used by wireless network | [optional] 
**Power** | Pointer to **string** | Strength of wireless signal | [optional] 
**Visible** | Pointer to **bool** | Whether the SSID is advertised or hidden | [optional] 
**Broadcasting** | Pointer to **bool** | Whether the SSID is broadcasting based on an availability schedule | [optional] 

## Methods

### NewGetDeviceWirelessStatus200ResponseBasicServiceSetsInner

`func NewGetDeviceWirelessStatus200ResponseBasicServiceSetsInner() *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner`

NewGetDeviceWirelessStatus200ResponseBasicServiceSetsInner instantiates a new GetDeviceWirelessStatus200ResponseBasicServiceSetsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeviceWirelessStatus200ResponseBasicServiceSetsInnerWithDefaults

`func NewGetDeviceWirelessStatus200ResponseBasicServiceSetsInnerWithDefaults() *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner`

NewGetDeviceWirelessStatus200ResponseBasicServiceSetsInnerWithDefaults instantiates a new GetDeviceWirelessStatus200ResponseBasicServiceSetsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidName

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.

### HasSsidName

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasSsidName() bool`

HasSsidName returns a boolean if a field has been set.

### GetSsidNumber

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetSsidNumber() int32`

GetSsidNumber returns the SsidNumber field if non-nil, zero value otherwise.

### GetSsidNumberOk

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetSsidNumberOk() (*int32, bool)`

GetSsidNumberOk returns a tuple with the SsidNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNumber

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetSsidNumber(v int32)`

SetSsidNumber sets SsidNumber field to given value.

### HasSsidNumber

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasSsidNumber() bool`

HasSsidNumber returns a boolean if a field has been set.

### GetEnabled

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetBand

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetBssid

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBssid() string`

GetBssid returns the Bssid field if non-nil, zero value otherwise.

### GetBssidOk

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBssidOk() (*string, bool)`

GetBssidOk returns a tuple with the Bssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssid

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetBssid(v string)`

SetBssid sets Bssid field to given value.

### HasBssid

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasBssid() bool`

HasBssid returns a boolean if a field has been set.

### GetChannel

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChannelWidth

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetChannelWidth() string`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetChannelWidthOk() (*string, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetChannelWidth(v string)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetPower

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetPower() string`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetPowerOk() (*string, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetPower(v string)`

SetPower sets Power field to given value.

### HasPower

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetVisible

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetBroadcasting

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBroadcasting() bool`

GetBroadcasting returns the Broadcasting field if non-nil, zero value otherwise.

### GetBroadcastingOk

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBroadcastingOk() (*bool, bool)`

GetBroadcastingOk returns a tuple with the Broadcasting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasting

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetBroadcasting(v bool)`

SetBroadcasting sets Broadcasting field to given value.

### HasBroadcasting

`func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasBroadcasting() bool`

HasBroadcasting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


