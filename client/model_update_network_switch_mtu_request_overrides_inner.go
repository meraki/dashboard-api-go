/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// UpdateNetworkSwitchMtuRequestOverridesInner struct for UpdateNetworkSwitchMtuRequestOverridesInner
type UpdateNetworkSwitchMtuRequestOverridesInner struct {
	// List of switch serials. Applicable only for switch network.
	Switches []string `json:"switches,omitempty"`
	// List of switch profile IDs. Applicable only for template network.
	SwitchProfiles []string `json:"switchProfiles,omitempty"`
	// MTU size for the switches or switch profiles.
	MtuSize int32 `json:"mtuSize"`
}

// NewUpdateNetworkSwitchMtuRequestOverridesInner instantiates a new UpdateNetworkSwitchMtuRequestOverridesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchMtuRequestOverridesInner(mtuSize int32) *UpdateNetworkSwitchMtuRequestOverridesInner {
	this := UpdateNetworkSwitchMtuRequestOverridesInner{}
	this.MtuSize = mtuSize
	return &this
}

// NewUpdateNetworkSwitchMtuRequestOverridesInnerWithDefaults instantiates a new UpdateNetworkSwitchMtuRequestOverridesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchMtuRequestOverridesInnerWithDefaults() *UpdateNetworkSwitchMtuRequestOverridesInner {
	this := UpdateNetworkSwitchMtuRequestOverridesInner{}
	return &this
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) GetSwitches() []string {
	if o == nil || isNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) GetSwitchesOk() ([]string, bool) {
	if o == nil || isNil(o.Switches) {
    return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) HasSwitches() bool {
	if o != nil && !isNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) SetSwitches(v []string) {
	o.Switches = v
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) GetSwitchProfiles() []string {
	if o == nil || isNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || isNil(o.SwitchProfiles) {
    return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) HasSwitchProfiles() bool {
	if o != nil && !isNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetMtuSize returns the MtuSize field value
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) GetMtuSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MtuSize
}

// GetMtuSizeOk returns a tuple with the MtuSize field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) GetMtuSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.MtuSize, true
}

// SetMtuSize sets field value
func (o *UpdateNetworkSwitchMtuRequestOverridesInner) SetMtuSize(v int32) {
	o.MtuSize = v
}

func (o UpdateNetworkSwitchMtuRequestOverridesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	if !isNil(o.SwitchProfiles) {
		toSerialize["switchProfiles"] = o.SwitchProfiles
	}
	if true {
		toSerialize["mtuSize"] = o.MtuSize
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSwitchMtuRequestOverridesInner struct {
	value *UpdateNetworkSwitchMtuRequestOverridesInner
	isSet bool
}

func (v NullableUpdateNetworkSwitchMtuRequestOverridesInner) Get() *UpdateNetworkSwitchMtuRequestOverridesInner {
	return v.value
}

func (v *NullableUpdateNetworkSwitchMtuRequestOverridesInner) Set(val *UpdateNetworkSwitchMtuRequestOverridesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchMtuRequestOverridesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchMtuRequestOverridesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchMtuRequestOverridesInner(val *UpdateNetworkSwitchMtuRequestOverridesInner) *NullableUpdateNetworkSwitchMtuRequestOverridesInner {
	return &NullableUpdateNetworkSwitchMtuRequestOverridesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchMtuRequestOverridesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchMtuRequestOverridesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


