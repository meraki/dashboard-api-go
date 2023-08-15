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

// checks if the UpdateNetworkSwitchRoutingMulticastRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkSwitchRoutingMulticastRequest{}

// UpdateNetworkSwitchRoutingMulticastRequest struct for UpdateNetworkSwitchRoutingMulticastRequest
type UpdateNetworkSwitchRoutingMulticastRequest struct {
	DefaultSettings *UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings `json:"defaultSettings,omitempty"`
	// Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings.
	Overrides []UpdateNetworkSwitchRoutingMulticastRequestOverridesInner `json:"overrides,omitempty"`
}

// NewUpdateNetworkSwitchRoutingMulticastRequest instantiates a new UpdateNetworkSwitchRoutingMulticastRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkSwitchRoutingMulticastRequest() *UpdateNetworkSwitchRoutingMulticastRequest {
	this := UpdateNetworkSwitchRoutingMulticastRequest{}
	return &this
}

// NewUpdateNetworkSwitchRoutingMulticastRequestWithDefaults instantiates a new UpdateNetworkSwitchRoutingMulticastRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkSwitchRoutingMulticastRequestWithDefaults() *UpdateNetworkSwitchRoutingMulticastRequest {
	this := UpdateNetworkSwitchRoutingMulticastRequest{}
	return &this
}

// GetDefaultSettings returns the DefaultSettings field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingMulticastRequest) GetDefaultSettings() UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings {
	if o == nil || IsNil(o.DefaultSettings) {
		var ret UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings
		return ret
	}
	return *o.DefaultSettings
}

// GetDefaultSettingsOk returns a tuple with the DefaultSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequest) GetDefaultSettingsOk() (*UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings, bool) {
	if o == nil || IsNil(o.DefaultSettings) {
		return nil, false
	}
	return o.DefaultSettings, true
}

// HasDefaultSettings returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequest) HasDefaultSettings() bool {
	if o != nil && !IsNil(o.DefaultSettings) {
		return true
	}

	return false
}

// SetDefaultSettings gets a reference to the given UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings and assigns it to the DefaultSettings field.
func (o *UpdateNetworkSwitchRoutingMulticastRequest) SetDefaultSettings(v UpdateNetworkSwitchRoutingMulticastRequestDefaultSettings) {
	o.DefaultSettings = &v
}

// GetOverrides returns the Overrides field value if set, zero value otherwise.
func (o *UpdateNetworkSwitchRoutingMulticastRequest) GetOverrides() []UpdateNetworkSwitchRoutingMulticastRequestOverridesInner {
	if o == nil || IsNil(o.Overrides) {
		var ret []UpdateNetworkSwitchRoutingMulticastRequestOverridesInner
		return ret
	}
	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequest) GetOverridesOk() ([]UpdateNetworkSwitchRoutingMulticastRequestOverridesInner, bool) {
	if o == nil || IsNil(o.Overrides) {
		return nil, false
	}
	return o.Overrides, true
}

// HasOverrides returns a boolean if a field has been set.
func (o *UpdateNetworkSwitchRoutingMulticastRequest) HasOverrides() bool {
	if o != nil && !IsNil(o.Overrides) {
		return true
	}

	return false
}

// SetOverrides gets a reference to the given []UpdateNetworkSwitchRoutingMulticastRequestOverridesInner and assigns it to the Overrides field.
func (o *UpdateNetworkSwitchRoutingMulticastRequest) SetOverrides(v []UpdateNetworkSwitchRoutingMulticastRequestOverridesInner) {
	o.Overrides = v
}

func (o UpdateNetworkSwitchRoutingMulticastRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkSwitchRoutingMulticastRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultSettings) {
		toSerialize["defaultSettings"] = o.DefaultSettings
	}
	if !IsNil(o.Overrides) {
		toSerialize["overrides"] = o.Overrides
	}
	return toSerialize, nil
}

type NullableUpdateNetworkSwitchRoutingMulticastRequest struct {
	value *UpdateNetworkSwitchRoutingMulticastRequest
	isSet bool
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRequest) Get() *UpdateNetworkSwitchRoutingMulticastRequest {
	return v.value
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRequest) Set(val *UpdateNetworkSwitchRoutingMulticastRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkSwitchRoutingMulticastRequest(val *UpdateNetworkSwitchRoutingMulticastRequest) *NullableUpdateNetworkSwitchRoutingMulticastRequest {
	return &NullableUpdateNetworkSwitchRoutingMulticastRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkSwitchRoutingMulticastRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkSwitchRoutingMulticastRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


