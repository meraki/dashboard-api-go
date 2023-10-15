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

// checks if the GetNetworkVlanProfiles200ResponseInnerVlanNamesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkVlanProfiles200ResponseInnerVlanNamesInner{}

// GetNetworkVlanProfiles200ResponseInnerVlanNamesInner struct for GetNetworkVlanProfiles200ResponseInnerVlanNamesInner
type GetNetworkVlanProfiles200ResponseInnerVlanNamesInner struct {
	// Name of the VLAN, string length must be from 1 to 32 characters
	Name *string `json:"name,omitempty"`
	// VLAN ID
	VlanId *string `json:"vlanId,omitempty"`
}

// NewGetNetworkVlanProfiles200ResponseInnerVlanNamesInner instantiates a new GetNetworkVlanProfiles200ResponseInnerVlanNamesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkVlanProfiles200ResponseInnerVlanNamesInner() *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner {
	this := GetNetworkVlanProfiles200ResponseInnerVlanNamesInner{}
	return &this
}

// NewGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerWithDefaults instantiates a new GetNetworkVlanProfiles200ResponseInnerVlanNamesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkVlanProfiles200ResponseInnerVlanNamesInnerWithDefaults() *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner {
	this := GetNetworkVlanProfiles200ResponseInnerVlanNamesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) SetName(v string) {
	o.Name = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) GetVlanId() string {
	if o == nil || IsNil(o.VlanId) {
		var ret string
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) GetVlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given string and assigns it to the VlanId field.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) SetVlanId(v string) {
	o.VlanId = &v
}

func (o GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlanId"] = o.VlanId
	}
	return toSerialize, nil
}

type NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInner struct {
	value *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner
	isSet bool
}

func (v NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInner) Get() *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner {
	return v.value
}

func (v *NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInner) Set(val *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInner(val *GetNetworkVlanProfiles200ResponseInnerVlanNamesInner) *NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInner {
	return &NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInner{value: val, isSet: true}
}

func (v NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkVlanProfiles200ResponseInnerVlanNamesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


