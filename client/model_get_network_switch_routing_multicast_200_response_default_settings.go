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

// GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings Default multicast setting for entire network. IGMP snooping and Flood unknown       multicast traffic settings are enabled by default.
type GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings struct {
	// IGMP snooping enabled for the entire network
	IgmpSnoopingEnabled *bool `json:"igmpSnoopingEnabled,omitempty"`
	// Flood unknown multicast traffic enabled for the entire network
	FloodUnknownMulticastTrafficEnabled *bool `json:"floodUnknownMulticastTrafficEnabled,omitempty"`
}

// NewGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings instantiates a new GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings() *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings {
	this := GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings{}
	return &this
}

// NewGetNetworkSwitchRoutingMulticast200ResponseDefaultSettingsWithDefaults instantiates a new GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchRoutingMulticast200ResponseDefaultSettingsWithDefaults() *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings {
	this := GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings{}
	return &this
}

// GetIgmpSnoopingEnabled returns the IgmpSnoopingEnabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) GetIgmpSnoopingEnabled() bool {
	if o == nil || isNil(o.IgmpSnoopingEnabled) {
		var ret bool
		return ret
	}
	return *o.IgmpSnoopingEnabled
}

// GetIgmpSnoopingEnabledOk returns a tuple with the IgmpSnoopingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) GetIgmpSnoopingEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.IgmpSnoopingEnabled) {
    return nil, false
	}
	return o.IgmpSnoopingEnabled, true
}

// HasIgmpSnoopingEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) HasIgmpSnoopingEnabled() bool {
	if o != nil && !isNil(o.IgmpSnoopingEnabled) {
		return true
	}

	return false
}

// SetIgmpSnoopingEnabled gets a reference to the given bool and assigns it to the IgmpSnoopingEnabled field.
func (o *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) SetIgmpSnoopingEnabled(v bool) {
	o.IgmpSnoopingEnabled = &v
}

// GetFloodUnknownMulticastTrafficEnabled returns the FloodUnknownMulticastTrafficEnabled field value if set, zero value otherwise.
func (o *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) GetFloodUnknownMulticastTrafficEnabled() bool {
	if o == nil || isNil(o.FloodUnknownMulticastTrafficEnabled) {
		var ret bool
		return ret
	}
	return *o.FloodUnknownMulticastTrafficEnabled
}

// GetFloodUnknownMulticastTrafficEnabledOk returns a tuple with the FloodUnknownMulticastTrafficEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) GetFloodUnknownMulticastTrafficEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.FloodUnknownMulticastTrafficEnabled) {
    return nil, false
	}
	return o.FloodUnknownMulticastTrafficEnabled, true
}

// HasFloodUnknownMulticastTrafficEnabled returns a boolean if a field has been set.
func (o *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) HasFloodUnknownMulticastTrafficEnabled() bool {
	if o != nil && !isNil(o.FloodUnknownMulticastTrafficEnabled) {
		return true
	}

	return false
}

// SetFloodUnknownMulticastTrafficEnabled gets a reference to the given bool and assigns it to the FloodUnknownMulticastTrafficEnabled field.
func (o *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) SetFloodUnknownMulticastTrafficEnabled(v bool) {
	o.FloodUnknownMulticastTrafficEnabled = &v
}

func (o GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IgmpSnoopingEnabled) {
		toSerialize["igmpSnoopingEnabled"] = o.IgmpSnoopingEnabled
	}
	if !isNil(o.FloodUnknownMulticastTrafficEnabled) {
		toSerialize["floodUnknownMulticastTrafficEnabled"] = o.FloodUnknownMulticastTrafficEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings struct {
	value *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings
	isSet bool
}

func (v NullableGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) Get() *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings {
	return v.value
}

func (v *NullableGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) Set(val *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings(val *GetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) *NullableGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings {
	return &NullableGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchRoutingMulticast200ResponseDefaultSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


