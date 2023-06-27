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

// UpdateNetworkSmDevicesFieldsRequestDeviceFields The new fields of the device. Each field of this object is optional.
type UpdateNetworkSmDevicesFieldsRequestDeviceFields struct {
	// New name for the device
	Name *string `json:"name,omitempty"`
	// New notes for the device
	Notes *string `json:"notes,omitempty"`
}

// NewUpdateNetworkSmDevicesFieldsRequestDeviceFields instantiates a new UpdateNetworkSmDevicesFieldsRequestDeviceFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSmDevicesFieldsRequestDeviceFields() *UpdateNetworkSmDevicesFieldsRequestDeviceFields {
	this := UpdateNetworkSmDevicesFieldsRequestDeviceFields{}
	return &this
}

// NewUpdateNetworkSmDevicesFieldsRequestDeviceFieldsWithDefaults instantiates a new UpdateNetworkSmDevicesFieldsRequestDeviceFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSmDevicesFieldsRequestDeviceFieldsWithDefaults() *UpdateNetworkSmDevicesFieldsRequestDeviceFields {
	this := UpdateNetworkSmDevicesFieldsRequestDeviceFields{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkSmDevicesFieldsRequestDeviceFields) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFieldsRequestDeviceFields) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkSmDevicesFieldsRequestDeviceFields) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkSmDevicesFieldsRequestDeviceFields) SetName(v string) {
	o.Name = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *UpdateNetworkSmDevicesFieldsRequestDeviceFields) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSmDevicesFieldsRequestDeviceFields) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
    return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *UpdateNetworkSmDevicesFieldsRequestDeviceFields) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *UpdateNetworkSmDevicesFieldsRequestDeviceFields) SetNotes(v string) {
	o.Notes = &v
}

func (o UpdateNetworkSmDevicesFieldsRequestDeviceFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkSmDevicesFieldsRequestDeviceFields struct {
	value *UpdateNetworkSmDevicesFieldsRequestDeviceFields
	isSet bool
}

func (v NullableUpdateNetworkSmDevicesFieldsRequestDeviceFields) Get() *UpdateNetworkSmDevicesFieldsRequestDeviceFields {
	return v.value
}

func (v *NullableUpdateNetworkSmDevicesFieldsRequestDeviceFields) Set(val *UpdateNetworkSmDevicesFieldsRequestDeviceFields) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSmDevicesFieldsRequestDeviceFields) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSmDevicesFieldsRequestDeviceFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSmDevicesFieldsRequestDeviceFields(val *UpdateNetworkSmDevicesFieldsRequestDeviceFields) *NullableUpdateNetworkSmDevicesFieldsRequestDeviceFields {
	return &NullableUpdateNetworkSmDevicesFieldsRequestDeviceFields{value: val, isSet: true}
}

func (v NullableUpdateNetworkSmDevicesFieldsRequestDeviceFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSmDevicesFieldsRequestDeviceFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

