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

// checks if the GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource{}

// GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource Source of the packet.
type GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource struct {
	// Source mac address of the packet.
	Mac *string `json:"mac,omitempty"`
	Ipv4 *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSourceIpv4 `json:"ipv4,omitempty"`
	// Source port of the packet.
	Port *int32 `json:"port,omitempty"`
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource{}
	return &this
}

// NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSourceWithDefaults instantiates a new GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSourceWithDefaults() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource {
	this := GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource{}
	return &this
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) SetMac(v string) {
	o.Mac = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) GetIpv4() GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSourceIpv4 {
	if o == nil || IsNil(o.Ipv4) {
		var ret GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSourceIpv4
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) GetIpv4Ok() (*GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSourceIpv4, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSourceIpv4 and assigns it to the Ipv4 field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) SetIpv4(v GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSourceIpv4) {
	o.Ipv4 = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) SetPort(v int32) {
	o.Port = &v
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource struct {
	value *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource
	isSet bool
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) Get() *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource {
	return v.value
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) Set(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource(val *GetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource {
	return &NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchDhcpV4ServersSeen200ResponseInnerLastPacketSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


