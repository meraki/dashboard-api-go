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

// CreateNetworkWirelessRfProfileRequestPerSsidSettings3 Settings for SSID 3
type CreateNetworkWirelessRfProfileRequestPerSsidSettings3 struct {
	// Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.
	MinBitrate *float32 `json:"minBitrate,omitempty"`
	// Choice between 'dual', '2.4ghz' or '5ghz'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewCreateNetworkWirelessRfProfileRequestPerSsidSettings3 instantiates a new CreateNetworkWirelessRfProfileRequestPerSsidSettings3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkWirelessRfProfileRequestPerSsidSettings3() *CreateNetworkWirelessRfProfileRequestPerSsidSettings3 {
	this := CreateNetworkWirelessRfProfileRequestPerSsidSettings3{}
	return &this
}

// NewCreateNetworkWirelessRfProfileRequestPerSsidSettings3WithDefaults instantiates a new CreateNetworkWirelessRfProfileRequestPerSsidSettings3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkWirelessRfProfileRequestPerSsidSettings3WithDefaults() *CreateNetworkWirelessRfProfileRequestPerSsidSettings3 {
	this := CreateNetworkWirelessRfProfileRequestPerSsidSettings3{}
	return &this
}

// GetMinBitrate returns the MinBitrate field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) GetMinBitrate() float32 {
	if o == nil || isNil(o.MinBitrate) {
		var ret float32
		return ret
	}
	return *o.MinBitrate
}

// GetMinBitrateOk returns a tuple with the MinBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) GetMinBitrateOk() (*float32, bool) {
	if o == nil || isNil(o.MinBitrate) {
    return nil, false
	}
	return o.MinBitrate, true
}

// HasMinBitrate returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) HasMinBitrate() bool {
	if o != nil && !isNil(o.MinBitrate) {
		return true
	}

	return false
}

// SetMinBitrate gets a reference to the given float32 and assigns it to the MinBitrate field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) SetMinBitrate(v float32) {
	o.MinBitrate = &v
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) GetBandOperationMode() string {
	if o == nil || isNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) GetBandOperationModeOk() (*string, bool) {
	if o == nil || isNil(o.BandOperationMode) {
    return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) HasBandOperationMode() bool {
	if o != nil && !isNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) GetBandSteeringEnabled() bool {
	if o == nil || isNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BandSteeringEnabled) {
    return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) HasBandSteeringEnabled() bool {
	if o != nil && !isNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o CreateNetworkWirelessRfProfileRequestPerSsidSettings3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MinBitrate) {
		toSerialize["minBitrate"] = o.MinBitrate
	}
	if !isNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !isNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings3 struct {
	value *CreateNetworkWirelessRfProfileRequestPerSsidSettings3
	isSet bool
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings3) Get() *CreateNetworkWirelessRfProfileRequestPerSsidSettings3 {
	return v.value
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings3) Set(val *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings3) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkWirelessRfProfileRequestPerSsidSettings3(val *CreateNetworkWirelessRfProfileRequestPerSsidSettings3) *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings3 {
	return &NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings3{value: val, isSet: true}
}

func (v NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkWirelessRfProfileRequestPerSsidSettings3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


