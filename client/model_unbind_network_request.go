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

// checks if the UnbindNetworkRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnbindNetworkRequest{}

// UnbindNetworkRequest struct for UnbindNetworkRequest
type UnbindNetworkRequest struct {
	// Optional boolean to retain all the current configs given by the template.
	RetainConfigs *bool `json:"retainConfigs,omitempty"`
}

// NewUnbindNetworkRequest instantiates a new UnbindNetworkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnbindNetworkRequest() *UnbindNetworkRequest {
	this := UnbindNetworkRequest{}
	return &this
}

// NewUnbindNetworkRequestWithDefaults instantiates a new UnbindNetworkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnbindNetworkRequestWithDefaults() *UnbindNetworkRequest {
	this := UnbindNetworkRequest{}
	return &this
}

// GetRetainConfigs returns the RetainConfigs field value if set, zero value otherwise.
func (o *UnbindNetworkRequest) GetRetainConfigs() bool {
	if o == nil || IsNil(o.RetainConfigs) {
		var ret bool
		return ret
	}
	return *o.RetainConfigs
}

// GetRetainConfigsOk returns a tuple with the RetainConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnbindNetworkRequest) GetRetainConfigsOk() (*bool, bool) {
	if o == nil || IsNil(o.RetainConfigs) {
		return nil, false
	}
	return o.RetainConfigs, true
}

// HasRetainConfigs returns a boolean if a field has been set.
func (o *UnbindNetworkRequest) HasRetainConfigs() bool {
	if o != nil && !IsNil(o.RetainConfigs) {
		return true
	}

	return false
}

// SetRetainConfigs gets a reference to the given bool and assigns it to the RetainConfigs field.
func (o *UnbindNetworkRequest) SetRetainConfigs(v bool) {
	o.RetainConfigs = &v
}

func (o UnbindNetworkRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnbindNetworkRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RetainConfigs) {
		toSerialize["retainConfigs"] = o.RetainConfigs
	}
	return toSerialize, nil
}

type NullableUnbindNetworkRequest struct {
	value *UnbindNetworkRequest
	isSet bool
}

func (v NullableUnbindNetworkRequest) Get() *UnbindNetworkRequest {
	return v.value
}

func (v *NullableUnbindNetworkRequest) Set(val *UnbindNetworkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnbindNetworkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnbindNetworkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnbindNetworkRequest(val *UnbindNetworkRequest) *NullableUnbindNetworkRequest {
	return &NullableUnbindNetworkRequest{value: val, isSet: true}
}

func (v NullableUnbindNetworkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnbindNetworkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


