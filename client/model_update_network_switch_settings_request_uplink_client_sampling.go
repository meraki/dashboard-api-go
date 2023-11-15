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

// checks if the UpdateNetworkSwitchSettingsRequestUplinkClientSampling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchSettingsRequestUplinkClientSampling{}

// UpdateNetworkSwitchSettingsRequestUplinkClientSampling Uplink client sampling
type UpdateNetworkSwitchSettingsRequestUplinkClientSampling struct {
	// Enable uplink client sampling
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkSwitchSettingsRequestUplinkClientSampling instantiates a new UpdateNetworkSwitchSettingsRequestUplinkClientSampling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchSettingsRequestUplinkClientSampling() *UpdateNetworkSwitchSettingsRequestUplinkClientSampling {
	this := UpdateNetworkSwitchSettingsRequestUplinkClientSampling{}
	return &this
}

// NewUpdateNetworkSwitchSettingsRequestUplinkClientSamplingWithDefaults instantiates a new UpdateNetworkSwitchSettingsRequestUplinkClientSampling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchSettingsRequestUplinkClientSamplingWithDefaults() *UpdateNetworkSwitchSettingsRequestUplinkClientSampling {
	this := UpdateNetworkSwitchSettingsRequestUplinkClientSampling{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchSettingsRequestUplinkClientSampling) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchSettingsRequestUplinkClientSampling) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchSettingsRequestUplinkClientSampling) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkSwitchSettingsRequestUplinkClientSampling) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkSwitchSettingsRequestUplinkClientSampling) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchSettingsRequestUplinkClientSampling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchSettingsRequestUplinkClientSampling struct {
	value *UpdateNetworkSwitchSettingsRequestUplinkClientSampling
	isSet bool
}

func (v NullableUpdateNetworkSwitchSettingsRequestUplinkClientSampling) Get() *UpdateNetworkSwitchSettingsRequestUplinkClientSampling {
	return v.value
}

func (v *NullableUpdateNetworkSwitchSettingsRequestUplinkClientSampling) Set(val *UpdateNetworkSwitchSettingsRequestUplinkClientSampling) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchSettingsRequestUplinkClientSampling) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchSettingsRequestUplinkClientSampling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchSettingsRequestUplinkClientSampling(val *UpdateNetworkSwitchSettingsRequestUplinkClientSampling) *NullableUpdateNetworkSwitchSettingsRequestUplinkClientSampling {
	return &NullableUpdateNetworkSwitchSettingsRequestUplinkClientSampling{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchSettingsRequestUplinkClientSampling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchSettingsRequestUplinkClientSampling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


