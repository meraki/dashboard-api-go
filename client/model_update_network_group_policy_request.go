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

// checks if the UpdateNetworkGroupPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNetworkGroupPolicyRequest{}

// UpdateNetworkGroupPolicyRequest struct for UpdateNetworkGroupPolicyRequest
type UpdateNetworkGroupPolicyRequest struct {
	// The name for your group policy.
	Name *string `json:"name,omitempty"`
	Scheduling *CreateNetworkGroupPolicyRequestScheduling `json:"scheduling,omitempty"`
	Bandwidth *CreateNetworkGroupPolicyRequestBandwidth `json:"bandwidth,omitempty"`
	FirewallAndTrafficShaping *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping `json:"firewallAndTrafficShaping,omitempty"`
	ContentFiltering *CreateNetworkGroupPolicyRequestContentFiltering `json:"contentFiltering,omitempty"`
	// Whether clients bound to your policy will bypass splash authorization or behave according to the network's rules. Can be one of 'network default' or 'bypass'. Only available if your network has a wireless configuration.
	SplashAuthSettings *string `json:"splashAuthSettings,omitempty"`
	VlanTagging *CreateNetworkGroupPolicyRequestVlanTagging `json:"vlanTagging,omitempty"`
	BonjourForwarding *CreateNetworkGroupPolicyRequestBonjourForwarding `json:"bonjourForwarding,omitempty"`
}

// NewUpdateNetworkGroupPolicyRequest instantiates a new UpdateNetworkGroupPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNetworkGroupPolicyRequest() *UpdateNetworkGroupPolicyRequest {
	this := UpdateNetworkGroupPolicyRequest{}
	return &this
}

// NewUpdateNetworkGroupPolicyRequestWithDefaults instantiates a new UpdateNetworkGroupPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNetworkGroupPolicyRequestWithDefaults() *UpdateNetworkGroupPolicyRequest {
	this := UpdateNetworkGroupPolicyRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNetworkGroupPolicyRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkGroupPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNetworkGroupPolicyRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNetworkGroupPolicyRequest) SetName(v string) {
	o.Name = &v
}

// GetScheduling returns the Scheduling field value if set, zero value otherwise.
func (o *UpdateNetworkGroupPolicyRequest) GetScheduling() CreateNetworkGroupPolicyRequestScheduling {
	if o == nil || IsNil(o.Scheduling) {
		var ret CreateNetworkGroupPolicyRequestScheduling
		return ret
	}
	return *o.Scheduling
}

// GetSchedulingOk returns a tuple with the Scheduling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkGroupPolicyRequest) GetSchedulingOk() (*CreateNetworkGroupPolicyRequestScheduling, bool) {
	if o == nil || IsNil(o.Scheduling) {
		return nil, false
	}
	return o.Scheduling, true
}

// HasScheduling returns a boolean if a field has been set.
func (o *UpdateNetworkGroupPolicyRequest) HasScheduling() bool {
	if o != nil && !IsNil(o.Scheduling) {
		return true
	}

	return false
}

// SetScheduling gets a reference to the given CreateNetworkGroupPolicyRequestScheduling and assigns it to the Scheduling field.
func (o *UpdateNetworkGroupPolicyRequest) SetScheduling(v CreateNetworkGroupPolicyRequestScheduling) {
	o.Scheduling = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *UpdateNetworkGroupPolicyRequest) GetBandwidth() CreateNetworkGroupPolicyRequestBandwidth {
	if o == nil || IsNil(o.Bandwidth) {
		var ret CreateNetworkGroupPolicyRequestBandwidth
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkGroupPolicyRequest) GetBandwidthOk() (*CreateNetworkGroupPolicyRequestBandwidth, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *UpdateNetworkGroupPolicyRequest) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given CreateNetworkGroupPolicyRequestBandwidth and assigns it to the Bandwidth field.
func (o *UpdateNetworkGroupPolicyRequest) SetBandwidth(v CreateNetworkGroupPolicyRequestBandwidth) {
	o.Bandwidth = &v
}

// GetFirewallAndTrafficShaping returns the FirewallAndTrafficShaping field value if set, zero value otherwise.
func (o *UpdateNetworkGroupPolicyRequest) GetFirewallAndTrafficShaping() CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping {
	if o == nil || IsNil(o.FirewallAndTrafficShaping) {
		var ret CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping
		return ret
	}
	return *o.FirewallAndTrafficShaping
}

// GetFirewallAndTrafficShapingOk returns a tuple with the FirewallAndTrafficShaping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkGroupPolicyRequest) GetFirewallAndTrafficShapingOk() (*CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping, bool) {
	if o == nil || IsNil(o.FirewallAndTrafficShaping) {
		return nil, false
	}
	return o.FirewallAndTrafficShaping, true
}

