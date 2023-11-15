/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetDeviceWirelessStatus200ResponseBasicServiceSetsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDeviceWirelessStatus200ResponseBasicServiceSetsInner{}

// GetDeviceWirelessStatus200ResponseBasicServiceSetsInner struct for GetDeviceWirelessStatus200ResponseBasicServiceSetsInner
type GetDeviceWirelessStatus200ResponseBasicServiceSetsInner struct {
	// Name of wireless network
	SsidName *string `json:"ssidName,omitempty"`
	// Unique identifier of wireless network
	SsidNumber *int32 `json:"ssidNumber,omitempty"`
	// Status of wireless network
	Enabled *bool `json:"enabled,omitempty"`
	// Frequency range used by wireless network
	Band *string `json:"band,omitempty"`
	// Unique identifier of wireless access point
	Bssid *string `json:"bssid,omitempty"`
	// Frequency channel used by wireless network
	Channel *int32 `json:"channel,omitempty"`
	// Width of frequency channel used by wireless network
	ChannelWidth *string `json:"channelWidth,omitempty"`
	// Strength of wireless signal
	Power *string `json:"power,omitempty"`
	// Whether the SSID is advertised or hidden
	Visible *bool `json:"visible,omitempty"`
	// Whether the SSID is broadcasting based on an availability schedule
	Broadcasting *bool `json:"broadcasting,omitempty"`
}

// NewGetDeviceWirelessStatus200ResponseBasicServiceSetsInner instantiates a new GetDeviceWirelessStatus200ResponseBasicServiceSetsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceWirelessStatus200ResponseBasicServiceSetsInner() *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner {
	this := GetDeviceWirelessStatus200ResponseBasicServiceSetsInner{}
	return &this
}

// NewGetDeviceWirelessStatus200ResponseBasicServiceSetsInnerWithDefaults instantiates a new GetDeviceWirelessStatus200ResponseBasicServiceSetsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceWirelessStatus200ResponseBasicServiceSetsInnerWithDefaults() *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner {
	this := GetDeviceWirelessStatus200ResponseBasicServiceSetsInner{}
	return &this
}

// GetSsidName returns the SsidName field value if set, zero value otherwise.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetSsidName() string {
	if o == nil || IsNil(o.SsidName) {
		var ret string
		return ret
	}
	return *o.SsidName
}

// GetSsidNameOk returns a tuple with the SsidName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetSsidNameOk() (*string, bool) {
	if o == nil || IsNil(o.SsidName) {
		return nil, false
	}
	return o.SsidName, true
}

// HasSsidName returns a boolean if a field has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasSsidName() bool {
	if o != nil && !IsNil(o.SsidName) {
		return true
	}

	return false
}

// SetSsidName gets a reference to the given string and assigns it to the SsidName field.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetSsidName(v string) {
	o.SsidName = &v
}

// GetSsidNumber returns the SsidNumber field value if set, zero value otherwise.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetSsidNumber() int32 {
	if o == nil || IsNil(o.SsidNumber) {
		var ret int32
		return ret
	}
	return *o.SsidNumber
}

// GetSsidNumberOk returns a tuple with the SsidNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetSsidNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.SsidNumber) {
		return nil, false
	}
	return o.SsidNumber, true
}

// HasSsidNumber returns a boolean if a field has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasSsidNumber() bool {
	if o != nil && !IsNil(o.SsidNumber) {
		return true
	}

	return false
}

// SetSsidNumber gets a reference to the given int32 and assigns it to the SsidNumber field.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetSsidNumber(v int32) {
	o.SsidNumber = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetBand returns the Band field value if set, zero value otherwise.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBand() string {
	if o == nil || IsNil(o.Band) {
		var ret string
		return ret
	}
	return *o.Band
}

// GetBandOk returns a tuple with the Band field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBandOk() (*string, bool) {
	if o == nil || IsNil(o.Band) {
		return nil, false
	}
	return o.Band, true
}

// HasBand returns a boolean if a field has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasBand() bool {
	if o != nil && !IsNil(o.Band) {
		return true
	}

	return false
}

// SetBand gets a reference to the given string and assigns it to the Band field.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetBand(v string) {
	o.Band = &v
}

// GetBssid returns the Bssid field value if set, zero value otherwise.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBssid() string {
	if o == nil || IsNil(o.Bssid) {
		var ret string
		return ret
	}
	return *o.Bssid
}

// GetBssidOk returns a tuple with the Bssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBssidOk() (*string, bool) {
	if o == nil || IsNil(o.Bssid) {
		return nil, false
	}
	return o.Bssid, true
}

