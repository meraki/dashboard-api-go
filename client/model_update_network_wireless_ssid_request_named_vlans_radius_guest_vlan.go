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

// checks if the UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan{}

// UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan Guest VLAN settings. Used to direct traffic to a guest VLAN when none of the RADIUS servers are reachable or a client receives access-reject from the RADIUS server.
type UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan struct {
	// Whether or not RADIUS guest named VLAN is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// RADIUS guest VLAN name.
	Name *string `json:"name,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan instantiates a new UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan() *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan {
	this := UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlanWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlanWithDefaults() *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan {
	this := UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) SetName(v string) {
	o.Name = &v
}

func (o UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan struct {
	value *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) Get() *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) Set(val *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan(val *UpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) *NullableUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan {
	return &NullableUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestNamedVlansRadiusGuestVlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


