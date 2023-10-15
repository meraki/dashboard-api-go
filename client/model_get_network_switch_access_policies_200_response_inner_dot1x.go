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

// checks if the GetNetworkSwitchAccessPolicies200ResponseInnerDot1x type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkSwitchAccessPolicies200ResponseInnerDot1x{}

// GetNetworkSwitchAccessPolicies200ResponseInnerDot1x 802.1x Settings
type GetNetworkSwitchAccessPolicies200ResponseInnerDot1x struct {
	// Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
	ControlDirection *string `json:"controlDirection,omitempty"`
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerDot1x instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerDot1x object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkSwitchAccessPolicies200ResponseInnerDot1x() *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x {
	this := GetNetworkSwitchAccessPolicies200ResponseInnerDot1x{}
	return &this
}

// NewGetNetworkSwitchAccessPolicies200ResponseInnerDot1xWithDefaults instantiates a new GetNetworkSwitchAccessPolicies200ResponseInnerDot1x object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkSwitchAccessPolicies200ResponseInnerDot1xWithDefaults() *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x {
	this := GetNetworkSwitchAccessPolicies200ResponseInnerDot1x{}
	return &this
}

// GetControlDirection returns the ControlDirection field value if set, zero value otherwise.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) GetControlDirection() string {
	if o == nil || IsNil(o.ControlDirection) {
		var ret string
		return ret
	}
	return *o.ControlDirection
}

// GetControlDirectionOk returns a tuple with the ControlDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) GetControlDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.ControlDirection) {
		return nil, false
	}
	return o.ControlDirection, true
}

// HasControlDirection returns a boolean if a field has been set.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) HasControlDirection() bool {
	if o != nil && !IsNil(o.ControlDirection) {
		return true
	}

	return false
}

// SetControlDirection gets a reference to the given string and assigns it to the ControlDirection field.
func (o *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) SetControlDirection(v string) {
	o.ControlDirection = &v
}

func (o GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ControlDirection) {
		toSerialize["controlDirection"] = o.ControlDirection
	}
	return toSerialize, nil
}

type NullableGetNetworkSwitchAccessPolicies200ResponseInnerDot1x struct {
	value *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x
	isSet bool
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerDot1x) Get() *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x {
	return v.value
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerDot1x) Set(val *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerDot1x) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerDot1x) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkSwitchAccessPolicies200ResponseInnerDot1x(val *GetNetworkSwitchAccessPolicies200ResponseInnerDot1x) *NullableGetNetworkSwitchAccessPolicies200ResponseInnerDot1x {
	return &NullableGetNetworkSwitchAccessPolicies200ResponseInnerDot1x{value: val, isSet: true}
}

func (v NullableGetNetworkSwitchAccessPolicies200ResponseInnerDot1x) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkSwitchAccessPolicies200ResponseInnerDot1x) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


