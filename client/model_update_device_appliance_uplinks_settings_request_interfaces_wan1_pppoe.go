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

// checks if the UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe{}

// UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe Configuration options for PPPoE.
type UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe struct {
	// Whether PPPoE is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	Authentication *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication `json:"authentication,omitempty"`
}

// NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe {
	this := UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe{}
	return &this
}

// NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeWithDefaults instantiates a new UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeWithDefaults() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe {
	this := UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) GetAuthentication() UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication {
	if o == nil || IsNil(o.Authentication) {
		var ret UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) GetAuthenticationOk() (*UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication and assigns it to the Authentication field.
func (o *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) SetAuthentication(v UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1PppoeAuthentication) {
	o.Authentication = &v
}

func (o UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return toSerialize, nil
}

type NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe struct {
	value *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe
	isSet bool
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) Get() *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe {
	return v.value
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) Set(val *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe(val *UpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe {
	return &NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe{value: val, isSet: true}
}

func (v NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceApplianceUplinksSettingsRequestInterfacesWan1Pppoe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


