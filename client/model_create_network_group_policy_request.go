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

// checks if the CreateNetworkGroupPolicyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkGroupPolicyRequest{}

// CreateNetworkGroupPolicyRequest struct for CreateNetworkGroupPolicyRequest
type CreateNetworkGroupPolicyRequest struct {
	// The name for your group policy. Required.
	Name string `json:"name"`
	Scheduling *CreateNetworkGroupPolicyRequestScheduling `json:"scheduling,omitempty"`
	Bandwidth *CreateNetworkGroupPolicyRequestBandwidth `json:"bandwidth,omitempty"`
	FirewallAndTrafficShaping *CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping `json:"firewallAndTrafficShaping,omitempty"`
	ContentFiltering *CreateNetworkGroupPolicyRequestContentFiltering `json:"contentFiltering,omitempty"`
	// Whether clients bound to your policy will bypass splash authorization or behave according to the network's rules. Can be one of 'network default' or 'bypass'. Only available if your network has a wireless configuration.
	SplashAuthSettings *string `json:"splashAuthSettings,omitempty"`
	VlanTagging *CreateNetworkGroupPolicyRequestVlanTagging `json:"vlanTagging,omitempty"`
	BonjourForwarding *CreateNetworkGroupPolicyRequestBonjourForwarding `json:"bonjourForwarding,omitempty"`
}

// NewCreateNetworkGroupPolicyRequest instantiates a new CreateNetworkGroupPolicyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkGroupPolicyRequest(name string) *CreateNetworkGroupPolicyRequest {
	this := CreateNetworkGroupPolicyRequest{}
	this.Name = name
	return &this
}

// NewCreateNetworkGroupPolicyRequestWithDefaults instantiates a new CreateNetworkGroupPolicyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkGroupPolicyRequestWithDefaults() *CreateNetworkGroupPolicyRequest {
	this := CreateNetworkGroupPolicyRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNetworkGroupPolicyRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNetworkGroupPolicyRequest) SetName(v string) {
	o.Name = v
}

// GetScheduling returns the Scheduling field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequest) GetScheduling() CreateNetworkGroupPolicyRequestScheduling {
	if o == nil || IsNil(o.Scheduling) {
		var ret CreateNetworkGroupPolicyRequestScheduling
		return ret
	}
	return *o.Scheduling
}

// GetSchedulingOk returns a tuple with the Scheduling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequest) GetSchedulingOk() (*CreateNetworkGroupPolicyRequestScheduling, bool) {
	if o == nil || IsNil(o.Scheduling) {
		return nil, false
	}
	return o.Scheduling, true
}

// HasScheduling returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequest) HasScheduling() bool {
	if o != nil && !IsNil(o.Scheduling) {
		return true
	}

	return false
}

// SetScheduling gets a reference to the given CreateNetworkGroupPolicyRequestScheduling and assigns it to the Scheduling field.
func (o *CreateNetworkGroupPolicyRequest) SetScheduling(v CreateNetworkGroupPolicyRequestScheduling) {
	o.Scheduling = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequest) GetBandwidth() CreateNetworkGroupPolicyRequestBandwidth {
	if o == nil || IsNil(o.Bandwidth) {
		var ret CreateNetworkGroupPolicyRequestBandwidth
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequest) GetBandwidthOk() (*CreateNetworkGroupPolicyRequestBandwidth, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequest) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given CreateNetworkGroupPolicyRequestBandwidth and assigns it to the Bandwidth field.
func (o *CreateNetworkGroupPolicyRequest) SetBandwidth(v CreateNetworkGroupPolicyRequestBandwidth) {
	o.Bandwidth = &v
}

// GetFirewallAndTrafficShaping returns the FirewallAndTrafficShaping field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequest) GetFirewallAndTrafficShaping() CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping {
	if o == nil || IsNil(o.FirewallAndTrafficShaping) {
		var ret CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping
		return ret
	}
	return *o.FirewallAndTrafficShaping
}

// GetFirewallAndTrafficShapingOk returns a tuple with the FirewallAndTrafficShaping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequest) GetFirewallAndTrafficShapingOk() (*CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping, bool) {
	if o == nil || IsNil(o.FirewallAndTrafficShaping) {
		return nil, false
	}
	return o.FirewallAndTrafficShaping, true
}

// HasFirewallAndTrafficShaping returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequest) HasFirewallAndTrafficShaping() bool {
	if o != nil && !IsNil(o.FirewallAndTrafficShaping) {
		return true
	}

	return false
}

// SetFirewallAndTrafficShaping gets a reference to the given CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping and assigns it to the FirewallAndTrafficShaping field.
func (o *CreateNetworkGroupPolicyRequest) SetFirewallAndTrafficShaping(v CreateNetworkGroupPolicyRequestFirewallAndTrafficShaping) {
	o.FirewallAndTrafficShaping = &v
}

// GetContentFiltering returns the ContentFiltering field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequest) GetContentFiltering() CreateNetworkGroupPolicyRequestContentFiltering {
	if o == nil || IsNil(o.ContentFiltering) {
		var ret CreateNetworkGroupPolicyRequestContentFiltering
		return ret
	}
	return *o.ContentFiltering
}