// HasFirewallAndTrafficShaping returns a boolean if a field has been set.
func (o *UpdateNetworkGroupPolicyRequest) HasFirewallAndTrafficShaping() bool {
	if o != nil && !IsNil(o.FirewallAndTrafficShaping) {
		return true
	}

	return false
}

// SetFirewallAndTrafficShaping gets a reference to the given CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping and assigns it to the FirewallAndTrafficShaping field.
func (o *UpdateNetworkGroupPolicyRequest) SetFirewallAndTrafficShaping(v CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) {
	o.FirewallAndTrafficShaping = &v
}

// GetContentFiltering returns the ContentFiltering field value if set, zero value otherwise.
func (o *UpdateNetworkGroupPolicyRequest) GetContentFiltering() CreateNetworkGroupPolicyRequestContentFiltering {
	if o == nil || IsNil(o.ContentFiltering) {
		var ret CreateNetworkGroupPolicyRequestContentFiltering
		return ret
	}
	return *o.ContentFiltering
}

// GetContentFilteringOk returns a tuple with the ContentFiltering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkGroupPolicyRequest) GetContentFilteringOk() (*CreateNetworkGroupPolicyRequestContentFiltering, bool) {
	if o == nil || IsNil(o.ContentFiltering) {
		return nil, false
	}
	return o.ContentFiltering, true
}

// HasContentFiltering returns a boolean if a field has been set.
func (o *UpdateNetworkGroupPolicyRequest) HasContentFiltering() bool {
	if o != nil && !IsNil(o.ContentFiltering) {
		return true
	}

	return false
}

// SetContentFiltering gets a reference to the given CreateNetworkGroupPolicyRequestContentFiltering and assigns it to the ContentFiltering field.
func (o *UpdateNetworkGroupPolicyRequest) SetContentFiltering(v CreateNetworkGroupPolicyRequestContentFiltering) {
	o.ContentFiltering = &v
}

// GetSplashAuthSettings returns the SplashAuthSettings field value if set, zero value otherwise.
func (o *UpdateNetworkGroupPolicyRequest) GetSplashAuthSettings() string {
	if o == nil || IsNil(o.SplashAuthSettings) {
		var ret string
		return ret
	}
	return *o.SplashAuthSettings
}

// GetSplashAuthSettingsOk returns a tuple with the SplashAuthSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkGroupPolicyRequest) GetSplashAuthSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.SplashAuthSettings) {
		return nil, false
	}
	return o.SplashAuthSettings, true
}

// HasSplashAuthSettings returns a boolean if a field has been set.
func (o *UpdateNetworkGroupPolicyRequest) HasSplashAuthSettings() bool {
	if o != nil && !IsNil(o.SplashAuthSettings) {
		return true
	}

	return false
}

// SetSplashAuthSettings gets a reference to the given string and assigns it to the SplashAuthSettings field.
func (o *UpdateNetworkGroupPolicyRequest) SetSplashAuthSettings(v string) {
	o.SplashAuthSettings = &v
}

// GetVlanTagging returns the VlanTagging field value if set, zero value otherwise.
func (o *UpdateNetworkGroupPolicyRequest) GetVlanTagging() CreateNetworkGroupPolicyRequestVlanTagging {
	if o == nil || IsNil(o.VlanTagging) {
		var ret CreateNetworkGroupPolicyRequestVlanTagging
		return ret
	}
	return *o.VlanTagging
}

