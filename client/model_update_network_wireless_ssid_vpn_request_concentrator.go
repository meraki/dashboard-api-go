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

// checks if the UpdateNetworkWirelessSsidVpnRequestConcentrator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidVpnRequestConcentrator{}

// UpdateNetworkWirelessSsidVpnRequestConcentrator The VPN concentrator settings for this SSID.
type UpdateNetworkWirelessSsidVpnRequestConcentrator struct {
	// The NAT ID of the concentrator that should be set.
	NetworkId *string `json:"networkId,omitempty"`
	// The VLAN that should be tagged for the concentrator.
	VlanId *int32 `json:"vlanId,omitempty"`
}

// NewUpdateNetworkWirelessSsidVpnRequestConcentrator instantiates a new UpdateNetworkWirelessSsidVpnRequestConcentrator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidVpnRequestConcentrator() *UpdateNetworkWirelessSsidVpnRequestConcentrator {
	this := UpdateNetworkWirelessSsidVpnRequestConcentrator{}
	return &this
}

// NewUpdateNetworkWirelessSsidVpnRequestConcentratorWithDefaults instantiates a new UpdateNetworkWirelessSsidVpnRequestConcentrator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidVpnRequestConcentratorWithDefaults() *UpdateNetworkWirelessSsidVpnRequestConcentrator {
	this := UpdateNetworkWirelessSsidVpnRequestConcentrator{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequestConcentrator) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestConcentrator) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestConcentrator) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *UpdateNetworkWirelessSsidVpnRequestConcentrator) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidVpnRequestConcentrator) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestConcentrator) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidVpnRequestConcentrator) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *UpdateNetworkWirelessSsidVpnRequestConcentrator) SetVlanId(v int32) {
	o.VlanId = &v
}

func (o UpdateNetworkWirelessSsidVpnRequestConcentrator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidVpnRequestConcentrator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidVpnRequestConcentrator struct {
	value *UpdateNetworkWirelessSsidVpnRequestConcentrator
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestConcentrator) Get() *UpdateNetworkWirelessSsidVpnRequestConcentrator {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestConcentrator) Set(val *UpdateNetworkWirelessSsidVpnRequestConcentrator) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestConcentrator) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestConcentrator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidVpnRequestConcentrator(val *UpdateNetworkWirelessSsidVpnRequestConcentrator) *NullableUpdateNetworkWirelessSsidVpnRequestConcentrator {
	return &NullableUpdateNetworkWirelessSsidVpnRequestConcentrator{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidVpnRequestConcentrator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidVpnRequestConcentrator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


