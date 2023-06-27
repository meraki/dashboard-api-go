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

// GetDeviceWirelessBluetoothSettings200Response struct for GetDeviceWirelessBluetoothSettings200Response
type GetDeviceWirelessBluetoothSettings200Response struct {
	// Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Uuid *string `json:"uuid,omitempty"`
	// Desired major value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Major *int32 `json:"major,omitempty"`
	// Desired minor value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Minor *int32 `json:"minor,omitempty"`
}

// NewGetDeviceWirelessBluetoothSettings200Response instantiates a new GetDeviceWirelessBluetoothSettings200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDeviceWirelessBluetoothSettings200Response() *GetDeviceWirelessBluetoothSettings200Response {
	this := GetDeviceWirelessBluetoothSettings200Response{}
	return &this
}

// NewGetDeviceWirelessBluetoothSettings200ResponseWithDefaults instantiates a new GetDeviceWirelessBluetoothSettings200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDeviceWirelessBluetoothSettings200ResponseWithDefaults() *GetDeviceWirelessBluetoothSettings200Response {
	this := GetDeviceWirelessBluetoothSettings200Response{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *GetDeviceWirelessBluetoothSettings200Response) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessBluetoothSettings200Response) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
    return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *GetDeviceWirelessBluetoothSettings200Response) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *GetDeviceWirelessBluetoothSettings200Response) SetUuid(v string) {
	o.Uuid = &v
}

// GetMajor returns the Major field value if set, zero value otherwise.
func (o *GetDeviceWirelessBluetoothSettings200Response) GetMajor() int32 {
	if o == nil || isNil(o.Major) {
		var ret int32
		return ret
	}
	return *o.Major
}

// GetMajorOk returns a tuple with the Major field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessBluetoothSettings200Response) GetMajorOk() (*int32, bool) {
	if o == nil || isNil(o.Major) {
    return nil, false
	}
	return o.Major, true
}

// HasMajor returns a boolean if a field has been set.
func (o *GetDeviceWirelessBluetoothSettings200Response) HasMajor() bool {
	if o != nil && !isNil(o.Major) {
		return true
	}

	return false
}

// SetMajor gets a reference to the given int32 and assigns it to the Major field.
func (o *GetDeviceWirelessBluetoothSettings200Response) SetMajor(v int32) {
	o.Major = &v
}

// GetMinor returns the Minor field value if set, zero value otherwise.
func (o *GetDeviceWirelessBluetoothSettings200Response) GetMinor() int32 {
	if o == nil || isNil(o.Minor) {
		var ret int32
		return ret
	}
	return *o.Minor
}

// GetMinorOk returns a tuple with the Minor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDeviceWirelessBluetoothSettings200Response) GetMinorOk() (*int32, bool) {
	if o == nil || isNil(o.Minor) {
    return nil, false
	}
	return o.Minor, true
}

// HasMinor returns a boolean if a field has been set.
func (o *GetDeviceWirelessBluetoothSettings200Response) HasMinor() bool {
	if o != nil && !isNil(o.Minor) {
		return true
	}

	return false
}

// SetMinor gets a reference to the given int32 and assigns it to the Minor field.
func (o *GetDeviceWirelessBluetoothSettings200Response) SetMinor(v int32) {
	o.Minor = &v
}

func (o GetDeviceWirelessBluetoothSettings200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !isNil(o.Major) {
		toSerialize["major"] = o.Major
	}
	if !isNil(o.Minor) {
		toSerialize["minor"] = o.Minor
	}
	return json.Marshal(toSerialize)
}

type NullableGetDeviceWirelessBluetoothSettings200Response struct {
	value *GetDeviceWirelessBluetoothSettings200Response
	isSet bool
}

func (v NullableGetDeviceWirelessBluetoothSettings200Response) Get() *GetDeviceWirelessBluetoothSettings200Response {
	return v.value
}

func (v *NullableGetDeviceWirelessBluetoothSettings200Response) Set(val *GetDeviceWirelessBluetoothSettings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDeviceWirelessBluetoothSettings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDeviceWirelessBluetoothSettings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDeviceWirelessBluetoothSettings200Response(val *GetDeviceWirelessBluetoothSettings200Response) *NullableGetDeviceWirelessBluetoothSettings200Response {
	return &NullableGetDeviceWirelessBluetoothSettings200Response{value: val, isSet: true}
}

func (v NullableGetDeviceWirelessBluetoothSettings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDeviceWirelessBluetoothSettings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


