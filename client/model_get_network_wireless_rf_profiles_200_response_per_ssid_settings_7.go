/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7{}

// GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 Settings for SSID 7
type GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 struct {
	// Name of SSID
	Name *string `json:"name,omitempty"`
	// Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *int32 `json:"minBitrate,omitempty"`
	// Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	Bands *GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands `json:"bands,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 instantiates a new GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7() *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 {
	this := GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7{}
	return &this
}

// NewGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7WithDefaults instantiates a new GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7WithDefaults() *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 {
	this := GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) SetName(v string) {
	o.Name = &v
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) GetMinBitrate() int32 {
	if o == nil || IsNil(o.MinBitrate) {
		var ret int32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) GetMinBitrateOk() (*int32, bool) {
	if o == nil || IsNil(o.MinBitrate) {
		return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) HasMinBitrate() bool {
	if o != nil && !IsNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given int32 and assigns it to the MinBitrate field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) SetMinBitrate(v int32) {
	o.MinBitrate = &v
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) GetBandOperationMode() string {
	if o == nil || IsNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) GetBandOperationModeOk() (*string, bool) {
	if o == nil || IsNil(o.BandOperationMode) {
		return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) HasBandOperationMode() bool {
	if o != nil && !IsNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBands returns the Bands field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) GetBands() GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands {
	if o == nil || IsNil(o.Bands) {
		var ret GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands
		return ret
	}
	return *o.Bands
}

// GetBandsOk returns a tuple with the Bands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) GetBandsOk() (*GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands, bool) {
	if o == nil || IsNil(o.Bands) {
		return nil, false
	}
	return o.Bands, true
}

// HasBands returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) HasBands() bool {
	if o != nil && !IsNil(o.Bands) {
		return true
	}

	return false
}

// SetBands gets a reference to the given GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands and assigns it to the Bands field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) SetBands(v GetNetworkWirelessRfProfiles200ResponseApBandSettingsBands) {
	o.Bands = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) GetBandSteeringEnabled() bool {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) HasBandSteeringEnabled() bool {
	if o != nil && !IsNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !IsNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !IsNil(o.Bands) {
		toSerialize["bands"] = o.Bands
	}
	if !IsNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return toSerialize, nil
}

type NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 struct {
	value *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7
	isSet bool
}

func (v NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) Get() *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 {
	return v.value
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) Set(val *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7(val *GetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) *NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7 {
	return &NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponsePerSsidSettings7) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


