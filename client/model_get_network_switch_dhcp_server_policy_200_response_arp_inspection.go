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

// checks if the GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection{}

// GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection Dynamic ARP Inspection settings
type GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection struct {
	// Enable or disable Dynamic ARP Inspection on the network. Default value is false.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection instantiates a new GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection() *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection {
	this := GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection{}
	return &this
}

// NewGetNetworkSwitchDhcpServerPolicy200ResponseArpInspectionWithDefaults instantiates a new GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpServerPolicy200ResponseArpInspectionWithDefaults() *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection {
	this := GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection struct {
	value *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) Get() *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) Set(val *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection(val *GetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) *NullableGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection {
	return &NullableGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpServerPolicy200ResponseArpInspection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


