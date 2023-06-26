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

// GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner struct for GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner
type GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner struct {
	// The name of the client which has fixed IP address
	Name *string `json:"name,omitempty"`
	// The MAC address of the client which has fixed IP address
	Mac *string `json:"mac,omitempty"`
	// The IP address of the client which has fixed IP address assigned to it
	Ip *string `json:"ip,omitempty"`
}

// NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner instantiates a new GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner() *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner {
	this := GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner{}
	return &this
}

// NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInnerWithDefaults instantiates a new GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInnerWithDefaults() *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner {
	this := GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) SetName(v string) {
	o.Name = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) GetMac() string {
	if o == nil || isNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) GetMacOk() (*string, bool) {
	if o == nil || isNil(o.Mac) {
    return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) HasMac() bool {
	if o != nil && !isNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) SetMac(v string) {
	o.Mac = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) SetIp(v string) {
	o.Ip = &v
}

func (o GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	return json.Marshal(toSerialize)
}

type NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner struct {
	value *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner
	isSet bool
}

func (v NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) Get() *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner {
	return v.value
}

func (v *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) Set(val *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner(val *GetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner {
	return &NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchStackRoutingInterfaceDhcp200ResponseFixedIpAssignmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