// GetContentFilteringOk returns a tuple with the ContentFiltering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequest) GetContentFilteringOk() (*CreateNetworkGroupPolicyRequestContentFiltering, bool) {
	if o == nil || IsNil(o.ContentFiltering) {
		return nil, false
	}
	return o.ContentFiltering, true
}

// HasContentFiltering returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequest) HasContentFiltering() bool {
	if o != nil && !IsNil(o.ContentFiltering) {
		return true
	}

	return false
}

// SetContentFiltering gets a reference to the given CreateNetworkGroupPolicyRequestContentFiltering and assigns it to the ContentFiltering field.
func (o *CreateNetworkGroupPolicyRequest) SetContentFiltering(v CreateNetworkGroupPolicyRequestContentFiltering) {
	o.ContentFiltering = &v
}

// GetSplashAuthSettings returns the SplashAuthSettings field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequest) GetSplashAuthSettings() string {
	if o == nil || IsNil(o.SplashAuthSettings) {
		var ret string
		return ret
	}
	return *o.SplashAuthSettings
}

// GetSplashAuthSettingsOk returns a tuple with the SplashAuthSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequest) GetSplashAuthSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.SplashAuthSettings) {
		return nil, false
	}
	return o.SplashAuthSettings, true
}

// HasSplashAuthSettings returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequest) HasSplashAuthSettings() bool {
	if o != nil && !IsNil(o.SplashAuthSettings) {
		return true
	}

	return false
}

// SetSplashAuthSettings gets a reference to the given string and assigns it to the SplashAuthSettings field.
func (o *CreateNetworkGroupPolicyRequest) SetSplashAuthSettings(v string) {
	o.SplashAuthSettings = &v
}

// GetVlanTagging returns the VlanTagging field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequest) GetVlanTagging() CreateNetworkGroupPolicyRequestVlanTagging {
	if o == nil || IsNil(o.VlanTagging) {
		var ret CreateNetworkGroupPolicyRequestVlanTagging
		return ret
	}
	return *o.VlanTagging
}

// GetVlanTaggingOk returns a tuple with the VlanTagging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequest) GetVlanTaggingOk() (*CreateNetworkGroupPolicyRequestVlanTagging, bool) {
	if o == nil || IsNil(o.VlanTagging) {
		return nil, false
	}
	return o.VlanTagging, true
}

// HasVlanTagging returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequest) HasVlanTagging() bool {
	if o != nil && !IsNil(o.VlanTagging) {
		return true
	}

	return false
}

// SetVlanTagging gets a reference to the given CreateNetworkGroupPolicyRequestVlanTagging and assigns it to the VlanTagging field.
func (o *CreateNetworkGroupPolicyRequest) SetVlanTagging(v CreateNetworkGroupPolicyRequestVlanTagging) {
	o.VlanTagging = &v
}

// GetBonjourForwarding returns the BonjourForwarding field value if set, zero value otherwise.
func (o *CreateNetworkGroupPolicyRequest) GetBonjourForwarding() CreateNetworkGroupPolicyRequestBonjourForwarding {
	if o == nil || IsNil(o.BonjourForwarding) {
		var ret CreateNetworkGroupPolicyRequestBonjourForwarding
		return ret
	}
	return *o.BonjourForwarding
}

// GetBonjourForwardingOk returns a tuple with the BonjourForwarding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkGroupPolicyRequest) GetBonjourForwardingOk() (*CreateNetworkGroupPolicyRequestBonjourForwarding, bool) {
	if o == nil || IsNil(o.BonjourForwarding) {
		return nil, false
	}
	return o.BonjourForwarding, true
}

// HasBonjourForwarding returns a boolean if a field has been set.
func (o *CreateNetworkGroupPolicyRequest) HasBonjourForwarding() bool {
	if o != nil && !IsNil(o.BonjourForwarding) {
		return true
	}

	return false
}

// SetBonjourForwarding gets a reference to the given CreateNetworkGroupPolicyRequestBonjourForwarding and assigns it to the BonjourForwarding field.
func (o *CreateNetworkGroupPolicyRequest) SetBonjourForwarding(v CreateNetworkGroupPolicyRequestBonjourForwarding) {
	o.BonjourForwarding = &v
}

func (o CreateNetworkGroupPolicyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkGroupPolicyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
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

type NullableCreateNetworkGroupPolicyRequest struct {
	value *CreateNetworkGroupPolicyRequest
	isSet bool
}

func (v NullableCreateNetworkGroupPolicyRequest) Get() *CreateNetworkGroupPolicyRequest {
	return v.value
}

func (v *NullableCreateNetworkGroupPolicyRequest) Set(val *CreateNetworkGroupPolicyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNetworkGroupPolicyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNetworkGroupPolicyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNetworkGroupPolicyRequest(val *CreateNetworkGroupPolicyRequest) *NullableCreateNetworkGroupPolicyRequest {
	return &NullableCreateNetworkGroupPolicyRequest{value: val, isSet: true}
}

func (v NullableCreateNetworkGroupPolicyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNetworkGroupPolicyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


