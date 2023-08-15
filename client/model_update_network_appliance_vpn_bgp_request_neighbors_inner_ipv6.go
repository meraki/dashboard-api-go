/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6{}

// UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 Information regarding IPv6 address of the neighbor, Required if `ip` is not present.
type UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 struct {
	// The IPv6 address of the neighbor.
	Address string `json:"address"`
}

// NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 instantiates a new UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6(address string) *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 {
	this := UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6{}
	this.Address = address
	return &this
}

// NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6WithDefaults instantiates a new UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6WithDefaults() *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 {
	this := UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6{}
	return &this
}

// GetAddress returns the Address field value
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) SetAddress(v string) {
	o.Address = v
}

func (o UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 struct {
	value *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6
	isSet bool
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) Get() *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 {
	return v.value
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) Set(val *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6(val *UpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6 {
	return &NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceVpnBgpRequestNeighborsInnerIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


