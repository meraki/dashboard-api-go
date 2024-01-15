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

// checks if the UpdateNetworkApplianceFirewallOneToOneNatRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkApplianceFirewallOneToOneNatRulesRequest{}

// UpdateNetworkApplianceFirewallOneToOneNatRulesRequest struct for UpdateNetworkApplianceFirewallOneToOneNatRulesRequest
type UpdateNetworkApplianceFirewallOneToOneNatRulesRequest struct {
	// An array of 1:1 nat rules
	Rules []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner `json:"rules"`
}

// NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequest instantiates a new UpdateNetworkApplianceFirewallOneToOneNatRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequest(rules []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) *UpdateNetworkApplianceFirewallOneToOneNatRulesRequest {
	this := UpdateNetworkApplianceFirewallOneToOneNatRulesRequest{}
	this.Rules = rules
	return &this
}

// NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestWithDefaults instantiates a new UpdateNetworkApplianceFirewallOneToOneNatRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestWithDefaults() *UpdateNetworkApplianceFirewallOneToOneNatRulesRequest {
	this := UpdateNetworkApplianceFirewallOneToOneNatRulesRequest{}
	return &this
}

// GetRules returns the Rules field value
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequest) GetRules() []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner {
	if o == nil {
		var ret []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequest) GetRulesOk() ([]UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesRequest) SetRules(v []UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner) {
	o.Rules = v
}

func (o UpdateNetworkApplianceFirewallOneToOneNatRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkApplianceFirewallOneToOneNatRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rules"] = o.Rules
	return toSerialize, nil
}

type NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequest struct {
	value *UpdateNetworkApplianceFirewallOneToOneNatRulesRequest
	isSet bool
}

func (v NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequest) Get() *UpdateNetworkApplianceFirewallOneToOneNatRulesRequest {
	return v.value
}

func (v *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequest) Set(val *UpdateNetworkApplianceFirewallOneToOneNatRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequest(val *UpdateNetworkApplianceFirewallOneToOneNatRulesRequest) *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequest {
	return &NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkApplianceFirewallOneToOneNatRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


