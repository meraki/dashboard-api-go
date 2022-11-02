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

// UpdateNetworkApplianceTrafficShapingRequest struct for UpdateNetworkApplianceTrafficShapingRequest
type UpdateNetworkApplianceTrafficShapingRequest struct {
	GlobalBandwidthLimits *UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits `json:"globalBandwidthLimits,omitempty"`
}

// NewUpdateNetworkApplianceTrafficShapingRequest instantiates a new UpdateNetworkApplianceTrafficShapingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceTrafficShapingRequest() *UpdateNetworkApplianceTrafficShapingRequest {
	this := UpdateNetworkApplianceTrafficShapingRequest{}
	return &this
}

// NewUpdateNetworkApplianceTrafficShapingRequestWithDefaults instantiates a new UpdateNetworkApplianceTrafficShapingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceTrafficShapingRequestWithDefaults() *UpdateNetworkApplianceTrafficShapingRequest {
	this := UpdateNetworkApplianceTrafficShapingRequest{}
	return &this
}

// GetGlobalBandwidthLimits returns the GlobalBandwidthLimits field value if set, zero value otherwise.
func (o *UpdateNetworkApplianceTrafficShapingRequest) GetGlobalBandwidthLimits() UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits {
	if o == nil || isNil(o.GlobalBandwidthLimits) {
		var ret UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits
		return ret
	}
	return *o.GlobalBandwidthLimits
}

// GetGlobalBandwidthLimitsOk returns a tuple with the GlobalBandwidthLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceTrafficShapingRequest) GetGlobalBandwidthLimitsOk() (*UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits, bool) {
	if o == nil || isNil(o.GlobalBandwidthLimits) {
    return nil, false
	}
	return o.GlobalBandwidthLimits, true
}

// HasGlobalBandwidthLimits returns a boolean if a field has been set.
func (o *UpdateNetworkApplianceTrafficShapingRequest) HasGlobalBandwidthLimits() bool {
	if o != nil && !isNil(o.GlobalBandwidthLimits) {
		return true
	}

	return false
}

// SetGlobalBandwidthLimits gets a reference to the given UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits and assigns it to the GlobalBandwidthLimits field.
func (o *UpdateNetworkApplianceTrafficShapingRequest) SetGlobalBandwidthLimits(v UpdateNetworkApplianceTrafficShapingRequestGlobalBandwidthLimits) {
	o.GlobalBandwidthLimits = &v
}

func (o UpdateNetworkApplianceTrafficShapingRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.GlobalBandwidthLimits) {
		toSerialize["globalBandwidthLimits"] = o.GlobalBandwidthLimits
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNetworkApplianceTrafficShapingRequest struct {
	value *UpdateNetworkApplianceTrafficShapingRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceTrafficShapingRequest) Get() *UpdateNetworkApplianceTrafficShapingRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRequest) Set(val *UpdateNetworkApplianceTrafficShapingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceTrafficShapingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceTrafficShapingRequest(val *UpdateNetworkApplianceTrafficShapingRequest) *NullableUpdateNetworkApplianceTrafficShapingRequest {
	return &NullableUpdateNetworkApplianceTrafficShapingRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceTrafficShapingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceTrafficShapingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


