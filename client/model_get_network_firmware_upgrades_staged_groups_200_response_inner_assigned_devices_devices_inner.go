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

// checks if the GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner{}

// GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner struct for GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner
type GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner struct {
	// Serial of the device
	Serial *string `json:"serial,omitempty"`
	// Name of the device
	Name *string `json:"name,omitempty"`
}

// NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner instantiates a new GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner {
	this := GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner{}
	return &this
}

// NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInnerWithDefaults instantiates a new GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInnerWithDefaults() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner {
	this := GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner{}
	return &this
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) SetSerial(v string) {
	o.Serial = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) SetName(v string) {
	o.Name = &v
}

func (o GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner struct {
	value *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner
	isSet bool
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) Get() *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner {
	return v.value
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) Set(val *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner(val *GetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner {
	return &NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner{value: val, isSet: true}
}

func (v NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNetworkFirmwareUpgradesStagedGroups200ResponseInnerAssignedDevicesDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


