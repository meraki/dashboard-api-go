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

// UpdateNetworkApplianceSettingsRequestDynamicDns Dynamic DNS settings for a network
type UpdateNetworkApplianceSettingsRequestDynamicDns struct {
	// Dynamic DNS url prefix. DDNS must be enabled to update
	Prefix *string `json:"prefix,omitempty"`
	// Dynamic DNS enabled
	Enabled *bool `json:"enabled,omitempty"`
}

// NewUpdateNetworkApplianceSettingsRequestDynamicDns instantiates a new UpdateNetworkApplianceSettingsRequestDynamicDns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceSettingsRequestDynamicDns() *UpdateNetworkApplianceSettingsRequestDynamicDns {
	this := UpdateNetworkApplianceSettingsRequestDynamicDns{}
	return &this
}

// NewUpdateNetworkApplianceSettingsRequestDynamicDnsWithDefaults instantiates a new UpdateNetworkApplianceSettingsRequestDynamicDns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceSettingsRequestDynamicDnsWithDefaults() *UpdateNetworkApplianceSettingsRequestDynamicDns {
	this := UpdateNetworkApplianceSettingsRequestDynamicDns{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSettingsRequestDynamicDns) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSettingsRequestDynamicDns) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSettingsRequestDynamicDns) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *UpdateNetworkApplianceSettingsRequestDynamicDns) SetPrefix(v string) {
	o.Prefix = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceSettingsRequestDynamicDns) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceSettingsRequestDynamicDns) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
    return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceSettingsRequestDynamicDns) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkApplianceSettingsRequestDynamicDns) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o UpdateNetworkApplianceSettingsRequestDynamicDns) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceSettingsRequestDynamicDns struct {
	value *UpdateNetworkApplianceSettingsRequestDynamicDns
	isSet bool
}

func (v NullableUpdateNetworkApplianceSettingsRequestDynamicDns) Get() *UpdateNetworkApplianceSettingsRequestDynamicDns {
	return v.value
}

func (v *NullableUpdateNetworkApplianceSettingsRequestDynamicDns) Set(val *UpdateNetworkApplianceSettingsRequestDynamicDns) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceSettingsRequestDynamicDns) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceSettingsRequestDynamicDns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceSettingsRequestDynamicDns(val *UpdateNetworkApplianceSettingsRequestDynamicDns) *NullableUpdateNetworkApplianceSettingsRequestDynamicDns {
	return &NullableUpdateNetworkApplianceSettingsRequestDynamicDns{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceSettingsRequestDynamicDns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceSettingsRequestDynamicDns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


