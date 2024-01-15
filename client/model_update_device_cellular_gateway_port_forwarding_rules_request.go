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

// checks if the UpdateDeviceCellularGatewayPortForwardingRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDeviceCellularGatewayPortForwardingRulesRequest{}

// UpdateDeviceCellularGatewayPortForwardingRulesRequest struct for UpdateDeviceCellularGatewayPortForwardingRulesRequest
type UpdateDeviceCellularGatewayPortForwardingRulesRequest struct {
	// An array of port forwarding params
	Rules []UpdateDeviceCellularGatewayPortForwardingRulesRequestRulesInner `json:"rules,omitempty"`
}

// NewUpdateDeviceCellularGatewayPortForwardingRulesRequest instantiates a new UpdateDeviceCellularGatewayPortForwardingRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceCellularGatewayPortForwardingRulesRequest() *UpdateDeviceCellularGatewayPortForwardingRulesRequest {
	this := UpdateDeviceCellularGatewayPortForwardingRulesRequest{}
	return &this
}

// NewUpdateDeviceCellularGatewayPortForwardingRulesRequestWithDefaults instantiates a new UpdateDeviceCellularGatewayPortForwardingRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceCellularGatewayPortForwardingRulesRequestWithDefaults() *UpdateDeviceCellularGatewayPortForwardingRulesRequest {
	this := UpdateDeviceCellularGatewayPortForwardingRulesRequest{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *UpdateDeviceCellularGatewayPortForwardingRulesRequest) GetRules() []UpdateDeviceCellularGatewayPortForwardingRulesRequestRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []UpdateDeviceCellularGatewayPortForwardingRulesRequestRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceCellularGatewayPortForwardingRulesRequest) GetRulesOk() ([]UpdateDeviceCellularGatewayPortForwardingRulesRequestRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *UpdateDeviceCellularGatewayPortForwardingRulesRequest) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []UpdateDeviceCellularGatewayPortForwardingRulesRequestRulesInner and assigns it to the Rules field.
func (o *UpdateDeviceCellularGatewayPortForwardingRulesRequest) SetRules(v []UpdateDeviceCellularGatewayPortForwardingRulesRequestRulesInner) {
	o.Rules = v
}

func (o UpdateDeviceCellularGatewayPortForwardingRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDeviceCellularGatewayPortForwardingRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableUpdateDeviceCellularGatewayPortForwardingRulesRequest struct {
	value *UpdateDeviceCellularGatewayPortForwardingRulesRequest
	isSet bool
}

func (v NullableUpdateDeviceCellularGatewayPortForwardingRulesRequest) Get() *UpdateDeviceCellularGatewayPortForwardingRulesRequest {
	return v.value
}

func (v *NullableUpdateDeviceCellularGatewayPortForwardingRulesRequest) Set(val *UpdateDeviceCellularGatewayPortForwardingRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceCellularGatewayPortForwardingRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceCellularGatewayPortForwardingRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceCellularGatewayPortForwardingRulesRequest(val *UpdateDeviceCellularGatewayPortForwardingRulesRequest) *NullableUpdateDeviceCellularGatewayPortForwardingRulesRequest {
	return &NullableUpdateDeviceCellularGatewayPortForwardingRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateDeviceCellularGatewayPortForwardingRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceCellularGatewayPortForwardingRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


