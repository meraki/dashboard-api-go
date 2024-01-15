/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 03 January, 2024 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.42.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest{}

// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct for UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest
type UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct {
	// If true, the SSID device type group policies are enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of device type policies.
	DeviceTypePolicies []UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner `json:"deviceTypePolicies,omitempty"`
}

// NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest instantiates a new UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest() *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest {
	this := UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest{}
	return &this
}

// NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestWithDefaults() *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest {
	this := UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDeviceTypePolicies returns the DeviceTypePolicies field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) GetDeviceTypePolicies() []UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner {
	if o == nil || IsNil(o.DeviceTypePolicies) {
		var ret []UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner
		return ret
	}
	return o.DeviceTypePolicies
}

// GetDeviceTypePoliciesOk returns a tuple with the DeviceTypePolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) GetDeviceTypePoliciesOk() ([]UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner, bool) {
	if o == nil || IsNil(o.DeviceTypePolicies) {
		return nil, false
	}
	return o.DeviceTypePolicies, true
}

// HasDeviceTypePolicies returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) HasDeviceTypePolicies() bool {
	if o != nil && !IsNil(o.DeviceTypePolicies) {
		return true
	}

	return false
}

// SetDeviceTypePolicies gets a reference to the given []UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner and assigns it to the DeviceTypePolicies field.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) SetDeviceTypePolicies(v []UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) {
	o.DeviceTypePolicies = v
}

func (o UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.DeviceTypePolicies) {
		toSerialize["deviceTypePolicies"] = o.DeviceTypePolicies
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest struct {
	value *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) Get() *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) Set(val *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest(val *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) *NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest {
	return &NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