// HasBssid returns a boolean if a field has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasBssid() bool {
	if o != nil && !IsNil(o.Bssid) {
		return true
	}

	return false
}

// SetBssid gets a reference to the given string and assigns it to the Bssid field.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetBssid(v string) {
	o.Bssid = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetChannel() int32 {
	if o == nil || IsNil(o.Channel) {
		var ret int32
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetChannelOk() (*int32, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given int32 and assigns it to the Channel field.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetChannel(v int32) {
	o.Channel = &v
}

// GetChannelWidth returns the ChannelWidth field value if set, zero value otherwise.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetChannelWidth() string {
	if o == nil || IsNil(o.ChannelWidth) {
		var ret string
		return ret
	}
	return *o.ChannelWidth
}

// GetChannelWidthOk returns a tuple with the ChannelWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetChannelWidthOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelWidth) {
		return nil, false
	}
	return o.ChannelWidth, true
}

// HasChannelWidth returns a boolean if a field has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasChannelWidth() bool {
	if o != nil && !IsNil(o.ChannelWidth) {
		return true
	}

	return false
}

// SetChannelWidth gets a reference to the given string and assigns it to the ChannelWidth field.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetChannelWidth(v string) {
	o.ChannelWidth = &v
}

// GetPower returns the Power field value if set, zero value otherwise.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetPower() string {
	if o == nil || IsNil(o.Power) {
		var ret string
		return ret
	}
	return *o.Power
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetPowerOk() (*string, bool) {
	if o == nil || IsNil(o.Power) {
		return nil, false
	}
	return o.Power, true
}

// HasPower returns a boolean if a field has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasPower() bool {
	if o != nil && !IsNil(o.Power) {
		return true
	}

	return false
}

// SetPower gets a reference to the given string and assigns it to the Power field.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetPower(v string) {
	o.Power = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetVisible(v bool) {
	o.Visible = &v
}

// GetBroadcasting returns the Broadcasting field value if set, zero value otherwise.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBroadcasting() bool {
	if o == nil || IsNil(o.Broadcasting) {
		var ret bool
		return ret
	}
	return *o.Broadcasting
}

// GetBroadcastingOk returns a tuple with the Broadcasting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) GetBroadcastingOk() (*bool, bool) {
	if o == nil || IsNil(o.Broadcasting) {
		return nil, false
	}
	return o.Broadcasting, true
}

// HasBroadcasting returns a boolean if a field has been set.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) HasBroadcasting() bool {
	if o != nil && !IsNil(o.Broadcasting) {
		return true
	}

	return false
}

// SetBroadcasting gets a reference to the given bool and assigns it to the Broadcasting field.
func (o *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) SetBroadcasting(v bool) {
	o.Broadcasting = &v
}

func (o GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SsidName) {
		toSerialize["ssidName"] = o.SsidName
	}
	if !IsNil(o.SsidNumber) {
		toSerialize["ssidNumber"] = o.SsidNumber
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Band) {
		toSerialize["band"] = o.Band
	}
	if !IsNil(o.Bssid) {
		toSerialize["bssid"] = o.Bssid
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.ChannelWidth) {
		toSerialize["channelWidth"] = o.ChannelWidth
	}
	if !IsNil(o.Power) {
		toSerialize["power"] = o.Power
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	if !IsNil(o.Broadcasting) {
		toSerialize["broadcasting"] = o.Broadcasting
	}
	return toSerialize, nil
}

type NullableGetDeviceWirelessStatus200ResponseBasicServiceSetsInner struct {
	value *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner
	isSet bool
}

func (v NullableGetDeviceWirelessStatus200ResponseBasicServiceSetsInner) Get() *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner {
	return v.value
}

func (v *NullableGetDeviceWirelessStatus200ResponseBasicServiceSetsInner) Set(val *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceWirelessStatus200ResponseBasicServiceSetsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceWirelessStatus200ResponseBasicServiceSetsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceWirelessStatus200ResponseBasicServiceSetsInner(val *GetDeviceWirelessStatus200ResponseBasicServiceSetsInner) *NullableGetDeviceWirelessStatus200ResponseBasicServiceSetsInner {
	return &NullableGetDeviceWirelessStatus200ResponseBasicServiceSetsInner{value: val, isSet: true}
}

func (v NullableGetDeviceWirelessStatus200ResponseBasicServiceSetsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceWirelessStatus200ResponseBasicServiceSetsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


