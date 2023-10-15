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

// checks if the UpdateNetworkWirelessSsidRequestGre type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestGre{}

// UpdateNetworkWirelessSsidRequestGre Ethernet over GRE settings
type UpdateNetworkWirelessSsidRequestGre struct {
	Concentrator *UpdateNetworkWirelessSsidRequestGreConcentrator `json:"concentrator,omitempty"`
	// Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel.
	Key *int32 `json:"key,omitempty"`
}

// NewUpdateNetworkWirelessSsidRequestGre instantiates a new UpdateNetworkWirelessSsidRequestGre object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestGre() *UpdateNetworkWirelessSsidRequestGre {
	this := UpdateNetworkWirelessSsidRequestGre{}
	return &this
}

// NewUpdateNetworkWirelessSsidRequestGreWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestGre object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestGreWithDefaults() *UpdateNetworkWirelessSsidRequestGre {
	this := UpdateNetworkWirelessSsidRequestGre{}
	return &this
}

// GetConcentrator returns the Concentrator field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestGre) GetConcentrator() UpdateNetworkWirelessSsidRequestGreConcentrator {
	if o == nil || IsNil(o.Concentrator) {
		var ret UpdateNetworkWirelessSsidRequestGreConcentrator
		return ret
	}
	return *o.Concentrator
}

// GetConcentratorOk returns a tuple with the Concentrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestGre) GetConcentratorOk() (*UpdateNetworkWirelessSsidRequestGreConcentrator, bool) {
	if o == nil || IsNil(o.Concentrator) {
		return nil, false
	}
	return o.Concentrator, true
}

// HasConcentrator returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestGre) HasConcentrator() bool {
	if o != nil && !IsNil(o.Concentrator) {
		return true
	}

	return false
}

// SetConcentrator gets a reference to the given UpdateNetworkWirelessSsidRequestGreConcentrator and assigns it to the Concentrator field.
func (o *UpdateNetworkWirelessSsidRequestGre) SetConcentrator(v UpdateNetworkWirelessSsidRequestGreConcentrator) {
	o.Concentrator = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidRequestGre) GetKey() int32 {
	if o == nil || IsNil(o.Key) {
		var ret int32
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestGre) GetKeyOk() (*int32, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidRequestGre) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given int32 and assigns it to the Key field.
func (o *UpdateNetworkWirelessSsidRequestGre) SetKey(v int32) {
	o.Key = &v
}

func (o UpdateNetworkWirelessSsidRequestGre) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestGre) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Concentrator) {
		toSerialize["concentrator"] = o.Concentrator
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestGre struct {
	value *UpdateNetworkWirelessSsidRequestGre
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestGre) Get() *UpdateNetworkWirelessSsidRequestGre {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestGre) Set(val *UpdateNetworkWirelessSsidRequestGre) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestGre) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestGre) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestGre(val *UpdateNetworkWirelessSsidRequestGre) *NullableUpdateNetworkWirelessSsidRequestGre {
	return &NullableUpdateNetworkWirelessSsidRequestGre{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestGre) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestGre) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


