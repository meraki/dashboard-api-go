/*
Meraki Dashboard API

A RESTful API to programmatically manage and monitor Cisco Meraki networks at scale.  > Date: 01 November, 2023 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.39.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest{}

// UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest struct for UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest
type UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest struct {
	// Custom VPN exclusion rules. Pass an empty array to clear existing rules.
	Custom []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner `json:"custom,omitempty"`
	// Major Application based VPN exclusion rules. Pass an empty array to clear existing rules.
	MajorApplications []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner `json:"majorApplications,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest() *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequestWithDefaults() *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest {
	this := UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest{}
	return &this
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) GetCustom() []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner {
	if o == nil || IsNil(o.Custom) {
		var ret []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner
		return ret
	}
	return o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) GetCustomOk() ([]UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner and assigns it to the Custom field.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) SetCustom(v []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestCustomInner) {
	o.Custom = v
}

// GetMajorApplications returns the MajorApplications field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) GetMajorApplications() []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner {
	if o == nil || IsNil(o.MajorApplications) {
		var ret []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner
		return ret
	}
	return o.MajorApplications
}

// GetMajorApplicationsOk returns a tuple with the MajorApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) GetMajorApplicationsOk() ([]UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner, bool) {
	if o == nil || IsNil(o.MajorApplications) {
		return nil, false
	}
	return o.MajorApplications, true
}

// HasMajorApplications returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) HasMajorApplications() bool {
	if o != nil && !IsNil(o.MajorApplications) {
		return true
	}

	return false
}

// SetMajorApplications gets a reference to the given []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner and assigns it to the MajorApplications field.
func (o *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) SetMajorApplications(v []UpdateNetworkApplianceTrafficShapingVpnExclusionsRequestMajorApplicationsInner) {
	o.MajorApplications = v
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.MajorApplications) {
		toSerialize["majorApplications"] = o.MajorApplications
	}
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest struct {
	value *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) Get() *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) Set(val *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest(val *UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest {
	return &NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


