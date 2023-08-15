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

// checks if the UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner{}

// UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner struct for UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner
type UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner struct {
	// The device type. Can be one of 'Android', 'BlackBerry', 'Chrome OS', 'iPad', 'iPhone', 'iPod', 'Mac OS X', 'Windows', 'Windows Phone', 'B&N Nook' or 'Other OS'
	DeviceType string `json:"deviceType"`
	// The device policy. Can be one of 'Allowed', 'Blocked' or 'Group policy'
	DevicePolicy string `json:"devicePolicy"`
	// ID of the group policy object.
	GroupPolicyId *int32 `json:"groupPolicyId,omitempty"`
}

// NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner instantiates a new UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner(deviceType string, devicePolicy string) *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner {
	this := UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner{}
	this.DeviceType = deviceType
	this.DevicePolicy = devicePolicy
	return &this
}

// NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInnerWithDefaults instantiates a new UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInnerWithDefaults() *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner {
	this := UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner{}
	return &this
}

// GetDeviceType returns the DeviceType field value
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) GetDeviceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) GetDeviceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) SetDeviceType(v string) {
	o.DeviceType = v
}

// GetDevicePolicy returns the DevicePolicy field value
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) GetDevicePolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevicePolicy
}

// GetDevicePolicyOk returns a tuple with the DevicePolicy field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) GetDevicePolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevicePolicy, true
}

// SetDevicePolicy sets field value
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) SetDevicePolicy(v string) {
	o.DevicePolicy = v
}

// GetGroupPolicyId returns the GroupPolicyId field value if set, zero value otherwise.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) GetGroupPolicyId() int32 {
	if o == nil || IsNil(o.GroupPolicyId) {
		var ret int32
		return ret
	}
	return *o.GroupPolicyId
}

// GetGroupPolicyIdOk returns a tuple with the GroupPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) GetGroupPolicyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.GroupPolicyId) {
		return nil, false
	}
	return o.GroupPolicyId, true
}

// HasGroupPolicyId returns a boolean if a field has been set.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) HasGroupPolicyId() bool {
	if o != nil && !IsNil(o.GroupPolicyId) {
		return true
	}

	return false
}

// SetGroupPolicyId gets a reference to the given int32 and assigns it to the GroupPolicyId field.
func (o *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) SetGroupPolicyId(v int32) {
	o.GroupPolicyId = &v
}

func (o UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deviceType"] = o.DeviceType
	toSerialize["devicePolicy"] = o.DevicePolicy
	if !IsNil(o.GroupPolicyId) {
		toSerialize["groupPolicyId"] = o.GroupPolicyId
	}
	return toSerialize, nil
}

type NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner struct {
	value *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner
	isSet bool
}

func (v NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) Get() *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner {
	return v.value
}

func (v *NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) Set(val *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner(val *UpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) *NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner {
	return &NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner{value: val, isSet: true}
}

func (v NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkWirelessSsidDeviceTypeGroupPoliciesRequestDeviceTypePoliciesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


