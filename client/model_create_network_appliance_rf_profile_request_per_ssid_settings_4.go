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

// CreateNetworkApplianceRfProfileRequestPerSsidSettings4 Settings for SSID 4
type CreateNetworkApplianceRfProfileRequestPerSsidSettings4 struct {
	// Choice between 'dual', '2.4ghz' or '5ghz'.
	BandOperationMode *string `json:"bandOperationMode,omitempty"`
	// Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.
	BandSteeringEnabled *bool `json:"bandSteeringEnabled,omitempty"`
}

// NewCreateNetworkApplianceRfProfileRequestPerSsidSettings4 instantiates a new CreateNetworkApplianceRfProfileRequestPerSsidSettings4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkApplianceRfProfileRequestPerSsidSettings4() *CreateNetworkApplianceRfProfileRequestPerSsidSettings4 {
	this := CreateNetworkApplianceRfProfileRequestPerSsidSettings4{}
	return &this
}

// NewCreateNetworkApplianceRfProfileRequestPerSsidSettings4WithDefaults instantiates a new CreateNetworkApplianceRfProfileRequestPerSsidSettings4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkApplianceRfProfileRequestPerSsidSettings4WithDefaults() *CreateNetworkApplianceRfProfileRequestPerSsidSettings4 {
	this := CreateNetworkApplianceRfProfileRequestPerSsidSettings4{}
	return &this
}

// GetBandOperationMode returns the BandOperationMode field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings4) GetBandOperationMode() string {
	if o == nil || isNil(o.BandOperationMode) {
		var ret string
		return ret
	}
	return *o.BandOperationMode
}

// GetBandOperationModeOk returns a tuple with the BandOperationMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings4) GetBandOperationModeOk() (*string, bool) {
	if o == nil || isNil(o.BandOperationMode) {
    return nil, false
	}
	return o.BandOperationMode, true
}

// HasBandOperationMode returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings4) HasBandOperationMode() bool {
	if o != nil && !isNil(o.BandOperationMode) {
		return true
	}

	return false
}

// SetBandOperationMode gets a reference to the given string and assigns it to the BandOperationMode field.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings4) SetBandOperationMode(v string) {
	o.BandOperationMode = &v
}

// GetBandSteeringEnabled returns the BandSteeringEnabled field value if set, zero value otherwise.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings4) GetBandSteeringEnabled() bool {
	if o == nil || isNil(o.BandSteeringEnabled) {
		var ret bool
		return ret
	}
	return *o.BandSteeringEnabled
}

// GetBandSteeringEnabledOk returns a tuple with the BandSteeringEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings4) GetBandSteeringEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.BandSteeringEnabled) {
    return nil, false
	}
	return o.BandSteeringEnabled, true
}

// HasBandSteeringEnabled returns a boolean if a field has been set.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings4) HasBandSteeringEnabled() bool {
	if o != nil && !isNil(o.BandSteeringEnabled) {
		return true
	}

	return false
}

// SetBandSteeringEnabled gets a reference to the given bool and assigns it to the BandSteeringEnabled field.
func (o *CreateNetworkApplianceRfProfileRequestPerSsidSettings4) SetBandSteeringEnabled(v bool) {
	o.BandSteeringEnabled = &v
}

func (o CreateNetworkApplianceRfProfileRequestPerSsidSettings4) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.BandOperationMode) {
		toSerialize["bandOperationMode"] = o.BandOperationMode
	}
	if !isNil(o.BandSteeringEnabled) {
		toSerialize["bandSteeringEnabled"] = o.BandSteeringEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings4 struct {
	value *CreateNetworkApplianceRfProfileRequestPerSsidSettings4
	isSet bool
}

func (v NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings4) Get() *CreateNetworkApplianceRfProfileRequestPerSsidSettings4 {
	return v.value
}

func (v *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings4) Set(val *CreateNetworkApplianceRfProfileRequestPerSsidSettings4) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings4) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkApplianceRfProfileRequestPerSsidSettings4(val *CreateNetworkApplianceRfProfileRequestPerSsidSettings4) *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings4 {
	return &NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings4{value: val, isSet: true}
}

func (v NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkApplianceRfProfileRequestPerSsidSettings4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


