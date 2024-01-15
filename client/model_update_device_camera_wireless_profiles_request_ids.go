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

// checks if the UpdateDeviceCameraWirelessProfilesRequestIds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCameraWirelessProfilesRequestIds{}

// UpdateDeviceCameraWirelessProfilesRequestIds The ids of the wireless profile to assign to the given camera
type UpdateDeviceCameraWirelessProfilesRequestIds struct {
	// The id of the primary wireless profile
	Primary *string `json:"primary,omitempty"`
	// The id of the secondary wireless profile
	Secondary *string `json:"secondary,omitempty"`
	// The id of the backup wireless profile
	Backup *string `json:"backup,omitempty"`
}

// NewUpdateDeviceCameraWirelessProfilesRequestIds instantiates a new UpdateDeviceCameraWirelessProfilesRequestIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCameraWirelessProfilesRequestIds() *UpdateDeviceCameraWirelessProfilesRequestIds {
	this := UpdateDeviceCameraWirelessProfilesRequestIds{}
	return &this
}

// NewUpdateDeviceCameraWirelessProfilesRequestIdsWithDefaults instantiates a new UpdateDeviceCameraWirelessProfilesRequestIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCameraWirelessProfilesRequestIdsWithDefaults() *UpdateDeviceCameraWirelessProfilesRequestIds {
	this := UpdateDeviceCameraWirelessProfilesRequestIds{}
	return &this
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) GetPrimary() string {
	if o == nil || IsNil(o.Primary) {
		var ret string
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) GetPrimaryOk() (*string, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given string and assigns it to the Primary field.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) SetPrimary(v string) {
	o.Primary = &v
}

// GetSecondary returns the Secondary field value if set, zero value otherwise.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) GetSecondary() string {
	if o == nil || IsNil(o.Secondary) {
		var ret string
		return ret
	}
	return *o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) GetSecondaryOk() (*string, bool) {
	if o == nil || IsNil(o.Secondary) {
		return nil, false
	}
	return o.Secondary, true
}

// HasSecondary returns a boolean if a field has been set.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) HasSecondary() bool {
	if o != nil && !IsNil(o.Secondary) {
		return true
	}

	return false
}

// SetSecondary gets a reference to the given string and assigns it to the Secondary field.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) SetSecondary(v string) {
	o.Secondary = &v
}

// GetBackup returns the Backup field value if set, zero value otherwise.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) GetBackup() string {
	if o == nil || IsNil(o.Backup) {
		var ret string
		return ret
	}
	return *o.Backup
}

// GetBackupOk returns a tuple with the Backup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) GetBackupOk() (*string, bool) {
	if o == nil || IsNil(o.Backup) {
		return nil, false
	}
	return o.Backup, true
}

// HasBackup returns a boolean if a field has been set.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) HasBackup() bool {
	if o != nil && !IsNil(o.Backup) {
		return true
	}

	return false
}

// SetBackup gets a reference to the given string and assigns it to the Backup field.
func (o *UpdateDeviceCameraWirelessProfilesRequestIds) SetBackup(v string) {
	o.Backup = &v
}

func (o UpdateDeviceCameraWirelessProfilesRequestIds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCameraWirelessProfilesRequestIds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.Secondary) {
		toSerialize["secondary"] = o.Secondary
	}
	if !IsNil(o.Backup) {
		toSerialize["backup"] = o.Backup
	}
	return toSerialize, nil
}

type NullableUpdateDeviceCameraWirelessProfilesRequestIds struct {
	value *UpdateDeviceCameraWirelessProfilesRequestIds
	isSet bool
}

func (v NullableUpdateDeviceCameraWirelessProfilesRequestIds) Get() *UpdateDeviceCameraWirelessProfilesRequestIds {
	return v.value
}

func (v *NullableUpdateDeviceCameraWirelessProfilesRequestIds) Set(val *UpdateDeviceCameraWirelessProfilesRequestIds) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCameraWirelessProfilesRequestIds) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCameraWirelessProfilesRequestIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCameraWirelessProfilesRequestIds(val *UpdateDeviceCameraWirelessProfilesRequestIds) *NullableUpdateDeviceCameraWirelessProfilesRequestIds {
	return &NullableUpdateDeviceCameraWirelessProfilesRequestIds{value: val, isSet: true}
}

func (v NullableUpdateDeviceCameraWirelessProfilesRequestIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCameraWirelessProfilesRequestIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


