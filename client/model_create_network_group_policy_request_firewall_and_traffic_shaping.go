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

// CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping     The firewall and traffic shaping rules and settings for your policy. 
type CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping struct {
	// How firewall and traffic shaping rules are enforced. Can be 'network default', 'ignore' or 'custom'.
	Settings *string `json:"settings,omitempty"`
	//     An array of traffic shaping rules. Rules are applied in the order that     they are specified in. An empty list (or null) means no rules. Note that     you are allowed a maximum of 8 rules. 
	TrafficShapingRules []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner `json:"trafficShapingRules,omitempty"`
	// An ordered array of the L3 firewall rules
	L3FirewallRules []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner `json:"l3FirewallRules,omitempty"`
	// An ordered array of L7 firewall rules
	L7FirewallRules []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner `json:"l7FirewallRules,omitempty"`
}

// NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping {
	this := CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping{}
	return &this
}

// NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingWithDefaults instantiates a new CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkGroupPolicyRequestFirewallAndTrafficShapingWithDefaults() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping {
	this := CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetSettings() string {
	if o == nil || isNil(o.Settings) {
		var ret string
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetSettingsOk() (*string, bool) {
	if o == nil || isNil(o.Settings) {
    return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) HasSettings() bool {
	if o != nil && !isNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given string and assigns it to the Settings field.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) SetSettings(v string) {
	o.Settings = &v
}

// GetTrafficShapingRules returns the TrafficShapingRules field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetTrafficShapingRules() []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner {
	if o == nil || isNil(o.TrafficShapingRules) {
		var ret []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner
		return ret
	}
	return o.TrafficShapingRules
}

// GetTrafficShapingRulesOk returns a tuple with the TrafficShapingRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetTrafficShapingRulesOk() ([]CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner, bool) {
	if o == nil || isNil(o.TrafficShapingRules) {
    return nil, false
	}
	return o.TrafficShapingRules, true
}

// HasTrafficShapingRules returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) HasTrafficShapingRules() bool {
	if o != nil && !isNil(o.TrafficShapingRules) {
		return true
	}

	return false
}

// SetTrafficShapingRules gets a reference to the given []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner and assigns it to the TrafficShapingRules field.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) SetTrafficShapingRules(v []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingTrafficShapingRulesInner) {
	o.TrafficShapingRules = v
}

// GetL3FirewallRules returns the L3FirewallRules field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetL3FirewallRules() []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner {
	if o == nil || isNil(o.L3FirewallRules) {
		var ret []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner
		return ret
	}
	return o.L3FirewallRules
}

// GetL3FirewallRulesOk returns a tuple with the L3FirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetL3FirewallRulesOk() ([]CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner, bool) {
	if o == nil || isNil(o.L3FirewallRules) {
    return nil, false
	}
	return o.L3FirewallRules, true
}

// HasL3FirewallRules returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) HasL3FirewallRules() bool {
	if o != nil && !isNil(o.L3FirewallRules) {
		return true
	}

	return false
}

// SetL3FirewallRules gets a reference to the given []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner and assigns it to the L3FirewallRules field.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) SetL3FirewallRules(v []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL3FirewallRulesInner) {
	o.L3FirewallRules = v
}

// GetL7FirewallRules returns the L7FirewallRules field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetL7FirewallRules() []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner {
	if o == nil || isNil(o.L7FirewallRules) {
		var ret []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner
		return ret
	}
	return o.L7FirewallRules
}

// GetL7FirewallRulesOk returns a tuple with the L7FirewallRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) GetL7FirewallRulesOk() ([]CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner, bool) {
	if o == nil || isNil(o.L7FirewallRules) {
    return nil, false
	}
	return o.L7FirewallRules, true
}

// HasL7FirewallRules returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) HasL7FirewallRules() bool {
	if o != nil && !isNil(o.L7FirewallRules) {
		return true
	}

	return false
}

// SetL7FirewallRules gets a reference to the given []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner and assigns it to the L7FirewallRules field.
func (o *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) SetL7FirewallRules(v []CreateNetworkGroupPolicyRequestFirewallAndTrafficShapingL7FirewallRulesInner) {
	o.L7FirewallRules = v
}

func (o CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	if !isNil(o.TrafficShapingRules) {
		toSerialize["trafficShapingRules"] = o.TrafficShapingRules
	}
	if !isNil(o.L3FirewallRules) {
		toSerialize["l3FirewallRules"] = o.L3FirewallRules
	}
	if !isNil(o.L7FirewallRules) {
		toSerialize["l7FirewallRules"] = o.L7FirewallRules
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping struct {
	value *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping
	isSet bool
}

func (v NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) Get() *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping {
	return v.value
}

func (v *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) Set(val *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping(val *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping {
	return &NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping{value: val, isSet: true}
}

func (v NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


