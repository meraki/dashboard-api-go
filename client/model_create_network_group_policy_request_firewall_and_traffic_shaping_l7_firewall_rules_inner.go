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

// CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner struct for CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner
type CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner struct {
	// The policy applied to matching traffic. Must be 'deny'.
	Policy *string `json:"policy,omitempty"`
	// Type of the L7 Rule. Must be 'application', 'applicationCategory', 'host', 'port' or 'ipRange'
	Type *string `json:"type,omitempty"`
	// The 'value' of what you want to block. If 'type' is 'host', 'port' or 'ipRange', 'value' must be a string matching either a hostname (e.g. somewhere.com), a port (e.g. 8080), or an IP range (e.g. 192.1.0.0/16). If 'type' is 'application' or 'applicationCategory', then 'value' must be an object with an ID for the application.
	Value *string `json:"value,omitempty"`
}

// NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner {
	this := CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner{}
	return &this
}

// NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInnerWithDefaults instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInnerWithDefaults() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner {
	this := CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetPolicy() string {
	if o == nil || isNil(o.Policy) {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetPolicyOk() (*string, bool) {
	if o == nil || isNil(o.Policy) {
    return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) HasPolicy() bool {
	if o != nil && !isNil(o.Policy) {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) SetPolicy(v string) {
	o.Policy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) SetValue(v string) {
	o.Value = &v
}

func (o CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Policy) {
		toSerialize["policy"] = o.Policy
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner struct {
	value *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner
	isSet bool
}

func (v NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) Get() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner {
	return v.value
}

func (v *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) Set(val *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner(val *CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner {
	return &NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner{value: val, isSet: true}
}

func (v NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


