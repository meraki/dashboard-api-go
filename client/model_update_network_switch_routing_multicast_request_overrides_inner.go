/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkSwitchRoutingMulticastRequestOverridesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchRoutingMulticastRequestOverridesInner{}

// UpdateNetworkSwitchRoutingMulticastRequestOverridesInner struct for UpdateNetworkSwitchRoutingMulticastRequestOverridesInner
type UpdateNetworkSwitchRoutingMulticastRequestOverridesInner struct {
	// List of switch templates ids for template network
	SwitchProfiles []string `json:"switchProfiles,omitempty"`
	// List of switch serials for non-template network
	Switches []string `json:"switches,omitempty"`
	// List of switch stack ids for non-template network
	Stacks []string `json:"stacks,omitempty"`
	// IGMP snooping setting for switches, switch stacks or switch templates
	IgmpSnoopingEnabled bool `json:"igmpSnoopingEnabled"`
	// Flood unknown multicast traffic setting for switches, switch stacks or switch templates
	FloodUnknownMulticastTrafficEnabled bool `json:"floodUnknownMulticastTrafficEnabled"`
}

// NewUpdateNetworkSwitchRoutingMulticastRequestOverridesInner instantiates a new UpdateNetworkSwitchRoutingMulticastRequestOverridesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchRoutingMulticastRequestOverridesInner(igmpSnoopingEnabled bool, floodUnknownMulticastTrafficEnabled bool) *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner {
	this := UpdateNetworkSwitchRoutingMulticastRequestOverridesInner{}
	this.IgmpSnoopingEnabled = igmpSnoopingEnabled
	this.FloodUnknownMulticastTrafficEnabled = floodUnknownMulticastTrafficEnabled
	return &this
}

// NewUpdateNetworkSwitchRoutingMulticastRequestOverridesInnerWithDefaults instantiates a new UpdateNetworkSwitchRoutingMulticastRequestOverridesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchRoutingMulticastRequestOverridesInnerWithDefaults() *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner {
	this := UpdateNetworkSwitchRoutingMulticastRequestOverridesInner{}
	return &this
}

// GetSwitchProfiles returns the SwitchProfiles field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetSwitchProfiles() []string {
	if o == nil || IsNil(o.SwitchProfiles) {
		var ret []string
		return ret
	}
	return o.SwitchProfiles
}

// GetSwitchProfilesOk returns a tuple with the SwitchProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetSwitchProfilesOk() ([]string, bool) {
	if o == nil || IsNil(o.SwitchProfiles) {
		return nil, false
	}
	return o.SwitchProfiles, true
}

// HasSwitchProfiles returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) HasSwitchProfiles() bool {
	if o != nil && !IsNil(o.SwitchProfiles) {
		return true
	}

	return false
}

// SetSwitchProfiles gets a reference to the given []string and assigns it to the SwitchProfiles field.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) SetSwitchProfiles(v []string) {
	o.SwitchProfiles = v
}

// GetSwitches returns the Switches field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetSwitches() []string {
	if o == nil || IsNil(o.Switches) {
		var ret []string
		return ret
	}
	return o.Switches
}

// GetSwitchesOk returns a tuple with the Switches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetSwitchesOk() ([]string, bool) {
	if o == nil || IsNil(o.Switches) {
		return nil, false
	}
	return o.Switches, true
}

// HasSwitches returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) HasSwitches() bool {
	if o != nil && !IsNil(o.Switches) {
		return true
	}

	return false
}

// SetSwitches gets a reference to the given []string and assigns it to the Switches field.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) SetSwitches(v []string) {
	o.Switches = v
}

// GetStacks returns the Stacks field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetStacks() []string {
	if o == nil || IsNil(o.Stacks) {
		var ret []string
		return ret
	}
	return o.Stacks
}

// GetStacksOk returns a tuple with the Stacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetStacksOk() ([]string, bool) {
	if o == nil || IsNil(o.Stacks) {
		return nil, false
	}
	return o.Stacks, true
}

// HasStacks returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) HasStacks() bool {
	if o != nil && !IsNil(o.Stacks) {
		return true
	}

	return false
}

// SetStacks gets a reference to the given []string and assigns it to the Stacks field.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) SetStacks(v []string) {
	o.Stacks = v
}

// GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field value
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetIgmpSnoopingEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgmpSnoopingEnabled
}

// GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetIgmpSnoopingEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgmpSnoopingEnabled, true
}

// SetIgmpSnoopingEnabled sets field value
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) SetIgmpSnoopingEnabled(v bool) {
	o.IgmpSnoopingEnabled = v
}

// GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field value
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetFloodUnknownMulticastTrafficEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FloodUnknownMulticastTrafficEnabled
}

// GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FloodUnknownMulticastTrafficEnabled, true
}

// SetFloodUnknownMulticastTrafficEnabled sets field value
func (o *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) SetFloodUnknownMulticastTrafficEnabled(v bool) {
	o.FloodUnknownMulticastTrafficEnabled = v
}

func (o UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SwitchProfiles) {
		toSerialize["switchProfiles"] = o.SwitchProfiles
	}
	if !IsNil(o.Switches) {
		toSerialize["switches"] = o.Switches
	}
	if !IsNil(o.Stacks) {
		toSerialize["stacks"] = o.Stacks
	}
	toSerialize["igmpSnoopingEnabled"] = o.IgmpSnoopingEnabled
	toSerialize["floodUnknownMulticastTrafficEnabled"] = o.FloodUnknownMulticastTrafficEnabled
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchRoutingMulticastRequestOverridesInner struct {
	value *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner
	isSet bool
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRequestOverridesInner) Get() *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner {
	return v.value
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRequestOverridesInner) Set(val *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRequestOverridesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRequestOverridesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchRoutingMulticastRequestOverridesInner(val *UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) *NullableUpdateNetworkSwitchRoutingMulticastRequestOverridesInner {
	return &NullableUpdateNetworkSwitchRoutingMulticastRequestOverridesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRequestOverridesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRequestOverridesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


