/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 04 October, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.38.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the CreateNetworkWirelessRfProfileRequestPerSsidSettings6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkWirelessRfProfileRequestPerSsidSettings6{}

// CreateNetworkWirelessRfProfileRequestPerSsidSettings6 Settings for SSID 6
type CreateNetworkWirelessRfProfileRequestPerSsidSettings6 struct {
	// Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	Bands *CreateNetworkWirelessRfProfileRequestApBandSettingsBands `json:"bands,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewCreateNetworkWirelessRfProfileRequestPerSsidSettings6 instantiates a new CreateNetworkWirelessRfProfileRequestPerSsidSettings6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfileRequestPerSsidSettings6() *CreateNetworkWirelessRfProfileRequestPerSsidSettings6 {
	this := CreateNetworkWirelessRfProfileRequestPerSsidSettings6{}
	return &this
}

// NewCreateNetworkWirelessRfProfileRequestPerSsidSettings6WithDefaults instantiates a new CreateNetworkWirelessRfProfileRequestPerSsidSettings6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfileRequestPerSsidSettings6WithDefaults() *CreateNetworkWirelessRfProfileRequestPerSsidSettings6 {
	this := CreateNetworkWirelessRfProfileRequestPerSsidSettings6{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) GetMinBitrate() float32 {
	if o == nil || IsNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) GetMinBitrateOk() (*float32, bool) {
	if o == nil || IsNil(o.MinBitrate) {
		return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) HasMinBitrate() bool {
	if o != nil && !IsNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) GetBandOperationMode() string {
	if o == nil || IsNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) GetBandOperationModeOk() (*string, bool) {
	if o == nil || IsNil(o.BandOperationMode) {
		return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) HasBandOperationMode() bool {
	if o != nil && !IsNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBands returns the Bands field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) GetBands() CreateNetworkWirelessRfProfileRequestApBandSettingsBands {
	if o == nil || IsNil(o.Bands) {
		var ret CreateNetworkWirelessRfProfileRequestApBandSettingsBands
		return ret
	}
	return *o.Bands
}

// GetBandsOk returns a tuple with the Bands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) GetBandsOk() (*CreateNetworkWirelessRfProfileRequestApBandSettingsBands, bool) {
	if o == nil || IsNil(o.Bands) {
		return nil, false
	}
	return o.Bands, true
}

// HasBands returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) HasBands() bool {
	if o != nil && !IsNil(o.Bands) {
		return true
	}

	return false
}

// SetBands gets a reference to the given CreateNetworkWirelessRfProfileRequestApBandSettingsBands and assigns it to the Bands field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) SetBands(v CreateNetworkWirelessRfProfileRequestApBandSettingsBands) {
	o.Bands = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) GetBandSteeringEnabled() bool {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BandSteeringEnabled) {
		return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) HasBandSteeringEnabled() bool {
	if o != nil && !IsNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o CreateNetworkWirelessRfProfileRequestPerSsidSettings6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkWirelessRfProfileRequestPerSsidSettings6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings6 struct {
	value *CreateNetworkWirelessRfProfileRequestPerSsidSettings6
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings6) Get() *CreateNetworkWirelessRfProfileRequestPerSsidSettings6 {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings6) Set(val *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings6) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfileRequestPerSsidSettings6(val *CreateNetworkWirelessRfProfileRequestPerSsidSettings6) *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings6 {
	return &NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings6{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


