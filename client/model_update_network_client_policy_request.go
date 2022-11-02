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

// UpdateNetworkClientPolicyRequest struct for UpdateNetworkClientPolicyRequest
type UpdateNetworkClientPolicyRequest struct {
	// The policy to assign. Can be 'Whitelisted', 'Blocked', 'Normal' or 'Group policy'. Required.
	DevicePolicy string `json:"devicePolicy"`
	// [optional] If 'devicePolicy' is set to 'Group policy' this param is used to specify the group policy ID.
	GroupPolicyId *string `json:"groupPolicyId,omitempty"`
}

// NewUpdateNetworkClientPolicyRequest instantiates a new UpdateNetworkClientPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkClientPolicyRequest(devicePolicy string) *UpdateNetworkClientPolicyRequest {
	this := UpdateNetworkClientPolicyRequest{}
	this.DevicePolicy = devicePolicy
	return &this
}

// NewUpdateNetworkClientPolicyRequestWithDefaults instantiates a new UpdateNetworkClientPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkClientPolicyRequestWithDefaults() *UpdateNetworkClientPolicyRequest {
	this := UpdateNetworkClientPolicyRequest{}
	return &this
}

// GetDevicePolicy returns the DevicePolicy field value
func (o *UpdateNetworkClientPolicyRequest) GetDevicePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientPolicyRequest) GetDevicePolicyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.DevicePolicy, true
}

// SetDevicePolicy sets field value
func (o *UpdateNetworkClientPolicyRequest) SetDevicePolicy(v string) {
	o.DevicePolicy = v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *UpdateNetworkClientPolicyRequest) GetGroupPolicyId() string {
	if o == nil || isNil(o.GroupPolicyId) {
		var ret string
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkClientPolicyRequest) GetGroupPolicyIdOk() (*string, bool) {
	if o == nil || isNil(o.GroupPolicyId) {
    return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *UpdateNetworkClientPolicyRequest) HasGroupPolicyId() bool {
	if o != nil && !isNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given string and assigns it to the GroupPolicyId field.
func (o *UpdateNetworkClientPolicyRequest) SetGroupPolicyId(v string) {
	o.GroupPolicyId = &v
}

func (o UpdateNetworkClientPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["devicePolicy"] = o.DevicePolicy
	}
	if !isNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkClientPolicyRequest struct {
	value *UpdateNetworkClientPolicyRequest
	isSet bool
}

func (v NullableUpdateNetworkClientPolicyRequest) Get() *UpdateNetworkClientPolicyRequest {
	return v.value
}

func (v *NullableUpdateNetworkClientPolicyRequest) Set(val *UpdateNetworkClientPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkClientPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkClientPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkClientPolicyRequest(val *UpdateNetworkClientPolicyRequest) *NullableUpdateNetworkClientPolicyRequest {
	return &NullableUpdateNetworkClientPolicyRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkClientPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkClientPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


