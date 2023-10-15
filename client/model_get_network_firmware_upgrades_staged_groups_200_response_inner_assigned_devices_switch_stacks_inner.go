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

// checks if the GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner{}

// GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner struct for GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner
type GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner struct {
	// ID of the Switch Stack
	Id *string `json:"id,omitempty"`
	// Name of the Switch Stack
	Name *string `json:"name,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner instantiates a new GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner {
	this := GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInnerWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInnerWithDefaults() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner {
	this := GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) SetName(v string) {
	o.Name = &v
}

func (o GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner struct {
	value *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) Get() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) Set(val *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner(val *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner {
	return &NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesSwitchStacksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


