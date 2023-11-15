/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UnenrollNetworkSmDevice200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnenrollNetworkSmDevice200Response{}

// UnenrollNetworkSmDevice200Response struct for UnenrollNetworkSmDevice200Response
type UnenrollNetworkSmDevice200Response struct {
	// Boolean indicating whether the operation was completed successfully.
	Success *bool `json:"success,omitempty"`
}

// NewUnenrollNetworkSmDevice200Response instantiates a new UnenrollNetworkSmDevice200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnenrollNetworkSmDevice200Response() *UnenrollNetworkSmDevice200Response {
	this := UnenrollNetworkSmDevice200Response{}
	return &this
}

// NewUnenrollNetworkSmDevice200ResponseWithDefaults instantiates a new UnenrollNetworkSmDevice200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnenrollNetworkSmDevice200ResponseWithDefaults() *UnenrollNetworkSmDevice200Response {
	this := UnenrollNetworkSmDevice200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UnenrollNetworkSmDevice200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnenrollNetworkSmDevice200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UnenrollNetworkSmDevice200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UnenrollNetworkSmDevice200Response) SetSuccess(v bool) {
	o.Success = &v
}

func (o UnenrollNetworkSmDevice200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnenrollNetworkSmDevice200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableUnenrollNetworkSmDevice200Response struct {
	value *UnenrollNetworkSmDevice200Response
	isSet bool
}

func (v NullableUnenrollNetworkSmDevice200Response) Get() *UnenrollNetworkSmDevice200Response {
	return v.value
}

func (v *NullableUnenrollNetworkSmDevice200Response) Set(val *UnenrollNetworkSmDevice200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUnenrollNetworkSmDevice200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUnenrollNetworkSmDevice200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnenrollNetworkSmDevice200Response(val *UnenrollNetworkSmDevice200Response) *NullableUnenrollNetworkSmDevice200Response {
	return &NullableUnenrollNetworkSmDevice200Response{value: val, isSet: true}
}

func (v NullableUnenrollNetworkSmDevice200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnenrollNetworkSmDevice200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


