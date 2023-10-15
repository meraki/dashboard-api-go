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

// checks if the UpdateDeviceApplianceUplinksSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceApplianceUplinksSettingsRequest{}

// UpdateDeviceApplianceUplinksSettingsRequest struct for UpdateDeviceApplianceUplinksSettingsRequest
type UpdateDeviceApplianceUplinksSettingsRequest struct {
	Interfaces UpdateDeviceApplianceUplinksSettingsRequestInterfaces `json:"interfaces"`
}

// NewUpdateDeviceApplianceUplinksSettingsRequest instantiates a new UpdateDeviceApplianceUplinksSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceApplianceUplinksSettingsRequest(interfaces UpdateDeviceApplianceUplinksSettingsRequestInterfaces) *UpdateDeviceApplianceUplinksSettingsRequest {
	this := UpdateDeviceApplianceUplinksSettingsRequest{}
	this.Interfaces = interfaces
	return &this
}

// NewUpdateDeviceApplianceUplinksSettingsRequestWithDefaults instantiates a new UpdateDeviceApplianceUplinksSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceApplianceUplinksSettingsRequestWithDefaults() *UpdateDeviceApplianceUplinksSettingsRequest {
	this := UpdateDeviceApplianceUplinksSettingsRequest{}
	return &this
}

// GetInterfaces returns the Interfaces field value
func (o *UpdateDeviceApplianceUplinksSettingsRequest) GetInterfaces() UpdateDeviceApplianceUplinksSettingsRequestInterfaces {
	if o == nil {
		var ret UpdateDeviceApplianceUplinksSettingsRequestInterfaces
		return ret
	}

	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequest) GetInterfacesOk() (*UpdateDeviceApplianceUplinksSettingsRequestInterfaces, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interfaces, true
}

// SetInterfaces sets field value
func (o *UpdateDeviceApplianceUplinksSettingsRequest) SetInterfaces(v UpdateDeviceApplianceUplinksSettingsRequestInterfaces) {
	o.Interfaces = v
}

func (o UpdateDeviceApplianceUplinksSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceApplianceUplinksSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interfaces"] = o.Interfaces
	return toSerialize, nil
}

type NullableUpdateDeviceApplianceUplinksSettingsRequest struct {
	value *UpdateDeviceApplianceUplinksSettingsRequest
	isSet bool
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequest) Get() *UpdateDeviceApplianceUplinksSettingsRequest {
	return v.value
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequest) Set(val *UpdateDeviceApplianceUplinksSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceApplianceUplinksSettingsRequest(val *UpdateDeviceApplianceUplinksSettingsRequest) *NullableUpdateDeviceApplianceUplinksSettingsRequest {
	return &NullableUpdateDeviceApplianceUplinksSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


