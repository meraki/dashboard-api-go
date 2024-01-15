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

// checks if the UpdateNetworkWirelessSsidRequestGreConcentrator type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidRequestGreConcentrator{}

// UpdateNetworkWirelessSsidRequestGreConcentrator The EoGRE concentrator's settings
type UpdateNetworkWirelessSsidRequestGreConcentrator struct {
	// The EoGRE concentrator's IP or FQDN. This param is required when ipAssignmentMode is 'Ethernet over GRE'.
	Host string `json:"host"`
}

// NewUpdateNetworkWirelessSsidRequestGreConcentrator instantiates a new UpdateNetworkWirelessSsidRequestGreConcentrator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidRequestGreConcentrator(host string) *UpdateNetworkWirelessSsidRequestGreConcentrator {
	this := UpdateNetworkWirelessSsidRequestGreConcentrator{}
	this.Host = host
	return &this
}

// NewUpdateNetworkWirelessSsidRequestGreConcentratorWithDefaults instantiates a new UpdateNetworkWirelessSsidRequestGreConcentrator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidRequestGreConcentratorWithDefaults() *UpdateNetworkWirelessSsidRequestGreConcentrator {
	this := UpdateNetworkWirelessSsidRequestGreConcentrator{}
	return &this
}

// GetHost returns the Host field value
func (o *UpdateNetworkWirelessSsidRequestGreConcentrator) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidRequestGreConcentrator) GetHostOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *UpdateNetworkWirelessSsidRequestGreConcentrator) SetHost(v string) {
	o.Host = v
}

func (o UpdateNetworkWirelessSsidRequestGreConcentrator) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidRequestGreConcentrator) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["host"] = o.Host
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidRequestGreConcentrator struct {
	value *UpdateNetworkWirelessSsidRequestGreConcentrator
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidRequestGreConcentrator) Get() *UpdateNetworkWirelessSsidRequestGreConcentrator {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidRequestGreConcentrator) Set(val *UpdateNetworkWirelessSsidRequestGreConcentrator) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidRequestGreConcentrator) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidRequestGreConcentrator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidRequestGreConcentrator(val *UpdateNetworkWirelessSsidRequestGreConcentrator) *NullableUpdateNetworkWirelessSsidRequestGreConcentrator {
	return &NullableUpdateNetworkWirelessSsidRequestGreConcentrator{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidRequestGreConcentrator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidRequestGreConcentrator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


