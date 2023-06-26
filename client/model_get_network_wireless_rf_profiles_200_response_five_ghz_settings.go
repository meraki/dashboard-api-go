/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 07 June, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.34.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings Settings related to 5Ghz band
type GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings struct {
	// Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.
	MaxPower *int32 `json:"maxPower,omitempty"`
	// Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.
	MinPower *int32 `json:"minPower,omitempty"`
	// Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Sets valid auto channels for 5Ghz band. Can be one of '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161' or '165'.Defaults to [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165].
	ValidAutoChannels []int32 `json:"validAutoChannels,omitempty"`
	// Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'. Defaults to auto.
	ChannelWidth *string `json:"channelWidth,omitempty"`
	// The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.
	Rxsop *int32 `json:"rxsop,omitempty"`
}

// NewGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings instantiates a new GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings() *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings {
	this := GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings{}
	return &this
}

// NewGetNetworkWirelessRfProfiles200ResponseFiveGhzSettingsWithDefaults instantiates a new GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessRfProfiles200ResponseFiveGhzSettingsWithDefaults() *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings {
	this := GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings{}
	return &this
}

// GetMaxPower returns the MaxPower field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetMaxPower() int32 {
	if o == nil || isNil(o.MaxPower) {
		var ret int32
		return ret
	}
	return *o.MaxPower
}

// GetMaxPowerOk returns a tuple with the MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetMaxPowerOk() (*int32, bool) {
	if o == nil || isNil(o.MaxPower) {
    return nil, false
	}
	return o.MaxPower, true
}

// HasMaxPower returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) HasMaxPower() bool {
	if o != nil && !isNil(o.MaxPower) {
		return true
	}

	return false
}

// SetMaxPower gets a reference to the given int32 and assigns it to the MaxPower field.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) SetMaxPower(v int32) {
	o.MaxPower = &v
}

// GetMinPower returns the MinPower field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetMinPower() int32 {
	if o == nil || isNil(o.MinPower) {
		var ret int32
		return ret
	}
	return *o.MinPower
}

// GetMinPowerOk returns a tuple with the MinPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetMinPowerOk() (*int32, bool) {
	if o == nil || isNil(o.MinPower) {
    return nil, false
	}
	return o.MinPower, true
}

// HasMinPower returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) HasMinPower() bool {
	if o != nil && !isNil(o.MinPower) {
		return true
	}

	return false
}

// SetMinPower gets a reference to the given int32 and assigns it to the MinPower field.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) SetMinPower(v int32) {
	o.MinPower = &v
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetMinBitrate() int32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetMinBitrateOk() (*int32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetValidAutoChannels returns the ValidAutoChannels field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetValidAutoChannels() []int32 {
	if o == nil || isNil(o.ValidAutoChannels) {
		var ret []int32
		return ret
	}
	return o.ValidAutoChannels
}

// GetValidAutoChannelsOk returns a tuple with the ValidAutoChannels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetValidAutoChannelsOk() ([]int32, bool) {
	if o == nil || isNil(o.ValidAutoChannels) {
    return nil, false
	}
	return o.ValidAutoChannels, true
}

// HasValidAutoChannels returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) HasValidAutoChannels() bool {
	if o != nil && !isNil(o.ValidAutoChannels) {
		return true
	}

	return false
}

// SetValidAutoChannels gets a reference to the given []int32 and assigns it to the ValidAutoChannels field.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) SetValidAutoChannels(v []int32) {
	o.ValidAutoChannels = v
}

// GetChannelWidth returns the ChannelWidth field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetChannelWidth() string {
	if o == nil || isNil(o.ChannelWidth) {
		var ret string
		return ret
	}
	return *o.ChannelWidth
}

// GetChannelWidthOk returns a tuple with the ChannelWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetChannelWidthOk() (*string, bool) {
	if o == nil || isNil(o.ChannelWidth) {
    return nil, false
	}
	return o.ChannelWidth, true
}

// HasChannelWidth returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) HasChannelWidth() bool {
	if o != nil && !isNil(o.ChannelWidth) {
		return true
	}

	return false
}

// SetChannelWidth gets a reference to the given string and assigns it to the ChannelWidth field.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) SetChannelWidth(v string) {
	o.ChannelWidth = &v
}

// GetRxsop returns the Rxsop field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetRxsop() int32 {
	if o == nil || isNil(o.Rxsop) {
		var ret int32
		return ret
	}
	return *o.Rxsop
}

// GetRxsopOk returns a tuple with the Rxsop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) GetRxsopOk() (*int32, bool) {
	if o == nil || isNil(o.Rxsop) {
    return nil, false
	}
	return o.Rxsop, true
}

// HasRxsop returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) HasRxsop() bool {
	if o != nil && !isNil(o.Rxsop) {
		return true
	}

	return false
}

// SetRxsop gets a reference to the given int32 and assigns it to the Rxsop field.
func (o *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) SetRxsop(v int32) {
	o.Rxsop = &v
}

func (o GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxPower) {
		toSerialize["maxPower"] = o.MaxPower
	}
	if !isNil(o.MinPower) {
		toSerialize["minPower"] = o.MinPower
	}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.ValidAutoChannels) {
		toSerialize["validAutoChannels"] = o.ValidAutoChannels
	}
	if !isNil(o.ChannelWidth) {
		toSerialize["channelWidth"] = o.ChannelWidth
	}
	if !isNil(o.Rxsop) {
		toSerialize["rxsop"] = o.Rxsop
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings struct {
	value *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings
	isSet bool
}

func (v NullableGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) Get() *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings {
	return v.value
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) Set(val *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings(val *GetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) *NullableGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings {
	return &NullableGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponseFiveGhzSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


