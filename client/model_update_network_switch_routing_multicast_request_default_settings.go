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

// checks if the UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings{}

// UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings Default multicast setting for entire network. IGMP snooping and Flood unknown multicast traffic settings are enabled by default.
type UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings struct {
	// IGMP snooping setting for entire network
	IgmpSnoopingEnabled *bool `json:"igmpSnoopingEnabled,omitempty"`
	// Flood unknown multicast traffic setting for entire network
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"`
}

// NewUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings instantiates a new UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings() *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings {
	this := UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings{}
	return &this
}

// NewUpdateNetworkSwitchRoutingMulticastRequestDefaultSettingsWithDefaults instantiates a new UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchRoutingMulticastRequestDefaultSettingsWithDefaults() *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings {
	this := UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings{}
	return &this
}

// GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) GetIgmpSnoopingEnabled() bool {
	if o == nil || IsNil(o.IgmpSnoopingEnabled) {
		var ret bool
		return ret
	}
	return *o.IgmpSnoopingEnabled
}

// GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) GetIgmpSnoopingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IgmpSnoopingEnabled) {
		return nil, false
	}
	return o.IgmpSnoopingEnabled, true
}

// HasIgmpSnoopingEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) HasIgmpSnoopingEnabled() bool {
	if o != nil && !IsNil(o.IgmpSnoopingEnabled) {
		return true
	}

	return false
}

// SetIgmpSnoopingEnabled gets a reference to the given bool and assigns it to the IgmpSnoopingEnabled field.
func (o *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) SetIgmpSnoopingEnabled(v bool) {
	o.IgmpSnoopingEnabled = &v
}

// GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) GetFloodUnknownMulticastTrafficEnabled() bool {
	if o == nil || IsNil(o.FloodUnknownMulticastTrafficEnabled) {
		var ret bool
		return ret
	}
	return *o.FloodUnknownMulticastTrafficEnabled
}

// GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FloodUnknownMulticastTrafficEnabled) {
		return nil, false
	}
	return o.FloodUnknownMulticastTrafficEnabled, true
}

// HasFloodUnknownMulticastTrafficEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) HasFloodUnknownMulticastTrafficEnabled() bool {
	if o != nil && !IsNil(o.FloodUnknownMulticastTrafficEnabled) {
		return true
	}

	return false
}

// SetFloodUnknownMulticastTrafficEnabled gets a reference to the given bool and assigns it to the FloodUnknownMulticastTrafficEnabled field.
func (o *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) SetFloodUnknownMulticastTrafficEnabled(v bool) {
	o.FloodUnknownMulticastTrafficEnabled = &v
}

func (o UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IgmpSnoopingEnabled) {
		toSerialize["igmpSnoopingEnabled"] = o.IgmpSnoopingEnabled
	}
	if !IsNil(o.FloodUnknownMulticastTrafficEnabled) {
		toSerialize["floodUnknownMulticastTrafficEnabled"] = o.FloodUnknownMulticastTrafficEnabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings struct {
	value *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings
	isSet bool
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) Get() *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings {
	return v.value
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) Set(val *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings(val *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) *NullableUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings {
	return &NullableUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


