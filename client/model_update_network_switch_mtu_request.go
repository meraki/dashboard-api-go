/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 02 August, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.36.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkSwitchMtuRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchMtuRequest{}

// UpdateNetworkSwitchMtuRequest struct for UpdateNetworkSwitchMtuRequest
type UpdateNetworkSwitchMtuRequest struct {
	// MTU size for the entire network. Default value is 9578.
	DefaultMtuSize *int32 `json:"defaultMtuSize,omitempty"`
	// Override MTU size for individual switches or switch templates. An empty array will clear overrides.
	Overrides []GetNetworkSwitchMtu200ResponseOverridesInner `json:"overrides,omitempty"`
}

// NewUpdateNetworkSwitchMtuRequest instantiates a new UpdateNetworkSwitchMtuRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchMtuRequest() *UpdateNetworkSwitchMtuRequest {
	this := UpdateNetworkSwitchMtuRequest{}
	return &this
}

// NewUpdateNetworkSwitchMtuRequestWithDefaults instantiates a new UpdateNetworkSwitchMtuRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchMtuRequestWithDefaults() *UpdateNetworkSwitchMtuRequest {
	this := UpdateNetworkSwitchMtuRequest{}
	return &this
}

// GetDefaultMtuSize returns the DefaultMtuSize field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchMtuRequest) GetDefaultMtuSize() int32 {
	if o == nil || IsNil(o.DefaultMtuSize) {
		var ret int32
		return ret
	}
	return *o.DefaultMtuSize
}

// GetDefaultMtuSizeOk returns a tuple with the DefaultMtuSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchMtuRequest) GetDefaultMtuSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultMtuSize) {
		return nil, false
	}
	return o.DefaultMtuSize, true
}

// HasDefaultMtuSize returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchMtuRequest) HasDefaultMtuSize() bool {
	if o != nil && !IsNil(o.DefaultMtuSize) {
		return true
	}

	return false
}

// SetDefaultMtuSize gets a reference to the given int32 and assigns it to the DefaultMtuSize field.
func (o *UpdateNetworkSwitchMtuRequest) SetDefaultMtuSize(v int32) {
	o.DefaultMtuSize = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchMtuRequest) GetOverrides() []GetNetworkSwitchMtu200ResponseOverridesInner {
	if o == nil || IsNil(o.Overrides) {
		var ret []GetNetworkSwitchMtu200ResponseOverridesInner
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchMtuRequest) GetOverridesOk() ([]GetNetworkSwitchMtu200ResponseOverridesInner, bool) {
	if o == nil || IsNil(o.Overrides) {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchMtuRequest) HasOverrides() bool {
	if o != nil && !IsNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []GetNetworkSwitchMtu200ResponseOverridesInner and assigns it to the Overrides field.
func (o *UpdateNetworkSwitchMtuRequest) SetOverrides(v []GetNetworkSwitchMtu200ResponseOverridesInner) {
	o.Overrides = v
}

func (o UpdateNetworkSwitchMtuRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchMtuRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultMtuSize) {
		toSerialize["defaultMtuSize"] = o.DefaultMtuSize
	}
	if !IsNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchMtuRequest struct {
	value *UpdateNetworkSwitchMtuRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchMtuRequest) Get() *UpdateNetworkSwitchMtuRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchMtuRequest) Set(val *UpdateNetworkSwitchMtuRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchMtuRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchMtuRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchMtuRequest(val *UpdateNetworkSwitchMtuRequest) *NullableUpdateNetworkSwitchMtuRequest {
	return &NullableUpdateNetworkSwitchMtuRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchMtuRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchMtuRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


