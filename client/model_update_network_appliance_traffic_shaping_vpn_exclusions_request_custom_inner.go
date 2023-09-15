/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 06 September, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner{}

// UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner struct for UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner
type UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner struct {
	// Protocol.
	Protocol string `json:"protocol"`
	// Destination address; hostname required for DNS, IPv4 otherwise.
	Destination *string `json:"destination,omitempty"`
	// Destination port.
	Port *string `json:"port,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner(protocol string) *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner{}
	this.Protocol = protocol
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInnerWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInnerWithDefaults() *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner{}
	return &this
}

// GetProtocol returns the Protocol field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) SetProtocol(v string) {
	o.Protocol = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetDestination() string {
	if o == nil || IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) SetDestination(v string) {
	o.Destination = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetPort() string {
	if o == nil || IsNil(o.Port) {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) GetPortOk() (*string, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) SetPort(v string) {
	o.Port = &v
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["protocol"] = o.Protocol
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner struct {
	value *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) Get() *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) Set(val *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner(val *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner {
	return &NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