// GetVlanTaggingOk returns a tuple with the VlanTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkGroupPolicyRequest) GetVlanTaggingOk() (*CreateNetworkGroupPolicyRequestVlanTagging, bool) {
	if o == nil || IsNil(o.VlanTagging) {
		return nil, false
	}
	return o.VlanTagging, true
}

// HasVlanTagging returns a boolean if a field has been set.
func (o *UpdateNetworkGroupPolicyRequest) HasVlanTagging() bool {
	if o != nil && !IsNil(o.VlanTagging) {
		return true
	}

	return false
}

// SetVlanTagging gets a reference to the given CreateNetworkGroupPolicyRequestVlanTagging and assigns it to the VlanTagging field.
func (o *UpdateNetworkGroupPolicyRequest) SetVlanTagging(v CreateNetworkGroupPolicyRequestVlanTagging) {
	o.VlanTagging = &v
}

// GetBonjourForwarding returns the BonjourForwarding field value if set, zero value otherwise.
func (o *UpdateNetworkGroupPolicyRequest) GetBonjourForwarding() CreateNetworkGroupPolicyRequestBonjourForwarding {
	if o == nil || IsNil(o.BonjourForwarding) {
		var ret CreateNetworkGroupPolicyRequestBonjourForwarding
		return ret
	}
	return *o.BonjourForwarding
}

// GetBonjourForwardingOk returns a tuple with the BonjourForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNetworkGroupPolicyRequest) GetBonjourForwardingOk() (*CreateNetworkGroupPolicyRequestBonjourForwarding, bool) {
	if o == nil || IsNil(o.BonjourForwarding) {
		return nil, false
	}
	return o.BonjourForwarding, true
}

// HasBonjourForwarding returns a boolean if a field has been set.
func (o *UpdateNetworkGroupPolicyRequest) HasBonjourForwarding() bool {
	if o != nil && !IsNil(o.BonjourForwarding) {
		return true
	}

	return false
}

// SetBonjourForwarding gets a reference to the given CreateNetworkGroupPolicyRequestBonjourForwarding and assigns it to the BonjourForwarding field.
func (o *UpdateNetworkGroupPolicyRequest) SetBonjourForwarding(v CreateNetworkGroupPolicyRequestBonjourForwarding) {
	o.BonjourForwarding = &v
}

func (o UpdateNetworkGroupPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNetworkGroupPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Scheduling) {
		toSerialize["scheduling"] = o.Scheduling
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if !IsNil(o.FirewallAndTrafficShaping) {
		toSerialize["firewallAndTrafficShaping"] = o.FirewallAndTrafficShaping
	}
	if !IsNil(o.ContentFiltering) {
		toSerialize["contentFiltering"] = o.ContentFiltering
	}
	if !IsNil(o.SplashAuthSettings) {
		toSerialize["splashAuthSettings"] = o.SplashAuthSettings
	}
	if !IsNil(o.VlanTagging) {
		toSerialize["vlanTagging"] = o.VlanTagging
	}
	if !IsNil(o.BonjourForwarding) {
		toSerialize["bonjourForwarding"] = o.BonjourForwarding
	}
	return toSerialize, nil
}

type NullableUpdateNetworkGroupPolicyRequest struct {
	value *UpdateNetworkGroupPolicyRequest
	isSet bool
}

func (v NullableUpdateNetworkGroupPolicyRequest) Get() *UpdateNetworkGroupPolicyRequest {
	return v.value
}

func (v *NullableUpdateNetworkGroupPolicyRequest) Set(val *UpdateNetworkGroupPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNetworkGroupPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNetworkGroupPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNetworkGroupPolicyRequest(val *UpdateNetworkGroupPolicyRequest) *NullableUpdateNetworkGroupPolicyRequest {
	return &NullableUpdateNetworkGroupPolicyRequest{value: val, isSet: true}
}

func (v NullableUpdateNetworkGroupPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNetworkGroupPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


