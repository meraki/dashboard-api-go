/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 02 November, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.27.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// RemoveNetworkDevicesRequest struct for RemoveNetworkDevicesRequest
type RemoveNetworkDevicesRequest struct {
	// The serial of a device
	Serial string `json:"serial"`
}

// NewRemoveNetworkDevicesRequest instantiates a new RemoveNetworkDevicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveNetworkDevicesRequest(serial string) *RemoveNetworkDevicesRequest {
	this := RemoveNetworkDevicesRequest{}
	this.Serial = serial
	return &this
}

// NewRemoveNetworkDevicesRequestWithDefaults instantiates a new RemoveNetworkDevicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveNetworkDevicesRequestWithDefaults() *RemoveNetworkDevicesRequest {
	this := RemoveNetworkDevicesRequest{}
	return &this
}

// GetSerial returns the Serial field value
func (o *RemoveNetworkDevicesRequest) GetSerial() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Serial
}

// GetSerialOk returns a tuple with the Serial field value
// and a boolean to check if the value has been set.
func (o *RemoveNetworkDevicesRequest) GetSerialOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Serial, true
}

// SetSerial sets field value
func (o *RemoveNetworkDevicesRequest) SetSerial(v string) {
	o.Serial = v
}

func (o RemoveNetworkDevicesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["serial"] = o.Serial
	}
	return json.Marshal(toSerialize)
}

type NullableRemoveNetworkDevicesRequest struct {
	value *RemoveNetworkDevicesRequest
	isSet bool
}

func (v NullableRemoveNetworkDevicesRequest) Get() *RemoveNetworkDevicesRequest {
	return v.value
}

func (v *NullableRemoveNetworkDevicesRequest) Set(val *RemoveNetworkDevicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveNetworkDevicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveNetworkDevicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveNetworkDevicesRequest(val *RemoveNetworkDevicesRequest) *NullableRemoveNetworkDevicesRequest {
	return &NullableRemoveNetworkDevicesRequest{value: val, isSet: true}
}

func (v NullableRemoveNetworkDevicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveNetworkDevicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


