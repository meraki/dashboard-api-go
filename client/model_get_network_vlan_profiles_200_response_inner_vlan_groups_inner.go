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

// checks if the GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner{}

// GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner struct for GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner
type GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner struct {
	// Name of the VLAN, string length must be from 1 to 32 characters
	Name *string `json:"name,omitempty"`
	// Comma-separated VLAN IDs or ID ranges
	VlanIds *string `json:"vlanIds,omitempty"`
}

// NewGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner instantiates a new GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner() *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner {
	this := GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner{}
	return &this
}

// NewGetNetworkVlanProfiles200ResponseInnerVlanGroupsInnerWithDefaults instantiates a new GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkVlanProfiles200ResponseInnerVlanGroupsInnerWithDefaults() *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner {
	this := GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) SetName(v string) {
	o.Name = &v
}

// GetVlanIds returns the VlanIds field value if set, zero value otherwise.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) GetVlanIds() string {
	if o == nil || IsNil(o.VlanIds) {
		var ret string
		return ret
	}
	return *o.VlanIds
}

// GetVlanIdsOk returns a tuple with the VlanIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) GetVlanIdsOk() (*string, bool) {
	if o == nil || IsNil(o.VlanIds) {
		return nil, false
	}
	return o.VlanIds, true
}

// HasVlanIds returns a boolean if a field has been set.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) HasVlanIds() bool {
	if o != nil && !IsNil(o.VlanIds) {
		return true
	}

	return false
}

// SetVlanIds gets a reference to the given string and assigns it to the VlanIds field.
func (o *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) SetVlanIds(v string) {
	o.VlanIds = &v
}

func (o GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.VlanIds) {
		toSerialize["vlanIds"] = o.VlanIds
	}
	return toSerialize, nil
}

type NullableGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner struct {
	value *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner
	isSet bool
}

func (v NullableGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) Get() *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner {
	return v.value
}

func (v *NullableGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) Set(val *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner(val *GetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) *NullableGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner {
	return &NullableGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner{value: val, isSet: true}
}

func (v NullableGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkVlanProfiles200ResponseInnerVlanGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


