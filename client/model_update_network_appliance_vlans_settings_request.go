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

// checks if the UpdateNetworkApplianceVlansSettingsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceVlansSettingsRequest{}

// UpdateNetworkApplianceVlansSettingsRequest struct for UpdateNetworkApplianceVlansSettingsRequest
type UpdateNetworkApplianceVlansSettingsRequest struct {
	// Boolean indicating whether to enable (true) or disable (false) VLANs for the network
	VlansEnabled *bool `json:"vlansEnabled,omitempty"`
}

// NewUpdateNetworkApplianceVlansSettingsRequest instantiates a new UpdateNetworkApplianceVlansSettingsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceVlansSettingsRequest() *UpdateNetworkApplianceVlansSettingsRequest {
	this := UpdateNetworkApplianceVlansSettingsRequest{}
	return &this
}

// NewUpdateNetworkApplianceVlansSettingsRequestWithDefaults instantiates a new UpdateNetworkApplianceVlansSettingsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceVlansSettingsRequestWithDefaults() *UpdateNetworkApplianceVlansSettingsRequest {
	this := UpdateNetworkApplianceVlansSettingsRequest{}
	return &this
}

// GetVlansEnabled returns the VlansEnabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceVlansSettingsRequest) GetVlansEnabled() bool {
	if o == nil || IsNil(o.VlansEnabled) {
		var ret bool
		return ret
	}
	return *o.VlansEnabled
}

// GetVlansEnabledOk returns a tuple with the VlansEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceVlansSettingsRequest) GetVlansEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.VlansEnabled) {
		return nil, false
	}
	return o.VlansEnabled, true
}

// HasVlansEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceVlansSettingsRequest) HasVlansEnabled() bool {
	if o != nil && !IsNil(o.VlansEnabled) {
		return true
	}

	return false
}

// SetVlansEnabled gets a reference to the given bool and assigns it to the VlansEnabled field.
func (o *UpdateNetworkApplianceVlansSettingsRequest) SetVlansEnabled(v bool) {
	o.VlansEnabled = &v
}

func (o UpdateNetworkApplianceVlansSettingsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceVlansSettingsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VlansEnabled) {
		toSerialize["vlansEnabled"] = o.VlansEnabled
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceVlansSettingsRequest struct {
	value *UpdateNetworkApplianceVlansSettingsRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceVlansSettingsRequest) Get() *UpdateNetworkApplianceVlansSettingsRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceVlansSettingsRequest) Set(val *UpdateNetworkApplianceVlansSettingsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceVlansSettingsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceVlansSettingsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceVlansSettingsRequest(val *UpdateNetworkApplianceVlansSettingsRequest) *NullableUpdateNetworkApplianceVlansSettingsRequest {
	return &NullableUpdateNetworkApplianceVlansSettingsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceVlansSettingsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceVlansSettingsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


