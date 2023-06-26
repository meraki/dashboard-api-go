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

// GetNetworkWirelessRfProfiles200ResponseTransmission Settings related to radio transmission.
type GetNetworkWirelessRfProfiles200ResponseTransmission struct {
	// Toggle for radio transmission. When false, radios will not transmit at all.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewGetNetworkWirelessRfProfiles200ResponseTransmission instantiates a new GetNetworkWirelessRfProfiles200ResponseTransmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkWirelessRfProfiles200ResponseTransmission() *GetNetworkWirelessRfProfiles200ResponseTransmission {
	this := GetNetworkWirelessRfProfiles200ResponseTransmission{}
	return &this
}

// NewGetNetworkWirelessRfProfiles200ResponseTransmissionWithDefaults instantiates a new GetNetworkWirelessRfProfiles200ResponseTransmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkWirelessRfProfiles200ResponseTransmissionWithDefaults() *GetNetworkWirelessRfProfiles200ResponseTransmission {
	this := GetNetworkWirelessRfProfiles200ResponseTransmission{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetNetworkWirelessRfProfiles200ResponseTransmission) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseTransmission) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetNetworkWirelessRfProfiles200ResponseTransmission) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetNetworkWirelessRfProfiles200ResponseTransmission) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o GetNetworkWirelessRfProfiles200ResponseTransmission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkWirelessRfProfiles200ResponseTransmission struct {
	value *GetNetworkWirelessRfProfiles200ResponseTransmission
	isSet bool
}

func (v NullableGetNetworkWirelessRfProfiles200ResponseTransmission) Get() *GetNetworkWirelessRfProfiles200ResponseTransmission {
	return v.value
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponseTransmission) Set(val *GetNetworkWirelessRfProfiles200ResponseTransmission) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkWirelessRfProfiles200ResponseTransmission) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponseTransmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkWirelessRfProfiles200ResponseTransmission(val *GetNetworkWirelessRfProfiles200ResponseTransmission) *NullableGetNetworkWirelessRfProfiles200ResponseTransmission {
	return &NullableGetNetworkWirelessRfProfiles200ResponseTransmission{value: val, isSet: true}
}

func (v NullableGetNetworkWirelessRfProfiles200ResponseTransmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkWirelessRfProfiles200ResponseTransmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


