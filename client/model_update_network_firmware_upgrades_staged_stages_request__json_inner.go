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

// checks if the UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner{}

// UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner struct for UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner
type UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner struct {
	Group *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup `json:"group,omitempty"`
}

// NewUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner instantiates a new UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner() *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner {
	this := UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner{}
	return &this
}

// NewUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerWithDefaults instantiates a new UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerWithDefaults() *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner {
	this := UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) GetGroup() UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup {
	if o == nil || IsNil(o.Group) {
		var ret UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) GetGroupOk() (*UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup and assigns it to the Group field.
func (o *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) SetGroup(v UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInnerGroup) {
	o.Group = &v
}

func (o UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	return toSerialize, nil
}

type NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner struct {
	value *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner
	isSet bool
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) Get() *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner {
	return v.value
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) Set(val *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner(val *UpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner {
	return &NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkFirmwareUpgradesStagedStagesRequestJsonInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


