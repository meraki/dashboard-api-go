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

// checks if the UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner{}

// UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner struct for UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner
type UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner struct {
	// The serial of the related device
	Serial string `json:"serial"`
}

// NewUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner instantiates a new UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner(serial string) *UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner {
	this := UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner{}
	this.Serial = serial
	return &this
}

// NewUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInnerWithDefaults instantiates a new UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInnerWithDefaults() *UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner {
	this := UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner{}
	return &this
}

// GetSerial returns the Serial field value
func (o *UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) GetSerialOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) SetSerial(v string) {
	o.Serial = v
}

func (o UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serial"] = o.Serial
	return toSerialize, nil
}

type NullableUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner struct {
	value *UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner
	isSet bool
}

func (v NullableUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) Get() *UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner {
	return v.value
}

func (v *NullableUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) Set(val *UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner(val *UpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) *NullableUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner {
	return &NullableUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner{value: val, isSet: true}
}

func (v NullableUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceSensorRelationshipsRequestLivestreamRelatedDevicesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


