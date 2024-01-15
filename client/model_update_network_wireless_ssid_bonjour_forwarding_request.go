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

// checks if the UpdateNetworkWirelessSsidBonjourForwardingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidBonjourForwardingRequest{}

// UpdateNetworkWirelessSsidBonjourForwardingRequest struct for UpdateNetworkWirelessSsidBonjourForwardingRequest
type UpdateNetworkWirelessSsidBonjourForwardingRequest struct {
	// If true, Bonjour forwarding is enabled on this SSID.
	Enabled *bool `json:"enabled,omitempty"`
	// List of bonjour forwarding rules.
	Rules []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner `json:"rules,omitempty"`
	Exception *UpdateNetworkWirelessSsidBonjourForwardingRequestException `json:"exception,omitempty"`
}

// NewUpdateNetworkWirelessSsidBonjourForwardingRequest instantiates a new UpdateNetworkWirelessSsidBonjourForwardingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidBonjourForwardingRequest() *UpdateNetworkWirelessSsidBonjourForwardingRequest {
	this := UpdateNetworkWirelessSsidBonjourForwardingRequest{}
	return &this
}

// NewUpdateNetworkWirelessSsidBonjourForwardingRequestWithDefaults instantiates a new UpdateNetworkWirelessSsidBonjourForwardingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidBonjourForwardingRequestWithDefaults() *UpdateNetworkWirelessSsidBonjourForwardingRequest {
	this := UpdateNetworkWirelessSsidBonjourForwardingRequest{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) GetRules() []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) GetRulesOk() ([]CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner and assigns it to the Rules field.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) SetRules(v []CreateNetworkGroupPolicyRequestBonjourForwardingRulesInner) {
	o.Rules = v
}

// GetException returns the Exception field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) GetException() UpdateNetworkWirelessSsidBonjourForwardingRequestException {
	if o == nil || IsNil(o.Exception) {
		var ret UpdateNetworkWirelessSsidBonjourForwardingRequestException
		return ret
	}
	return *o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) GetExceptionOk() (*UpdateNetworkWirelessSsidBonjourForwardingRequestException, bool) {
	if o == nil || IsNil(o.Exception) {
		return nil, false
	}
	return o.Exception, true
}

// HasException returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) HasException() bool {
	if o != nil && !IsNil(o.Exception) {
		return true
	}

	return false
}

// SetException gets a reference to the given UpdateNetworkWirelessSsidBonjourForwardingRequestException and assigns it to the Exception field.
func (o *UpdateNetworkWirelessSsidBonjourForwardingRequest) SetException(v UpdateNetworkWirelessSsidBonjourForwardingRequestException) {
	o.Exception = &v
}

func (o UpdateNetworkWirelessSsidBonjourForwardingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidBonjourForwardingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Exception) {
		toSerialize["exception"] = o.Exception
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidBonjourForwardingRequest struct {
	value *UpdateNetworkWirelessSsidBonjourForwardingRequest
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidBonjourForwardingRequest) Get() *UpdateNetworkWirelessSsidBonjourForwardingRequest {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidBonjourForwardingRequest) Set(val *UpdateNetworkWirelessSsidBonjourForwardingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidBonjourForwardingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidBonjourForwardingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidBonjourForwardingRequest(val *UpdateNetworkWirelessSsidBonjourForwardingRequest) *NullableUpdateNetworkWirelessSsidBonjourForwardingRequest {
	return &NullableUpdateNetworkWirelessSsidBonjourForwardingRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidBonjourForwardingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidBonjourForwardingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


